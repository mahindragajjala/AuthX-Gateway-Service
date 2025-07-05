package controllers

import (
	"authserver/apigateway/proto"
	auth "authserver/authentication"
	"authserver/db/mongodb"
	"authserver/db/postgres"
	"authserver/models"
	"authserver/utils"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceServer struct {
	proto.UnimplementedUserServiceServer
}

func (s *UserServiceServer) SendUser(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	fmt.Println("Received SendUser:", req.Email)
	return &proto.UserResponse{
		Message: "User received: " + req.Email,
	}, nil
}

func (s *UserServiceServer) SignupUser(ctx context.Context, req *proto.SignupRequest) (*proto.SignupResponse, error) {
	fmt.Println("Received SignupUser:", req.Email)

	// Validate input
	if req.Email == "" || req.Password == "" {
		return &proto.SignupResponse{
			Message: "Email or password is missing",
		}, nil
	}

	if !auth.IsValidEmail(req.Email) {
		fmt.Println("Invalid email format")
		return &proto.SignupResponse{
			Message: "Invalid email format: " + req.Email,
		}, nil
	}

	if len(req.Password) < 6 {
		fmt.Println("Password is too short")
		return &proto.SignupResponse{
			Message: "Password must be at least 6 characters",
		}, nil
	}

	// Check if email already exists
	exists, err := auth.IsEmailExists(req.Email)
	if err != nil {
		fmt.Println("Database error while checking email:", err)
		return &proto.SignupResponse{
			Message: "Internal server error while checking email",
		}, nil
	}

	if exists {
		fmt.Println("User already exists:", req.Email)
		return &proto.SignupResponse{
			Message: "User already exists with email: " + req.Email,
		}, nil
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return &proto.SignupResponse{
			Message: "Internal server error while hashing password",
		}, nil
	}

	// Insert new user
	userID := uuid.New()
	now := time.Now()

	query := `
		INSERT INTO users (
			id, email, password_hash, created_at, updated_at, is_verified, status, last_login, login_count, role
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err = postgres.DB.Exec(query,
		userID,
		req.Email,
		string(hashedPassword),
		now,
		now,
		false,       // is_verified
		"active",    // status
		time.Time{}, // last_login
		0,           // login_count
		"user",      // role
	)

	if err != nil {
		fmt.Println("Database error while inserting user:", err)
		return &proto.SignupResponse{
			Message: "Failed to insert user into database",
		}, nil
	}

	//JWT Token generation
	accessToken, refreshToken, err := utils.SignupHandler(req)
	if err != nil {
		fmt.Println("Creation of the access token and refresh error - signup")
	}

	fmt.Println("User successfully signed up:", req.Email)

	responsedata := &proto.SignupResponse{
		Message:      "User signed up successfully: " + req.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	fmt.Println(responsedata)

	// Build audit log
	log := models.AuditLog{
		Email:     req.Email,
		Action:    "signup",
		Timestamp: time.Now(),
	}

	err = mongodb.InsertAuditLog(log)
	if err != nil {
		fmt.Println("Error in the Inserting the monogdb...")
	}

	return &proto.SignupResponse{
		Message:      "Signup Successfull: " + req.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *UserServiceServer) LoginUser(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {

	var JWT_TOKEN string

	// Check if email exists
	exist, _ := postgres.IsEmailExists(req.Email)
	if !exist {
		fmt.Println("Email does not exist:", req.Email)
		return &proto.LoginResponse{
			Message:     "Login Un-Successful - Please Signup Again",
			AccessToken: "NOT CREATED",
		}, nil
	}

	// Get user data from DB
	userdata, err := postgres.GetUserByEmail(req.Email)
	if err != nil {
		fmt.Println("Error retrieving user data:", err)
		return &proto.LoginResponse{
			Message:     "Internal Server Error while retrieving user",
			AccessToken: "NOT CREATED",
		}, err
	}

	// Compare hashed password with user input
	err = bcrypt.CompareHashAndPassword([]byte(userdata.PasswordHash), []byte(req.Password))
	if err != nil {
		fmt.Println("Password mismatch for email:", req.Email)
		return &proto.LoginResponse{
			Message:     "Login Failed - Invalid Credentials",
			AccessToken: "NOT CREATED",
		}, nil
	}

	// Passwords match → generate JWT
	fmt.Println("Passwords match for email:", req.Email)

	accessToken, refreshToken, err := utils.LoginHandler(req)
	if err != nil {
		fmt.Println("Creation of the access token and refresh error - login")
	}
	JWT_TOKEN = accessToken
	fmt.Println("Generated JWT Token:", JWT_TOKEN)
	fmt.Println("User successfully logged in:", req.Email)

	// Build audit log
	log := models.AuditLog{
		Email:     req.Email,
		Action:    "login",
		Timestamp: time.Now(),
	}

	err = mongodb.InsertAuditLog(log)
	if err != nil {
		fmt.Println("Error in the Inserting the monogdb...")
	}

	// Final successful response
	return &proto.LoginResponse{
		Message:      "Login Successful: " + req.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
