// Code generated by "go generate". Please don't change this file directly.

package codebuild

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// Project_CloudWatchLogsConfig AWS CloudFormation Resource (AWS::CodeBuild::Project.CloudWatchLogsConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-cloudwatchlogsconfig.html
type Project_CloudWatchLogsConfig struct {

	// GroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-cloudwatchlogsconfig.html#cfn-codebuild-project-cloudwatchlogsconfig-groupname
	GroupName *string `json:"GroupName,omitempty"`

	// Status AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-cloudwatchlogsconfig.html#cfn-codebuild-project-cloudwatchlogsconfig-status
	Status string `json:"Status"`

	// StreamName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-cloudwatchlogsconfig.html#cfn-codebuild-project-cloudwatchlogsconfig-streamname
	StreamName *string `json:"StreamName,omitempty"`

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
func (r *Project_CloudWatchLogsConfig) AWSCloudFormationType() string {
	return "AWS::CodeBuild::Project.CloudWatchLogsConfig"
}
