/*
Copyright 2017 The Kubernetes Authors.
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
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CustomedHorizontalPodAutoscaler is a specification for a CustomedHorizontalPodAutoscaler resource
type CustomedHorizontalPodAutoscaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec CustomedHorizontalPodAutoscalerSpec `json:"spec"`
	//Status CustomedHorizontalPodAutoscalerStatus `json:"status,omitempty"`
}

type CustomedHorizontalPodAutoscalerSpec struct {
	CoolDownTime   string         `json:"coolDownTime"`
	MaxReplicas    *int32         `json:"maxReplicas"`
	MinReplicas    *int32         `json:"minReplicas"`
	Rules          []Rule         `json:"rules"`
	ScaleTargetRef ScaleTargetRef `json:"scaleTargetRef"`
}

type Rule struct {
	Actions       []Action      `json:"actions"`
	Disable       *bool         `json:"disable"`
	MetricTrigger MetricTrigger `json:"metricTrigger"`
	RuleName      string        `json:"ruleName"`
	RuleType      string        `json:"ruleType"`
}

type Action struct {
	MetricRange    string `json:"metricRange"`
	OperationType  string `json:"operationType"`
	OperationUnit  string `json:"operationUnit"`
	OperationValue *int32 `json:"operationValue"`
}

type MetricTrigger struct {
	HitThreshold    *int32   `json:"hitThreshold"`
	MetricName      string   `json:"metricName"`
	MetricOperation string   `json:"metricOperation"`
	MetricValue     *float32 `json:"metricValue"`
	PeriodSeconds   *int32   `json:"periodSeconds"`
	Statistic       string   `json:"statistic"`
}

type ScaleTargetRef struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CustomedHorizontalPodAutoscalerList is a list of CustomedHorizontalPodAutoscalerList resources
type CustomedHorizontalPodAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []CustomedHorizontalPodAutoscaler `json:"items"`
}
