package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{
		{Name: "John", Age: 20},
		{Name: "Mary", Age: 17},
		{Name: "Peter", Age: 18},
	}

	var wg sync.WaitGroup
	wg.Add(len(users))

	for i, user := range users {
		go func(i int, user User) {
			addUser(i, user)
			wg.Done()
		}(i, user)
	}
	wg.Wait()
}

func addUser(i int, user User) {
	if !validateAge(user) {
		fmt.Println(i, "User not added")
		return
	}
	fmt.Println(i, "User added")
}

func validateAge(user User) bool {
	return user.Age > 18
}
