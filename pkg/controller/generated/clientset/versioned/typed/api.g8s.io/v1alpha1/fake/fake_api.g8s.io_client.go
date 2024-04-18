/*
Copyright 2024 James Riley O'Donnell.

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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/jrodonnell/g8s/pkg/controller/generated/clientset/versioned/typed/api.g8s.io/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeApiV1alpha1 struct {
	*testing.Fake
}

func (c *FakeApiV1alpha1) Allowlists() v1alpha1.AllowlistInterface {
	return &FakeAllowlists{c}
}

func (c *FakeApiV1alpha1) Logins(namespace string) v1alpha1.LoginInterface {
	return &FakeLogins{c, namespace}
}

func (c *FakeApiV1alpha1) SSHKeyPairs(namespace string) v1alpha1.SSHKeyPairInterface {
	return &FakeSSHKeyPairs{c, namespace}
}

func (c *FakeApiV1alpha1) SelfSignedTLSBundles(namespace string) v1alpha1.SelfSignedTLSBundleInterface {
	return &FakeSelfSignedTLSBundles{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeApiV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}