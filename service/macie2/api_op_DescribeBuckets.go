// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Specifies criteria for filtering, sorting, and paginating the results of
// a query for information about S3 buckets.
type DescribeBucketsInput struct {
	_ struct{} `type:"structure"`

	// Specifies, as a map, one or more attribute-based conditions that filter the
	// results of a query for information about S3 buckets.
	Criteria map[string]BucketCriteriaAdditionalProperties `locationName:"criteria" type:"map"`

	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	NextToken *string `locationName:"nextToken" type:"string"`

	// Specifies criteria for sorting the results of a query for information about
	// S3 buckets.
	SortCriteria *BucketSortCriteria `locationName:"sortCriteria" type:"structure"`
}

// String returns the string representation
func (s DescribeBucketsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeBucketsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Criteria != nil {
		v := s.Criteria

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "criteria", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SortCriteria != nil {
		v := s.SortCriteria

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "sortCriteria", v, metadata)
	}
	return nil
}

// Provides the results of a query that retrieved statistical data and other
// information about one or more S3 buckets that Amazon Macie monitors and analyzes.
type DescribeBucketsOutput struct {
	_ struct{} `type:"structure"`

	Buckets []BucketMetadata `locationName:"buckets" type:"list"`

	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeBucketsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeBucketsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Buckets != nil {
		v := s.Buckets

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "buckets", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeBuckets = "DescribeBuckets"

// DescribeBucketsRequest returns a request value for making API operation for
// Amazon Macie 2.
//
// Retrieves (queries) statistical data and other information about one or more
// S3 buckets that Amazon Macie monitors and analyzes.
//
//    // Example sending a request using DescribeBucketsRequest.
//    req := client.DescribeBucketsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie2-2020-01-01/DescribeBuckets
func (c *Client) DescribeBucketsRequest(input *DescribeBucketsInput) DescribeBucketsRequest {
	op := &aws.Operation{
		Name:       opDescribeBuckets,
		HTTPMethod: "POST",
		HTTPPath:   "/datasources/s3",
	}

	if input == nil {
		input = &DescribeBucketsInput{}
	}

	req := c.newRequest(op, input, &DescribeBucketsOutput{})

	return DescribeBucketsRequest{Request: req, Input: input, Copy: c.DescribeBucketsRequest}
}

// DescribeBucketsRequest is the request type for the
// DescribeBuckets API operation.
type DescribeBucketsRequest struct {
	*aws.Request
	Input *DescribeBucketsInput
	Copy  func(*DescribeBucketsInput) DescribeBucketsRequest
}

// Send marshals and sends the DescribeBuckets API request.
func (r DescribeBucketsRequest) Send(ctx context.Context) (*DescribeBucketsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeBucketsResponse{
		DescribeBucketsOutput: r.Request.Data.(*DescribeBucketsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeBucketsResponse is the response type for the
// DescribeBuckets API operation.
type DescribeBucketsResponse struct {
	*DescribeBucketsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeBuckets request.
func (r *DescribeBucketsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
