package handler

import (
	"net/http"
	"strconv"

	"github.com/HOangAG2207/GoBe-K03/internal/service"
	"github.com/labstack/echo/v4"
)

// defaultPasswordLength defines the fallback password length
// used when the client does not provide a custom length
const defaultPasswordLength = 12

// GenPassword defines the contract for password HTTP handler
type GenPassword interface {
	// GeneratePassword handles HTTP request to generate password
	// Input:
	//   - ectx: echo context containing request/response data
	// Output:
	//   - error: returned if request handling fails
	GeneratePassword(ectx echo.Context) error
}

// genPasswordHandler is the concrete implementation of GenPassword
type genPasswordHandler struct {
	// genPasswordService is the dependency from service layer
	// responsible for generating passwords
	genPasswordService service.GenPassword
}

// NewGenPassword creates a new handler instance
// Input:
//   - svc: password generation service
//
// Output:
//   - GenPassword: handler interface
func NewGenPassword(svc service.GenPassword) GenPassword {
	return &genPasswordHandler{
		genPasswordService: svc,
	}
}

// ===== Response DTO =====

// GenPasswordResponse represents success response
// Returned when password is generated successfully
type GenPasswordResponse struct {
	Password string `json:"password"`
}

// ErrorResponse represents error response
// Used for returning error messages to client
type ErrorResponse struct {
	Error string `json:"error"`
}

// GeneratePassword generates a password.
// @Summary Generate random password
// @Description Generate a random password with optional custom length. If length is not provided, default is 12.
// @Tags password generation
// @Accept json
// @Produce json
// @Param length query int false "Password length (must be > 0, default = 12)"
// @Success 200 {object} GenPasswordResponse "Password generated successfully"
// @Failure 400 {object} ErrorResponse "Invalid length parameter"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/gen-password [get]
func (h *genPasswordHandler) GeneratePassword(c echo.Context) error {

	// ===== 1. Determine password length =====

	// default length is used if no query param is provided
	length := defaultPasswordLength

	// check if client provides "length" query parameter
	if l := c.QueryParam("length"); l != "" {

		// convert string to integer
		parsed, err := strconv.Atoi(l)

		// validate:
		// - must be a number
		// - must be greater than 0
		if err != nil || parsed <= 0 {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Error: "invalid length",
			})
		}

		// override default length
		length = parsed
	}

	// ===== 2. Call service layer =====

	// generate password using service
	pass, err := h.genPasswordService.GeneratePassword(length)
	if err != nil {

		// return generic error to avoid leaking internal details
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Internal Server Error",
		})
	}

	// ===== 3. Return success response =====

	return c.JSON(http.StatusOK, GenPasswordResponse{
		Password: pass,
	})
}
