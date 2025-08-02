package branch

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (b *branchService) Create(name, location, branchManagerId string) *error_wrapper.ErrorWrapper {
	// find the proper branch manager via role and id
	branchManager, err := b.userDomain.FindById(branchManagerId)
	if err != nil {
		return err
	}

	// Todo: pindahin enum
	if branchManager.Role != model.RoleBranchManager && branchManager.Role != model.RoleBusinessOwner {
		return error_wrapper.New(model.SErrUserNotBranchManager, "User not authorized")
	}

	err = b.branchDomain.Create(name, location, branchManagerId)
	if err != nil {
		return err
	}

	return nil
}

func (b *branchService) FindAll() ([]model.Branch, *error_wrapper.ErrorWrapper) {
	branches, err := b.branchDomain.FindAll()
	if err != nil {
		return nil, err
	}

	return branches, nil
}

func (b *branchService) FindByID(id string) (*model.Branch, *error_wrapper.ErrorWrapper) {
	branch, err := b.branchDomain.FindByID(id)
	if err != nil {
		return nil, err
	}

	return branch, nil
}

func (b *branchService) Update(id, name, location, branchManagerId string) *error_wrapper.ErrorWrapper {
	branchManager, err := b.userDomain.FindById(branchManagerId)
	if err != nil {
		return err
	}

	err = b.branchDomain.Update(id, name, location, branchManager.UUID)
	if err != nil {
		return err
	}

	return nil
}

func (b *branchService) Delete(id string) *error_wrapper.ErrorWrapper {
	err := b.branchDomain.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
