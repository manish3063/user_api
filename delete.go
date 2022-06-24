package main

import "fmt"

func deleteuserID(userID string) string {
	message := "user not found"
	user := User{}
	delete(Data, userID)
	if user, ok := Data[userID]; ok {

		message = "user delated success"
		fmt.Println(user)
		return message
	}
	fmt.Println(user)
	return message
}
