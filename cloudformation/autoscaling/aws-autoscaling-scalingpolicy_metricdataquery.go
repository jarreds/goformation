// Code generated by "go generate". Please don't change this file directly.

package autoscaling

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// ScalingPolicy_MetricDataQuery AWS CloudFormation Resource (AWS::AutoScaling::ScalingPolicy.MetricDataQuery)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-metricdataquery.html
type ScalingPolicy_MetricDataQuery struct {

	// Expression AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-metricdataquery.html#cfn-autoscaling-scalingpolicy-metricdataquery-expression
	Expression *string `json:"Expression,omitempty"`

	// Id AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-metricdataquery.html#cfn-autoscaling-scalingpolicy-metricdataquery-id
	Id string `json:"Id"`

	// Label AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-metricdataquery.html#cfn-autoscaling-scalingpolicy-metricdataquery-label
	Label *string `json:"Label,omitempty"`

	// MetricStat AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-metricdataquery.html#cfn-autoscaling-scalingpolicy-metricdataquery-metricstat
	MetricStat *ScalingPolicy_MetricStat `json:"MetricStat,omitempty"`

	// ReturnData AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-metricdataquery.html#cfn-autoscaling-scalingpolicy-metricdataquery-returndata
	ReturnData *bool `json:"ReturnData,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *ScalingPolicy_MetricDataQuery) AWSCloudFormationType() string {
	return "AWS::AutoScaling::ScalingPolicy.MetricDataQuery"
}
