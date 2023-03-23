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
