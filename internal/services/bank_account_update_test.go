package services

import (
	"financial/internal/errors"
	"financial/internal/models"

	"testing"
)

func generateUpdateBankAccount() (UpdateBankAccount, *BankAccountRepositoryMock) {
	repo := &BankAccountRepositoryMock{}
	return NewUpdateBankAccount(repo), repo
}

func TestUpdateBankAccountSuccessfully(t *testing.T) {
	service, bankAccRepo := generateUpdateBankAccount()

	returnedBankAcc := models.BankAccount{
		ID:          1,
		UserID:      1,
		Name:        "test",
		Description: "some description",
	}
	updatedName := "test"
	updatedDescription := "some description"
	dto := UpdateBankAccountDto{
		Name:        &updatedName,
		Description: &updatedDescription,
	}

	FindByIDCalledTimes := 0
	bankAccRepo.FindByIDFunc = func(ID, userId uint) (models.BankAccount, error) {
		FindByIDCalledTimes += 1
		return returnedBankAcc, nil
	}

	UpdateCalledTimes := 0
	bankAccRepo.UpdateFunc = func(bankAccount models.BankAccount) error {
		UpdateCalledTimes += 1
		if bankAccount.Name != updatedName {
			t.Error("expect bank account name to be: " + updatedName + "but got: " + bankAccount.Name)
		}
		if bankAccount.Description != updatedDescription {
			t.Error("expect bank account description to be: " + updatedDescription + "but got: " + bankAccount.Description)
		}
		return nil
	}

	if err := service.Run(1, dto, 1); err != nil {
		t.Error("expected no errors but got: ", err.Error())
	}

	if FindByIDCalledTimes != 1 {
		t.Error("FindByID to have beein called 1 time but was called ",
			FindByIDCalledTimes)
	}

	if UpdateCalledTimes != 1 {
		t.Error("Update to have been called 1 time but was called ",
			UpdateCalledTimes)
	}
}

func TestUpdateBankAccountShouldReturnErrorWhenFindByIDReturnError(t *testing.T) {
	service, bankAccRepo := generateUpdateBankAccount()
	expectedError := errors.NotFoundError()

	updatedName := "test"
	updatedDescription := "some description"
	dto := UpdateBankAccountDto{
		Name:        &updatedName,
		Description: &updatedDescription,
	}

	FindByIDCalledTimes := 0
	bankAccRepo.FindByIDFunc = func(ID, userId uint) (models.BankAccount, error) {
		FindByIDCalledTimes += 1
		return models.BankAccount{}, expectedError
	}

	UpdateCalledTimes := 0
	bankAccRepo.UpdateFunc = func(bankAccount models.BankAccount) error {
		UpdateCalledTimes += 1
		return nil
	}

	if err := service.Run(1, dto, 1); err == nil {
		t.Error("expected error " + expectedError.Error() + " but got nil")
	}

	if FindByIDCalledTimes != 1 {
		t.Error("FindByID to have been called 1 time but was called ",
			FindByIDCalledTimes)
	}

	if UpdateCalledTimes != 0 {
		t.Error("Update to have been called 0 times but was called ",
			UpdateCalledTimes)
	}
}

func TestUpdateBankAccountShouldReturnErrorWhenUpdateReturnError(t *testing.T) {
	service, bankAccRepo := generateUpdateBankAccount()
	expectedError := errors.NotFoundError()

	updatedName := "test"
	updatedDescription := "some description"
	dto := UpdateBankAccountDto{
		Name:        &updatedName,
		Description: &updatedDescription,
	}

	returnedBankAcc := models.BankAccount{
		ID:          1,
		UserID:      1,
		Name:        "test",
		Description: "some description",
	}
	FindByIDCalledTimes := 0
	bankAccRepo.FindByIDFunc = func(ID, userId uint) (models.BankAccount, error) {
		FindByIDCalledTimes += 1
		return returnedBankAcc, nil
	}

	UpdateCalledTimes := 0
	bankAccRepo.UpdateFunc = func(bankAccount models.BankAccount) error {
		UpdateCalledTimes += 1
		return expectedError
	}

	if err := service.Run(1, dto, 1); err == nil {
		t.Errorf("expected error " + expectedError.Error() + " but got nil")
	}

	if FindByIDCalledTimes != 1 {
		t.Error("FindByID to have been called 1 time but was called ",
			FindByIDCalledTimes)
	}

	if UpdateCalledTimes != 1 {
		t.Error("Update to have been called 1 times but was called ",
			UpdateCalledTimes)
	}
}
