package shorten_url_handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HOangAG2207/GoBe-K03/internal/service/shorten_url/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type reqBody struct {
	URL   string `json:"url"`
	ExpIn int    `json:"exp"`
}

func newCtx(body []byte) (echo.Context, *httptest.ResponseRecorder, *http.Request) {
	req := httptest.NewRequest(http.MethodPost, "/short", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return c, rec, req
}

func TestHandler_ShortenUrl(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string

		input reqBody

		mock func() *mocks.Service

		expectStatus int
		check        func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "success",
			input: reqBody{
				URL:   "https://google.com",
				ExpIn: 10,
			},
			mock: func() *mocks.Service {
				m := new(mocks.Service)

				m.On("ShortenURL",
					mock.Anything,
					"https://google.com",
					10,
				).Return("abc123", nil)

				return m
			},
			expectStatus: http.StatusCreated,
			check: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var res map[string]interface{}
				err := json.Unmarshal(rec.Body.Bytes(), &res)

				assert.NoError(t, err)
				assert.Equal(t, "abc123", res["code"])
				assert.Equal(t, "Shorten URL generated successfully!", res["message"])
			},
		},

		{
			name: "invalid request",
			input: reqBody{
				URL:   "",
				ExpIn: 0,
			},
			mock: func() *mocks.Service {
				return new(mocks.Service)
			},
			expectStatus: http.StatusBadRequest,
			check: func(t *testing.T, rec *httptest.ResponseRecorder) {
				assert.Contains(t, rec.Body.String(), "invalid request payload")
			},
		},

		{
			name: "service error",
			input: reqBody{
				URL:   "https://google.com",
				ExpIn: 10,
			},
			mock: func() *mocks.Service {
				m := new(mocks.Service)

				m.On("ShortenURL",
					mock.Anything,
					"https://google.com",
					10,
				).Return("", errors.New("db error"))

				return m
			},
			expectStatus: http.StatusInternalServerError,
			check: func(t *testing.T, rec *httptest.ResponseRecorder) {
				assert.Contains(t, rec.Body.String(), "internal server error")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			body, _ := json.Marshal(tc.input)

			c, rec, req := newCtx(body)

			mockSvc := tc.mock()
			defer mockSvc.AssertExpectations(t)

			h := NewUrlHandler(mockSvc)

			c.SetRequest(req)

			err := h.ShortURL(c)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectStatus, rec.Code)

			tc.check(t, rec)
		})
	}
}
