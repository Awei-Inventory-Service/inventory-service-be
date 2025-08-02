package branch

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (b *branchDomain) Create(name, location, branchManagerId string) *error_wrapper.ErrorWrapper {
	newBranch := model.Branch{
		Name:            name,
		Location:        location,
		BranchManagerID: branchManagerId,
	}

	return b.branchResource.Create(newBranch)
}

func (b *branchDomain) FindAll() ([]model.Branch, *error_wrapper.ErrorWrapper) {
	return b.branchResource.FindAll()
}

func (b *branchDomain) FindByID(id string) (*model.Branch, *error_wrapper.ErrorWrapper) {
	return b.branchResource.FindByID(id)
}

func (b *branchDomain) Update(id, name, location, branchManagerId string) *error_wrapper.ErrorWrapper {
	newBranch := model.Branch{
		Name:            name,
		Location:        location,
		BranchManagerID: branchManagerId,
	}

	return b.branchResource.Update(id, newBranch)
}

func (b *branchDomain) Delete(id string) *error_wrapper.ErrorWrapper {
	return b.branchResource.Delete(id)
}
