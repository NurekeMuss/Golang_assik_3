package group

import "clean_code/services/contact/internal/useCase/adapters/storage"

type UseCase struct {
	adapterStorage storage.Group
}

func New(storage storage.Group) *UseCase {
	var uc = &UseCase{
		adapterStorage: storage,
	}
	return uc
}
