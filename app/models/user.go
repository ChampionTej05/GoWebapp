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

// The users variable is a slice of pointers to User objects, which stores all the users in the system.
// The nextID variable is an integer that keeps track of the next ID to be assigned to a new user.
var (
	users  []*User
	nextID = 1
)

// The GetUsers() function returns a list of all users in the system by simply returning the users slice.
func GetUsers() []*User {
	fmt.Println("Users, ", users)
	return users
}

// The AddUser(user User) (User, error) function takes a User object as input and adds it to the users slice. It also assigns an ID to the user, and returns the added User object and an error (if any).
func AddUser(user User) (User, error) {
	if user.Id != 0 {
		// can't return NIL as value is expected and not pointer
		return User{}, errors.New("User should not have any ID, it is provided by Controller")
	}
	user.Id = nextID
	nextID++
	users = append(users, &user)
	fmt.Println("User is added successfully")
	fmt.Printf("Users after addition %+v", users)
	return user, nil
}

// The GetUserByID(id int) (User, error) function takes an integer ID as input and returns the User object with the matching ID. It searches the users slice for a user with a matching ID, and returns the user if found.
func GetUserByID(id int) (User, error) {
	for _, user := range users {
		if user.Id == id {
			// users[] stores pointer of users, so dereference it to send the value
			return *user, nil
		}
	}
	return User{}, fmt.Errorf("User with id %v is not found", id)
}

// The UpdateUserByID(newUser User) (User, error) function takes a User object as input and updates the user in the users slice with the same ID. It returns the updated User object and an error (if any).
func UpdateUserByID(newUser User) (User, error) {
	for idx, user := range users {
		if user.Id == newUser.Id {
			users[idx] = &newUser
			return newUser, nil

		}
	}
	return User{}, fmt.Errorf("User with id %v is not found", newUser.Id)
}

// The RemoveUserByID(id int) error function takes an integer ID as input and removes the user with the matching ID from the users slice. It returns an error if the user is not found.
func RemoveUserByID(id int) error {
	for idx, user := range users {
		if user.Id == id {
			users = append(users[:idx], users[idx+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with id %v is not found", id)
}

// Remove all users from the slice
func RemoveAllUsers() {
	users = nil
}
