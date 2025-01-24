package branch

import (
	"errors"

	"github.com/inventory-service/internal/model"
)

func (b *branchService) Create(name, location, branchManagerId string) error {
	// find the proper branch manager via role and id
	branchManager, err := b.userRepository.FindById(branchManagerId)
	if err != nil {
		return err
	}

	// Todo: pindahin enum
	if branchManager.Role != "branch_manager" {
		return errors.New("user is not a branch manager")
	}

	err = b.branchRepository.Create(name, location, branchManagerId)
	if err != nil {
		return err
	}

	return nil
}

func (b *branchService) FindAll() ([]model.Branch, error) {
	branches, err := b.branchRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return branches, nil
}

func (b *branchService) FindByID(id string) (*model.Branch, error) {
	branch, err := b.branchRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return branch, nil
}

func (b *branchService) Update(id, name, location, branchManagerId string) error {
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

func (b *branchService) Delete(id string) error {
	err := b.branchRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
