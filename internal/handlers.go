package internal

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jsatof/listenbuddy/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type User = models.User
type SongRequest = models.SongRequest

var db *pgx.Conn

func insertSongRequest() {

}

func searchSongRequest() {

}

func deleteSongRequest() {

}

func updateSongRequest() {

}

func updateUser() {

}

func deleteUser() {

}

func searchUser() {

}

func insertUser(user User) int64 {
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
	var user User

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

func GenerateHash(s string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return ""
	}

	return string(hash)
}
