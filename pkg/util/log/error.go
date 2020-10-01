package log

/*
Utility functions to abstract error logging.

*/

import (
	"fmt"
)

func LogError(source, contextMessage string, err error) {
	if err != nil {
		fmt.Println("\n<<Error!>> Source: ", source, "Context: ", contextMessage, "Message: ", err.Error())
	}
}

func LogOptimizationError(source, contextMessage string) {
		fmt.Println("\n<<Optimization Error!>> Source: ", source, "Context: ", contextMessage)
}
