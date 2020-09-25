package log

import (
	"fmt"
)

type InfoMessage struct {
	source, message string
	data            map[string]interface{}
}

func LogInfo(infoMsg InfoMessage) {
	fmt.Println(infoMsg)
}
