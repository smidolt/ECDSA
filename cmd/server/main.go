package main

import (
    "log"
    "net/http"
    "my_project/internal"
    "my_project/internal/server"
    "github.com/julienschmidt/httprouter"
)

func main() {
    db, err := internal.ConnectDB()
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    srv := server.NewServer(db)

    router := httprouter.New()
    router.POST("/api/status", srv.CreateStatus)
    router.GET("/api/status/:statusId/:index", srv.GetStatus)
    router.PUT("/api/status/:statusId/:index", srv.SetStatus)
    router.DELETE("/api/status/:statusId/:index", srv.DeleteStatus)
    router.GET("/api/status", srv.GetAllStatuses)

    log.Fatal(http.ListenAndServe(":8080", router))
}
