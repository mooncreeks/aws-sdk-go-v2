// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package dynamodbstreamsiface provides an interface to enable mocking the Amazon DynamoDB Streams service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package dynamodbstreamsiface

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
)

// DynamoDBStreamsAPI provides an interface to enable mocking the
// dynamodbstreams.DynamoDBStreams service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon DynamoDB Streams.
//    func myFunc(svc dynamodbstreamsiface.DynamoDBStreamsAPI) bool {
//        // Make svc.DescribeStream request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := dynamodbstreams.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockDynamoDBStreamsClient struct {
//        dynamodbstreamsiface.DynamoDBStreamsAPI
//    }
//    func (m *mockDynamoDBStreamsClient) DescribeStream(input *dynamodbstreams.DescribeStreamInput) (*dynamodbstreams.DescribeStreamOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockDynamoDBStreamsClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type DynamoDBStreamsAPI interface {
	DescribeStreamRequest(*dynamodbstreams.DescribeStreamInput) dynamodbstreams.DescribeStreamRequest

	GetRecordsRequest(*dynamodbstreams.GetRecordsInput) dynamodbstreams.GetRecordsRequest

	GetShardIteratorRequest(*dynamodbstreams.GetShardIteratorInput) dynamodbstreams.GetShardIteratorRequest

	ListStreamsRequest(*dynamodbstreams.ListStreamsInput) dynamodbstreams.ListStreamsRequest
}

var _ DynamoDBStreamsAPI = (*dynamodbstreams.DynamoDBStreams)(nil)
