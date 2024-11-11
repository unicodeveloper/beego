package main

import (
	"fmt"
	_ "helloworld/routers"
	"log"

	_ "github.com/lib/pq"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

// Users -
type Users struct {
	ID        int    `orm:"column(id)"`
	FirstName string `orm:"column(first_name)"`
	LastName  string `orm:"column(last_name)"`
}

func init() {
	// set default database
	orm.RegisterDriver("postgres", orm.DRPostgres)

	// Grab the db url from the app config
	dbURL, err := beego.AppConfig.String("db_url")
	if err != nil {
		log.Fatal("Error getting database URL: ", err)
	}

	// set default database
	orm.RegisterDataBase("default", "postgres", dbURL)

	// register model
	orm.RegisterModel(new(Users))

	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	// Create a slice of Users to insert
	users := []Users{
		{FirstName: "John", LastName: "Doe"},
		{FirstName: "Jane", LastName: "Doe"},
		{FirstName: "Railway", LastName: "Deploy Beego"},
	}

	// Iterate over the slice and insert each user
	for _, user := range users {
		id, err := o.Insert(&user)
		if err != nil {
			fmt.Printf("Failed to insert user %s %s: %v\n", user.FirstName, user.LastName, err)
		} else {
			fmt.Printf("Inserted user ID: %d, Name: %s %s\n", id, user.FirstName, user.LastName)
		}
	}

	beego.Run()
}
