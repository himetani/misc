package main

import "fmt"

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type User struct {
	ID int32
}

type UserList []User

type UserListProxy struct {
	SomeDatavase UserList
	StackCache UserList
	StackCapacity int
	DidDidLastSearchUsedCache bool
}

type 
func main() {
	fmt.Println("vim-go")
}
