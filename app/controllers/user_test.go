package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ChampionTej05/GoWebapp/models"
)

// func Test_userController_ServeHTTP(t *testing.T) {
// 	type fields struct {
// 		userIDPattern *regexp.Regexp
// 	}
// 	type args struct {
// 		w http.ResponseWriter
// 		r *http.Request
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			uc := userController{
// 				userIDPattern: tt.fields.userIDPattern,
// 			}
// 			uc.ServeHTTP(tt.args.w, tt.args.r)
// 		})
// 	}
// }

func TestGetAllUsers(t *testing.T) {
	models.RemoveAllUsers()
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Logf("Error in creating request %v", err)
		t.Fatal(err)
	}
	handler := http.HandlerFunc(newUserController().ServeHTTP)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var users []models.User
	json.Unmarshal(rr.Body.Bytes(), &users)
	if len(users) != 0 {
		t.Errorf("Expected an empty array but got %v", users)
	}
}

func TestGetUserByID(t *testing.T) {
	models.RemoveAllUsers()
	user, err := models.AddUser(models.User{FirstName: "John", LastName: "Doe"})
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler := http.HandlerFunc(newUserController().ServeHTTP)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var responseUser models.User
	json.Unmarshal(rr.Body.Bytes(), &responseUser)
	if responseUser.Id != user.Id {
		t.Errorf("Expected user ID %v but got %v", user.Id, responseUser.Id)
	}
}

func TestAddUser(t *testing.T) {
	models.RemoveAllUsers()
	user := models.User{
		FirstName: "John",
		LastName:  "Doe",
	}
	userJson, _ := json.Marshal(user)
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(userJson))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	handler := http.HandlerFunc(newUserController().ServeHTTP)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var responseUser models.User
	json.Unmarshal(rr.Body.Bytes(), &responseUser)
	if responseUser.Id == 0 {
		t.Errorf("Expected a user ID but got %v", responseUser.Id)
	}
}
