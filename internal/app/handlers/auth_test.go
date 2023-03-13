package handlers

import (
	"bytes"
	"errors"
	"github.com/HeadGardener/game-tracker/internal/app/models"
	"github.com/HeadGardener/game-tracker/internal/app/services"
	mock_services "github.com/HeadGardener/game-tracker/internal/app/services/mocks"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_services.MockAuthorization, user models.RegUserInput)

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            models.RegUserInput
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			inputBody: `{"name":"ok","username":"tester","password":"12345"}`,
			inputUser: models.RegUserInput{
				Name:     "ok",
				Username: "tester",
				Password: "12345",
			},
			mockBehavior: func(s *mock_services.MockAuthorization, user models.RegUserInput) {
				s.EXPECT().Create(user).Return(1, nil)
			},
			expectedStatusCode:   201,
			expectedResponseBody: "{\"id\":1}\n",
		},
		{
			name:                 "empty name field",
			inputBody:            `{"username":"tester","password":"12345"}`,
			mockBehavior:         func(s *mock_services.MockAuthorization, user models.RegUserInput) {},
			expectedStatusCode:   400,
			expectedResponseBody: "{\"message\":\"there can't be empty fields in user struct\"}\n",
		},
		{
			name:      "service failure",
			inputBody: `{"name":"test","username":"tester","password":"12345"}`,
			inputUser: models.RegUserInput{
				Name:     "test",
				Username: "tester",
				Password: "12345",
			},
			mockBehavior: func(s *mock_services.MockAuthorization, user models.RegUserInput) {
				s.EXPECT().Create(user).Return(0, errors.New("service failure"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: "{\"message\":\"service failure\"}\n",
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_services.NewMockAuthorization(c)
			tc.mockBehavior(auth, tc.inputUser)

			service := &services.Service{Authorization: auth}
			handler := NewHandler(service)

			router := chi.NewRouter()
			router.Post("/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(tc.inputBody))

			router.ServeHTTP(w, r)

			assert.Equal(t, tc.expectedStatusCode, w.Code)
			assert.Equal(t, tc.expectedResponseBody, w.Body.String())
		})
	}
}
