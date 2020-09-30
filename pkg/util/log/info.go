package log

import (
	"fmt"
)


func LogInfo(source, message string, data map[string]interface{}) {
	fmt.Println("\n<<INFO LOG>> SOURCE: ", source, " MESSAGE: ", message, "DATA:\n", data)
}
