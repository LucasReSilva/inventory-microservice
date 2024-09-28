package main

import (
	"context"

	"github.com/Azure/go-amqp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var session *amqp.Session

func main() {
	db, _ := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	db.AutoMigrate(&Stock{})

	conn, err := amqp.Dial(context.TODO(), AMQP_CON, &amqp.ConnOptions{ContainerID: "GoCon"})
	if err != nil {
		// handle error
	}

	session, err := conn.NewSession(context.TODO(), nil)
	if err != nil {
		// handle error
	}
	receber(session)
}
