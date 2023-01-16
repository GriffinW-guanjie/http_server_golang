package database

import (
	"encoding/json"
	"os"
	"time"
)

const defaultPath = "./db.json"
type Client struct {
	location string
}
type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}

type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}


type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}

func NewClient(path string) Client {
	return Client{location: path}
}

func (c Client) createDB() error {
	result,_ := json.Marshal(`databaseSchema{ 
	}`)
	err := os.WriteFile(c.location, result, 0666)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) EnsureDB() error {
	_,err := os.ReadFile(c.location)
	if err != nil {
		c.createDB()
		return nil
	} else {
		return nil
	}
}