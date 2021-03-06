// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateVariableInput struct {
	_ struct{} `type:"structure"`

	// The source of the data.
	//
	// DataSource is a required field
	DataSource DataSource `locationName:"dataSource" type:"string" required:"true" enum:"true"`

	// The data type.
	//
	// DataType is a required field
	DataType DataType `locationName:"dataType" type:"string" required:"true" enum:"true"`

	// The default value for the variable when no value is received.
	//
	// DefaultValue is a required field
	DefaultValue *string `locationName:"defaultValue" type:"string" required:"true"`

	// The description.
	Description *string `locationName:"description" type:"string"`

	// The name of the variable.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// The variable type.
	VariableType *string `locationName:"variableType" type:"string"`
}

// String returns the string representation
func (s CreateVariableInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVariableInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVariableInput"}
	if len(s.DataSource) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("DataSource"))
	}
	if len(s.DataType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("DataType"))
	}

	if s.DefaultValue == nil {
		invalidParams.Add(aws.NewErrParamRequired("DefaultValue"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateVariableOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateVariableOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateVariable = "CreateVariable"

// CreateVariableRequest returns a request value for making API operation for
// Amazon Fraud Detector.
//
// Creates a variable.
//
//    // Example sending a request using CreateVariableRequest.
//    req := client.CreateVariableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/frauddetector-2019-11-15/CreateVariable
func (c *Client) CreateVariableRequest(input *CreateVariableInput) CreateVariableRequest {
	op := &aws.Operation{
		Name:       opCreateVariable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateVariableInput{}
	}

	req := c.newRequest(op, input, &CreateVariableOutput{})

	return CreateVariableRequest{Request: req, Input: input, Copy: c.CreateVariableRequest}
}

// CreateVariableRequest is the request type for the
// CreateVariable API operation.
type CreateVariableRequest struct {
	*aws.Request
	Input *CreateVariableInput
	Copy  func(*CreateVariableInput) CreateVariableRequest
}

// Send marshals and sends the CreateVariable API request.
func (r CreateVariableRequest) Send(ctx context.Context) (*CreateVariableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVariableResponse{
		CreateVariableOutput: r.Request.Data.(*CreateVariableOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVariableResponse is the response type for the
// CreateVariable API operation.
type CreateVariableResponse struct {
	*CreateVariableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVariable request.
func (r *CreateVariableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
