package services

import (
	"fastmath/repositories"
)

type AllRepository struct {
	ApiRepository repositories.ApiRepositoryContext
}

type AllService struct {
	ApiService ApiService
}

type AllConfigFromDB struct {
}
