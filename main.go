package main

import (
    "database/sql"
    "fmt"
    "net/http"
    "log"
    "os"
    _ "github.com/lib/pq"
    "time"
)

func sayhelloWorld(w http.ResponseWriter, r *http.Request) {
    log.Println("Responding to request")
    DB_ADDR := os.Getenv("PG_ADDR")
    DB_PORT := os.Getenv("PG_PORT")
    DB_USER := os.Getenv("PG_USER")
    DB_PASS := os.Getenv("PG_PASS")

    dbinfo := fmt.Sprintf(
      "host=%s port=%s user=%s password=%s sslmode=disable",
      DB_ADDR,
      DB_PORT,
      DB_USER,
      DB_PASS,
    )
    db, err := sql.Open("postgres", dbinfo)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    rows, err := db.Query("SELECT now()")
    if err != nil {
        panic(err)
    }
    for rows.Next() {
        var now time.Time
        err = rows.Scan(&now)
        if err != nil {
            panic(err)
        }
        fmt.Fprintf(w, "The time is: %s\n", now)
    }
}

func main() {
    log.Println("PG_ADDR:", os.Getenv("PG_ADDR"))
    log.Println("PG_PORT:", os.Getenv("PG_PORT"))
    log.Println("PG_USER:", os.Getenv("PG_USER"))
    log.Println("PG_PASS:", os.Getenv("PG_PASS"))

    http.HandleFunc("/", sayhelloWorld)
    err := http.ListenAndServe(":8888", nil)
    if err != nil {
        log.Fatal(err)
    }
}
