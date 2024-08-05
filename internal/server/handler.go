package server

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "my_project/internal"
    "github.com/julienschmidt/httprouter"
)

type Server struct {
    DB *sql.DB
}

func NewServer(db *sql.DB) *Server {
    return &Server{DB: db}
}

func (s *Server) CreateStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    statuses := internal.NewStatuses()
    encoded, err := statuses.EncodeBase64()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var id int
    err = s.DB.QueryRow("INSERT INTO statuses (data) VALUES ($1) RETURNING id", encoded).Scan(&id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (s *Server) GetStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id := ps.ByName("statusId")
    index, err := strconv.Atoi(ps.ByName("index"))
    if err != nil {
        http.Error(w, "Invalid index", http.StatusBadRequest)
        return
    }

    var data string
    err = s.DB.QueryRow("SELECT data FROM statuses WHERE id = $1", id).Scan(&data)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Status not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    statuses := &internal.Statuses{}
    err = statuses.DecodeBase64(data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    value, err := statuses.Get(index)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := map[string]interface{}{
        "id":    id,
        "index": index,
        "value": value,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func (s *Server) SetStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id := ps.ByName("statusId")
    index, err := strconv.Atoi(ps.ByName("index"))
    if err != nil {
        http.Error(w, "Invalid index", http.StatusBadRequest)
        return
    }

    log.Printf("Updating status for ID: %s, Index: %d", id, index)

    var data string
    err = s.DB.QueryRow("SELECT data FROM statuses WHERE id = $1", id).Scan(&data)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Status not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    statuses := &internal.Statuses{}
    err = statuses.DecodeBase64(data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := statuses.Set(index, true); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    log.Printf("Updated Statuses Data: %v", statuses.Data)

    encoded, err := statuses.EncodeBase64()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    log.Printf("Encoded Data: %s", encoded)

    _, err = s.DB.Exec("UPDATE statuses SET data = $1 WHERE id = $2", encoded, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    log.Println("Status updated successfully")
    w.WriteHeader(http.StatusNoContent)
}

func (s *Server) DeleteStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id := ps.ByName("statusId")
    index, err := strconv.Atoi(ps.ByName("index"))
    if err != nil {
        http.Error(w, "Invalid index", http.StatusBadRequest)
        return
    }

    log.Printf("Deleting status for ID: %s, Index: %d", id, index)

    var data string
    err = s.DB.QueryRow("SELECT data FROM statuses WHERE id = $1", id).Scan(&data)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Status not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    statuses := &internal.Statuses{}
    err = statuses.DecodeBase64(data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := statuses.Set(index, false); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    encoded, err := statuses.EncodeBase64()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    _, err = s.DB.Exec("UPDATE statuses SET data = $1 WHERE id = $2", encoded, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    log.Println("Status deleted successfully")
    w.WriteHeader(http.StatusNoContent)
}

func (s *Server) GetAllStatuses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    rows, err := s.DB.Query("SELECT id FROM statuses")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var ids []int
    for rows.Next() {
        var id int
        if err := rows.Scan(&id); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        ids = append(ids, id)
    }

    if err := rows.Err(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string][]int{"ids": ids})
}
