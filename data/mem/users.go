// +build !mgo

package mem

import (
	"basaigbook/data/model"
)

type Users struct {
	users []model.Account
}

func (u *Users) SignUp(email, password string) (*model.Account, error) {
	acctID := len(u.users) + 1
	userID := acctID * 200

	acct := model.Account{
		ID:    acctID,
		Email: email,
		Users: []model.User{{
			Email:    email,
			ID:       userID,
			Password: password,
		}},
	}

	u.users = append(u.users, acct)
	return &acct, nil
}

func (u *Users) GetDetail(id model.Key) (*model.Account, error) {
	var user model.Account
	for _, acct := range u.users {
		if acct.ID == id {
			user = acct
			break
		}
	}
	return &user, nil
}

func (u *Users) RefreshSession(conn *bool, dbName string) {
	u.users = append(u.users, model.Account{
		ID:    1,
		Email: "test@domain.com",
		Users: []model.User{{
			Email:    "test@domain.com",
			ID:       1,
			Password: "unittest",
		}},
	})
}
