package ec2

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/cawhite/eventstore-tool/backup/internal"
)

var ec2Client *ec2.EC2

func init() {
	ec2Client = ec2.New(internal.AWSSession)
}

func GetInstances(env string) {
	describeInstancesInput := &ec2.DescribeInstancesInput{}
	if env != "" {
		filter := &ec2.Filter{Name:
		aws.String("tag:environment"),
			Values: []*string{aws.String(env)},
		}
		describeInstancesInput.Filters = []*ec2.Filter{filter}
	}
	ec2Client.DescribeInstancesPages(describeInstancesInput, func(output *ec2.DescribeInstancesOutput, b bool) bool {
		for _, reservation := range output.Reservations {
			for _, instance := range reservation.Instances {
				fmt.Println(*instance.InstanceId)
			}
		}
		return true
	})
}
