// +build ignore

package main

import (
	"fmt"
	"sync"
)

var (
	names = []string{
		"Pepe",
		"Gozalo",
		"Juan",
		"Carolina",
	}
	lastNames = []string{
		"Escobar",
		"Sierra",
		"Velez",
		"Mejia",
	}
	usernames = []string{
		"pep66",
		"jsi3rra",
		"jvlez8",
		"caro27",
	}
	emails = []string{
		"pepe27@gmail.com",
		"gozalosierra@gmail.com",
		"juanv@gmail.com",
		"carolina@gmail.com",
	}
	passwords = []string{
		"qwerty",
		"123456",
		"AeIoU!@",
		"S3CUR3P455W0RD!\"#$%&/()=",
	}
)

// User model
type User struct {
	Username     string
	Email        string
	LastName     string
	Name         string
	PasswordHash string
}

func makeUsers() []User {
	users := []User{}
	for i := 0; i < 10; i++ {
		u := User{
			Username:     usernames[i%4],
			Email:        emails[i%4],
			LastName:     lastNames[i%4],
			Name:         names[i%4],
			PasswordHash: passwords[i%4],
		}
		users = append(users, u)
	}
	return users
}

// START PARALLELFORLOOP OMIT
func parallelLoop() {
	users := makeUsers()
	var wg sync.WaitGroup
	for _, u := range users {
		wg.Add(1)
		go func(u User) {
			fmt.Printf("%s: (%s)\n", u.Email, u.Username)
			wg.Done()
		}(u)
	}
	wg.Wait()
}

// END PARALLELFORLOOP OMIT
func main() {
	parallelLoop()
}
