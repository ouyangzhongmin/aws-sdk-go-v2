// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetDatabaseInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog in which the database resides. If none is provided,
	// the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The name of the database to retrieve. For Hive compatibility, this should
	// be all lowercase.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDatabaseInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDatabaseInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDatabaseInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetDatabaseOutput struct {
	_ struct{} `type:"structure"`

	// The definition of the specified database in the Data Catalog.
	Database *Database `type:"structure"`
}

// String returns the string representation
func (s GetDatabaseOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDatabase = "GetDatabase"

// GetDatabaseRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves the definition of a specified database.
//
//    // Example sending a request using GetDatabaseRequest.
//    req := client.GetDatabaseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetDatabase
func (c *Client) GetDatabaseRequest(input *GetDatabaseInput) GetDatabaseRequest {
	op := &aws.Operation{
		Name:       opGetDatabase,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDatabaseInput{}
	}

	req := c.newRequest(op, input, &GetDatabaseOutput{})

	return GetDatabaseRequest{Request: req, Input: input, Copy: c.GetDatabaseRequest}
}

// GetDatabaseRequest is the request type for the
// GetDatabase API operation.
type GetDatabaseRequest struct {
	*aws.Request
	Input *GetDatabaseInput
	Copy  func(*GetDatabaseInput) GetDatabaseRequest
}

// Send marshals and sends the GetDatabase API request.
func (r GetDatabaseRequest) Send(ctx context.Context) (*GetDatabaseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDatabaseResponse{
		GetDatabaseOutput: r.Request.Data.(*GetDatabaseOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDatabaseResponse is the response type for the
// GetDatabase API operation.
type GetDatabaseResponse struct {
	*GetDatabaseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDatabase request.
func (r *GetDatabaseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
