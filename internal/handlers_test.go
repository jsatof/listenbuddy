package internal

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func TestDBConnection(t *testing.T) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}
	conn.Close(context.Background())
}

func TestDatabaseCRUD(t *testing.T) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS testusers (id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY, username TEXT NOT NULL, password TEXT NOT NULL)")
	if err != nil {
		log.Fatalln("Could not create test user table", err)
	}

	password, err := bcrypt.GenerateFromPassword([]byte("mypassword1234"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln("Failed to generate password hash:", err)
	}
	userToInsert := [][]interface{}{
		{15, "testchungus", password},
	}

	numRows, err := conn.CopyFrom(context.Background(), pgx.Identifier{"testusers"}, []string{"id", "username", "password"}, pgx.CopyFromRows(userToInsert))
	if err != nil {
		log.Fatalln("Could not insert testchungus user into database:", err)
	}
	if numRows != 1 {
		log.Fatalln("Could not insert testchungus user into datatbase: numRows != 1")
	}

	var user User
	err = conn.QueryRow(context.Background(), "SELECT (id,username,password) FROM testusers WHERE username = 'testchungus'").Scan(&user)
	if err == pgx.ErrNoRows {
		log.Fatalln("Could not find the user testchungus in database testusers:", err)
	} else if err != nil && err != pgx.ErrNoRows {
		log.Fatalln("Could not scan the row result into User struct:", err)
	}
	log.Println("Found user testchungus: {", user.ID, user.Username, user.Password, "}")

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("mypassword1234")); err != nil {
		log.Fatalln("Password Check failed:", err)
	}

	ct, err := conn.Exec(context.Background(), "DELETE FROM testusers WHERE username = 'testchungus'")
	if err != nil {
		log.Fatalln("Deletion failed:", err)
	}
	if ct.RowsAffected() <= 0 {
		log.Fatalln("Deletion affected 0 rows")
	}
}
