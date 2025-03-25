package service

import (
	"encoding/base64"

	"github.com/YuriGarciaRibeiro/url-shortener/internal/app/model"
	"github.com/YuriGarciaRibeiro/url-shortener/internal/storage/postgres"
)

type ShortenerService struct {
	repo *postgres.URLRepository
}

func NewShortenerService(repo *postgres.URLRepository) *ShortenerService {
	return &ShortenerService{repo: repo}
}

func (s *ShortenerService) Shorten(url string) (*model.Url, error) {
	// Gera um hash único (ex: usando uma função de hashing)
	hash := generateHash(url)

	// Cria a URL no banco de dados
	newURL := &model.Url{
		Hash:        hash,
		OriginalURL: url,
	}
	if err := s.repo.Save(newURL); err != nil {
		return nil, err
	}

	return newURL, nil
}

func (s *ShortenerService) FindByHash(hash string) (*model.Url, error) {
	return s.repo.FindByHash(hash)
}

func (s *ShortenerService) IncrementClicks(hash string) error {
	return s.repo.IncrementClicks(hash)
}

func (s *ShortenerService) GetAll() ([]model.Url, error) {
	return s.repo.GetAll()
}

// Função para gerar hash (exemplo simples)
func generateHash(input string) string {
	base64Input := base64.StdEncoding.EncodeToString([]byte(input))
	base64Input = base64.URLEncoding.EncodeToString([]byte(base64Input))
	return base64Input[:8]
}
