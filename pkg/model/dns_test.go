/*
Copyright 2016 The Kubernetes Authors.

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

package model

import (
	"testing"

	"k8s.io/kops/pkg/apis/kops"
)

func TestHasNonGossipPublicName(t *testing.T) {
	hostnameTests := []struct {
		masterPublicHostname  string
		bastionPublicHostname string
		expectedResult        bool
	}{
		{"api.example.com.k8s.local", "", false},
		{"api.example.com.k8s.local", "bastion.example.com.k8s.local", false},
		{"api.example.com", "bastion.example.com.k8s.local", true},
		{"api.example.com.k8s.local", "bastion.example.com", true},
		{"api.example.com", "bastion.example.com", true},
	}
	for _, test := range hostnameTests {
		spec := kops.ClusterSpec{
			MasterPublicName: test.masterPublicHostname,
			Topology: &kops.TopologySpec{
				Bastion: &kops.BastionSpec{
					BastionPublicName: test.bastionPublicHostname,
				},
			},
		}
		actualResult := hasNonGossipPublicName(spec)
		if actualResult != test.expectedResult {
			t.Fatalf("expected hasNonGossipPublicName([%s, %s]) to be %t",
				test.masterPublicHostname, test.bastionPublicHostname, test.expectedResult)
		}
	}
}
