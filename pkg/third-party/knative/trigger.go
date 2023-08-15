package knative

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kEventing "knative.dev/eventing/pkg/apis/eventing/v1"
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

type TriggerParameters struct {
	TriggerName     string
	Namespace       string
	BrokerName      string
	SubscriberName  string
	OwnerReferences []v1.OwnerReference
	Attributes      map[string]string
	Labels          map[string]string
}

func NewTrigger(triggerParameters TriggerParameters) kEventing.Trigger {
	return kEventing.Trigger{
		TypeMeta: v1.TypeMeta{
			Kind:       "Trigger",
			APIVersion: "eventing.knative.dev/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:            triggerParameters.TriggerName,
			Namespace:       triggerParameters.Namespace,
			Labels:          triggerParameters.Labels,
			OwnerReferences: triggerParameters.OwnerReferences,
		},
		Spec: kEventing.TriggerSpec{
			Broker: triggerParameters.BrokerName,
			Filter: &kEventing.TriggerFilter{
				Attributes: triggerParameters.Attributes,
			},
			Subscriber: duckv1.Destination{
				Ref: &duckv1.KReference{
					Kind:       "Service",
					APIVersion: "serving.knative.dev/v1",
					Name:       triggerParameters.SubscriberName,
				},
				URI: &apis.URL{
					Path: "/api/v1/twin-events",
				},
			},
		},
	}
}