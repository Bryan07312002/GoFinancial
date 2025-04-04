package handlers

import (
	"financial/internal/services"

	"net/http"
)

type UpdateBadgeFactory interface {
	CreateUpdateBadge() services.UpdateBadge
}

type UpdateBadgeHandler struct {
	factory UpdateBadgeFactory
}

func NewUpdateBadgeHandler(factory UpdateBadgeFactory) http.Handler {
	return &UpdateBadgeHandler{factory}
}

type UpdateBadgeRequestBody struct {
	Name  *string `json:"name"`
	Color *string `json:"color"`
}

func (u *UpdateBadgeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var dto UpdateBadgeRequestBody
	userID, err := extractBodyAndUserId(r, &dto)
	if err != nil {
		writeError(err, w)
		return
	}

	badgeID, err := extractUintFromUrl(r, "id")
	if err != nil {
    println(err.Error())
		writeError(err, w)
		return
	}

	if err := u.factory.CreateUpdateBadge().Run(
		uint(badgeID),
		services.UpdateBadgeDto{
			Name:  dto.Name,
			Color: dto.Color,
		},
		userID,
	); err != nil {
		writeError(err, w)
		return
	}
}
