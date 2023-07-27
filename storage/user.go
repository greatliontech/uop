package storage

import (
	"crypto/rsa"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type User struct {
	ID                string `yaml:"id"`
	Username          string `yaml:"username"`
	Password          string `yaml:"password"`
	FirstName         string `yaml:"firstName"`
	LastName          string `yaml:"lastName"`
	Email             string `yaml:"email"`
	EmailVerified     bool   `yaml:"emailVerified"`
	Phone             string `yaml:"phone"`
	PhoneVerified     bool   `yaml:"phoneVerified"`
	PreferredLanguage string `yaml:"preferredLanguage"`
	IsAdmin           bool   `yaml:"isAdmin"`
}

type Service struct {
	keys map[string]*rsa.PublicKey
}

type UserStore interface {
	GetUserByID(string) *User
	GetUserByUsername(string) *User
	ExampleClientID() string
}

type userStore struct {
	users map[string]*User
}

func NewUserStore(file string) (UserStore, error) {

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(data))

	users := make(map[string]*User)

	if err := yaml.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}

	return &userStore{
		users: users,
	}, nil
}

// ExampleClientID is only used in the example server
func (u userStore) ExampleClientID() string {
	return "service"
}

func (u userStore) GetUserByID(id string) *User {
	return u.users[id]
}

func (u userStore) GetUserByUsername(username string) *User {
	for _, user := range u.users {
		if user.Username == username {
			return user
		}
	}
	return nil
}
