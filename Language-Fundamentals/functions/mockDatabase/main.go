package main

import (
	"fmt"
	"log"
)

// User represents a user with an id and first name.
type User struct {
	ID    int
	First string
}

// `DataStore` defines an interface fo storing retrievable data.
// Any Type that implements this interfcase (has these two methods) is also of Type `DataStore`.
// This means any value of Type `MockDatastore` is also type `DataStore`.
// This means we could have a value of Type `*sql.Db` and it can also be of type `Datastore`.
// Tis means we can write functions to take Type `Datastore` and use either of these:
// -- `MockDatastore` --
// -- `*sql.DB` --
type DataStore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

// Servicerepresents a service that stores retrievable data.
type Service struct {
	ds DataStore
}

func (s Service) GetUser(id int) (User, error) {
	return s.ds.GetUser(id)
}

func (s Service) SaveUser(u User) error {
	return s.ds.SaveUser(u)
}

// MockDataStore is temporaty service that stores retrievable data.
type MockDataStore struct {
	Users map[int]User
}

func (md MockDataStore) GetUser(id int) (User, error) {
	user, ok := md.Users[id]
	if !ok {
		return User{}, fmt.Errorf("User %v not found", id)
	}
	return user, nil
}
func (md MockDataStore) SaveUser(u User) error {
	_, ok := md.Users[u.ID]
	if ok {
		return fmt.Errorf("User %v already exists", u.ID)
	}
	md.Users[u.ID] = u
	return nil
}

func main() {
	db := MockDataStore{
		Users: make(map[int]User),
	}

	srvc := Service{
		ds: db,
	}

	u1 := User{
		ID:    1,
		First: "James",
	}

	err := srvc.SaveUser(u1)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	returnedUser, err := srvc.GetUser(u1.ID)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	fmt.Println(u1)
	fmt.Println(returnedUser)
}
