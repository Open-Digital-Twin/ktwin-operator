//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v0

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommandRequest) DeepCopyInto(out *CommandRequest) {
	*out = *in
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(TwinSchema)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommandRequest.
func (in *CommandRequest) DeepCopy() *CommandRequest {
	if in == nil {
		return nil
	}
	out := new(CommandRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommandResponse) DeepCopyInto(out *CommandResponse) {
	*out = *in
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(TwinSchema)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommandResponse.
func (in *CommandResponse) DeepCopy() *CommandResponse {
	if in == nil {
		return nil
	}
	out := new(CommandResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinCommand) DeepCopyInto(out *TwinCommand) {
	*out = *in
	in.Request.DeepCopyInto(&out.Request)
	in.Response.DeepCopyInto(&out.Response)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinCommand.
func (in *TwinCommand) DeepCopy() *TwinCommand {
	if in == nil {
		return nil
	}
	out := new(TwinCommand)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinComplexType) DeepCopyInto(out *TwinComplexType) {
	*out = *in
	if in.Fields != nil {
		in, out := &in.Fields, &out.Fields
		*out = make([]TwinComplexTypeFields, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinComplexType.
func (in *TwinComplexType) DeepCopy() *TwinComplexType {
	if in == nil {
		return nil
	}
	out := new(TwinComplexType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinComplexTypeFields) DeepCopyInto(out *TwinComplexTypeFields) {
	*out = *in
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(TwinComplexTypeSchema)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinComplexTypeFields.
func (in *TwinComplexTypeFields) DeepCopy() *TwinComplexTypeFields {
	if in == nil {
		return nil
	}
	out := new(TwinComplexTypeFields)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinComplexTypeSchema) DeepCopyInto(out *TwinComplexTypeSchema) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinComplexTypeSchema.
func (in *TwinComplexTypeSchema) DeepCopy() *TwinComplexTypeSchema {
	if in == nil {
		return nil
	}
	out := new(TwinComplexTypeSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinEnumSchema) DeepCopyInto(out *TwinEnumSchema) {
	*out = *in
	if in.EnumValues != nil {
		in, out := &in.EnumValues, &out.EnumValues
		*out = make([]TwinEnumSchemaValues, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinEnumSchema.
func (in *TwinEnumSchema) DeepCopy() *TwinEnumSchema {
	if in == nil {
		return nil
	}
	out := new(TwinEnumSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinEnumSchemaValues) DeepCopyInto(out *TwinEnumSchemaValues) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinEnumSchemaValues.
func (in *TwinEnumSchemaValues) DeepCopy() *TwinEnumSchemaValues {
	if in == nil {
		return nil
	}
	out := new(TwinEnumSchemaValues)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInstance) DeepCopyInto(out *TwinInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInstance.
func (in *TwinInstance) DeepCopy() *TwinInstance {
	if in == nil {
		return nil
	}
	out := new(TwinInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TwinInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInstanceAmqpEndpointSettings) DeepCopyInto(out *TwinInstanceAmqpEndpointSettings) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInstanceAmqpEndpointSettings.
func (in *TwinInstanceAmqpEndpointSettings) DeepCopy() *TwinInstanceAmqpEndpointSettings {
	if in == nil {
		return nil
	}
	out := new(TwinInstanceAmqpEndpointSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInstanceDataSpec) DeepCopyInto(out *TwinInstanceDataSpec) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]TwinInstancePropertyData, len(*in))
		copy(*out, *in)
	}
	if in.Telemetries != nil {
		in, out := &in.Telemetries, &out.Telemetries
		*out = make([]TwinInstanceTelemetryData, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInstanceDataSpec.
func (in *TwinInstanceDataSpec) DeepCopy() *TwinInstanceDataSpec {
	if in == nil {
		return nil
	}
	out := new(TwinInstanceDataSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInstanceEndpointSettings) DeepCopyInto(out *TwinInstanceEndpointSettings) {
	*out = *in
	if in.HttpEndpoint != nil {
		in, out := &in.HttpEndpoint, &out.HttpEndpoint
		*out = new(TwinInstanceHttpEndpointSettings)
		**out = **in
	}
	if in.MqttEndpoint != nil {
		in, out := &in.MqttEndpoint, &out.MqttEndpoint
		*out = new(TwinInstanceMqttEndpointSettings)
		**out = **in
	}
	if in.AmqpEndpoint != nil {
		in, out := &in.AmqpEndpoint, &out.AmqpEndpoint
		*out = new(TwinInstanceAmqpEndpointSettings)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInstanceEndpointSettings.
func (in *TwinInstanceEndpointSettings) DeepCopy() *TwinInstanceEndpointSettings {
	if in == nil {
		return nil
	}
	out := new(TwinInstanceEndpointSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInstanceEvents) DeepCopyInto(out *TwinInstanceEvents) {
	*out = *in
	in.Filters.DeepCopyInto(&out.Filters)
	out.Sink = in.Sink
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInstanceEvents.
func (in *TwinInstanceEvents) DeepCopy() *TwinInstanceEvents {
	if in == nil {
		return nil
	}
	out := new(TwinInstanceEvents)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInstanceEventsFilters) DeepCopyInto(out *TwinInstanceEventsFilters) {
	*out = *in
	if in.Exact != nil {
		in, out := &in.Exact, &out.Exact
		*out = make(TwinInstanceEventsFiltersAttributes, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInstanceEventsFilters.
func (in *TwinInstanceEventsFilters) DeepCopy() *TwinInstanceEventsFilters {
	if in == nil {
		return nil
	}
	out := new(TwinInstanceEventsFilters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in TwinInstanceEventsFiltersAttributes) DeepCopyInto(out *TwinInstanceEventsFiltersAttributes) {
	{
		in := &in
		*out = make(TwinInstanceEventsFiltersAttributes, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInstanceEventsFiltersAttributes.
func (in TwinInstanceEventsFiltersAttributes) DeepCopy() TwinInstanceEventsFiltersAttributes {
	if in == nil {
		return nil
	}
	out := new(TwinInstanceEventsFiltersAttributes)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInstanceHttpEndpointSettings) DeepCopyInto(out *TwinInstanceHttpEndpointSettings) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInstanceHttpEndpointSettings.
func (in *TwinInstanceHttpEndpointSettings) DeepCopy() *TwinInstanceHttpEndpointSettings {
	if in == nil {
		return nil
	}
	out := new(TwinInstanceHttpEndpointSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInstanceList) DeepCopyInto(out *TwinInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TwinInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInstanceList.
func (in *TwinInstanceList) DeepCopy() *TwinInstanceList {
	if in == nil {
		return nil
	}
	out := new(TwinInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TwinInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInstanceMqttEndpointSettings) DeepCopyInto(out *TwinInstanceMqttEndpointSettings) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInstanceMqttEndpointSettings.
func (in *TwinInstanceMqttEndpointSettings) DeepCopy() *TwinInstanceMqttEndpointSettings {
	if in == nil {
		return nil
	}
	out := new(TwinInstanceMqttEndpointSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInstancePropertyData) DeepCopyInto(out *TwinInstancePropertyData) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInstancePropertyData.
func (in *TwinInstancePropertyData) DeepCopy() *TwinInstancePropertyData {
	if in == nil {
		return nil
	}
	out := new(TwinInstancePropertyData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInstanceRelationship) DeepCopyInto(out *TwinInstanceRelationship) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInstanceRelationship.
func (in *TwinInstanceRelationship) DeepCopy() *TwinInstanceRelationship {
	if in == nil {
		return nil
	}
	out := new(TwinInstanceRelationship)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInstanceSpec) DeepCopyInto(out *TwinInstanceSpec) {
	*out = *in
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = make([]TwinInstanceEvents, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EndpointSettings != nil {
		in, out := &in.EndpointSettings, &out.EndpointSettings
		*out = new(TwinInstanceEndpointSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(TwinInstanceDataSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.TwinInstanceRelationships != nil {
		in, out := &in.TwinInstanceRelationships, &out.TwinInstanceRelationships
		*out = make([]TwinInstanceRelationship, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInstanceSpec.
func (in *TwinInstanceSpec) DeepCopy() *TwinInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(TwinInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInstanceStatus) DeepCopyInto(out *TwinInstanceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInstanceStatus.
func (in *TwinInstanceStatus) DeepCopy() *TwinInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(TwinInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInstanceTelemetryData) DeepCopyInto(out *TwinInstanceTelemetryData) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInstanceTelemetryData.
func (in *TwinInstanceTelemetryData) DeepCopy() *TwinInstanceTelemetryData {
	if in == nil {
		return nil
	}
	out := new(TwinInstanceTelemetryData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInterface) DeepCopyInto(out *TwinInterface) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInterface.
func (in *TwinInterface) DeepCopy() *TwinInterface {
	if in == nil {
		return nil
	}
	out := new(TwinInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TwinInterface) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInterfaceAutoScaling) DeepCopyInto(out *TwinInterfaceAutoScaling) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInterfaceAutoScaling.
func (in *TwinInterfaceAutoScaling) DeepCopy() *TwinInterfaceAutoScaling {
	if in == nil {
		return nil
	}
	out := new(TwinInterfaceAutoScaling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInterfaceEventStore) DeepCopyInto(out *TwinInterfaceEventStore) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInterfaceEventStore.
func (in *TwinInterfaceEventStore) DeepCopy() *TwinInterfaceEventStore {
	if in == nil {
		return nil
	}
	out := new(TwinInterfaceEventStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInterfaceEventsSink) DeepCopyInto(out *TwinInterfaceEventsSink) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInterfaceEventsSink.
func (in *TwinInterfaceEventsSink) DeepCopy() *TwinInterfaceEventsSink {
	if in == nil {
		return nil
	}
	out := new(TwinInterfaceEventsSink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInterfaceList) DeepCopyInto(out *TwinInterfaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TwinInterface, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInterfaceList.
func (in *TwinInterfaceList) DeepCopy() *TwinInterfaceList {
	if in == nil {
		return nil
	}
	out := new(TwinInterfaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TwinInterfaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInterfaceService) DeepCopyInto(out *TwinInterfaceService) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.AutoScaling.DeepCopyInto(&out.AutoScaling)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInterfaceService.
func (in *TwinInterfaceService) DeepCopy() *TwinInterfaceService {
	if in == nil {
		return nil
	}
	out := new(TwinInterfaceService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInterfaceSpec) DeepCopyInto(out *TwinInterfaceSpec) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]TwinProperty, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Commands != nil {
		in, out := &in.Commands, &out.Commands
		*out = make([]TwinCommand, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Relationships != nil {
		in, out := &in.Relationships, &out.Relationships
		*out = make([]TwinRelationship, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Telemetries != nil {
		in, out := &in.Telemetries, &out.Telemetries
		*out = make([]TwinTelemetry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.EventStore = in.EventStore
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(TwinInterfaceService)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInterfaceSpec.
func (in *TwinInterfaceSpec) DeepCopy() *TwinInterfaceSpec {
	if in == nil {
		return nil
	}
	out := new(TwinInterfaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinInterfaceStatus) DeepCopyInto(out *TwinInterfaceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinInterfaceStatus.
func (in *TwinInterfaceStatus) DeepCopy() *TwinInterfaceStatus {
	if in == nil {
		return nil
	}
	out := new(TwinInterfaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinObjectSchema) DeepCopyInto(out *TwinObjectSchema) {
	*out = *in
	if in.EnumValues != nil {
		in, out := &in.EnumValues, &out.EnumValues
		*out = make([]TwinEnumSchemaValues, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinObjectSchema.
func (in *TwinObjectSchema) DeepCopy() *TwinObjectSchema {
	if in == nil {
		return nil
	}
	out := new(TwinObjectSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinObjectSchemaFields) DeepCopyInto(out *TwinObjectSchemaFields) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinObjectSchemaFields.
func (in *TwinObjectSchemaFields) DeepCopy() *TwinObjectSchemaFields {
	if in == nil {
		return nil
	}
	out := new(TwinObjectSchemaFields)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinProperty) DeepCopyInto(out *TwinProperty) {
	*out = *in
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(TwinSchema)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinProperty.
func (in *TwinProperty) DeepCopy() *TwinProperty {
	if in == nil {
		return nil
	}
	out := new(TwinProperty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinRelationship) DeepCopyInto(out *TwinRelationship) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]TwinProperty, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(TwinSchema)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinRelationship.
func (in *TwinRelationship) DeepCopy() *TwinRelationship {
	if in == nil {
		return nil
	}
	out := new(TwinRelationship)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinSchema) DeepCopyInto(out *TwinSchema) {
	*out = *in
	if in.ComplexType != nil {
		in, out := &in.ComplexType, &out.ComplexType
		*out = new(TwinComplexType)
		(*in).DeepCopyInto(*out)
	}
	if in.EnumType != nil {
		in, out := &in.EnumType, &out.EnumType
		*out = new(TwinEnumSchema)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinSchema.
func (in *TwinSchema) DeepCopy() *TwinSchema {
	if in == nil {
		return nil
	}
	out := new(TwinSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinTelemetry) DeepCopyInto(out *TwinTelemetry) {
	*out = *in
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(TwinSchema)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinTelemetry.
func (in *TwinTelemetry) DeepCopy() *TwinTelemetry {
	if in == nil {
		return nil
	}
	out := new(TwinTelemetry)
	in.DeepCopyInto(out)
	return out
}
