/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package runtime

import (
	"k8s.io/kubernetes/pkg/api/unversioned"
)

// SetGroupVersionKind satisfies the ObjectKind interface for all objects that embed PluginBase
func (obj *PluginBase) SetGroupVersionKind(gvk *unversioned.GroupVersionKind) {
	_, obj.Kind = gvk.ToAPIVersionAndKind()
}

// GroupVersionKind satisfies the ObjectKind interface for all objects that embed PluginBase
func (obj *PluginBase) GroupVersionKind() *unversioned.GroupVersionKind {
	return unversioned.FromAPIVersionAndKind("", obj.Kind)
}

// SetGroupVersionKind satisfies the ObjectKind interface for all objects that embed TypeMeta
func (obj *TypeMeta) SetGroupVersionKind(gvk *unversioned.GroupVersionKind) {
	obj.APIVersion, obj.Kind = gvk.ToAPIVersionAndKind()
}

// GroupVersionKind satisfies the ObjectKind interface for all objects that embed TypeMeta
func (obj *TypeMeta) GroupVersionKind() *unversioned.GroupVersionKind {
	return unversioned.FromAPIVersionAndKind(obj.APIVersion, obj.Kind)
}

func (obj *Unknown) GetObjectKind() unversioned.ObjectKind      { return &obj.TypeMeta }
func (obj *Unstructured) GetObjectKind() unversioned.ObjectKind { return &obj.TypeMeta }
