package postgres

import (
	"fmt"
	"time"

	_ "github.com/lib/pq" // for PostgreSQL; change depending on your DB
)

func Delete_Table_Data() {
	DB.Exec("DELETE FROM users")
	time.Sleep(2 * time.Millisecond)
	fmt.Println("Table data deleted successfully.")
}
