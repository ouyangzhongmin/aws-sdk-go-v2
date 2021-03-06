// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListServerNeighborsInput struct {
	_ struct{} `type:"structure"`

	// Configuration ID of the server for which neighbors are being listed.
	//
	// ConfigurationId is a required field
	ConfigurationId *string `locationName:"configurationId" type:"string" required:"true"`

	// Maximum number of results to return in a single page of output.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// List of configuration IDs to test for one-hop-away.
	NeighborConfigurationIds []string `locationName:"neighborConfigurationIds" type:"list"`

	// Token to retrieve the next set of results. For example, if you previously
	// specified 100 IDs for ListServerNeighborsRequest$neighborConfigurationIds
	// but set ListServerNeighborsRequest$maxResults to 10, you received a set of
	// 10 results along with a token. Use that token in this query to get the next
	// set of 10.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Flag to indicate if port and protocol information is needed as part of the
	// response.
	PortInformationNeeded *bool `locationName:"portInformationNeeded" type:"boolean"`
}

// String returns the string representation
func (s ListServerNeighborsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListServerNeighborsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListServerNeighborsInput"}

	if s.ConfigurationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListServerNeighborsOutput struct {
	_ struct{} `type:"structure"`

	// Count of distinct servers that are one hop away from the given server.
	KnownDependencyCount *int64 `locationName:"knownDependencyCount" type:"long"`

	// List of distinct servers that are one hop away from the given server.
	//
	// Neighbors is a required field
	Neighbors []NeighborConnectionDetail `locationName:"neighbors" type:"list" required:"true"`

	// Token to retrieve the next set of results. For example, if you specified
	// 100 IDs for ListServerNeighborsRequest$neighborConfigurationIds but set ListServerNeighborsRequest$maxResults
	// to 10, you received a set of 10 results along with this token. Use this token
	// in the next query to retrieve the next set of 10.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListServerNeighborsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListServerNeighbors = "ListServerNeighbors"

// ListServerNeighborsRequest returns a request value for making API operation for
// AWS Application Discovery Service.
//
// Retrieves a list of servers that are one network hop away from a specified
// server.
//
//    // Example sending a request using ListServerNeighborsRequest.
//    req := client.ListServerNeighborsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/ListServerNeighbors
func (c *Client) ListServerNeighborsRequest(input *ListServerNeighborsInput) ListServerNeighborsRequest {
	op := &aws.Operation{
		Name:       opListServerNeighbors,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListServerNeighborsInput{}
	}

	req := c.newRequest(op, input, &ListServerNeighborsOutput{})

	return ListServerNeighborsRequest{Request: req, Input: input, Copy: c.ListServerNeighborsRequest}
}

// ListServerNeighborsRequest is the request type for the
// ListServerNeighbors API operation.
type ListServerNeighborsRequest struct {
	*aws.Request
	Input *ListServerNeighborsInput
	Copy  func(*ListServerNeighborsInput) ListServerNeighborsRequest
}

// Send marshals and sends the ListServerNeighbors API request.
func (r ListServerNeighborsRequest) Send(ctx context.Context) (*ListServerNeighborsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListServerNeighborsResponse{
		ListServerNeighborsOutput: r.Request.Data.(*ListServerNeighborsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListServerNeighborsResponse is the response type for the
// ListServerNeighbors API operation.
type ListServerNeighborsResponse struct {
	*ListServerNeighborsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListServerNeighbors request.
func (r *ListServerNeighborsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
