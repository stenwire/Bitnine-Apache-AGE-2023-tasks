package main
// querying all rows from the database with Go lang sql package
import (
    "database/sql"
    "fmt"
	"os"
	"log"
	"strconv"
	"github.com/joho/godotenv"
    _ "github.com/lib/pq"
)
const (
    host = "localhost"
    port = 5432
    user = "postgres"
    password = "despicable01"
    dbname = "bitnine"
)
var (
    ename string
    sal int
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func goDotEnvVariableInt(key string) int {

	// load .env file
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	// return os.Getenv(key)

	val := goDotEnvVariableInt(key)
	ret, err := strconv.Atoi(val)
    if err != nil {
        panic(fmt.Sprintf("some error"))
    }
    return ret

}

func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        goDotEnvVariable(host), goDotEnvVariableInt(port_go), goDotEnvVariable(user), goDotEnvVariable(password), goDotEnvVariable(dbname))

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    defer db.Close()
    rows, err := db.Query("SELECT user_id, name, age, phone FROM user_table ORDER BY user_id")
    if err != nil {
        panic(err)
    }
    defer rows.Close()
    for rows.Next() {
        err := rows.Scan(&ename, &sal)
        if err != nil {
            panic(err)
        }
        fmt.Println("\n", ename, sal)
    }
    err = rows.Err()
    if err != nil {
        panic(err)
    }
}
