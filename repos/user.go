package repos

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/romanceresnak/go-grpc/types"
)

//UsersRepo - the user repo interface
type UsersRepo interface {
	Create(user *types.User) error
	FindById(int64) (*types.User, error)
	FindByEmail(string) (*types.User, error)
	Update(user *types.User) error
}

//NewUsersRepo - returns a new user repo
func NewUsersRepo(db *xorm.Engine) UsersRepo {
	return &userdRepo{db: db}
}

type userdRepo struct {
	db *xorm.Engine
}

func (u userdRepo) Create(user *types.User) (err error) {
	if err = types.Validate(user); err != nil {
		return
	}

	if _, err = u.db.Insert(user); err != nil {
		return
	}

	return
}

func (u userdRepo) FindByEmail(email string) (user *types.User, err error) {
	if len(email) <= 0 {
		err = fmt.Errorf("valid email is required to find a user")
	}

	user = new(types.User)
	user.Email = email

	has, err := u.db.Get(user)

	if !has {
		err = fmt.Errorf("unable to find user")
		return
	}
	return
}

func (u userdRepo) FindById(id int64) (user *types.User, err error) {
	if id <= 0 {
		err = fmt.Errorf("valid positive id is required to find a user")
	}

	user = new(types.User)

	has, err := u.db.Id(id).Get(user)

	if !has {
		err = fmt.Errorf("unable to find user")
		return
	}
	return
}

func (u userdRepo) Update(user *types.User) (err error) {
	if user == nil || user.ID <= 0 {
		return fmt.Errorf("invalid user passed in")
	}

	if _, err = u.db.Id(user.ID).Update(user); err != nil {
		return
	}

	return
}
