package postgres

import "clean_code/services/contact/internal/domain/group"

func (r *Repository) CreateGroup(group *group.Group) (*group.Group, error) {
	panic("implement later")
}

func (r *Repository) ReadGroupByID(groupID int) (*group.Group, error) {
	panic("implement later")
}

func (r *Repository) UpdateGroup(groupID int, group *group.Group) (*group.Group, error) {
	panic("implement later")
}

func (r *Repository) DeleteGroup(groupID int) error {
	panic("implement later")
}
