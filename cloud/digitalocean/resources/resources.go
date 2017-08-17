// Copyright © 2017 The Kubicorn Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resources

import (
	"github.com/kris-nova/kubicorn/cloud"
	"github.com/kris-nova/kubicorn/cloud/digitalocean/godoSdk"
)

var Sdk *godoSdk.Sdk

type Shared struct {
	CloudID        string
	Name           string `json:"name,omitempty"`
	TagResource    cloud.Resource
	Tags           []string
	CachedActual   cloud.Resource
	CachedExpected cloud.Resource
}

func (s Shared) getCachedActual() cloud.Resource {
	if s.CachedActual != nil {
		return s.CachedActual
	}
	return nil
}

func (s Shared) getCachedExpected() cloud.Resource {
	if s.CachedExpected != nil {
		return s.CachedExpected
	}
	return nil
}
