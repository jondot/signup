package main

import (
	"launchpad.net/mgo"
	"launchpad.net/mgo/bson"
	"time"
)

type Contact struct {
	ForProject string
	Email      string
	Created_at int64
}

func (ent Contact) Create(project string, email string) {
	sess, err := mgo.Dial(mongo)
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	c := sess.DB(mongo_db).C("contacts")
	t := time.Now()
	_, err = c.Upsert(bson.M{"email": email, "forproject": project}, &Contact{project, email, t.Unix()})
	if err != nil {
		panic(err)
	}
}
