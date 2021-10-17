package models

import (
	"errors"
	"fmt"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	fmt.Println("Users, ", users)
	return users
}

func AddUser(user User) (User, error) {
	if user.Id != 0 {
		// can't return NIL as value is expected and not pointer
		return User{}, errors.New("User should not have any ID, it is provided by Controller")
	}
	user.Id = nextID
	nextID++
	users = append(users, &user)
	fmt.Println("User is added successfully")
	fmt.Println("Users after addition ", users)
	return user, nil
}

func GetUserByID(id int) (User, error) {
	for _, user := range users {
		if user.Id == id {
			// users[] stores pointer of users, so dereference it to send the value
			return *user, nil
		}
	}
	return User{}, fmt.Errorf("User with id %v is not found", id)
}

func UpdateUserByID(newUser User) (User, error) {
	for idx, user := range users {
		if user.Id == newUser.Id {
			users[idx] = &newUser
			return newUser, nil

		}
	}
	return User{}, fmt.Errorf("User with id %v is not found", newUser.Id)
}

func RemoveUserByID(id int) error {
	for idx, user := range users {
		if user.Id == id {
			users = append(users[:idx], users[idx+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with id %v is not found", id)
}
