package shorten_url_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type urlRequest struct {
	URL   string `json:"url" binding:"required,url"`
	ExpIn int    `json:"exp" binding:"required,min=1"`
}
type urlResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ShortURL handles HTTP POST
//
// @Summary      Shorten a URL
// @Description  Generates a short code for URL and stores it in Redis
// @Tags         links
// @Accept       json
// @Produce      json
// @Param        request  body      urlRequest          true  "Shorten URL request"
// @Success      201      {object}  urlResponse
// @Failure      400      {object}  urlResponse
// @Failure      500      {object}  urlResponse
// @Router       /api/shorten-url [post]
func (h *urlHandler) ShortURL(c echo.Context) error {
	input := new(urlRequest)

	// ❌ FIX: chỉ bind 1 lần
	if err := c.Bind(input); err != nil || input.URL == "" || input.ExpIn <= 0 {
		return c.JSON(http.StatusBadRequest, urlResponse{
			Message: InValidRequestPayload.Error(),
		})
	}

	code, err := h.urlService.ShortenURL(
		c.Request().Context(),
		input.URL,
		input.ExpIn,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, urlResponse{
			Message: InternalServerError.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    code,
		"message": "Shorten URL generated successfully!",
	})
}
