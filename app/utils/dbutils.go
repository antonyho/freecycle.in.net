package utils

import (
	"github.com/revel/revel"
	"labix.org/v2/mgo"
	//	"labix.org/v2/mgo/bson"
)

type DbUtils struct {
}

func (u *DbUtils) GetSession() (*mgo.Session, *mgo.Database) {
	// read database config
	dsn, dsnFound := revel.Config.String("db.dsn")
	if dsnFound {
		session, err := mgo.Dial(dsn)
		if err == nil {
			db := session.DB("freecycle")

			return session, db
		} else {
			revel.ERROR.Fatal("Cannot establish database connection")
		}
	} else {
		revel.ERROR.Fatal("Cannot find DSN from config")
	}

	return nil, nil
}
