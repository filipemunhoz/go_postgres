package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	// connect to a database
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=postgres user=postgres password=postgres sslmode=disable")

	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect: %v\n", err))
	}
	defer conn.Close()

	log.Println("Connected to database")

	// test database connection
	err = conn.Ping()
	if err != nil {
		log.Fatal("Cannot ping database.")
	}
	log.Println("Pinged the database")

	// insert a row
	query := `insert into app.users (name, email, password) values ($1, $2, $3)`
	_, err = conn.Exec(query, "Apolo", "apollo@gmail.com", "apo1345")
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot insert into table: %v\n", err))
	}
	log.Println("Inserted Row")

	// get rows from table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(fmt.Sprintf("get rows from table: %v\n", err))
	}

	// update a row
	query = `update app.users set name = $1 where name = $2`
	_, err = conn.Exec(query, "NameUpdated", "Apolo")
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot update table: %v\n", err))
	}
	log.Println("Updated Row")

	// get rows from table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(fmt.Sprintf("get rows from table: %v\n", err))
	}

	// get one row by id
	query = `select id, name, email, password from app.users where id = $1`

	var name, email, password string
	var id int

	row := conn.QueryRow(query, 1)
	err = row.Scan(&id, &name, &email, &password)
	if err != nil {
		log.Fatal(fmt.Sprintf("cannot get row from table: %v\n", err))
	}

	// delete a row
	query = `delete from app.users where id = $1`
	_, err = conn.Exec(query, 5)
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot delete row: %v\n", err))
	}
	log.Println("Deleted Row")

}

func getAllRows(conn *sql.DB) error {

	rows, err := conn.Query("select id, name, email, password from app.users")

	if err != nil {
		log.Print(fmt.Sprintf("getAllRows(): %v\n", err))
		return err
	}
	defer rows.Close()

	var id int
	var name, email, password string

	for rows.Next() {
		err := rows.Scan(&id, &name, &email, &password)
		if err != nil {
			log.Print(fmt.Sprintf("rows.Scan(): %v\n", err))
			return err
		}
		fmt.Println("Record is: ", id, name, email, password)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
	}
	fmt.Println("-------------------------------------------------------")

	return nil

}
