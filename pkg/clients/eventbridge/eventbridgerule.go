package eventbridge

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/siddharthlal-cn/scratch-provider/apis/eventbridge/v1alpha1"
)

// RuleClient is the external client used for AWS Event Bridge
type RuleClient interface {
	PutRuleRequest(*eventbridge.PutRuleInput) eventbridge.PutRuleRequest
	DeleteRuleRequest(*eventbridge.DeleteRuleInput) eventbridge.DeleteRuleRequest
	DescribeRuleRequest(*eventbridge.DescribeRuleInput) eventbridge.DescribeRuleRequest
}

// NewRuleClient returns a new client using AWS credentials as JSON encoded data.
func NewRuleClient(cfg aws.Config) RuleClient {
	return eventbridge.New(cfg)
}

// GeneratePutRuleInput prepares input for PutRuleInput
func GeneratePutRuleInput(p *v1alpha1.EventBridgeRuleParameters) *eventbridge.PutRuleInput {
	input := &eventbridge.PutRuleInput{}

	input.Name = aws.String(p.Name)

	if p.ScheduleExpression != nil {
		input.ScheduleExpression = p.ScheduleExpression
	}

	if len(p.Tags) != 0 {
		input.Tags = make([]eventbridge.Tag, len(p.Tags))
		for i, val := range p.Tags {
			input.Tags[i] = eventbridge.Tag{
				Key:   aws.String(val.Key),
				Value: val.Value,
			}
		}
	}

	return input
}
