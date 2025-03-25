package postgres

import (
	"github.com/YuriGarciaRibeiro/url-shortener/internal/app/model"
	"gorm.io/gorm"
)

type URLRepository struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) *URLRepository {
	return &URLRepository{db: db}
}

// Salva uma URL no banco de dados
func (r *URLRepository) Save(url *model.Url) error {
	return r.db.Create(url).Error
}

// Busca uma URL pelo hash
func (r *URLRepository) FindByHash(hash string) (*model.Url, error) {
	var url model.Url
	err := r.db.Where("hash = ?", hash).First(&url).Error
	return &url, err
}

// Incrementa o contador de cliques
func (r *URLRepository) IncrementClicks(hash string) error {
	return r.db.Model(&model.Url{}).Where("hash = ?", hash).Update("clicks", gorm.Expr("clicks + 1")).Error
}

func (r *URLRepository) GetAll() ([]model.Url, error) {
	var urls []model.Url
	err := r.db.Find(&urls).Error
	return urls, err
}
