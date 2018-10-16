package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	Group            = "policy.experiments.k8s.io"
	Version          = "v1alpha1"
	ResourcePlural   = "podschedulingpolicyreviews"
	ResourceSingular = "podschedulingpolicyreview"
	ResourceKind     = "PodSchedulingPolicyReview"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SchedulingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []SchedulingPolicy `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SchedulingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              SchedulingPolicySpec   `json:"spec"`
	Status            SchedulingPolicyStatus `json:"status,omitempty"`
}

// SchedulingPolicySpec describes the desired scheduling policy
type SchedulingPolicySpec struct {
	// Priority of this policy compared to other matching policies. The higher
	// the value, the greater priority. Policies are evaluated against a pod in
	// priority order, with the first matching policy's action returned.
	Priority int32 `json:"priority"`
	// Either "allow" or "deny":
	//   "allow"  - matching pods are admitted
	//   "deny"   - matching pods are denied admittance
	Action Action `json:"action"`
	// A label query to identify which namespaces are within this policy's scope
	NamespaceSelector *metav1.LabelSelector `json:"namespaceSelector"`
	// A label query to select which pods are within this policy's scope
	PodSelector *metav1.LabelSelector `json:"podSelector"`
	//
	Criteria []Criteria `json:"criteria"`
}

// Action specifies what action, allow or deny, to take on matching pods
type Action string

const (
	// ActionAllow means matching pods should be admitted
	ActionAllow Action = "allow"
	// ActionDeny means matching pods should be denied admittance
	ActionDeny Action = "deny"
)

// Criteria contains scheduling rules to match the pod against
type Criteria struct {
	SchedulerNames []SchedulerNameMatcher `json:"schedulerNames"`
	// TODO
	// PriorityClassNames []PriorityClassMatcher `json:"priorityClassNames"`
	// Tolerations []TolerationsMatcher `json:"tolerationsMatcher"`
}

// SchedulerNameMatcher
type SchedulerNameMatcher struct {
	Operator CriteriaOperator
	Values   []string
}

// CriteriaOperator are supported set-based operators
type CriteriaOperator string

const (
	// CriteriaOpExists requires that a field or key has some value set
	CriteriaOpExists CriteriaOperator = "exists"
	// CriteriaOpIn requires that the value of the field or key be one of the ones listed
	CriteriaOpIn CriteriaOperator = "in"
	// CriteriaOpNotIn requires that the value of the field or key not be one of those listed
	CriteriaOpNotIn CriteriaOperator = "notin"
)

// SchedulingPolicyStatus contains information from the controller processing the policy resource
type SchedulingPolicyStatus struct {
	// Fill me
}
