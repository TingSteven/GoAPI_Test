package main

import (
    "net/http"

    "database/sql"
    "fmt"
    "log"
    "github.com/go-sql-driver/mysql"
    "github.com/gin-gonic/gin"
)

var db *sql.DB

type UserAccount struct {
    Id string       `json:"id"`
    Uname string    `json:"uname"`
    Memo string     `json:"memo"`
}

func main() {
    // Capture connection properties.
    cfg := mysql.Config{
        User:   "root",
        Passwd: "",
        Net:    "tcp",
        //Addr:   "127.0.0.1:3306",
        Addr:   "host.docker.internal",
        DBName: "test",
        AllowNativePasswords: true,
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    //fmt.Println("Connected!")

    router := gin.Default()
    router.GET("/testGet", getUserAccounts)

    router.Run("localhost:8080")
    //uAccounts, err := uAccountsByInfo("080158")
    //fmt.Println("ua", getUserAccounts)
}

func uAccountsByInfo(id string) ([]UserAccount, error) {
    // An useraccount slice to hold data from returned rows.
    var uAccounts []UserAccount

	rows, err := db.Query("SELECT * FROM user_account WHERE id = ?", id)
    if err != nil {
		return nil, fmt.Errorf("user_account %q: %v", id, err)
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var uac UserAccount
        if err := rows.Scan(&uac.Id, &uac.Uname, &uac.Memo); err != nil {
            return nil, fmt.Errorf("user_account %q: %v", id, err)
        }
        uAccounts = append(uAccounts, uac)
    }
    if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("user_account %q: %v", id, err)
    }
    return uAccounts, nil
}

func getUserAccounts(c *gin.Context) {
    uAccounts, err := uAccountsByInfo("080158")
    if err != nil {
       log.Fatal(err)
    }
    c.IndentedJSON(http.StatusOK, uAccounts)
}