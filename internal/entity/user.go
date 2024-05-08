package entity

import "errors"

type User struct {
	id            int
	name          string
	balance       int
	versionUpdate int
}

func NewUser(id int, name string, balance int, versionUpdate int) User {
	return User{
		id:            id,
		name:          name,
		balance:       balance,
		versionUpdate: versionUpdate,
	}
}

func (u *User) SubtractAmount(sum int) error {
	if u.balance < sum {
		return errors.New("not enough money")
	}

	u.balance -= sum

	return nil
}

func (u *User) GetBalance() int {
	return u.balance
}

func (u *User) GetId() int {
	return u.id
}

func (u *User) GetVersionUpdate() int {
	return u.versionUpdate
}
