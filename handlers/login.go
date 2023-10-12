package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/AkifhanIlgaz/jwt-auth/mock"
	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserClaims struct {
	Uid  string `json:"uid"`
	Role string `json:"admin"`
	jwt.RegisteredClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	// TODO: Read email and password from request
	loginRequest, err := readFromForm(r)
	if err != nil {
		fmt.Println("login: %w", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// TODO: Hash the password with bcrypt and check if the email and password is correct
	user, err := mock.GetUser(loginRequest.Email, loginRequest.Password)
	if err != nil {
		fmt.Println("login: %w", err)
		http.NotFound(w, r)
		return
	}

	fmt.Println(user)
	// ! If the credentials are correct

	// TODO: Then, create Access Token with uid, role, exp claims
	token, err := createAccessToken(UserClaims{
		Uid:  user.Uid,
		Role: "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	})

	fmt.Println(token)

	// TODO: Create refresh token
}

/*
let formData = new FormData();
formData.append('name', 'John');
formData.append('password', 'John123');

fetch("api/SampleData",

	{
	    body: formData,
	    method: "post"
	});
*/
func readFromForm(r *http.Request) (*LoginRequest, error) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		return nil, fmt.Errorf("read from form: form values cannot be blank")
	}

	return &LoginRequest{
		Email:    email,
		Password: password,
	}, nil
}

func readFromBody(r *http.Request) (*LoginRequest, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("read from body: %w", err)
	}

	var loginRequest LoginRequest

	err = json.Unmarshal(body, &loginRequest)
	if err != nil {
		return nil, fmt.Errorf("read from body | unmarshal: %w", err)
	}

	return &loginRequest, nil
}

func createAccessToken(claims UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("secret"))
}

func parseAccessToken(token string) *UserClaims {
	parsed, _ := jwt.ParseWithClaims(token, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	return parsed.Claims.(*UserClaims)
}
