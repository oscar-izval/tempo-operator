//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationSpec) DeepCopyInto(out *AuthenticationSpec) {
	*out = *in
	if in.OIDC != nil {
		in, out := &in.OIDC, &out.OIDC
		*out = new(OIDCSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationSpec.
func (in *AuthenticationSpec) DeepCopy() *AuthenticationSpec {
	if in == nil {
		return nil
	}
	out := new(AuthenticationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorizationSpec) DeepCopyInto(out *AuthorizationSpec) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]RoleSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RoleBindings != nil {
		in, out := &in.RoleBindings, &out.RoleBindings
		*out = make([]RoleBindingsSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationSpec.
func (in *AuthorizationSpec) DeepCopy() *AuthorizationSpec {
	if in == nil {
		return nil
	}
	out := new(AuthorizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentStatus) DeepCopyInto(out *ComponentStatus) {
	*out = *in
	if in.Compactor != nil {
		in, out := &in.Compactor, &out.Compactor
		*out = make(PodStatusMap, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.Distributor != nil {
		in, out := &in.Distributor, &out.Distributor
		*out = make(PodStatusMap, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.Ingester != nil {
		in, out := &in.Ingester, &out.Ingester
		*out = make(PodStatusMap, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.Querier != nil {
		in, out := &in.Querier, &out.Querier
		*out = make(PodStatusMap, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.QueryFrontend != nil {
		in, out := &in.QueryFrontend, &out.QueryFrontend
		*out = make(PodStatusMap, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = make(PodStatusMap, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentStatus.
func (in *ComponentStatus) DeepCopy() *ComponentStatus {
	if in == nil {
		return nil
	}
	out := new(ComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Defaulter) DeepCopyInto(out *Defaulter) {
	*out = *in
	in.ctrlConfig.DeepCopyInto(&out.ctrlConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Defaulter.
func (in *Defaulter) DeepCopy() *Defaulter {
	if in == nil {
		return nil
	}
	out := new(Defaulter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngestionLimitSpec) DeepCopyInto(out *IngestionLimitSpec) {
	*out = *in
	if in.IngestionBurstSizeBytes != nil {
		in, out := &in.IngestionBurstSizeBytes, &out.IngestionBurstSizeBytes
		*out = new(int)
		**out = **in
	}
	if in.IngestionRateLimitBytes != nil {
		in, out := &in.IngestionRateLimitBytes, &out.IngestionRateLimitBytes
		*out = new(int)
		**out = **in
	}
	if in.MaxBytesPerTrace != nil {
		in, out := &in.MaxBytesPerTrace, &out.MaxBytesPerTrace
		*out = new(int)
		**out = **in
	}
	if in.MaxTracesPerUser != nil {
		in, out := &in.MaxTracesPerUser, &out.MaxTracesPerUser
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngestionLimitSpec.
func (in *IngestionLimitSpec) DeepCopy() *IngestionLimitSpec {
	if in == nil {
		return nil
	}
	out := new(IngestionLimitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressSpec) DeepCopyInto(out *IngressSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.IngressClassName != nil {
		in, out := &in.IngressClassName, &out.IngressClassName
		*out = new(string)
		**out = **in
	}
	out.Route = in.Route
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
func (in *JaegerQuerySpec) DeepCopyInto(out *JaegerQuerySpec) {
	*out = *in
	in.Ingress.DeepCopyInto(&out.Ingress)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerQuerySpec.
func (in *JaegerQuerySpec) DeepCopy() *JaegerQuerySpec {
	if in == nil {
		return nil
	}
	out := new(JaegerQuerySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LimitSpec) DeepCopyInto(out *LimitSpec) {
	*out = *in
	if in.PerTenant != nil {
		in, out := &in.PerTenant, &out.PerTenant
		*out = make(map[string]RateLimitSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	in.Global.DeepCopyInto(&out.Global)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LimitSpec.
func (in *LimitSpec) DeepCopy() *LimitSpec {
	if in == nil {
		return nil
	}
	out := new(LimitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsConfigSpec) DeepCopyInto(out *MetricsConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsConfigSpec.
func (in *MetricsConfigSpec) DeepCopy() *MetricsConfigSpec {
	if in == nil {
		return nil
	}
	out := new(MetricsConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDCSpec) DeepCopyInto(out *OIDCSpec) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(TenantSecretSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDCSpec.
func (in *OIDCSpec) DeepCopy() *OIDCSpec {
	if in == nil {
		return nil
	}
	out := new(OIDCSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStorageSecretSpec) DeepCopyInto(out *ObjectStorageSecretSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStorageSecretSpec.
func (in *ObjectStorageSecretSpec) DeepCopy() *ObjectStorageSecretSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectStorageSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStorageSpec) DeepCopyInto(out *ObjectStorageSpec) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(ObjectStorageTLSSpec)
		**out = **in
	}
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStorageSpec.
func (in *ObjectStorageSpec) DeepCopy() *ObjectStorageSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStorageTLSSpec) DeepCopyInto(out *ObjectStorageTLSSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStorageTLSSpec.
func (in *ObjectStorageTLSSpec) DeepCopy() *ObjectStorageTLSSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectStorageTLSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilitySpec) DeepCopyInto(out *ObservabilitySpec) {
	*out = *in
	out.Metrics = in.Metrics
	out.Tracing = in.Tracing
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilitySpec.
func (in *ObservabilitySpec) DeepCopy() *ObservabilitySpec {
	if in == nil {
		return nil
	}
	out := new(ObservabilitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PodStatusMap) DeepCopyInto(out *PodStatusMap) {
	{
		in := &in
		*out = make(PodStatusMap, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodStatusMap.
func (in PodStatusMap) DeepCopy() PodStatusMap {
	if in == nil {
		return nil
	}
	out := new(PodStatusMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryLimit) DeepCopyInto(out *QueryLimit) {
	*out = *in
	if in.MaxBytesPerTagValues != nil {
		in, out := &in.MaxBytesPerTagValues, &out.MaxBytesPerTagValues
		*out = new(int)
		**out = **in
	}
	if in.MaxSearchBytesPerTrace != nil {
		in, out := &in.MaxSearchBytesPerTrace, &out.MaxSearchBytesPerTrace
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryLimit.
func (in *QueryLimit) DeepCopy() *QueryLimit {
	if in == nil {
		return nil
	}
	out := new(QueryLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitSpec) DeepCopyInto(out *RateLimitSpec) {
	*out = *in
	in.Ingestion.DeepCopyInto(&out.Ingestion)
	in.Query.DeepCopyInto(&out.Query)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitSpec.
func (in *RateLimitSpec) DeepCopy() *RateLimitSpec {
	if in == nil {
		return nil
	}
	out := new(RateLimitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	if in.Total != nil {
		in, out := &in.Total, &out.Total
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionConfig) DeepCopyInto(out *RetentionConfig) {
	*out = *in
	out.Traces = in.Traces
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionConfig.
func (in *RetentionConfig) DeepCopy() *RetentionConfig {
	if in == nil {
		return nil
	}
	out := new(RetentionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionSpec) DeepCopyInto(out *RetentionSpec) {
	*out = *in
	if in.PerTenant != nil {
		in, out := &in.PerTenant, &out.PerTenant
		*out = make(map[string]RetentionConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Global = in.Global
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionSpec.
func (in *RetentionSpec) DeepCopy() *RetentionSpec {
	if in == nil {
		return nil
	}
	out := new(RetentionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleBindingsSpec) DeepCopyInto(out *RoleBindingsSpec) {
	*out = *in
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]Subject, len(*in))
		copy(*out, *in)
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleBindingsSpec.
func (in *RoleBindingsSpec) DeepCopy() *RoleBindingsSpec {
	if in == nil {
		return nil
	}
	out := new(RoleBindingsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleSpec) DeepCopyInto(out *RoleSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tenants != nil {
		in, out := &in.Tenants, &out.Tenants
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]PermissionType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleSpec.
func (in *RoleSpec) DeepCopy() *RoleSpec {
	if in == nil {
		return nil
	}
	out := new(RoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteSpec) DeepCopyInto(out *RouteSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSpec.
func (in *RouteSpec) DeepCopy() *RouteSpec {
	if in == nil {
		return nil
	}
	out := new(RouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchSpec) DeepCopyInto(out *SearchSpec) {
	*out = *in
	if in.DefaultResultLimit != nil {
		in, out := &in.DefaultResultLimit, &out.DefaultResultLimit
		*out = new(int)
		**out = **in
	}
	out.MaxDuration = in.MaxDuration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchSpec.
func (in *SearchSpec) DeepCopy() *SearchSpec {
	if in == nil {
		return nil
	}
	out := new(SearchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subject) DeepCopyInto(out *Subject) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subject.
func (in *Subject) DeepCopy() *Subject {
	if in == nil {
		return nil
	}
	out := new(Subject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TempoComponentSpec) DeepCopyInto(out *TempoComponentSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TempoComponentSpec.
func (in *TempoComponentSpec) DeepCopy() *TempoComponentSpec {
	if in == nil {
		return nil
	}
	out := new(TempoComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TempoGatewaySpec) DeepCopyInto(out *TempoGatewaySpec) {
	*out = *in
	in.TempoComponentSpec.DeepCopyInto(&out.TempoComponentSpec)
	in.Ingress.DeepCopyInto(&out.Ingress)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TempoGatewaySpec.
func (in *TempoGatewaySpec) DeepCopy() *TempoGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(TempoGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TempoQueryFrontendSpec) DeepCopyInto(out *TempoQueryFrontendSpec) {
	*out = *in
	in.TempoComponentSpec.DeepCopyInto(&out.TempoComponentSpec)
	in.JaegerQuery.DeepCopyInto(&out.JaegerQuery)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TempoQueryFrontendSpec.
func (in *TempoQueryFrontendSpec) DeepCopy() *TempoQueryFrontendSpec {
	if in == nil {
		return nil
	}
	out := new(TempoQueryFrontendSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TempoStack) DeepCopyInto(out *TempoStack) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TempoStack.
func (in *TempoStack) DeepCopy() *TempoStack {
	if in == nil {
		return nil
	}
	out := new(TempoStack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TempoStack) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TempoStackList) DeepCopyInto(out *TempoStackList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TempoStack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TempoStackList.
func (in *TempoStackList) DeepCopy() *TempoStackList {
	if in == nil {
		return nil
	}
	out := new(TempoStackList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TempoStackList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TempoStackSpec) DeepCopyInto(out *TempoStackSpec) {
	*out = *in
	in.LimitSpec.DeepCopyInto(&out.LimitSpec)
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
	out.StorageSize = in.StorageSize.DeepCopy()
	out.Images = in.Images
	in.Storage.DeepCopyInto(&out.Storage)
	in.Retention.DeepCopyInto(&out.Retention)
	in.SearchSpec.DeepCopyInto(&out.SearchSpec)
	in.Template.DeepCopyInto(&out.Template)
	if in.Tenants != nil {
		in, out := &in.Tenants, &out.Tenants
		*out = new(TenantsSpec)
		(*in).DeepCopyInto(*out)
	}
	out.Observability = in.Observability
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TempoStackSpec.
func (in *TempoStackSpec) DeepCopy() *TempoStackSpec {
	if in == nil {
		return nil
	}
	out := new(TempoStackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TempoStackStatus) DeepCopyInto(out *TempoStackStatus) {
	*out = *in
	in.Components.DeepCopyInto(&out.Components)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TempoStackStatus.
func (in *TempoStackStatus) DeepCopy() *TempoStackStatus {
	if in == nil {
		return nil
	}
	out := new(TempoStackStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TempoTemplateSpec) DeepCopyInto(out *TempoTemplateSpec) {
	*out = *in
	in.Distributor.DeepCopyInto(&out.Distributor)
	in.Ingester.DeepCopyInto(&out.Ingester)
	in.Compactor.DeepCopyInto(&out.Compactor)
	in.Querier.DeepCopyInto(&out.Querier)
	in.QueryFrontend.DeepCopyInto(&out.QueryFrontend)
	in.Gateway.DeepCopyInto(&out.Gateway)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TempoTemplateSpec.
func (in *TempoTemplateSpec) DeepCopy() *TempoTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(TempoTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantSecretSpec) DeepCopyInto(out *TenantSecretSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantSecretSpec.
func (in *TenantSecretSpec) DeepCopy() *TenantSecretSpec {
	if in == nil {
		return nil
	}
	out := new(TenantSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantsSpec) DeepCopyInto(out *TenantsSpec) {
	*out = *in
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = make([]AuthenticationSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Authorization != nil {
		in, out := &in.Authorization, &out.Authorization
		*out = new(AuthorizationSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantsSpec.
func (in *TenantsSpec) DeepCopy() *TenantsSpec {
	if in == nil {
		return nil
	}
	out := new(TenantsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracingConfigSpec) DeepCopyInto(out *TracingConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracingConfigSpec.
func (in *TracingConfigSpec) DeepCopy() *TracingConfigSpec {
	if in == nil {
		return nil
	}
	out := new(TracingConfigSpec)
	in.DeepCopyInto(out)
	return out
}
