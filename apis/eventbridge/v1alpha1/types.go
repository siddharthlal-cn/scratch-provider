/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// Tag represent a user-provided metadata that can be associated with a
// SNS Topic. For more information about tagging,
// see Tagging SNS Topics (https://docs.aws.amazon.com/sns/latest/dg/sns-tags.html)
// in the SNS User Guide.
type Tag struct {

	// The key name that can be used to look up or retrieve the associated value.
	// For example, Department or Cost Center are common choices.
	Key string `json:"key"`

	// The value associated with this tag. For example, tags with a key name of
	// Department could have values such as Human Resources, Accounting, and Support.
	// Tags with a key name of Cost Center might have values that consist of the
	// number associated with the different cost centers in your company. Typically,
	// many resources have tags with the same key name but with different values.
	//
	// AWS always interprets the tag Value as a single string. If you need to store
	// an array, you can store comma-separated values in the string. However, you
	// must interpret the value in your code.
	// +optional
	Value *string `json:"value,omitempty"`
}

// EventBridgeRuleParameters are the configurable fields of a EventBridgeRule.
type EventBridgeRuleParameters struct {
	// Region is the region you'd like your SNSTopic to be created in.
	Region string `json:"region"`

	// Name refers to the name of the AWS SNS Topic
	// +immutable
	Name string `json:"name"`

	// The display name to use for a topic with SNS subscriptions.
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	// Setting this enables server side encryption at-rest to your topic.
	// The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK
	//
	// For more examples, see KeyId (https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters)
	// in the AWS Key Management Service API Reference.
	// +optional
	KMSMasterKeyID *string `json:"kmsMasterKeyId,omitempty"`

	// A description of the rule.
	Description *string `json:"description"`

	// The name or ARN of the event bus to associate with this rule. If you omit
	// this, the default event bus is used.
	// +optional
	EventBusName *string `json:"eventBusName,omitempty"`

	// The event pattern. For more information, see Events and Event Patterns (https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html)
	// in the Amazon EventBridge User Guide.
	// +optional
	EventPattern *string `json:"eventPattern,omitempty"`

	// The Amazon Resource Name (ARN) of the IAM role associated with the rule.
	// +optional
	RoleArn *string `json:"roleArn,omitempty"`

	// The scheduling expression. For example, "cron(0 20 * * ? *)" or "rate(5 minutes)".
	ScheduleExpression *string `json:"scheduleExpression"`

	// Indicates whether the rule is enabled or disabled.
	// +optional
	State *string `json:"state"`

	// Tags represetnt a list of user-provided metadata that can be associated with a
	// SNS Topic. For more information about tagging,
	// see Tagging SNS Topics (https://docs.aws.amazon.com/sns/latest/dg/sns-tags.html)
	// in the SNS User Guide.
	// +immutable
	// +optional
	Tags []Tag `json:"tags,omitempty"`
}

// EventBridgeRuleObservation are the observable fields of a EventBridgeRule.
type EventBridgeRuleObservation struct {
	// ARN is the Amazon Resource Name (ARN) specifying the SNS Topic.
	ARN string `json:"arn"`
}

// A EventBridgeRuleSpec defines the desired state of a EventBridgeRule.
type EventBridgeRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EventBridgeRuleParameters `json:"forProvider"`
}

// A EventBridgeRuleStatus represents the observed state of a EventBridgeRule.
type EventBridgeRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EventBridgeRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A EventBridgeRule is an example API type
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.bindingPhase"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.atProvider.state"
// +kubebuilder:printcolumn:name="CLASS",type="string",JSONPath=".spec.classRef.name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster
type EventBridgeRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EventBridgeRuleSpec   `json:"spec"`
	Status EventBridgeRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventBridgeRuleList contains a list of EventBridgeRule
type EventBridgeRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventBridgeRule `json:"items"`
}
