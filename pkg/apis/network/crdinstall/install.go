/*
Copyright 2020 KubeSphere Authors

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

package crdinstall

import (
	networkv1alpha1 "github.com/fearlesschenc/kubesphere/pkg/apis/network/v1alpha1"
	k8sruntime "k8s.io/apimachinery/pkg/runtime"
	urlruntime "k8s.io/apimachinery/pkg/util/runtime"
)

func Install(scheme *k8sruntime.Scheme) {
	urlruntime.Must(networkv1alpha1.AddToScheme(scheme))
	urlruntime.Must(scheme.SetVersionPriority(networkv1alpha1.SchemeGroupVersion))
}
