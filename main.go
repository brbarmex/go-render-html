package main

import (
    "encoding/json"
    "net/http"
    "time"
)

type Auditoria struct {
    Versao   string    `json:"versao"`
    DataHora time.Time `json:"data_hora"`
    Executor string    `json:"executor"`
}

func getAuditorias(w http.ResponseWriter, r *http.Request) {

    auditorias := []Auditoria{
        {"1.0.0", time.Now(), "Executor1"},
        {"1.0.1", time.Now().Add(time.Hour), "Executor2"},
        {"1.0.2", time.Now().Add(2 * time.Hour), "Executor3"},
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(auditorias)
}

func main() {
	
    http.HandleFunc("/api/auditorias", getAuditorias)
    http.Handle("/", http.FileServer(http.Dir("./static")))

    http.ListenAndServe(":8080", nil)
}
