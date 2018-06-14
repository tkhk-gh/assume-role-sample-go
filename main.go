package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	s := getSessionAssumeRole()
	// s := getDefaultSession()
	svc := ec2.New(s)

	out, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{})
	if err != nil {
		log.Fatal("Error: ", err)
	}

	count := 0
	for _, r := range out.Reservations {
		count += len(r.Instances)
	}
	fmt.Printf("Instance count: %d\n", count)
}

func getSessionAssumeRole() *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
}

func getDefaultSession() *session.Session {
	return session.Must(session.NewSession())
}
