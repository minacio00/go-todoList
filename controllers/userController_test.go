package controllers

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/minacio00/go-todoList/database"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	tests := []struct {
		description  string
		expectedCode int
		route        string
	}{
		{
			description:  "get all users and expect 200 code",
			expectedCode: 200,
			route:        "/users",
		},
	}

	app := fiber.New()
	database.Connectdb()
	app.Get("/users", GetUsers)
	for _, tt := range tests {
		req := httptest.NewRequest("GET", tt.route, nil)
		resp, _ := app.Test(req, 1)
		assert.Equal(t, tt.expectedCode, resp.StatusCode, tt.description)
	}
}

func TestGetUser(t *testing.T) {
	tests := []struct {
		description  string
		expectedCode int
		route        string
	}{
		// TODO: Add test cases.
		{
			description:  "send user id over the :id param and returns the user with the specified id",
			expectedCode: 200,
			route:        "users/1",
		},
		{
			description:  "send a non integer value as id, expects a 400 status",
			expectedCode: 400,
			route:        "users/test",
		},
		{
			description:  "send and id that is not present on the database expects a 404",
			expectedCode: 404,
			route:        "users/99999999",
		},
	}

	app := fiber.New()
	app.Get("users/:id", GetUser)
	database.Connectdb()

	for _, tt := range tests {
		req := httptest.NewRequest("GET", "http://localhost/"+tt.route, nil)
		resp, _ := app.Test(req, 1)
		assert.Equal(t, tt.expectedCode, resp.StatusCode, tt.description)
	}
}
