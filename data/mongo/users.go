// +build !mem

package mongo

import (
	"basaigbook/data/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Users struct {
	DB *mgo.Database
}

func (u *Users) SignUp(email, password string) (*model.Account, error) {
	acct := model.Account{ID: bson.NewObjectId(), Email: email}
	acct.Users = append(acct.Users, model.User{ID: bson.NewObjectId(), Email: email, Password: password})
	if err := u.DB.C("users").Insert(acct); err != nil {
		return nil, err
	}
	return &acct, nil
}

func (u *Users) GetDetail(id model.Key) (*model.Account, error) {
	var acct model.Account
	where := bson.M{"_id": id}
	if err := u.DB.C("users").Find(where).One(&acct); err != nil {
		return nil, err
	}
	return &acct, nil
}

func (u *Users) RefreshSession(s *mgo.Session, dbName string) {
	u.DB = s.Copy().DB(dbName)
}