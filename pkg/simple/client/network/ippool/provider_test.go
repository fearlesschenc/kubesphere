/*
Copyright 2020 The KubeSphere authors.

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

package ippool

import (
	"testing"

	"github.com/fearlesschenc/kubesphere/pkg/apis/network/v1alpha1"
	fakeks "github.com/fearlesschenc/kubesphere/pkg/client/clientset/versioned/fake"
	"github.com/fearlesschenc/kubesphere/pkg/simple/client/network/ippool/ipam"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func testNewProvider() provider {
	return NewProvider(fakeks.NewSimpleClientset(), Options{})
}

func TestProvider_GetIPPoolStats(t *testing.T) {
	p := testNewProvider()

	pool := v1alpha1.IPPool{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name: "testippool",
			Labels: map[string]string{
				v1alpha1.IPPoolTypeLabel: v1alpha1.VLAN,
			},
		},
		Spec: v1alpha1.IPPoolSpec{
			Type: v1alpha1.VLAN,
			CIDR: "192.168.0.0/24",
		},
		Status: v1alpha1.IPPoolStatus{},
	}

	_, err := p.kubesphereClient.NetworkV1alpha1().IPPools().Create(&pool)
	if err != nil {
		t.FailNow()
	}

	p.ipamclient.AutoAssign(ipam.AutoAssignArgs{
		HandleID: "testhandle",
		Attrs:    nil,
		Pool:     "testippool",
	})
	stat, err := p.GetIPPoolStats(&pool)
	if err != nil {
		t.FailNow()
	}
	if stat.Status.Unallocated != pool.NumAddresses()-1 || stat.Status.Reserved != 0 ||
		stat.Status.Capacity != pool.NumAddresses() || stat.Status.Allocations != 1 {
		t.FailNow()
	}
}
