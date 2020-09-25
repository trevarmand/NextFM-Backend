package main

import (
	"fmt"
	"github.com/trevarmand/nextfm-backend/pkg/util/aws"
)


func main() {
	fmt.Println("Starting NextFm Backend")

	sess, err := aws.GetSession()
	if(err != nil) {
		fmt.Println("Failed to establish session, exiting")
		return
	}

	fmt.Println("successfully established connection!", sess)
} 