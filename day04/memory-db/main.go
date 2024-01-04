// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

// User represents an entity with ID and Name fields.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// DataManager encapsulates CRUD operations on the in-memory storage.
type DataManager struct {
	Storage map[int]User
}

// NewDataManager creates a new DataManager instance.
func NewDataManager() *DataManager {
	return &DataManager{
		Storage: make(map[int]User),
	}
}

// LoadDataFromFile loads initial data from a JSON file.
func (dm *DataManager) LoadDataFromFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		return err
	}

	for _, user := range users {
		dm.Storage[user.ID] = user
	}

	return nil
}

// SaveDataToFile saves the current data to a JSON file.
func (dm *DataManager) SaveDataToFile(filename string) error {
	panic("TODO #1 implement")
	// Mashal
}

// Create adds a new User to the in-memory storage.
func (dm *DataManager) Create(name string) error {
	panic("TODO #2 implement: generate IDs in an incremental manner")
}

// Read retrieves and prints the details of a User based on its ID.
func (dm *DataManager) Read(id int) (User, error) {
	panic("TODO #3 implement")
}

// Update updates the details of an existing User based on its ID.
func (dm *DataManager) Update(user User) error {
	panic("TODO #4 implement")
}

// Delete removes a User from the in-memory storage based on its ID.
func (dm *DataManager) Delete(id int) error {
	panic("TODO #5 implement: delete and save it to data_deleted.json")
}

/*
	Requirement:
		* Implement TODO

		*  Ensure safe concurrent access to the in-memory data
			+ https://gobyexample.com/mutexes

		*  Handle the case when a user does not exist or duplicate

		*  Handle panic: https://gobyexample.com/recover

*/

func main() {
	// Example usage:
	dataManager := NewDataManager()

	// Load initial data from a file
	if err := dataManager.LoadDataFromFile("data.json"); err != nil {
		fmt.Println("Error loading data:", err)
		return
	}

	defer func() {
		if err := dataManager.SaveDataToFile("data.json"); err != nil {
			panic("dataManager.SaveDataToFile: " + err.Error())
		}
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		wg.Done()
		user, err := dataManager.Read(1)
		if err != nil {
			panic("dataManager.Read: " + err.Error())
		}
		fmt.Println(user)
	}()

	wg.Add(1)
	go func() {
		wg.Done()
		user := User{ID: 100, Name: "This is new name"}
		if err := dataManager.Update(user); err != nil {
			panic("dataManager.Update: " + err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		wg.Done()
		if err := dataManager.Delete(2); err != nil {
			panic("dataManager.Delete: " + err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		wg.Done()
		if err := dataManager.Create("Hello 3"); err != nil {
			panic("dataManager.Create" + err.Error())
		}
	}()

	wg.Wait()
	fmt.Println("------- DONE ------")
}
