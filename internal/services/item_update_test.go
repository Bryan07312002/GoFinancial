package services

import (
	"financial/internal/errors"
	"financial/internal/models"
	"reflect"

	"strconv"
	"testing"

	"github.com/shopspring/decimal"
)

func generateUpdateItem() (UpdateItem, *ItemRepositoryMock) {
	repo := &ItemRepositoryMock{}
	return NewUpdateItem(repo), repo
}

func mockFindByIDReturn() models.Item {
	return models.Item{
		ID:            1,
		TransactionID: 1,
		Name:          "fake item",
		Value:         decimal.NewFromInt32(10),
		Quantity:      2,
	}
}

func TestUpdateItemUpdatesSuccessffuly(t *testing.T) {
	service, itemRepo := generateUpdateItem()

	newName := "mock name 2"
	newQuantity := uint(10)
	newValue := decimal.NewFromInt32(9999)

	userID := uint(1)
	itemID := uint(1)

	FindByIDFnCalledTimes := 0
	expectedFindByIDFnCalledTimes := 1
	itemRepo.FindByIDFn = func(ID uint) (models.Item, error) {
		FindByIDFnCalledTimes += 1

		if ID != itemID {
			t.Error("expect: " + strconv.FormatUint(uint64(itemID), 10) +
				" got: " + strconv.FormatUint(uint64(ID), 10))
		}
		return mockFindByIDReturn(), nil
	}

	UpdateFnCalledTimes := 0
	expectedUpdateFnCalledTimes := 1
	itemRepo.UpdateFn = func(item models.Item, removeBagdes []uint, addBagdes []uint) error {
		UpdateFnCalledTimes += 1
		if item.Name != newName {
			t.Error("expect: " + newName + " got: " + item.Name)
		}

		if item.Value != newValue {
			t.Error("expect: " + newValue.String() + " got: " + item.Value.String())
		}

		if item.Quantity != newQuantity {
			t.Error("expect: " + strconv.FormatUint(uint64(newQuantity), 10) +
				" got: " + strconv.FormatUint(uint64(item.Quantity), 10))
		}

		return nil
	}

	if err := service.Run(itemID, UpdateItemDto{
		Name:     &newName,
		Quantity: &newQuantity,
		Value:    &newValue,
	}, userID); err != nil {
		t.Error("expect nil but got: " + err.Error())
	}

	if FindByIDFnCalledTimes != expectedFindByIDFnCalledTimes {
		t.Errorf("expected FindByID to be called " + strconv.Itoa(expectedFindByIDFnCalledTimes) + " times " +
			" but was called " + strconv.Itoa(FindByIDFnCalledTimes) + " times")
	}

	if UpdateFnCalledTimes != expectedUpdateFnCalledTimes {
		t.Errorf("expected Update to be called " + strconv.Itoa(expectedUpdateFnCalledTimes) + " times " +
			" but was called " + strconv.Itoa(UpdateFnCalledTimes) + " times")
	}
}

func TestUpdateItemShouldReturnErrorIfFindByIDReturnError(t *testing.T) {
	service, itemRepo := generateUpdateItem()
	expectedError := errors.NotFoundError()

	userID := uint(1)
	itemID := uint(1)

	FindByIDFnCalledTimes := 0
	expectedFindByIDFnCalledTimes := 1
	itemRepo.FindByIDFn = func(ID uint) (models.Item, error) {
		FindByIDFnCalledTimes += 1
		return models.Item{}, expectedError
	}

	UpdateFnCalledTimes := 0
	expectedUpdateFnCalledTimes := 0
	itemRepo.UpdateFn = func(item models.Item, removeBagdes []uint, addBagdes []uint) error {
		UpdateFnCalledTimes += 1
		return nil
	}

	err := service.Run(itemID, UpdateItemDto{}, userID)
	if err == nil {
		t.Error("expect" + expectedError.Error() + " but got: nil")
	}

	if err != nil && err.Error() != expectedError.Error() {
		t.Error("expect" + expectedError.Error() + " but got: " + err.Error())
	}

	if FindByIDFnCalledTimes != expectedFindByIDFnCalledTimes {
		t.Errorf("expected FindByID to be called " + strconv.Itoa(expectedFindByIDFnCalledTimes) + " times " +
			" but was called " + strconv.Itoa(FindByIDFnCalledTimes) + " times")
	}

	if UpdateFnCalledTimes != expectedUpdateFnCalledTimes {
		t.Errorf("expected Update to be called " + strconv.Itoa(expectedUpdateFnCalledTimes) + " times " +
			" but was called " + strconv.Itoa(UpdateFnCalledTimes) + " times")
	}
}

