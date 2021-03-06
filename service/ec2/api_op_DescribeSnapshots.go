// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeSnapshotsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The filters.
	//
	//    * description - A description of the snapshot.
	//
	//    * encrypted - Indicates whether the snapshot is encrypted (true | false)
	//
	//    * owner-alias - Value from an Amazon-maintained list (amazon | self |
	//    all | aws-marketplace | microsoft) of snapshot owners. Not to be confused
	//    with the user-configured AWS account alias, which is set from the IAM
	//    console.
	//
	//    * owner-id - The ID of the AWS account that owns the snapshot.
	//
	//    * progress - The progress of the snapshot, as a percentage (for example,
	//    80%).
	//
	//    * snapshot-id - The snapshot ID.
	//
	//    * start-time - The time stamp when the snapshot was initiated.
	//
	//    * status - The status of the snapshot (pending | completed | error).
	//
	//    * tag:<key> - The key/value combination of a tag assigned to the resource.
	//    Use the tag key in the filter name and the tag value as the filter value.
	//    For example, to find all resources that have a tag with the key Owner
	//    and the value TeamA, specify tag:Owner for the filter name and TeamA for
	//    the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. Use this filter
	//    to find all resources assigned a tag with a specific key, regardless of
	//    the tag value.
	//
	//    * volume-id - The ID of the volume the snapshot is for.
	//
	//    * volume-size - The size of the volume, in GiB.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of snapshot results returned by DescribeSnapshots in paginated
	// output. When this parameter is used, DescribeSnapshots only returns MaxResults
	// results in a single page along with a NextToken response element. The remaining
	// results of the initial request can be seen by sending another DescribeSnapshots
	// request with the returned NextToken value. This value can be between 5 and
	// 1000; if MaxResults is given a value larger than 1000, only 1000 results
	// are returned. If this parameter is not used, then DescribeSnapshots returns
	// all results. You cannot specify this parameter and the snapshot IDs parameter
	// in the same request.
	MaxResults *int64 `type:"integer"`

	// The NextToken value returned from a previous paginated DescribeSnapshots
	// request where MaxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the NextToken value. This value is null when there are no more results
	// to return.
	NextToken *string `type:"string"`

	// Describes the snapshots owned by these owners.
	OwnerIds []string `locationName:"Owner" locationNameList:"Owner" type:"list"`

	// The IDs of the AWS accounts that can create volumes from the snapshot.
	RestorableByUserIds []string `locationName:"RestorableBy" type:"list"`

	// The snapshot IDs.
	//
	// Default: Describes the snapshots for which you have create volume permissions.
	SnapshotIds []string `locationName:"SnapshotId" locationNameList:"SnapshotId" type:"list"`
}

// String returns the string representation
func (s DescribeSnapshotsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeSnapshotsOutput struct {
	_ struct{} `type:"structure"`

	// The NextToken value to include in a future DescribeSnapshots request. When
	// the results of a DescribeSnapshots request exceed MaxResults, this value
	// can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the snapshots.
	Snapshots []Snapshot `locationName:"snapshotSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeSnapshotsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeSnapshots = "DescribeSnapshots"

// DescribeSnapshotsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified EBS snapshots available to you or all of the EBS
// snapshots available to you.
//
// The snapshots available to you include public snapshots, private snapshots
// that you own, and private snapshots owned by other AWS accounts for which
// you have explicit create volume permissions.
//
// The create volume permissions fall into the following categories:
//
//    * public: The owner of the snapshot granted create volume permissions
//    for the snapshot to the all group. All AWS accounts have create volume
//    permissions for these snapshots.
//
//    * explicit: The owner of the snapshot granted create volume permissions
//    to a specific AWS account.
//
//    * implicit: An AWS account has implicit create volume permissions for
//    all snapshots it owns.
//
// The list of snapshots returned can be filtered by specifying snapshot IDs,
// snapshot owners, or AWS accounts with create volume permissions. If no options
// are specified, Amazon EC2 returns all snapshots for which you have create
// volume permissions.
//
// If you specify one or more snapshot IDs, only snapshots that have the specified
// IDs are returned. If you specify an invalid snapshot ID, an error is returned.
// If you specify a snapshot ID for which you do not have access, it is not
// included in the returned results.
//
// If you specify one or more snapshot owners using the OwnerIds option, only
// snapshots from the specified owners and for which you have access are returned.
// The results can include the AWS account IDs of the specified owners, amazon
// for snapshots owned by Amazon, or self for snapshots that you own.
//
// If you specify a list of restorable users, only snapshots with create snapshot
// permissions for those users are returned. You can specify AWS account IDs
// (if you own the snapshots), self for snapshots for which you own or have
// explicit permissions, or all for public snapshots.
//
// If you are describing a long list of snapshots, you can paginate the output
// to make the list more manageable. The MaxResults parameter sets the maximum
// number of results returned in a single page. If the list of results exceeds
// your MaxResults value, then that number of results is returned along with
// a NextToken value that can be passed to a subsequent DescribeSnapshots request
// to retrieve the remaining results.
//
// To get the state of fast snapshot restores for a snapshot, use DescribeFastSnapshotRestores.
//
// For more information about EBS snapshots, see Amazon EBS Snapshots (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSSnapshots.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DescribeSnapshotsRequest.
//    req := client.DescribeSnapshotsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeSnapshots
func (c *Client) DescribeSnapshotsRequest(input *DescribeSnapshotsInput) DescribeSnapshotsRequest {
	op := &aws.Operation{
		Name:       opDescribeSnapshots,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeSnapshotsInput{}
	}

	req := c.newRequest(op, input, &DescribeSnapshotsOutput{})

	return DescribeSnapshotsRequest{Request: req, Input: input, Copy: c.DescribeSnapshotsRequest}
}

// DescribeSnapshotsRequest is the request type for the
// DescribeSnapshots API operation.
type DescribeSnapshotsRequest struct {
	*aws.Request
	Input *DescribeSnapshotsInput
	Copy  func(*DescribeSnapshotsInput) DescribeSnapshotsRequest
}

// Send marshals and sends the DescribeSnapshots API request.
func (r DescribeSnapshotsRequest) Send(ctx context.Context) (*DescribeSnapshotsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSnapshotsResponse{
		DescribeSnapshotsOutput: r.Request.Data.(*DescribeSnapshotsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeSnapshotsRequestPaginator returns a paginator for DescribeSnapshots.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeSnapshotsRequest(input)
//   p := ec2.NewDescribeSnapshotsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeSnapshotsPaginator(req DescribeSnapshotsRequest) DescribeSnapshotsPaginator {
	return DescribeSnapshotsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeSnapshotsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeSnapshotsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeSnapshotsPaginator struct {
	aws.Pager
}

func (p *DescribeSnapshotsPaginator) CurrentPage() *DescribeSnapshotsOutput {
	return p.Pager.CurrentPage().(*DescribeSnapshotsOutput)
}

// DescribeSnapshotsResponse is the response type for the
// DescribeSnapshots API operation.
type DescribeSnapshotsResponse struct {
	*DescribeSnapshotsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSnapshots request.
func (r *DescribeSnapshotsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
