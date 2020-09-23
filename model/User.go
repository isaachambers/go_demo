package model

import (
	"errors"
	"fmt"
)

/*
In Go, an identifier that starts with a capital letter is exported from the package,
and can be accessed by anyone outside the package that declares it it.
If an identifier starts with a lower case letter, it can only be accessed from within the package
*/
type User struct {
	Id         int
	FirstName  string
	SecondName string
	Location   string
}

var (
	//users keeps a list of pointers to users
	users  []*User
	nextId = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(user User) (User, error) {
	if user.Id != 0 {
		return User{}, errors.New("new user must have an empty Id")
	}
	user.Id = nextId
	nextId++
	//add the pointer to users.
	users = append(users, &user)
	return user, nil
}

func GetUserById(id int) (User, error) {
	for _, u := range users {
		if u.Id == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("user with id %v not found", id)
}

func UpdateUser(user User) (User, error) {
	for i, u := range users {
		if u.Id == user.Id {
			users[i] = &user
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user with ID %v not found", user.Id)
}

func DeleteUserById(id int) error {
	user, err := GetUserById(id)
	if err != nil {
		return err
	}
	for i, u := range users {
		if u.Id == user.Id {
			//Get all users in slice upto the current one but not including it
			slice1 := users[:i]
			//Get all users after the current object until the end
			slice2 := users[i+1:]
			users = append(slice1, slice2...)
			return nil
		}
	}
	return fmt.Errorf("user with id %v could not be found", user.Id)
}
