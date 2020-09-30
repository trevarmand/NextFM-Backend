package log

/*
Utility functions to abstract error logging.

*/

import (
	"fmt"
)

func LogError(source, contextMessage string, err error) {
	if err != nil {
		fmt.Println("Error! Source: ", source, "Context: ", contextMessage, "Message: ", err.Error())
	}
}
