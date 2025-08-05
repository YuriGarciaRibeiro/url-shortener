package handlers

import (
	"net/http"

	"github.com/YuriGarciaRibeiro/url-shortener/internal/app/service"
	"github.com/gin-gonic/gin"
)

// ShortenURLRequest representa o corpo da requisição para encurtar a URL
type ShortenURLRequest struct {
	URL string `json:"url" binding:"required,url"`
}

// ShortenURLResponse representa a resposta de uma URL encurtada
type ShortenURLResponse struct {
	Hash     string `json:"hash"`
	ShortURL string `json:"short_url"`
}

// ErrorResponse representa uma resposta de erro genérica
type ErrorResponse struct {
	Error string `json:"error"`
}

// URLHandler representa o handler das rotas de URL
type URLHandler struct {
	service *service.ShortenerService
}

// NewURLHandler cria um novo URLHandler
func NewURLHandler(s *service.ShortenerService) *URLHandler {
	return &URLHandler{service: s}
}

// ShortenURL godoc
// @Summary Shorten a URL
// @Description Shortens a given URL and returns the shortened version
// @Tags URLs
// @Accept json
// @Produce json
// @Param request body ShortenURLRequest true "URL to shorten"
// @Success 200 {object} ShortenURLResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/shorten [post]
func (h *URLHandler) ShortenURL(c *gin.Context) {
	var request ShortenURLRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	url, err := h.service.Shorten(request.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ShortenURLResponse{
		Hash:     url.Hash,
		ShortURL: "http://localhost:8080/" + url.Hash,
	})
}

// RedirectURL godoc
// @Summary Redirect to original URL
// @Description Redirects a hash to the original URL
// @Tags URLs
// @Produce plain
// @Param hash path string true "Shortened URL hash"
// @Success 301
// @Failure 404 {object} ErrorResponse
// @Router /{hash} [get]
func (h *URLHandler) RedirectURL(c *gin.Context) {
	hash := c.Param("hash")
	url, err := h.service.FindByHash(hash)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "URL not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
	h.service.IncrementClicks(hash)
}

// GetAll godoc
// @Summary Get all shortened URLs
// @Tags URLs
// @Produce json
// @Success 200 {array} ShortenURLResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/urls [get]
func (h *URLHandler) GetAll(c *gin.Context) {
	urls, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	var response []ShortenURLResponse
	for _, url := range urls {
		response = append(response, ShortenURLResponse{
			Hash:     url.Hash,
			ShortURL: "http://localhost:8080/" + url.Hash,
		})
	}

	c.JSON(http.StatusOK, response)
}

// RegisterURLRoutes registra as rotas no router do Gin
func RegisterURLRoutes(r *gin.Engine, service *service.ShortenerService) {
	handler := NewURLHandler(service)
	api := r.Group("/api/v1")
	{
		api.POST("/shorten", handler.ShortenURL)
		api.GET("/urls", handler.GetAll)
	}

	r.GET("/:hash", handler.RedirectURL)
}
