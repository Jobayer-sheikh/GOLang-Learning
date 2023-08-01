package rabbitmq

import (
	"bytes"
	"encoding/binary"
	"fiber/learning/handler/ordermessage"
	"fmt"
	"log"
)

func ParseCommandAction(actionType any, body []byte) {

	var commandType string = actionType.(string)
	if commandType == "OrderMessageCommand" {

		var command = new(ordermessage.OrderMessageCommand)
		if generateCommand(commandType, body, command) {
			command.HandleInternal()
		}
	}
}

func generateCommand(actionType string, body []byte, command any) bool {
	buf := new(bytes.Buffer)
	buf.Write(body)

	if err := binary.Read(buf, binary.BigEndian, command); err != nil {
		log.Println(fmt.Sprintf("Can't generate command of type %s", actionType))
		return false
	}

	return true
}
