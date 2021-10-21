/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "k8s.io/application-aware-controller/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AppawareClusterAutoscalers returns a AppawareClusterAutoscalerInformer.
	AppawareClusterAutoscalers() AppawareClusterAutoscalerInformer
	// AppawareHorizontalPodAutoscalers returns a AppawareHorizontalPodAutoscalerInformer.
	AppawareHorizontalPodAutoscalers() AppawareHorizontalPodAutoscalerInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AppawareClusterAutoscalers returns a AppawareClusterAutoscalerInformer.
func (v *version) AppawareClusterAutoscalers() AppawareClusterAutoscalerInformer {
	return &appawareClusterAutoscalerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AppawareHorizontalPodAutoscalers returns a AppawareHorizontalPodAutoscalerInformer.
func (v *version) AppawareHorizontalPodAutoscalers() AppawareHorizontalPodAutoscalerInformer {
	return &appawareHorizontalPodAutoscalerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
