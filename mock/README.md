
# Gin Login Mock Test Example

## Steps

1. Run `mockgen` to generate mocks:
   ```sh
   mockgen -source=auth/auth.go -destination=auth/mocks/mock_auth.go -package=mocks
   ```

2. Run the tests:
   ```sh
   go test ./handlers
   ```
