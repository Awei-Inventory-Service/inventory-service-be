package branch

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (b *branchService) Create(name, location, branchManagerId string) *error_wrapper.ErrorWrapper {
	// find the proper branch manager via role and id
	branchManager, err := b.userRepository.FindById(branchManagerId)
	if err != nil {
		return err
	}

	// Todo: pindahin enum
	if branchManager.Role != model.RoleBranchManager && branchManager.Role != model.RoleBusinessOwner {
		return error_wrapper.New(model.SErrUserNotBranchManager, "User not authorized")
	}

	err = b.branchRepository.Create(name, location, branchManagerId)
	if err != nil {
		return err
	}

	return nil
}

func (b *branchService) FindAll() ([]model.Branch, *error_wrapper.ErrorWrapper) {
	branches, err := b.branchRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return branches, nil
}

func (b *branchService) FindByID(id string) (*model.Branch, *error_wrapper.ErrorWrapper) {
	branch, err := b.branchRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return branch, nil
}

func (b *branchService) Update(id, name, location, branchManagerId string) *error_wrapper.ErrorWrapper {
	branchManager, err := b.userRepository.FindById(branchManagerId)
	if err != nil {
		return err
	}

	err = b.branchRepository.Update(id, name, location, branchManager.UUID)
	if err != nil {
		return err
	}

	return nil
}

func (b *branchService) Delete(id string) *error_wrapper.ErrorWrapper {
	err := b.branchRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
