package main

import "fmt"

func deleteuserID(userID string) string {
	message := "user not found"
	user := User{}
	if user, ok := Data[userID]; ok {
		delete(Data, userID)
		message = "user delated success"
		fmt.Println(user)
		return message
	}
	fmt.Println(user)
	return message
}
