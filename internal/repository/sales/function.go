package sales

import (
	"github.com/inventory-service/internal/model"
)

func (s *salesRepository) Create(sale model.Sales) error {
	result := s.db.Create(&sale)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *salesRepository) FindAll() ([]model.Sales, error) {
	var sales []model.Sales
	result := s.db.Find(&sales)
	if result.Error != nil {
		return nil, result.Error
	}

	return sales, nil
}

func (s *salesRepository) FindByID(id string) (*model.Sales, error) {
	var sale model.Sales
	result := s.db.Where("uuid = ?", id).First(&sale)
	if result.Error != nil {
		return nil, result.Error
	}

	return &sale, nil
}

func (s *salesRepository) Update(id string, sale model.Sales) error {
	result := s.db.Where("uuid = ?", id).Updates(&sale)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *salesRepository) Delete(id string) error {
	result := s.db.Where("uuid = ?", id).Delete(&model.Sales{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
