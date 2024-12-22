package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-rest-api/types"
	"github.com/gorilla/mux"
)

func TestUseServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandles(userStore)

	t.Run("Should fail if the payload is invalid!", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "Subomi",
			LastName:  "Subby",
			Email:     "subby123@mail.com",
			Password:  "12345678",
		}
		marsh, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marsh))
		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(res, req)

		// if the request doesn't give a bad code
		if res.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Must run correctly to register user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "vv",
			LastName:  "vvv",
			Email:     "validate@mail.com",
			Password:  "123456789",
		}
		marsh, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marsh))
		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(res, req)

		// if the request doesn't give a bad code
		if res.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, res.Code)
		}
	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.Users, error) {
	return nil, fmt.Errorf("User not found!")
}

func (m *mockUserStore) GetUserByID(id int) (*types.Users, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(types.Users) error {
	return nil
}
