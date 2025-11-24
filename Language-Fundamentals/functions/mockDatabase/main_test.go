package main

import "testing"

func TestGetUser(t *testing.T) {
	md := &MockDataStore{
		Users: map[int]User{
			2: {ID: 2, First: "Jimmy"},
		},
	}

	s := &Service{
		ds: md,
	}
	u, err := s.GetUser(2)

	if err != nil {
		t.Errorf("Got an error: %v", err)
	}

	if u.First != "Jimmy" {
		t.Errorf("Got: %v, want: %s", u.First, "Jimmy")
	}
}
