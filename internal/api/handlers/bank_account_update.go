package handlers

import (
	"financial/internal/services"
	"net/http"
)

type UpdateBankAccountFactory interface {
	CreateUpdateBankAccount() services.UpdateBankAccount
}

type UpdateBankAccountHandler struct {
	factory UpdateBankAccountFactory
}

func NewUpdateBankAccountHandler(factory UpdateBankAccountFactory) http.Handler {
	return &UpdateBankAccountHandler{factory}
}

type UpdateBankAccountRequestBody struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

func (u *UpdateBankAccountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var dto UpdateBankAccountRequestBody
	userID, err := extractBodyAndUserId(r, &dto)
	if err != nil {
		writeError(err, w)
		return
	}

	bankAccountId, err := extractUintFromUrl(r, "id")
	if err != nil {
		writeError(err, w)
		return
	}

	if err := u.factory.CreateUpdateBankAccount().Run(
		uint(bankAccountId),
		services.UpdateBankAccountDto{
			Name:        dto.Name,
			Description: dto.Description,
		},
		userID,
	); err != nil {
		writeError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
