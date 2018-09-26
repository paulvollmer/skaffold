/*
Copyright 2018 The Skaffold Authors

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

package deploy

import (
	"testing"

	"github.com/GoogleContainerTools/skaffold/testutil"
)

func TestDependenciesForKustomization(t *testing.T) {
	tests := []struct {
		description string
		yaml        string
		expected    []string
		shouldErr   bool
	}{
		{
			description: "resources",
			yaml:        `resources: [pod1.yaml, pod2.yaml]`,
			expected:    []string{"kustomization.yaml", "pod1.yaml", "pod2.yaml"},
		},
		{
			description: "paches",
			yaml:        `patches: [patch1.yaml, patch2.yaml]`,
			expected:    []string{"kustomization.yaml", "patch1.yaml", "patch2.yaml"},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			tmp, cleanup := testutil.NewTempDir(t)
			defer cleanup()

			tmp.Write("kustomization.yaml", test.yaml)

			deps, err := dependenciesForKustomization(tmp.Root())

			testutil.CheckErrorAndDeepEqual(t, false, err, joinPaths(tmp.Root(), test.expected), deps)
		})
	}
}
