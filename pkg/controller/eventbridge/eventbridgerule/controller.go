package eventbridgerule

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	ebv2 "github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	awscommon "github.com/siddharthlal-cn/scratch-provider/pkg/clients"

	"github.com/siddharthlal-cn/scratch-provider/apis/eventbridge/v1alpha1"
	eventbridgeclient "github.com/siddharthlal-cn/scratch-provider/pkg/clients/eventbridge"
)

const (
	errUnexpectedObject = "the managed resource is not a EventBridgeRule resource"
	errCreate           = "failed to create the Event Bridge Rule"
)

// SetupEventBridgeRule adds a controller that reconciles EventBridgeRule.
func SetupEventBridgeRule(mgr ctrl.Manager, l logging.Logger) error {
	name := managed.ControllerName(v1alpha1.EventBridgeRuleGroupKind)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		For(&v1alpha1.EventBridgeRule{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(v1alpha1.EventBridgeRuleGroupVersionKind),
			managed.WithExternalConnecter(&connector{kube: mgr.GetClient(), newClientFn: eventbridgeclient.NewRuleClient}),
			managed.WithReferenceResolver(managed.NewAPISimpleReferenceResolver(mgr.GetClient())),
			managed.WithInitializers(managed.NewDefaultProviderConfig(mgr.GetClient())),
			managed.WithConnectionPublishers(),
			managed.WithLogger(l.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

type connector struct {
	kube        client.Client
	newClientFn func(config aws.Config) eventbridgeclient.RuleClient
}

func (c *connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*v1alpha1.EventBridgeRule)
	if !ok {
		return nil, errors.New(errUnexpectedObject)
	}
	cfg, err := awscommon.GetConfig(ctx, c.kube, mg, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, err
	}

	return &external{client: ebv2.New(*cfg), kube: c.kube}, nil
}

type external struct {
	client eventbridgeclient.RuleClient
	kube   client.Client
}

func (e *external) Observe(ctx context.Context, mgd resource.Managed) (managed.ExternalObservation, error) {
	return managed.ExternalObservation{}, nil
}

func (e *external) Create(ctx context.Context, mgd resource.Managed) (managed.ExternalCreation, error) {

	cr, ok := mgd.(*v1alpha1.EventBridgeRule)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	// cr.SetConditions()

	resp, err := e.client.PutRuleRequest(eventbridgeclient.GeneratePutRuleInput(&cr.Spec.ForProvider)).Send(ctx)

	if err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, errCreate)
	}

	meta.SetExternalName(cr, aws.StringValue(resp.PutRuleOutput.RuleArn))
	return managed.ExternalCreation{ExternalNameAssigned: true}, nil
}

func (e *external) Update(ctx context.Context, mgd resource.Managed) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, nil
}

func (e *external) Delete(ctx context.Context, mgd resource.Managed) error {
	// cr, ok := mgd.(*v1alpha1.EventBridgeRule)
	// if !ok {
	// 	return errors.New(errUnexpectedObject)
	// }

	// cr.Status.SetConditions(xpv1.Deleting())

	// _, err := e.client.DeleteRuleRequest(&ebv2.DeleteRuleInput{
	// 	Name: aws.String(meta.GetExternalName(cr)),
	// }).Send(ctx)

	// return errors.Wrap(resource.Ignore(ebv2.Found, err), errDelete)
	return nil
}
