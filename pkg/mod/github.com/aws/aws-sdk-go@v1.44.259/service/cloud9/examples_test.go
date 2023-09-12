// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloud9_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloud9"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// CreateEnvironmentEC2
//

func ExampleCloud9_CreateEnvironmentEC2_shared00() {
	svc := cloud9.New(session.New())
	input := &cloud9.CreateEnvironmentEC2Input{
		AutomaticStopTimeMinutes: aws.Int64(60),
		Description:              aws.String("This is my demonstration environment."),
		InstanceType:             aws.String("t2.micro"),
		Name:                     aws.String("my-demo-environment"),
		OwnerArn:                 aws.String("arn:aws:iam::123456789012:user/MyDemoUser"),
		SubnetId:                 aws.String("subnet-6300cd1b"),
	}

	result, err := svc.CreateEnvironmentEC2(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloud9.ErrCodeBadRequestException:
				fmt.Println(cloud9.ErrCodeBadRequestException, aerr.Error())
			case cloud9.ErrCodeConflictException:
				fmt.Println(cloud9.ErrCodeConflictException, aerr.Error())
			case cloud9.ErrCodeNotFoundException:
				fmt.Println(cloud9.ErrCodeNotFoundException, aerr.Error())
			case cloud9.ErrCodeForbiddenException:
				fmt.Println(cloud9.ErrCodeForbiddenException, aerr.Error())
			case cloud9.ErrCodeTooManyRequestsException:
				fmt.Println(cloud9.ErrCodeTooManyRequestsException, aerr.Error())
			case cloud9.ErrCodeLimitExceededException:
				fmt.Println(cloud9.ErrCodeLimitExceededException, aerr.Error())
			case cloud9.ErrCodeInternalServerErrorException:
				fmt.Println(cloud9.ErrCodeInternalServerErrorException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// CreateEnvironmentMembership
//

func ExampleCloud9_CreateEnvironmentMembership_shared00() {
	svc := cloud9.New(session.New())
	input := &cloud9.CreateEnvironmentMembershipInput{
		EnvironmentId: aws.String("8d9967e2f0624182b74e7690ad69ebEX"),
		Permissions:   aws.String("read-write"),
		UserArn:       aws.String("arn:aws:iam::123456789012:user/AnotherDemoUser"),
	}

	result, err := svc.CreateEnvironmentMembership(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloud9.ErrCodeBadRequestException:
				fmt.Println(cloud9.ErrCodeBadRequestException, aerr.Error())
			case cloud9.ErrCodeConflictException:
				fmt.Println(cloud9.ErrCodeConflictException, aerr.Error())
			case cloud9.ErrCodeNotFoundException:
				fmt.Println(cloud9.ErrCodeNotFoundException, aerr.Error())
			case cloud9.ErrCodeForbiddenException:
				fmt.Println(cloud9.ErrCodeForbiddenException, aerr.Error())
			case cloud9.ErrCodeTooManyRequestsException:
				fmt.Println(cloud9.ErrCodeTooManyRequestsException, aerr.Error())
			case cloud9.ErrCodeLimitExceededException:
				fmt.Println(cloud9.ErrCodeLimitExceededException, aerr.Error())
			case cloud9.ErrCodeInternalServerErrorException:
				fmt.Println(cloud9.ErrCodeInternalServerErrorException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// DeleteEnvironment
//

func ExampleCloud9_DeleteEnvironment_shared00() {
	svc := cloud9.New(session.New())
	input := &cloud9.DeleteEnvironmentInput{
		EnvironmentId: aws.String("8d9967e2f0624182b74e7690ad69ebEX"),
	}

	result, err := svc.DeleteEnvironment(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloud9.ErrCodeBadRequestException:
				fmt.Println(cloud9.ErrCodeBadRequestException, aerr.Error())
			case cloud9.ErrCodeConflictException:
				fmt.Println(cloud9.ErrCodeConflictException, aerr.Error())
			case cloud9.ErrCodeNotFoundException:
				fmt.Println(cloud9.ErrCodeNotFoundException, aerr.Error())
			case cloud9.ErrCodeForbiddenException:
				fmt.Println(cloud9.ErrCodeForbiddenException, aerr.Error())
			case cloud9.ErrCodeTooManyRequestsException:
				fmt.Println(cloud9.ErrCodeTooManyRequestsException, aerr.Error())
			case cloud9.ErrCodeLimitExceededException:
				fmt.Println(cloud9.ErrCodeLimitExceededException, aerr.Error())
			case cloud9.ErrCodeInternalServerErrorException:
				fmt.Println(cloud9.ErrCodeInternalServerErrorException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// DeleteEnvironmentMembership
//

func ExampleCloud9_DeleteEnvironmentMembership_shared00() {
	svc := cloud9.New(session.New())
	input := &cloud9.DeleteEnvironmentMembershipInput{
		EnvironmentId: aws.String("8d9967e2f0624182b74e7690ad69ebEX"),
		UserArn:       aws.String("arn:aws:iam::123456789012:user/AnotherDemoUser"),
	}

	result, err := svc.DeleteEnvironmentMembership(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloud9.ErrCodeBadRequestException:
				fmt.Println(cloud9.ErrCodeBadRequestException, aerr.Error())
			case cloud9.ErrCodeConflictException:
				fmt.Println(cloud9.ErrCodeConflictException, aerr.Error())
			case cloud9.ErrCodeNotFoundException:
				fmt.Println(cloud9.ErrCodeNotFoundException, aerr.Error())
			case cloud9.ErrCodeForbiddenException:
				fmt.Println(cloud9.ErrCodeForbiddenException, aerr.Error())
			case cloud9.ErrCodeTooManyRequestsException:
				fmt.Println(cloud9.ErrCodeTooManyRequestsException, aerr.Error())
			case cloud9.ErrCodeLimitExceededException:
				fmt.Println(cloud9.ErrCodeLimitExceededException, aerr.Error())
			case cloud9.ErrCodeInternalServerErrorException:
				fmt.Println(cloud9.ErrCodeInternalServerErrorException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// DescribeEnvironmentMemberships1
// The following example gets information about all of the environment members for the
// specified development environment.
func ExampleCloud9_DescribeEnvironmentMemberships_shared00() {
	svc := cloud9.New(session.New())
	input := &cloud9.DescribeEnvironmentMembershipsInput{
		EnvironmentId: aws.String("8d9967e2f0624182b74e7690ad69ebEX"),
	}

	result, err := svc.DescribeEnvironmentMemberships(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloud9.ErrCodeBadRequestException:
				fmt.Println(cloud9.ErrCodeBadRequestException, aerr.Error())
			case cloud9.ErrCodeConflictException:
				fmt.Println(cloud9.ErrCodeConflictException, aerr.Error())
			case cloud9.ErrCodeNotFoundException:
				fmt.Println(cloud9.ErrCodeNotFoundException, aerr.Error())
			case cloud9.ErrCodeForbiddenException:
				fmt.Println(cloud9.ErrCodeForbiddenException, aerr.Error())
			case cloud9.ErrCodeTooManyRequestsException:
				fmt.Println(cloud9.ErrCodeTooManyRequestsException, aerr.Error())
			case cloud9.ErrCodeLimitExceededException:
				fmt.Println(cloud9.ErrCodeLimitExceededException, aerr.Error())
			case cloud9.ErrCodeInternalServerErrorException:
				fmt.Println(cloud9.ErrCodeInternalServerErrorException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// DescribeEnvironmentMemberships2
// The following example gets information about the owner of the specified development
// environment.
func ExampleCloud9_DescribeEnvironmentMemberships_shared01() {
	svc := cloud9.New(session.New())
	input := &cloud9.DescribeEnvironmentMembershipsInput{
		EnvironmentId: aws.String("8d9967e2f0624182b74e7690ad69ebEX"),
		Permissions: []*string{
			aws.String("owner"),
		},
	}

	result, err := svc.DescribeEnvironmentMemberships(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloud9.ErrCodeBadRequestException:
				fmt.Println(cloud9.ErrCodeBadRequestException, aerr.Error())
			case cloud9.ErrCodeConflictException:
				fmt.Println(cloud9.ErrCodeConflictException, aerr.Error())
			case cloud9.ErrCodeNotFoundException:
				fmt.Println(cloud9.ErrCodeNotFoundException, aerr.Error())
			case cloud9.ErrCodeForbiddenException:
				fmt.Println(cloud9.ErrCodeForbiddenException, aerr.Error())
			case cloud9.ErrCodeTooManyRequestsException:
				fmt.Println(cloud9.ErrCodeTooManyRequestsException, aerr.Error())
			case cloud9.ErrCodeLimitExceededException:
				fmt.Println(cloud9.ErrCodeLimitExceededException, aerr.Error())
			case cloud9.ErrCodeInternalServerErrorException:
				fmt.Println(cloud9.ErrCodeInternalServerErrorException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// DescribeEnvironmentMemberships3
// The following example gets development environment membership information for the
// specified user.
func ExampleCloud9_DescribeEnvironmentMemberships_shared02() {
	svc := cloud9.New(session.New())
	input := &cloud9.DescribeEnvironmentMembershipsInput{
		UserArn: aws.String("arn:aws:iam::123456789012:user/MyDemoUser"),
	}

	result, err := svc.DescribeEnvironmentMemberships(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloud9.ErrCodeBadRequestException:
				fmt.Println(cloud9.ErrCodeBadRequestException, aerr.Error())
			case cloud9.ErrCodeConflictException:
				fmt.Println(cloud9.ErrCodeConflictException, aerr.Error())
			case cloud9.ErrCodeNotFoundException:
				fmt.Println(cloud9.ErrCodeNotFoundException, aerr.Error())
			case cloud9.ErrCodeForbiddenException:
				fmt.Println(cloud9.ErrCodeForbiddenException, aerr.Error())
			case cloud9.ErrCodeTooManyRequestsException:
				fmt.Println(cloud9.ErrCodeTooManyRequestsException, aerr.Error())
			case cloud9.ErrCodeLimitExceededException:
				fmt.Println(cloud9.ErrCodeLimitExceededException, aerr.Error())
			case cloud9.ErrCodeInternalServerErrorException:
				fmt.Println(cloud9.ErrCodeInternalServerErrorException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// DescribeEnvironmentStatus
//

func ExampleCloud9_DescribeEnvironmentStatus_shared00() {
	svc := cloud9.New(session.New())
	input := &cloud9.DescribeEnvironmentStatusInput{
		EnvironmentId: aws.String("8d9967e2f0624182b74e7690ad69ebEX"),
	}

	result, err := svc.DescribeEnvironmentStatus(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloud9.ErrCodeBadRequestException:
				fmt.Println(cloud9.ErrCodeBadRequestException, aerr.Error())
			case cloud9.ErrCodeConflictException:
				fmt.Println(cloud9.ErrCodeConflictException, aerr.Error())
			case cloud9.ErrCodeNotFoundException:
				fmt.Println(cloud9.ErrCodeNotFoundException, aerr.Error())
			case cloud9.ErrCodeForbiddenException:
				fmt.Println(cloud9.ErrCodeForbiddenException, aerr.Error())
			case cloud9.ErrCodeTooManyRequestsException:
				fmt.Println(cloud9.ErrCodeTooManyRequestsException, aerr.Error())
			case cloud9.ErrCodeLimitExceededException:
				fmt.Println(cloud9.ErrCodeLimitExceededException, aerr.Error())
			case cloud9.ErrCodeInternalServerErrorException:
				fmt.Println(cloud9.ErrCodeInternalServerErrorException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// DescribeEnvironments
//

func ExampleCloud9_DescribeEnvironments_shared00() {
	svc := cloud9.New(session.New())
	input := &cloud9.DescribeEnvironmentsInput{
		EnvironmentIds: []*string{
			aws.String("8d9967e2f0624182b74e7690ad69ebEX"),
			aws.String("349c86d4579e4e7298d500ff57a6b2EX"),
		},
	}

	result, err := svc.DescribeEnvironments(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloud9.ErrCodeBadRequestException:
				fmt.Println(cloud9.ErrCodeBadRequestException, aerr.Error())
			case cloud9.ErrCodeConflictException:
				fmt.Println(cloud9.ErrCodeConflictException, aerr.Error())
			case cloud9.ErrCodeNotFoundException:
				fmt.Println(cloud9.ErrCodeNotFoundException, aerr.Error())
			case cloud9.ErrCodeForbiddenException:
				fmt.Println(cloud9.ErrCodeForbiddenException, aerr.Error())
			case cloud9.ErrCodeTooManyRequestsException:
				fmt.Println(cloud9.ErrCodeTooManyRequestsException, aerr.Error())
			case cloud9.ErrCodeLimitExceededException:
				fmt.Println(cloud9.ErrCodeLimitExceededException, aerr.Error())
			case cloud9.ErrCodeInternalServerErrorException:
				fmt.Println(cloud9.ErrCodeInternalServerErrorException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// ListEnvironments
//

func ExampleCloud9_ListEnvironments_shared00() {
	svc := cloud9.New(session.New())
	input := &cloud9.ListEnvironmentsInput{}

	result, err := svc.ListEnvironments(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloud9.ErrCodeBadRequestException:
				fmt.Println(cloud9.ErrCodeBadRequestException, aerr.Error())
			case cloud9.ErrCodeConflictException:
				fmt.Println(cloud9.ErrCodeConflictException, aerr.Error())
			case cloud9.ErrCodeNotFoundException:
				fmt.Println(cloud9.ErrCodeNotFoundException, aerr.Error())
			case cloud9.ErrCodeForbiddenException:
				fmt.Println(cloud9.ErrCodeForbiddenException, aerr.Error())
			case cloud9.ErrCodeTooManyRequestsException:
				fmt.Println(cloud9.ErrCodeTooManyRequestsException, aerr.Error())
			case cloud9.ErrCodeLimitExceededException:
				fmt.Println(cloud9.ErrCodeLimitExceededException, aerr.Error())
			case cloud9.ErrCodeInternalServerErrorException:
				fmt.Println(cloud9.ErrCodeInternalServerErrorException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// UpdateEnvironment
//

func ExampleCloud9_UpdateEnvironment_shared00() {
	svc := cloud9.New(session.New())
	input := &cloud9.UpdateEnvironmentInput{
		Description:   aws.String("This is my changed demonstration environment."),
		EnvironmentId: aws.String("8d9967e2f0624182b74e7690ad69ebEX"),
		Name:          aws.String("my-changed-demo-environment"),
	}

	result, err := svc.UpdateEnvironment(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloud9.ErrCodeBadRequestException:
				fmt.Println(cloud9.ErrCodeBadRequestException, aerr.Error())
			case cloud9.ErrCodeConflictException:
				fmt.Println(cloud9.ErrCodeConflictException, aerr.Error())
			case cloud9.ErrCodeNotFoundException:
				fmt.Println(cloud9.ErrCodeNotFoundException, aerr.Error())
			case cloud9.ErrCodeForbiddenException:
				fmt.Println(cloud9.ErrCodeForbiddenException, aerr.Error())
			case cloud9.ErrCodeTooManyRequestsException:
				fmt.Println(cloud9.ErrCodeTooManyRequestsException, aerr.Error())
			case cloud9.ErrCodeLimitExceededException:
				fmt.Println(cloud9.ErrCodeLimitExceededException, aerr.Error())
			case cloud9.ErrCodeInternalServerErrorException:
				fmt.Println(cloud9.ErrCodeInternalServerErrorException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// UpdateEnvironmentMembership
//

func ExampleCloud9_UpdateEnvironmentMembership_shared00() {
	svc := cloud9.New(session.New())
	input := &cloud9.UpdateEnvironmentMembershipInput{
		EnvironmentId: aws.String("8d9967e2f0624182b74e7690ad69ebEX"),
		Permissions:   aws.String("read-only"),
		UserArn:       aws.String("arn:aws:iam::123456789012:user/AnotherDemoUser"),
	}

	result, err := svc.UpdateEnvironmentMembership(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloud9.ErrCodeBadRequestException:
				fmt.Println(cloud9.ErrCodeBadRequestException, aerr.Error())
			case cloud9.ErrCodeConflictException:
				fmt.Println(cloud9.ErrCodeConflictException, aerr.Error())
			case cloud9.ErrCodeNotFoundException:
				fmt.Println(cloud9.ErrCodeNotFoundException, aerr.Error())
			case cloud9.ErrCodeForbiddenException:
				fmt.Println(cloud9.ErrCodeForbiddenException, aerr.Error())
			case cloud9.ErrCodeTooManyRequestsException:
				fmt.Println(cloud9.ErrCodeTooManyRequestsException, aerr.Error())
			case cloud9.ErrCodeLimitExceededException:
				fmt.Println(cloud9.ErrCodeLimitExceededException, aerr.Error())
			case cloud9.ErrCodeInternalServerErrorException:
				fmt.Println(cloud9.ErrCodeInternalServerErrorException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
