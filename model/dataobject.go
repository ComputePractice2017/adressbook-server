package model

import (
	"os"

	r "gopkg.in/gorethink/gorethink.v3"
)

//Person is a struct to store personal information
type Person struct {
	ID    string `json:"id",gorethink:"id"`
	Name  string `json:"name",gorethink:"name"`
	Email string `json:"email",gorethink:"email"`
}

var session *r.Session

func InitSesson() error {
	dbaddress := os.Getenv("RETHINKDB_HOST")
	if dbaddress != "" {
		dbaddress = "localhost"
	}
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address: dbaddress,
	})
	if err != nil {
		return err
	}

	err = CreateDBIfNotExist()
	if err != nil {
		return err
	}

	err = CreateTableIfNotExist()

	return err
}

func CreateDBIfNotExist() error {
	res, err := r.DBList().Run(session)
	if err != nil {
		return err
	}

	var dbList []string
	err = res.All(&dbList)
	if err != nil {
		return err
	}

	for _, item := range dbList {
		if item == "address" {
			return nil
		}
	}

	_, err = r.DBCreate("address").Run(session)
	if err != nil {
		return err
	}

	return nil
}

func CreateTableIfNotExist() error {
	res, err := r.DB("address").TableList().Run(session)
	if err != nil {
		return err
	}

	var tableList []string
	err = res.All(&tableList)
	if err != nil {
		return err
	}

	for _, item := range tableList {
		if item == "address" {
			return nil
		}
	}

	_, err = r.DB("address").TableCreate("address", r.TableCreateOpts{PrimaryKey: "ID"}).Run(session)

	return err
}

func GetPesrons() ([]Person, error) {
	res, err := r.DB("address").Table("address").Run(session)
	if err != nil {
		return nil, err
	}

	var response []Person
	err = res.All(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewPerson(p Person) (Person, error) {
	res, err := r.UUID().Run(session)
	if err != nil {
		return p, err
	}

	var UUID string
	err = res.One(&UUID)
	if err != nil {
		return p, err
	}

	p.ID = UUID

	res, err = r.DB("address").Table("address").Insert(p).Run(session)
	if err != nil {
		return p, err
	}

	return p, nil
}

func EditPerson(p Person) error {
	_, err := r.DB("address").Table("address").Get(p.ID).Replace(p).Run(session)
	if err != nil {
		return err
	}
	return nil
}

func DeletePerson(id string) error {
	_, err := r.DB("address").Table("address").Get(id).Delete().Run(session)
	if err != nil {
		return err
	}
	return nil
}
