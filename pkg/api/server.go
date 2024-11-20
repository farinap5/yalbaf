package api

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
)

type Server struct {
	db *sql.DB
}

func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {id = "1"}

	data, err := s.QueryToJson(id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		x := `{"err":"` + err.Error() + `"}`
		w.Write([]byte(x))
	} else {
		x := `{"data":` + data + `}`
		w.Write([]byte(x))
	}
}

func StartServer() {
	server()
}

func server() {
	mux := http.NewServeMux()
	s := new(Server)
	err := s.createDatabase()
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = s.createData()
	if err != nil {
		log.Println(err.Error())
		return
	}

	mux.HandleFunc("/", s.home)

	host := ":8081"
	log.Println("Starting server on " + host)
	err = http.ListenAndServe(host, mux)
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server got closed\n")
	} else if err != nil {
		log.Printf("error listening for server: %s\n", err)
	}
	s.StopDatabase()
}