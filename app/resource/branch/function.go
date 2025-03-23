package branch

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (b *branchResource) Create(newBranch model.Branch) *error_wrapper.ErrorWrapper {

	result := b.db.Create(&newBranch)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return nil
}

func (b *branchResource) FindAll() ([]model.Branch, *error_wrapper.ErrorWrapper) {
	var branches []model.Branch
	result := b.db.Preload("BranchManager").Find(&branches)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return branches, nil
}

func (b *branchResource) FindByID(id string) (*model.Branch, *error_wrapper.ErrorWrapper) {
	var branch model.Branch
	result := b.db.Preload("BranchManager").Where("uuid = ?", id).First(&branch)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &branch, nil
}

func (b *branchResource) Update(id string, branch model.Branch) *error_wrapper.ErrorWrapper {

	result := b.db.Where("uuid = ?", id).Updates(&branch)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return nil
}

func (b *branchResource) Delete(id string) *error_wrapper.ErrorWrapper {
	result := b.db.Where("uuid = ?", id).Delete(&model.Branch{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}
