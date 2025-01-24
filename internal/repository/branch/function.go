package branch

import (
	"github.com/inventory-service/internal/model"
)

func (b *branchRepository) Create(name, location, branchManagerId string) error {
	branch := model.Branch{
		Name:            name,
		Location:        location,
		BranchManagerID: branchManagerId,
	}

	result := b.db.Create(&branch)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *branchRepository) FindAll() ([]model.Branch, error) {
	var branches []model.Branch
	result := b.db.Find(&branches)
	if result.Error != nil {
		return nil, result.Error
	}

	return branches, nil
}

func (b *branchRepository) FindByID(id string) (*model.Branch, error) {
	var branch model.Branch
	result := b.db.Where("uuid = ?", id).First(&branch)
	if result.Error != nil {
		return nil, result.Error
	}

	return &branch, nil
}

func (b *branchRepository) Update(id, name, location, branchManagerId string) error {
	branch := model.Branch{
		Name:            name,
		Location:        location,
		BranchManagerID: branchManagerId,
	}

	result := b.db.Where("uuid = ?", id).Updates(&branch)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *branchRepository) Delete(id string) error {
	result := b.db.Where("uuid = ?", id).Delete(&model.Branch{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
