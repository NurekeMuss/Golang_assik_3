package postgres

import (
	"clean_code/services/contact/internal/domain/contact"
)

func (r *Repository) CreateContantIntoGroup(groupID int, contact *contact.Contact) (*contact.Contact, error) {
	panic("implement later")
}

func (r *Repository) AddContactToGroup(groupID int, contactID int) error {
	panic("implement later")
}

func (r *Repository) DeleteContantFromGroup(groupID int, contactID int) error {
	panic("implement later")
}
