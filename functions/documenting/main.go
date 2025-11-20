package main

// User represents a user with an id and first name.
type User struct {
	ID    int
	First string
}

// DataStore defines an interface fo storing retrievable data.
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

func main() {

}
