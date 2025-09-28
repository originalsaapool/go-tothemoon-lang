package repository

import "todorestapi/internal/model"

type UserRepository struct {
	users []model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{users: []model.User{}}
}

func (r *UserRepository) GetAll() []model.User {
	return r.users
}

func (r *UserRepository) GetByID(id int) *model.User {
	for _, u := range r.users {
		if u.ID == id {
			return &u
		}
	}
	return nil
}

func (r *UserRepository) Save(u model.User) model.User {
	u.ID = len(r.users) + 1
	r.users = append(r.users, u)
	return u
}
