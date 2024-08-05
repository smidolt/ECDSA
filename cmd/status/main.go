package main

import (
    "log"
    "net/http"
    "database/sql"
    "my_project/internal/server"
    "github.com/julienschmidt/httprouter"
    _ "github.com/lib/pq"
)

func main() {
    db, err := sql.Open("postgres", "user=myusername password=mypassword dbname=status_db sslmode=disable")
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    s := server.NewServer(db)
    router := httprouter.New()

    router.POST("/api/status", s.CreateStatus)
    router.GET("/api/status/:statusId/:index", s.GetStatus)
    router.PUT("/api/status/:statusId/:index", s.SetStatus)
    router.DELETE("/api/status/:statusId/:index", s.DeleteStatus)
    router.GET("/api/status", s.GetAllStatuses)

    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
