package aws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/trevarmand/nextfm-backend/pkg/util/log"
)

/*
	Fetch an AWS Session based on AWS credential environment variables.
*/
func GetSession() (*session.Session, error) {
	os.Getenv("")
	sess, err := session.NewSession()

	if err != nil {
		log.LogError("util:aws:AwsSession:GetSession", err)
		return nil, err
	}
	return sess, nil
}
