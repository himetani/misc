package main

import (
	"math/rand"
	"testing"
)

func Test_UserListProxy(t *testing.T) {
	someDatabase := UserList{}

	rand.Seed(2342342)
	for i :=0; i< 1000000; i++ {
		n:=rand.Int31
		someDatabase = append(someDatabase, User{ID: n})
	}

	proxy := UserListProxy{
		SomeDatabase: &SomeDatabase: &someDatabase,
		StackCapacity: 2,
		StackCache: UserList{}
	}

	knownIDs := [3]int32 {someDatabase[3].ID, someDatabase[4].ID,someDatabase[5].ID}

	t.Run("FindUser - Empty cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}
	}
}
