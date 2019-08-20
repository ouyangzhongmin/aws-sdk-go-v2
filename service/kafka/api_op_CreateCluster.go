// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Creates a cluster.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/CreateClusterRequest
type CreateClusterInput struct {
	_ struct{} `type:"structure"`

	// Information about the broker nodes in the cluster.
	//
	// BrokerNodeGroupInfo is a required field
	BrokerNodeGroupInfo *BrokerNodeGroupInfo `locationName:"brokerNodeGroupInfo" type:"structure" required:"true"`

	// Includes all client authentication related information.
	ClientAuthentication *Authentication `locationName:"clientAuthentication" type:"structure"`

	// The name of the cluster.
	//
	// ClusterName is a required field
	ClusterName *string `locationName:"clusterName" min:"1" type:"string" required:"true"`

	// Represents the configuration that you want MSK to use for the brokers in
	// a cluster.
	ConfigurationInfo *ConfigurationInfo `locationName:"configurationInfo" type:"structure"`

	// Includes all encryption-related information.
	EncryptionInfo *EncryptionInfo `locationName:"encryptionInfo" type:"structure"`

	// Specifies the level of monitoring for the MSK cluster. The possible values
	// are DEFAULT, PER_BROKER, and PER_TOPIC_PER_BROKER.
	EnhancedMonitoring EnhancedMonitoring `locationName:"enhancedMonitoring" type:"string" enum:"true"`

	// The version of Apache Kafka.
	//
	// KafkaVersion is a required field
	KafkaVersion *string `locationName:"kafkaVersion" min:"1" type:"string" required:"true"`

	// The number of Kafka broker nodes in the Amazon MSK cluster.
	//
	// NumberOfBrokerNodes is a required field
	NumberOfBrokerNodes *int64 `locationName:"numberOfBrokerNodes" min:"1" type:"integer" required:"true"`

	// Create tags when creating the cluster.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateClusterInput"}

	if s.BrokerNodeGroupInfo == nil {
		invalidParams.Add(aws.NewErrParamRequired("BrokerNodeGroupInfo"))
	}

	if s.ClusterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterName"))
	}
	if s.ClusterName != nil && len(*s.ClusterName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClusterName", 1))
	}

	if s.KafkaVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("KafkaVersion"))
	}
	if s.KafkaVersion != nil && len(*s.KafkaVersion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KafkaVersion", 1))
	}

	if s.NumberOfBrokerNodes == nil {
		invalidParams.Add(aws.NewErrParamRequired("NumberOfBrokerNodes"))
	}
	if s.NumberOfBrokerNodes != nil && *s.NumberOfBrokerNodes < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("NumberOfBrokerNodes", 1))
	}
	if s.BrokerNodeGroupInfo != nil {
		if err := s.BrokerNodeGroupInfo.Validate(); err != nil {
			invalidParams.AddNested("BrokerNodeGroupInfo", err.(aws.ErrInvalidParams))
		}
	}
	if s.ConfigurationInfo != nil {
		if err := s.ConfigurationInfo.Validate(); err != nil {
			invalidParams.AddNested("ConfigurationInfo", err.(aws.ErrInvalidParams))
		}
	}
	if s.EncryptionInfo != nil {
		if err := s.EncryptionInfo.Validate(); err != nil {
			invalidParams.AddNested("EncryptionInfo", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateClusterInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.BrokerNodeGroupInfo != nil {
		v := s.BrokerNodeGroupInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "brokerNodeGroupInfo", v, metadata)
	}
	if s.ClientAuthentication != nil {
		v := s.ClientAuthentication

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "clientAuthentication", v, metadata)
	}
	if s.ClusterName != nil {
		v := *s.ClusterName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clusterName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConfigurationInfo != nil {
		v := s.ConfigurationInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "configurationInfo", v, metadata)
	}
	if s.EncryptionInfo != nil {
		v := s.EncryptionInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "encryptionInfo", v, metadata)
	}
	if len(s.EnhancedMonitoring) > 0 {
		v := s.EnhancedMonitoring

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enhancedMonitoring", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.KafkaVersion != nil {
		v := *s.KafkaVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "kafkaVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NumberOfBrokerNodes != nil {
		v := *s.NumberOfBrokerNodes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "numberOfBrokerNodes", protocol.Int64Value(v), metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

// Returns information about the created cluster.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/CreateClusterResponse
type CreateClusterOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string `locationName:"clusterArn" type:"string"`

	// The name of the MSK cluster.
	ClusterName *string `locationName:"clusterName" type:"string"`

	// The state of the cluster. The possible states are CREATING, ACTIVE, and FAILED.
	State ClusterState `locationName:"state" type:"string" enum:"true"`
}

// String returns the string representation
func (s CreateClusterOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateClusterOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClusterArn != nil {
		v := *s.ClusterArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clusterArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ClusterName != nil {
		v := *s.ClusterName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clusterName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.State) > 0 {
		v := s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "state", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opCreateCluster = "CreateCluster"

// CreateClusterRequest returns a request value for making API operation for
// Managed Streaming for Kafka.
//
// Creates a new MSK cluster.
//
//    // Example sending a request using CreateClusterRequest.
//    req := client.CreateClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/CreateCluster
func (c *Client) CreateClusterRequest(input *CreateClusterInput) CreateClusterRequest {
	op := &aws.Operation{
		Name:       opCreateCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/clusters",
	}

	if input == nil {
		input = &CreateClusterInput{}
	}

	req := c.newRequest(op, input, &CreateClusterOutput{})
	return CreateClusterRequest{Request: req, Input: input, Copy: c.CreateClusterRequest}
}

// CreateClusterRequest is the request type for the
// CreateCluster API operation.
type CreateClusterRequest struct {
	*aws.Request
	Input *CreateClusterInput
	Copy  func(*CreateClusterInput) CreateClusterRequest
}

// Send marshals and sends the CreateCluster API request.
func (r CreateClusterRequest) Send(ctx context.Context) (*CreateClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateClusterResponse{
		CreateClusterOutput: r.Request.Data.(*CreateClusterOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateClusterResponse is the response type for the
// CreateCluster API operation.
type CreateClusterResponse struct {
	*CreateClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCluster request.
func (r *CreateClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
