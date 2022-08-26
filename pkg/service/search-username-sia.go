package service

import (
	"database/sql"
	"errors"

	database "github.com/alifudin-a/go-search-username-sia/pkg/database/mysql"
	"github.com/alifudin-a/go-search-username-sia/pkg/domain/models"
	"github.com/alifudin-a/go-search-username-sia/pkg/repository"
)

type serviceSearch struct{}

func NewSearchUsernameSIARepository() repository.SearchUsernameSIARepository {
	return &serviceSearch{}
}

func (*serviceSearch) Search(arg repository.FullNameSIAParam) ([]models.UserSIA, error) {

	var db = database.DB2
	var user []models.UserSIA

	err := db.Select(&user, "SELECT u.username, u.full_name, u.email FROM users u WHERE u.full_name LIKE ?", "%"+arg.FullName+"%")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("User tidak ditemukan!")
		}
		return nil, err
	}

	return user, nil
}
