// base
package models

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Doc struct {
	Id    bson.ObjectId `bson:"_id"`
	Value interface{}   `bson:"value"`
}

type Name struct {
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
}

const CollectionName = "names"

/*
func main() {
	fmt.Println("Hello World!")
}
*/

func GetDB() (*mgo.Session, *mgo.Database, error) {
	session, err1 := mgo.Dial("localhost")
	if err1 != nil {
		fmt.Println("Session error:", err1)
		return session, nil, err1
	}
	db := session.DB("nameweb1")
	return session, db, nil
}

func (this Name) GetDuplicateCount() (int, error) {
	session, db, sessionErr := GetDB()
	if sessionErr != nil {
		fmt.Println("session error:", sessionErr)
		return 0, sessionErr
	}
	defer session.Close()
	c := db.C(CollectionName)
	count, err := c.Find(bson.M{"value": this}).Count()
	if err != nil {
		fmt.Println("Find error:", err)
		return count, err
	}
	return count, nil
}

func (this Name) AddName() error {
	session, db, sessionErr := GetDB()
	if sessionErr != nil {
		fmt.Println("session error:", sessionErr)
		return sessionErr
	}
	defer session.Close()
	c := db.C(CollectionName)
	obj := Doc{Id: bson.NewObjectId(), Value: this}
	if err := c.Insert(obj); err != nil {
		return err
	}
	return nil
}

func GetAllNames() ([]Doc, error) {
	session, db, sessionErr := GetDB()
	if sessionErr != nil {
		fmt.Println("session error:", sessionErr)
		return nil, sessionErr
	}
	defer session.Close()
	c := db.C(CollectionName)
	var results []Doc
	err := c.Find(nil).All(&results)
	return results, err
}
