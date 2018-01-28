// +build !ignore_autogenerated

/*
Copyright 2018 The Voyager Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package voyager

import (
	core_v1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ACMECertificateDetails).DeepCopyInto(out.(*ACMECertificateDetails))
			return nil
		}, InType: reflect.TypeOf(&ACMECertificateDetails{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AuthOption).DeepCopyInto(out.(*AuthOption))
			return nil
		}, InType: reflect.TypeOf(&AuthOption{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BasicAuth).DeepCopyInto(out.(*BasicAuth))
			return nil
		}, InType: reflect.TypeOf(&BasicAuth{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Certificate).DeepCopyInto(out.(*Certificate))
			return nil
		}, InType: reflect.TypeOf(&Certificate{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CertificateCondition).DeepCopyInto(out.(*CertificateCondition))
			return nil
		}, InType: reflect.TypeOf(&CertificateCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CertificateDetails).DeepCopyInto(out.(*CertificateDetails))
			return nil
		}, InType: reflect.TypeOf(&CertificateDetails{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CertificateList).DeepCopyInto(out.(*CertificateList))
			return nil
		}, InType: reflect.TypeOf(&CertificateList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CertificateSpec).DeepCopyInto(out.(*CertificateSpec))
			return nil
		}, InType: reflect.TypeOf(&CertificateSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CertificateStatus).DeepCopyInto(out.(*CertificateStatus))
			return nil
		}, InType: reflect.TypeOf(&CertificateStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CertificateStorage).DeepCopyInto(out.(*CertificateStorage))
			return nil
		}, InType: reflect.TypeOf(&CertificateStorage{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ChallengeProvider).DeepCopyInto(out.(*ChallengeProvider))
			return nil
		}, InType: reflect.TypeOf(&ChallengeProvider{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DNSChallengeProvider).DeepCopyInto(out.(*DNSChallengeProvider))
			return nil
		}, InType: reflect.TypeOf(&DNSChallengeProvider{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*FrontendRule).DeepCopyInto(out.(*FrontendRule))
			return nil
		}, InType: reflect.TypeOf(&FrontendRule{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*HTTPChallengeProvider).DeepCopyInto(out.(*HTTPChallengeProvider))
			return nil
		}, InType: reflect.TypeOf(&HTTPChallengeProvider{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*HTTPIngressBackend).DeepCopyInto(out.(*HTTPIngressBackend))
			return nil
		}, InType: reflect.TypeOf(&HTTPIngressBackend{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*HTTPIngressPath).DeepCopyInto(out.(*HTTPIngressPath))
			return nil
		}, InType: reflect.TypeOf(&HTTPIngressPath{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*HTTPIngressRuleValue).DeepCopyInto(out.(*HTTPIngressRuleValue))
			return nil
		}, InType: reflect.TypeOf(&HTTPIngressRuleValue{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Ingress).DeepCopyInto(out.(*Ingress))
			return nil
		}, InType: reflect.TypeOf(&Ingress{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*IngressBackend).DeepCopyInto(out.(*IngressBackend))
			return nil
		}, InType: reflect.TypeOf(&IngressBackend{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*IngressList).DeepCopyInto(out.(*IngressList))
			return nil
		}, InType: reflect.TypeOf(&IngressList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*IngressRef).DeepCopyInto(out.(*IngressRef))
			return nil
		}, InType: reflect.TypeOf(&IngressRef{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*IngressRule).DeepCopyInto(out.(*IngressRule))
			return nil
		}, InType: reflect.TypeOf(&IngressRule{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*IngressRuleValue).DeepCopyInto(out.(*IngressRuleValue))
			return nil
		}, InType: reflect.TypeOf(&IngressRuleValue{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*IngressSpec).DeepCopyInto(out.(*IngressSpec))
			return nil
		}, InType: reflect.TypeOf(&IngressSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*IngressStatus).DeepCopyInto(out.(*IngressStatus))
			return nil
		}, InType: reflect.TypeOf(&IngressStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*IngressTLS).DeepCopyInto(out.(*IngressTLS))
			return nil
		}, InType: reflect.TypeOf(&IngressTLS{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*LocalTypedReference).DeepCopyInto(out.(*LocalTypedReference))
			return nil
		}, InType: reflect.TypeOf(&LocalTypedReference{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SecretStore).DeepCopyInto(out.(*SecretStore))
			return nil
		}, InType: reflect.TypeOf(&SecretStore{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TCPIngressRuleValue).DeepCopyInto(out.(*TCPIngressRuleValue))
			return nil
		}, InType: reflect.TypeOf(&TCPIngressRuleValue{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TLSAuth).DeepCopyInto(out.(*TLSAuth))
			return nil
		}, InType: reflect.TypeOf(&TLSAuth{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*VaultStore).DeepCopyInto(out.(*VaultStore))
			return nil
		}, InType: reflect.TypeOf(&VaultStore{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMECertificateDetails) DeepCopyInto(out *ACMECertificateDetails) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMECertificateDetails.
func (in *ACMECertificateDetails) DeepCopy() *ACMECertificateDetails {
	if in == nil {
		return nil
	}
	out := new(ACMECertificateDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthOption) DeepCopyInto(out *AuthOption) {
	*out = *in
	if in.Basic != nil {
		in, out := &in.Basic, &out.Basic
		if *in == nil {
			*out = nil
		} else {
			*out = new(BasicAuth)
			**out = **in
		}
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		if *in == nil {
			*out = nil
		} else {
			*out = new(TLSAuth)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthOption.
func (in *AuthOption) DeepCopy() *AuthOption {
	if in == nil {
		return nil
	}
	out := new(AuthOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicAuth) DeepCopyInto(out *BasicAuth) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicAuth.
func (in *BasicAuth) DeepCopy() *BasicAuth {
	if in == nil {
		return nil
	}
	out := new(BasicAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Certificate) DeepCopyInto(out *Certificate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Certificate.
func (in *Certificate) DeepCopy() *Certificate {
	if in == nil {
		return nil
	}
	out := new(Certificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Certificate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateCondition) DeepCopyInto(out *CertificateCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateCondition.
func (in *CertificateCondition) DeepCopy() *CertificateCondition {
	if in == nil {
		return nil
	}
	out := new(CertificateCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateDetails) DeepCopyInto(out *CertificateDetails) {
	*out = *in
	in.NotBefore.DeepCopyInto(&out.NotBefore)
	in.NotAfter.DeepCopyInto(&out.NotAfter)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateDetails.
func (in *CertificateDetails) DeepCopy() *CertificateDetails {
	if in == nil {
		return nil
	}
	out := new(CertificateDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateList) DeepCopyInto(out *CertificateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Certificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateList.
func (in *CertificateList) DeepCopy() *CertificateList {
	if in == nil {
		return nil
	}
	out := new(CertificateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSpec) DeepCopyInto(out *CertificateSpec) {
	*out = *in
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.ChallengeProvider.DeepCopyInto(&out.ChallengeProvider)
	in.Storage.DeepCopyInto(&out.Storage)
	out.HTTPProviderIngressReference = in.HTTPProviderIngressReference
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSpec.
func (in *CertificateSpec) DeepCopy() *CertificateSpec {
	if in == nil {
		return nil
	}
	out := new(CertificateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateStatus) DeepCopyInto(out *CertificateStatus) {
	*out = *in
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CertificateCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastIssuedCertificate != nil {
		in, out := &in.LastIssuedCertificate, &out.LastIssuedCertificate
		if *in == nil {
			*out = nil
		} else {
			*out = new(CertificateDetails)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		if *in == nil {
			*out = nil
		} else {
			*out = new(ACMECertificateDetails)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateStatus.
func (in *CertificateStatus) DeepCopy() *CertificateStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateStorage) DeepCopyInto(out *CertificateStorage) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.LocalObjectReference)
			**out = **in
		}
	}
	if in.Vault != nil {
		in, out := &in.Vault, &out.Vault
		if *in == nil {
			*out = nil
		} else {
			*out = new(VaultStore)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateStorage.
func (in *CertificateStorage) DeepCopy() *CertificateStorage {
	if in == nil {
		return nil
	}
	out := new(CertificateStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChallengeProvider) DeepCopyInto(out *ChallengeProvider) {
	*out = *in
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		if *in == nil {
			*out = nil
		} else {
			*out = new(HTTPChallengeProvider)
			**out = **in
		}
	}
	if in.DNS != nil {
		in, out := &in.DNS, &out.DNS
		if *in == nil {
			*out = nil
		} else {
			*out = new(DNSChallengeProvider)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChallengeProvider.
func (in *ChallengeProvider) DeepCopy() *ChallengeProvider {
	if in == nil {
		return nil
	}
	out := new(ChallengeProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSChallengeProvider) DeepCopyInto(out *DNSChallengeProvider) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSChallengeProvider.
func (in *DNSChallengeProvider) DeepCopy() *DNSChallengeProvider {
	if in == nil {
		return nil
	}
	out := new(DNSChallengeProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrontendRule) DeepCopyInto(out *FrontendRule) {
	*out = *in
	out.Port = in.Port
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		if *in == nil {
			*out = nil
		} else {
			*out = new(AuthOption)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrontendRule.
func (in *FrontendRule) DeepCopy() *FrontendRule {
	if in == nil {
		return nil
	}
	out := new(FrontendRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPChallengeProvider) DeepCopyInto(out *HTTPChallengeProvider) {
	*out = *in
	out.Ingress = in.Ingress
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPChallengeProvider.
func (in *HTTPChallengeProvider) DeepCopy() *HTTPChallengeProvider {
	if in == nil {
		return nil
	}
	out := new(HTTPChallengeProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPIngressBackend) DeepCopyInto(out *HTTPIngressBackend) {
	*out = *in
	in.IngressBackend.DeepCopyInto(&out.IngressBackend)
	if in.RewriteRule != nil {
		in, out := &in.RewriteRule, &out.RewriteRule
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HeaderRule != nil {
		in, out := &in.HeaderRule, &out.HeaderRule
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPIngressBackend.
func (in *HTTPIngressBackend) DeepCopy() *HTTPIngressBackend {
	if in == nil {
		return nil
	}
	out := new(HTTPIngressBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPIngressPath) DeepCopyInto(out *HTTPIngressPath) {
	*out = *in
	in.Backend.DeepCopyInto(&out.Backend)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPIngressPath.
func (in *HTTPIngressPath) DeepCopy() *HTTPIngressPath {
	if in == nil {
		return nil
	}
	out := new(HTTPIngressPath)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPIngressRuleValue) DeepCopyInto(out *HTTPIngressRuleValue) {
	*out = *in
	out.Port = in.Port
	out.NodePort = in.NodePort
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]HTTPIngressPath, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPIngressRuleValue.
func (in *HTTPIngressRuleValue) DeepCopy() *HTTPIngressRuleValue {
	if in == nil {
		return nil
	}
	out := new(HTTPIngressRuleValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ingress) DeepCopyInto(out *Ingress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ingress.
func (in *Ingress) DeepCopy() *Ingress {
	if in == nil {
		return nil
	}
	out := new(Ingress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ingress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressBackend) DeepCopyInto(out *IngressBackend) {
	*out = *in
	if in.HostNames != nil {
		in, out := &in.HostNames, &out.HostNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.ServicePort = in.ServicePort
	if in.BackendRule != nil {
		in, out := &in.BackendRule, &out.BackendRule
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressBackend.
func (in *IngressBackend) DeepCopy() *IngressBackend {
	if in == nil {
		return nil
	}
	out := new(IngressBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressList) DeepCopyInto(out *IngressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ingress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressList.
func (in *IngressList) DeepCopy() *IngressList {
	if in == nil {
		return nil
	}
	out := new(IngressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRef) DeepCopyInto(out *IngressRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRef.
func (in *IngressRef) DeepCopy() *IngressRef {
	if in == nil {
		return nil
	}
	out := new(IngressRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRule) DeepCopyInto(out *IngressRule) {
	*out = *in
	in.IngressRuleValue.DeepCopyInto(&out.IngressRuleValue)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRule.
func (in *IngressRule) DeepCopy() *IngressRule {
	if in == nil {
		return nil
	}
	out := new(IngressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRuleValue) DeepCopyInto(out *IngressRuleValue) {
	*out = *in
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		if *in == nil {
			*out = nil
		} else {
			*out = new(HTTPIngressRuleValue)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.TCP != nil {
		in, out := &in.TCP, &out.TCP
		if *in == nil {
			*out = nil
		} else {
			*out = new(TCPIngressRuleValue)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRuleValue.
func (in *IngressRuleValue) DeepCopy() *IngressRuleValue {
	if in == nil {
		return nil
	}
	out := new(IngressRuleValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressSpec) DeepCopyInto(out *IngressSpec) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		if *in == nil {
			*out = nil
		} else {
			*out = new(HTTPIngressBackend)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = make([]IngressTLS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FrontendRules != nil {
		in, out := &in.FrontendRules, &out.FrontendRules
		*out = make([]FrontendRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]IngressRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LoadBalancerSourceRanges != nil {
		in, out := &in.LoadBalancerSourceRanges, &out.LoadBalancerSourceRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.Affinity)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]core_v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]core_v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.ExternalIPs != nil {
		in, out := &in.ExternalIPs, &out.ExternalIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressSpec.
func (in *IngressSpec) DeepCopy() *IngressSpec {
	if in == nil {
		return nil
	}
	out := new(IngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressStatus) DeepCopyInto(out *IngressStatus) {
	*out = *in
	in.LoadBalancer.DeepCopyInto(&out.LoadBalancer)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressStatus.
func (in *IngressStatus) DeepCopy() *IngressStatus {
	if in == nil {
		return nil
	}
	out := new(IngressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressTLS) DeepCopyInto(out *IngressTLS) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ref != nil {
		in, out := &in.Ref, &out.Ref
		if *in == nil {
			*out = nil
		} else {
			*out = new(LocalTypedReference)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressTLS.
func (in *IngressTLS) DeepCopy() *IngressTLS {
	if in == nil {
		return nil
	}
	out := new(IngressTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalTypedReference) DeepCopyInto(out *LocalTypedReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalTypedReference.
func (in *LocalTypedReference) DeepCopy() *LocalTypedReference {
	if in == nil {
		return nil
	}
	out := new(LocalTypedReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStore) DeepCopyInto(out *SecretStore) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStore.
func (in *SecretStore) DeepCopy() *SecretStore {
	if in == nil {
		return nil
	}
	out := new(SecretStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPIngressRuleValue) DeepCopyInto(out *TCPIngressRuleValue) {
	*out = *in
	out.Port = in.Port
	out.NodePort = in.NodePort
	in.Backend.DeepCopyInto(&out.Backend)
	if in.ALPN != nil {
		in, out := &in.ALPN, &out.ALPN
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPIngressRuleValue.
func (in *TCPIngressRuleValue) DeepCopy() *TCPIngressRuleValue {
	if in == nil {
		return nil
	}
	out := new(TCPIngressRuleValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSAuth) DeepCopyInto(out *TLSAuth) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSAuth.
func (in *TLSAuth) DeepCopy() *TLSAuth {
	if in == nil {
		return nil
	}
	out := new(TLSAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultStore) DeepCopyInto(out *VaultStore) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultStore.
func (in *VaultStore) DeepCopy() *VaultStore {
	if in == nil {
		return nil
	}
	out := new(VaultStore)
	in.DeepCopyInto(out)
	return out
}
