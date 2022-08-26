package rest

import (
	"encoding/json"
	"net/http"

	"github.com/alifudin-a/go-search-username-sia/pkg/domain/models"
	"github.com/alifudin-a/go-search-username-sia/pkg/repository"
	"github.com/alifudin-a/go-search-username-sia/pkg/service"

	"github.com/gofiber/fiber/v2"
)

type search struct{}

func NewSearchUsernameSIAHandler() *search {
	return &search{}
}

func (*search) SearchUsernameSIAHandler(c *fiber.Ctx) (err error) {

	var resp models.Response
	var user []models.UserSIA

	var arr []string
	_ = json.Unmarshal(c.Body(), &arr)

	srvc := service.NewSearchUsernameSIARepository()

	arg := repository.FullNameSIAParam{
		FullName: arr[0],
	}

	user, err = srvc.Search(arg)
	if err != nil {
		resp.Code = http.StatusNotFound
		resp.Message = "User tidak ditemukan!"
		return c.JSON(resp)
	}

	if user == nil {
		resp.Code = http.StatusNotFound
		resp.Message = "User tidak ditemukan!"
		return c.JSON(resp)
	}

	return c.JSON(fiber.Map{
		"data": user,
	})
}
