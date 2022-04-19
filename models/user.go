package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User //a slice that holds pointers to User objects
	nextID = 1     // at package level, you don't need the colon operator to get implicit specification
)

func GetUsers() []*User { // function to look at slice and return users
	return users
}

func AddUser(u User) (User, error) { // new function AddUser, with two parameters for User object, return two objects - User that was created and potential error object
	if u.ID != 0 {
		return User{}, errors.New("new user must not include ID or be zero") // Since our function defines User and error, we have to return a User, so we supply an empty
	}
	u.ID = nextID             // grab the next ID variable to have a starting point
	nextID++                  // increment it
	users = append(users, &u) // append the users to users collection, use addressOf operator to grab a reference to the user object that came in (u)
	return u, nil             // return user object I've created, and nil for error
}

func GetUserByID(id int) (User, error) {
	for _, u := range users { // we don't need index so we use read only variable, grab users one at a time using range keyword to loop over users collection
		if u.ID == id {
			return *u, nil // return derferenced user, with nil error
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users { // we want index AND values
		if candidate.ID == u.ID { // if the IDs match, we are going to use that index to replace that ID entry in the slice
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...) //we are taking a slice of users up to index i, and then append to it all the users from the index after the match onwards
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}
