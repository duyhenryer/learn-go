package server

import (
	"encoding/json"
	"net/http"
	"time"
)

func StartServer() {
	http.HandleFunc("/api/public/register", register)
	http.HandleFunc("/api/public/login", login)
	http.HandleFunc("/api/private/self", self)

	http.HandleFunc("/api/public/log/register", LogWrapper(register))
	http.HandleFunc("/api/public/log/login", LogWrapper(login))
	http.HandleFunc("/api/private/log/self", LogWrapper(self))

	http.ListenAndServe(":8090", nil)
}

/*
		TODO #2:
		- implement the logic to register a new user (Username, Password, full_name, address)
	  	- Validate Username (not empty and unique)
	  	- Validate Password (length should at least 8)
*/
func register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} // Cái nào != nil thì mới call được err.Error()
	if req.Username == "" {
		http.Error(w, "Username is empty", http.StatusBadRequest)
		return
	}
	if userStore.Exits(req.Username) {
		http.Error(w, "Username is registered", http.StatusBadRequest)
		return
	}
	if len(req.Password) < 8 {
		http.Error(w, "Passwd too short. At least 8", http.StatusBadRequest)
		return
	}
	if err := userStore.Save(UserInfo{
		username: req.Username,
		info:     req.Address,
		password: req.Password,
		address:  req.Address,
	}); err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

type RegisterRequest struct {
	Username string
	Password string
	FullName string
	Address  string
}

/*
		TODO #3:
		- implement the logic to login
		- validate the user's credentials (Username, Password)
	  	- Return JWT token to client
*/
func login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := userStore.Get(req.Username)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	if req.Password != user.password {
		http.Error(w, "No correct passwd", http.StatusBadRequest)
		return
	}

	token, err := GenerateToken("", 24*time.Second)
	if err != nil {
		return
	}

	resp := LoginResponse{Token: token}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(resp.Token))
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	return
}

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	Token string
}

/*
TODO #4:
- implement the logic to get user info
- Extract the JWT token from the header
- Validate Token
- Return user info`
*/
func self(w http.ResponseWriter, r *http.Request) {
	var req GetUserInfoRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = VerifyToken(req.Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//authHeader := r.Header.Get("Authorization")
	//
	//extractUserNameFn := func(authenticationHeader string) (string, error) {
	//	panic("implement me")
	//}
	//
	//Username, err := extractUserNameFn(authHeader)
	//if err != nil {
	//	return
	//}
	//
	//user, _ := userStore.Get(Username)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}

type GetUserInfoRequest struct {
	Token    string
	Username string
}

type GetUserInfoResponse struct {
	Username string
	Info     string
}

/*
TODO: extra wrapper
Print some logs to console
  - Path
  - Http Status code
  - Time start, Duration
*/
func LogWrapper(handler http.HandlerFunc) http.HandlerFunc {
	//panic("TODO implement me")
	return handler
}

/*
	TODO #1: implement in-memory user store
	TODO #2: implement register handler
	TODO #3: implement login handler
	TODO #4: implement self handler

	Extra: implement log handler
*/
