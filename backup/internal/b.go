package internal

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

var AWSSession *session.Session

func init() {
	AWSSession = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
}