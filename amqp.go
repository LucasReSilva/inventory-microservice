package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/go-amqp"
)

const AMQP_CON string = "amqp://guest:guest@localhost:5672"
const TOPIC string = "order"
const SUB string = "inventory"
const RECEIVER_NAME string = "Team-01"

func receber(session *amqp.Session) {

	receiver, err := session.NewReceiver(context.TODO(), TOPIC+"::"+SUB, &amqp.ReceiverOptions{Name: RECEIVER_NAME})
	if err != nil {
		// handle error
	}

	// receive the next message
	msg, err := receiver.Receive(context.TODO(), nil)
	if err != nil {
		// handle error
	}

	str := fmt.Sprintf("%v", msg.Value)
	var payload Withdrawal
	json.Unmarshal([]byte(str), &payload)
	_ = receiver.AcceptMessage(context.TODO(), msg)
}
