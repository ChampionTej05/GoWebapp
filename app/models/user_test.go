package models

import (
	"testing"
)

func TestAddUser(t *testing.T) {
	// Test adding a user
	user := User{FirstName: "John", LastName: "Doe"}
	addedUser, err := AddUser(user)
	if err != nil {
		t.Errorf("Error adding user: %v", err)
	}
	if addedUser.Id != 1 {
		t.Errorf("Expected ID to be 1, got %d", addedUser.Id)
	}

	// Test adding a user with an ID
	userWithID := User{Id: 2, FirstName: "Jane", LastName: "Doe"}
	_, err = AddUser(userWithID)
	if err == nil {
		t.Errorf("Expected error when adding user with ID, got none")
	}

	// remove all users from the database
	RemoveAllUsers()
}

func TestGetUserByID(t *testing.T) {
	// Test getting an existing user
	user := User{FirstName: "John", LastName: "Doe"}
	addedUser, _ := AddUser(user)
	foundUser, err := GetUserByID(addedUser.Id)
	if err != nil {
		t.Errorf("Error getting user: %v", err)
	}
	if foundUser.FirstName != user.FirstName || foundUser.LastName != user.LastName {
		t.Errorf("Expected user %v, got %v", user, foundUser)
	}

	// Test getting a non-existent user
	_, err = GetUserByID(999)
	if err == nil {
		t.Errorf("Expected error when getting non-existent user, got none")
	}
	// remove all users from the database
	RemoveAllUsers()
}

func TestUpdateUserByID(t *testing.T) {
	// Test updating an existing user
	user := User{FirstName: "John", LastName: "Doe"}
	addedUser, _ := AddUser(user)
	updatedUser := User{Id: addedUser.Id, FirstName: "Jane", LastName: "Doe"}
	updated, err := UpdateUserByID(updatedUser)
	if err != nil {
		t.Errorf("Error updating user: %v", err)
	}
	if updated.FirstName != updatedUser.FirstName || updated.LastName != updatedUser.LastName {
		t.Errorf("Expected updated user %v, got %v", updatedUser, updated)
	}

	// Test updating a non-existent user
	userWithoutID := User{FirstName: "John", LastName: "Doe"}
	_, err = UpdateUserByID(userWithoutID)
	if err == nil {
		t.Errorf("Expected error when updating non-existent user, got none")
	}
	// remove all users from the database
	RemoveAllUsers()
}

func TestRemoveUserByID(t *testing.T) {
	// Test removing an existing user
	user := User{FirstName: "John", LastName: "Doe"}
	addedUser, _ := AddUser(user)
	err := RemoveUserByID(addedUser.Id)
	if err != nil {
		t.Errorf("Error removing user: %v", err)
	}
	users := GetUsers()
	if len(users) != 0 {
		t.Logf("Users: %v", users)
		t.Errorf("Expected users to be empty, got %v", users)
	}

	// Test removing a non-existent user
	err = RemoveUserByID(999)
	if err == nil {
		t.Errorf("Expected error when removing non-existent user, got none")
	}
	// remove all users from the database
	RemoveAllUsers()
}
