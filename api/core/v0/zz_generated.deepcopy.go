//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v0

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventStore) DeepCopyInto(out *EventStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventStore.
func (in *EventStore) DeepCopy() *EventStore {
	if in == nil {
		return nil
	}
	out := new(EventStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventStoreAutoScaling) DeepCopyInto(out *EventStoreAutoScaling) {
	*out = *in
	if in.MinScale != nil {
		in, out := &in.MinScale, &out.MinScale
		*out = new(int)
		**out = **in
	}
	if in.MaxScale != nil {
		in, out := &in.MaxScale, &out.MaxScale
		*out = new(int)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(int)
		**out = **in
	}
	if in.TargetUtilizationPercentage != nil {
		in, out := &in.TargetUtilizationPercentage, &out.TargetUtilizationPercentage
		*out = new(int)
		**out = **in
	}
	if in.Parallelism != nil {
		in, out := &in.Parallelism, &out.Parallelism
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventStoreAutoScaling.
func (in *EventStoreAutoScaling) DeepCopy() *EventStoreAutoScaling {
	if in == nil {
		return nil
	}
	out := new(EventStoreAutoScaling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventStoreList) DeepCopyInto(out *EventStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventStoreList.
func (in *EventStoreList) DeepCopy() *EventStoreList {
	if in == nil {
		return nil
	}
	out := new(EventStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventStoreSpec) DeepCopyInto(out *EventStoreSpec) {
	*out = *in
	in.AutoScaling.DeepCopyInto(&out.AutoScaling)
	in.Resources.DeepCopyInto(&out.Resources)
	in.DispatcherResources.DeepCopyInto(&out.DispatcherResources)
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventStoreSpec.
func (in *EventStoreSpec) DeepCopy() *EventStoreSpec {
	if in == nil {
		return nil
	}
	out := new(EventStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventStoreStatus) DeepCopyInto(out *EventStoreStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventStoreStatus.
func (in *EventStoreStatus) DeepCopy() *EventStoreStatus {
	if in == nil {
		return nil
	}
	out := new(EventStoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gateway) DeepCopyInto(out *Gateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gateway.
func (in *Gateway) DeepCopy() *Gateway {
	if in == nil {
		return nil
	}
	out := new(Gateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Gateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayList) DeepCopyInto(out *GatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Gateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayList.
func (in *GatewayList) DeepCopy() *GatewayList {
	if in == nil {
		return nil
	}
	out := new(GatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewaySpec) DeepCopyInto(out *GatewaySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewaySpec.
func (in *GatewaySpec) DeepCopy() *GatewaySpec {
	if in == nil {
		return nil
	}
	out := new(GatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayStatus) DeepCopyInto(out *GatewayStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayStatus.
func (in *GatewayStatus) DeepCopy() *GatewayStatus {
	if in == nil {
		return nil
	}
	out := new(GatewayStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MQTTTrigger) DeepCopyInto(out *MQTTTrigger) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MQTTTrigger.
func (in *MQTTTrigger) DeepCopy() *MQTTTrigger {
	if in == nil {
		return nil
	}
	out := new(MQTTTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MQTTTrigger) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MQTTTriggerList) DeepCopyInto(out *MQTTTriggerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MQTTTrigger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MQTTTriggerList.
func (in *MQTTTriggerList) DeepCopy() *MQTTTriggerList {
	if in == nil {
		return nil
	}
	out := new(MQTTTriggerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MQTTTriggerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MQTTTriggerSpec) DeepCopyInto(out *MQTTTriggerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MQTTTriggerSpec.
func (in *MQTTTriggerSpec) DeepCopy() *MQTTTriggerSpec {
	if in == nil {
		return nil
	}
	out := new(MQTTTriggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MQTTTriggerStatus) DeepCopyInto(out *MQTTTriggerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MQTTTriggerStatus.
func (in *MQTTTriggerStatus) DeepCopy() *MQTTTriggerStatus {
	if in == nil {
		return nil
	}
	out := new(MQTTTriggerStatus)
	in.DeepCopyInto(out)
	return out
}
