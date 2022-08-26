package repository

import "github.com/alifudin-a/go-search-username-sia/pkg/domain/models"

type SearchUsernameSIARepository interface {
	Search(arg FullNameSIAParam) ([]models.UserSIA, error)
}

type FullNameSIAParam struct {
	FullName string
}
