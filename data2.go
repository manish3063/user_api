package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func getUserBybID(userID string) User {
	user := User{}
	if user, ok := Data[userID]; ok {
		return user
	}
	return user
}

func getUserByIDFromDB(userID string) (User, error) {
	var err error
	DB, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer DB.Close()
	user := User{}
	userSQL := "SELECT id, name, email, phone,userid, city, password FROM users WHERE userid = $1"

	err = DB.QueryRow(userSQL, userID).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.UserID, &user.City, &user.Password)
	if err != nil {
		log.Println("Failed to execute query: ", err)
	}
	return user, err
}

//insert query...
func insertUserinDB(reqBody User) bool {
	var result = true
	var err error
	DB, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer DB.Close()
	//user := User{}

	userSQL := `INSERT INTO users (id,name, email, phone, userid, city, password) VALUES ($1, $2, $3, $4, $5, $6,$7)`

	_, err = DB.Exec(userSQL, reqBody.ID, reqBody.Name, reqBody.Email, reqBody.Phone, reqBody.UserID, reqBody.City, reqBody.Password)

	if err != nil {
		result = false
		log.Println("Failed to execute query: ", err)
	}
	return result
}

//update useer query..

func updateUserinDB(reqBody User) bool {
	var result = true
	var err error
	DB, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer DB.Close()
	//user := User{}

	userSQL := `UPDATE users SET name = $2,email=$3,phone=$4,userid=$5,city=$6,password=$7 WHERE id = $1`

	_, err = DB.Exec(userSQL, reqBody.ID, reqBody.Name, reqBody.Email, reqBody.Phone, reqBody.UserID, reqBody.City, reqBody.Password)

	if err != nil {
		result = false
		log.Println("Failed to execute query: ", err)
	}
	return result
}

//delete user by userid
func deleteUserinDB(userid string) bool {
	var result = true
	var err error
	DB, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer DB.Close()
	//user := User{}

	userSQL := `DELETE FROM users WHERE id = $1`

	_, err = DB.Exec(userSQL, userid)

	if err != nil {
		result = false
		log.Println("Failed to execute query: ", err)
	}
	return result
}