func TestUpdateItemShouldReturnErrorIfUpdateReturnError(t *testing.T) {
	service, itemRepo := generateUpdateItem()
	expectedError := errors.ConflictError()

	userID := uint(1)
	itemID := uint(1)

	FindByIDFnCalledTimes := 0
	expectedFindByIDFnCalledTimes := 1
	itemRepo.FindByIDFn = func(ID uint) (models.Item, error) {
		FindByIDFnCalledTimes += 1
		return mockFindByIDReturn(), nil
	}

	UpdateFnCalledTimes := 0
	expectedUpdateFnCalledTimes := 1
	itemRepo.UpdateFn = func(item models.Item, removeBagdes []uint, addBagdes []uint) error {
		UpdateFnCalledTimes += 1
		return expectedError
	}

	if err := service.Run(itemID, UpdateItemDto{}, userID); err != nil {
		t.Error("expect nil but got: " + err.Error())
	}

	if FindByIDFnCalledTimes != expectedFindByIDFnCalledTimes {
		t.Errorf("expected FindByID to be called " + strconv.Itoa(expectedFindByIDFnCalledTimes) + " times " +
			" but was called " + strconv.Itoa(FindByIDFnCalledTimes) + " times")
	}

	if UpdateFnCalledTimes != expectedUpdateFnCalledTimes {
		t.Errorf("expected Update to be called " + strconv.Itoa(expectedUpdateFnCalledTimes) + " times " +
			" but was called " + strconv.Itoa(UpdateFnCalledTimes) + " times")
	}
}

func TestUpdateItemShouldRemoveOrAddBadgesOnUpdate(t *testing.T) {
	service, itemRepo := generateUpdateItem()

	userID := uint(1)
	itemID := uint(1)

	addBadgesInput := []uint{1, 2}
	removeBadgesInput := []uint{3, 4}

	FindByIDFnCalledTimes := 0
	expectedFindByIDFnCalledTimes := 1
	itemRepo.FindByIDFn = func(ID uint) (models.Item, error) {
		FindByIDFnCalledTimes += 1
		return mockFindByIDReturn(), nil
	}

	UpdateFnCalledTimes := 0
	expectedUpdateFnCalledTimes := 1
	itemRepo.UpdateFn = func(item models.Item, removeBagdes []uint, addBagdes []uint) error {
		UpdateFnCalledTimes += 1
		if !reflect.DeepEqual(removeBadgesInput, removeBagdes) {
			t.Error("dto input should be the same as given to remove badges")
		}

		if !reflect.DeepEqual(addBadgesInput, addBagdes) {
			t.Error("dto input should be the same as given to add badges")
		}
		return nil
	}

	if err := service.Run(itemID, UpdateItemDto{
		addBadges:    &addBadgesInput,
		removeBadges: &removeBadgesInput,
	}, userID); err != nil {
		t.Error("expect nil but got: " + err.Error())
	}

	if FindByIDFnCalledTimes != expectedFindByIDFnCalledTimes {
		t.Errorf("expected FindByID to be called " + strconv.Itoa(expectedFindByIDFnCalledTimes) + " times " +
			" but was called " + strconv.Itoa(FindByIDFnCalledTimes) + " times")
	}

	if UpdateFnCalledTimes != expectedUpdateFnCalledTimes {
		t.Errorf("expected Update to be called " + strconv.Itoa(expectedUpdateFnCalledTimes) + " times " +
			" but was called " + strconv.Itoa(UpdateFnCalledTimes) + " times")
	}
}
