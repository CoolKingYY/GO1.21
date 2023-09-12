// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudtraildataiface provides an interface to enable mocking the AWS CloudTrail Data Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudtraildataiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudtraildata"
)

// CloudTrailDataAPI provides an interface to enable mocking the
// cloudtraildata.CloudTrailData service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS CloudTrail Data Service.
//	func myFunc(svc cloudtraildataiface.CloudTrailDataAPI) bool {
//	    // Make svc.PutAuditEvents request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := cloudtraildata.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockCloudTrailDataClient struct {
//	    cloudtraildataiface.CloudTrailDataAPI
//	}
//	func (m *mockCloudTrailDataClient) PutAuditEvents(input *cloudtraildata.PutAuditEventsInput) (*cloudtraildata.PutAuditEventsOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockCloudTrailDataClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type CloudTrailDataAPI interface {
	PutAuditEvents(*cloudtraildata.PutAuditEventsInput) (*cloudtraildata.PutAuditEventsOutput, error)
	PutAuditEventsWithContext(aws.Context, *cloudtraildata.PutAuditEventsInput, ...request.Option) (*cloudtraildata.PutAuditEventsOutput, error)
	PutAuditEventsRequest(*cloudtraildata.PutAuditEventsInput) (*request.Request, *cloudtraildata.PutAuditEventsOutput)
}

var _ CloudTrailDataAPI = (*cloudtraildata.CloudTrailData)(nil)
