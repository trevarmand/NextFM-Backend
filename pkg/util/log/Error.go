package log

/*
Utility functions to abstract error logging.

*/

import (
	"fmt"
)

func LogError(source string, err error) {
	fmt.Println("Error! Source: ", source, "Message:", err.Error())
}
