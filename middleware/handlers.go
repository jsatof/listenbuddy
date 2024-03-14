package middleware

import (
	"context"
	"encoding/json"
	"listenbud/models"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn

func MakeDBConnection() error {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		return err
	}
	db = conn
	return nil
}

func insertUser(user models.User) int64 {
	_, err := db.Exec(context.Background(), "insert into listenuser (id,username,password) values ($1,$2,$3)", user.ID, user.Username, user.Password)
	if err != nil {
		log.Fatalf("Could not insert value (%v,%v)\n%v", user.ID, user.Username, err)
		return -1
	}
	return user.ID
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty user of type models.User
	var user models.User

	// decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call insert user function and pass the user
	insertID := insertUser(user)

	// format a response object
	type Response struct {
		ID      int64
		Message string
	}

	res := Response{
		ID:      insertID,
		Message: "User created successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}
