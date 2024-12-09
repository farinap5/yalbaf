package api

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// https://stackoverflow.com/questions/42774467/how-to-convert-sql-rows-to-typed-json-in-golang
func rowsToJSON(rows *sql.Rows) (string, error) {
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}

	var results []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))

		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return "", err
		}

		rowMap := make(map[string]interface{})
		for i, col := range columns {
			val := values[i]

			if b, ok := val.([]byte); ok {
				val = string(b)
			}
			rowMap[col] = val
		}

		results = append(results, rowMap)
	}

	jsonBytes, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func (s *Server)QueryToJson(id string) (string, error) {
	q := "SELECT brand, model FROM cars WHERE id=" + id + ";"
	log.Printf("q='%s'", q)
	
	rows, err := s.db.Query(q)
	if err != nil {
		return "",err
	}
	return rowsToJSON(rows)
}

func (s *Server)createData() error {
	usersTable := `
	CREATE TABLE users (
		id 		INTEGER PRIMARY KEY AUTOINCREMENT,
		name 	TEXT NOT NULL,
		email 	TEXT UNIQUE NOT NULL,
		passwd 	TEXT NOT NULL
	);`
	if _, err := s.db.Exec(usersTable); err != nil {
		return err
	}

	carsTable := `
	CREATE TABLE cars (
		id 			INTEGER PRIMARY KEY AUTOINCREMENT,
		brand 		TEXT NOT NULL,
		model 		TEXT NOT NULL,
		owner_id 	INTEGER,
		FOREIGN KEY (owner_id) REFERENCES users(id)
	);`
	if _, err := s.db.Exec(carsTable); err != nil {
		return err
	}

	insertUser := `
	INSERT INTO users (name, email, passwd)
	VALUES
		('Alice', 'alice@mail.com', 'key123#'),
		('Bob', 'bob@mail.com', 'key321#');`
	if _, err := s.db.Exec(insertUser); err != nil {
		return err
	}

	insertCars := `
	INSERT INTO cars (brand, model, owner_id)
	VALUES
		('Toyota', 'Corolla', 1),
		('Honda', 'Civic', 2);`
	if _, err := s.db.Exec(insertCars); err != nil {
		return err
	}
	

	return nil
}

func (s *Server)createDatabase() error {
	db, err := sql.Open("sqlite3", ":memory:")
	s.db = db
	return err
}

func (s *Server)StopDatabase() error {
	return s.db.Close()
}