package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone" binding:"required"`
	UserID   string `json:"user_id" binding:"required"`
	City     string `json:"city"`
	Password string `json:"password"`
}

var (
	Data map[string]User
	DB   *sql.DB
)

func main() {
	Data = make(map[string]User)
	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupRoutes(r *gin.Engine) {
	r.GET("/user/:user_id", GetUserByUserID)
	r.GET("/user", GetAllUser)
	r.POST("/user", CreateUser)
	r.PUT("/user/:user_id", UpdateUser)
	r.DELETE("/user/:user_id", DeleteUser)

}

//GetUserByUserID function
func GetUserByUserID(c *gin.Context) {
	//records := readCsvFile("./movies.csv")
	userID, ok := c.Params.Get("user_id")
	if ok == false {
		res := gin.H{
			"error": "user_id is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}
	user, _ := getUserByIDFromDB(userID)
	res := gin.H{
		"user": user,
	}
	c.JSON(http.StatusOK, res)
}

//GetAllUser function
func GetAllUser(c *gin.Context) {
	res := gin.H{
		"user": Data,
	}
	c.JSON(http.StatusOK, res)
}

//CreateUser POST
func CreateUser(c *gin.Context) {
	reqBody := User{}
	err := c.Bind(&reqBody)
	if err != nil {
		res := gin.H{
			"error": err,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if reqBody.UserID == "" {
		res := gin.H{
			"error": "UserId must not be empty",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	insert := insertUserinDB(reqBody)
	//Data[reqBody.UserID] = reqBody
	res := gin.H{
		"result": insert,
		"user":   reqBody,
	}
	c.JSON(http.StatusOK, res)
	return

}

//Update User PUT
func UpdateUser(c *gin.Context) {
	userID, ok := c.Params.Get("user_id")
	if ok == false {
		res := gin.H{
			"error": "user_id is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}

	reqBody := User{}

	err := c.Bind(&reqBody)
	if err != nil {
		res := gin.H{
			"error": err,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if reqBody.UserID == "" {
		res := gin.H{
			"error": "UserId must not be empty",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if reqBody.UserID != userID {
		res := gin.H{
			"error": "UserId cannot be updated",
		}
		c.JSON(http.StatusBadRequest, res)
		return

	}
	//password function call
	password := c.GetHeader("password")

	if !checkpassword(userID, password) {
		res := gin.H{
			"error": "invalid password",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	//
	update := updateUserinDB(reqBody)
	//Data[reqBody.UserID] = reqBody
	res := gin.H{
		"result": update,
		"user":   reqBody,
	}
	c.JSON(http.StatusOK, res)
	return

}

//Delete user by id
func DeleteUser(c *gin.Context) {
	userID, ok := c.Params.Get("user_id")
	if ok == false {
		res := gin.H{
			"error": "user_id is missing",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	reqBody := User{}

	err := c.Bind(&reqBody)
	if err != nil {
		res := gin.H{
			"error": err,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if reqBody.UserID == "" {
		res := gin.H{
			"error": "UserId must not be empty",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	//Data[userID] = reqBody
	// result := deleteuserID(userID)
	// res := gin.H{
	// 	"success": true,
	// 	"user":    reqBody,
	// 	"message": result,
	// }
	// c.JSON(http.StatusOK, res)
	// return

	delete := deleteUserinDB(userID)
	//Data[reqBody.UserID] = reqBody
	res := gin.H{
		"result": delete,
		"user":   reqBody,
	}
	c.JSON(http.StatusOK, res)
	return

}

// to check the password....
func checkpassword(username, password string) bool {
	user, err := getUserByIDFromDB(username)
	fmt.Println(user, err, username, password)
	if user.Password == password {
		return true
	}
	return false
}
