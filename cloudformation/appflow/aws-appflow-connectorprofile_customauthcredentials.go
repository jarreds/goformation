// Code generated by "go generate". Please don't change this file directly.

package appflow

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// ConnectorProfile_CustomAuthCredentials AWS CloudFormation Resource (AWS::AppFlow::ConnectorProfile.CustomAuthCredentials)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customauthcredentials.html
type ConnectorProfile_CustomAuthCredentials struct {

	// CredentialsMap AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customauthcredentials.html#cfn-appflow-connectorprofile-customauthcredentials-credentialsmap
	CredentialsMap *ConnectorProfile_CredentialsMap `json:"CredentialsMap,omitempty"`

	// CustomAuthenticationType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customauthcredentials.html#cfn-appflow-connectorprofile-customauthcredentials-customauthenticationtype
	CustomAuthenticationType string `json:"CustomAuthenticationType"`

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
func (r *ConnectorProfile_CustomAuthCredentials) AWSCloudFormationType() string {
	return "AWS::AppFlow::ConnectorProfile.CustomAuthCredentials"
}
