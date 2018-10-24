/*
Copyright 2018 The Conductor Authors.

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

package util

import (
	"fmt"
	"testing"

	"k8s.io/apimachinery/pkg/runtime"

	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	ting "k8s.io/client-go/testing"
)

// TestApplySecretError applying a secret and return error
func TestApplySecretError(t *testing.T) {
	g := NewGomegaWithT(t)
	mk := fake.NewSimpleClientset()

	mk.PrependReactor("get", "secrets", func(action ting.Action) (handled bool, ret runtime.Object, err error) {
		return true, nil, fmt.Errorf("test-error")
	})
	ex, err := ApplySecret(mk, &corev1.Secret{})
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(Equal("test-error"))
	g.Expect(ex).To(BeNil())
	a := mk.Actions()
	g.Expect(len(a)).To(Equal(1))
	g.Expect(a[0].GetVerb()).To(Equal("get"))
}

// TestApplySecretCreate applying a secret that does not exist - expected action: create
func TestApplySecretCreate(t *testing.T) {
	g := NewGomegaWithT(t)
	mk := fake.NewSimpleClientset()

	cs := &corev1.Secret{}
	ex, err := ApplySecret(mk, cs)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(ex).NotTo(BeNil())

	a := mk.Actions()
	g.Expect(len(a)).To(Equal(2))
	g.Expect(a[0].GetVerb()).To(Equal("get"))
	g.Expect(a[1].GetVerb()).To(Equal("create"))

	mk.ClearActions()
}

// TestApplySecretUpdate applying a secret that already exists - expected action: update
func TestApplySecretUpdate(t *testing.T) {
	g := NewGomegaWithT(t)

	cs := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "bar",
		},
	}

	mk := fake.NewSimpleClientset(cs)

	ex, err := ApplySecret(mk, cs)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(ex).NotTo(BeNil())

	a := mk.Actions()
	g.Expect(len(a)).To(Equal(2))
	g.Expect(a[0].GetVerb()).To(Equal("get"))
	g.Expect(a[1].GetVerb()).To(Equal("update"))
}
