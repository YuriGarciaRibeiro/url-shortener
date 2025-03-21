package handlers

import (
	"net/http"

	"github.com/YuriGarciaRibeiro/url-shortener/internal/app/service"
	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	service *service.ShortenerService
}

func NewURLHandler(s *service.ShortenerService) *URLHandler {
	return &URLHandler{service: s}
}

func (h *URLHandler) ShortenURL(c *gin.Context) {
	var request struct {
		URL string `json:"url" binding:"required,url"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url, err := h.service.Shorten(request.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to shorten URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"hash":      url.Hash,
		"short_url": "http://localhost:8080/" + url.Hash,
	})
}

func (h *URLHandler) RedirectURL(c *gin.Context) {
	hash := c.Param("hash")
	url, err := h.service.FindByHash(hash)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}

func RegisterURLRoutes(r *gin.Engine, service *service.ShortenerService) {
	handler := NewURLHandler(service)
	api := r.Group("/api/v1")
	{
		api.POST("/shorten", handler.ShortenURL)
	}

	r.GET("/:hash", handler.RedirectURL)
}
