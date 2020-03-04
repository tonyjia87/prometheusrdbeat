// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeStackSetOperationInput
type DescribeStackSetOperationInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the stack set operation.
	//
	// OperationId is a required field
	OperationId *string `min:"1" type:"string" required:"true"`

	// The name or the unique stack ID of the stack set for the stack operation.
	//
	// StackSetName is a required field
	StackSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeStackSetOperationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStackSetOperationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStackSetOperationInput"}

	if s.OperationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OperationId"))
	}
	if s.OperationId != nil && len(*s.OperationId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OperationId", 1))
	}

	if s.StackSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeStackSetOperationOutput
type DescribeStackSetOperationOutput struct {
	_ struct{} `type:"structure"`

	// The specified stack set operation.
	StackSetOperation *StackSetOperation `type:"structure"`
}

// String returns the string representation
func (s DescribeStackSetOperationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeStackSetOperation = "DescribeStackSetOperation"

// DescribeStackSetOperationRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns the description of the specified stack set operation.
//
//    // Example sending a request using DescribeStackSetOperationRequest.
//    req := client.DescribeStackSetOperationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeStackSetOperation
func (c *Client) DescribeStackSetOperationRequest(input *DescribeStackSetOperationInput) DescribeStackSetOperationRequest {
	op := &aws.Operation{
		Name:       opDescribeStackSetOperation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeStackSetOperationInput{}
	}

	req := c.newRequest(op, input, &DescribeStackSetOperationOutput{})
	return DescribeStackSetOperationRequest{Request: req, Input: input, Copy: c.DescribeStackSetOperationRequest}
}

// DescribeStackSetOperationRequest is the request type for the
// DescribeStackSetOperation API operation.
type DescribeStackSetOperationRequest struct {
	*aws.Request
	Input *DescribeStackSetOperationInput
	Copy  func(*DescribeStackSetOperationInput) DescribeStackSetOperationRequest
}

// Send marshals and sends the DescribeStackSetOperation API request.
func (r DescribeStackSetOperationRequest) Send(ctx context.Context) (*DescribeStackSetOperationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStackSetOperationResponse{
		DescribeStackSetOperationOutput: r.Request.Data.(*DescribeStackSetOperationOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeStackSetOperationResponse is the response type for the
// DescribeStackSetOperation API operation.
type DescribeStackSetOperationResponse struct {
	*DescribeStackSetOperationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStackSetOperation request.
func (r *DescribeStackSetOperationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
