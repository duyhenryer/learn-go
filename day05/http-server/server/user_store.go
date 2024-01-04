package server

import "sync"

// TODO #1: implement in-memory user store

var userStore = NewUserStore()

func NewUserStore() *UserStore {
	return &UserStore{data: make(map[string]UserInfo)}
}

type UserStore struct {
	mu   sync.Mutex
	data map[string]UserInfo
}

func (u *UserStore) Save(info UserInfo) error {
	//panic("TODO: implement me")
	u.mu.Lock() // Lock: Anyone to access to the data map
	defer u.mu.Unlock()

	u.data[info.username] = info // Save the user information in the data map
	return nil                   // Return nil to indicate success (no error)
}

func (u *UserStore) Exits(username string) bool {
	u.mu.Lock()
	defer u.mu.Unlock()

	if _, found := u.data[username]; found {
		return true
	}
	return false
}

func (u *UserStore) Get(username string) (UserInfo, error) {
	//panic("TODO: implement me")
	u.mu.Lock()
	defer u.mu.Unlock()

	info, ok := u.data[username] // Retrieve user information from the data map
	if !ok {
		return UserInfo{}, nil
	}
	return info, nil
}

type UserInfo struct {
	// TODO: implement me
	username string
	info     string
	password string
	address  string
}
