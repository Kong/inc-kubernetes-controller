//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	"github.com/kong/go-kong/kong"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigSource) DeepCopyInto(out *ConfigSource) {
	*out = *in
	out.SecretValue = in.SecretValue
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigSource.
func (in *ConfigSource) DeepCopy() *ConfigSource {
	if in == nil {
		return nil
	}
	out := new(ConfigSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongPlugin) DeepCopyInto(out *KongPlugin) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Config.DeepCopyInto(&out.Config)
	if in.ConfigFrom != nil {
		in, out := &in.ConfigFrom, &out.ConfigFrom
		*out = new(ConfigSource)
		**out = **in
	}
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]KongProtocol, len(*in))
		copy(*out, *in)
	}
	if in.Ordering != nil {
		in, out := &in.Ordering, &out.Ordering
		*out = new(kong.PluginOrdering)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongPlugin.
func (in *KongPlugin) DeepCopy() *KongPlugin {
	if in == nil {
		return nil
	}
	out := new(KongPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongPlugin) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongPluginList) DeepCopyInto(out *KongPluginList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KongPlugin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongPluginList.
func (in *KongPluginList) DeepCopy() *KongPluginList {
	if in == nil {
		return nil
	}
	out := new(KongPluginList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongPluginList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedConfigSource) DeepCopyInto(out *NamespacedConfigSource) {
	*out = *in
	out.SecretValue = in.SecretValue
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedConfigSource.
func (in *NamespacedConfigSource) DeepCopy() *NamespacedConfigSource {
	if in == nil {
		return nil
	}
	out := new(NamespacedConfigSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedSecretValueFromSource) DeepCopyInto(out *NamespacedSecretValueFromSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedSecretValueFromSource.
func (in *NamespacedSecretValueFromSource) DeepCopy() *NamespacedSecretValueFromSource {
	if in == nil {
		return nil
	}
	out := new(NamespacedSecretValueFromSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretValueFromSource) DeepCopyInto(out *SecretValueFromSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretValueFromSource.
func (in *SecretValueFromSource) DeepCopy() *SecretValueFromSource {
	if in == nil {
		return nil
	}
	out := new(SecretValueFromSource)
	in.DeepCopyInto(out)
	return out
}