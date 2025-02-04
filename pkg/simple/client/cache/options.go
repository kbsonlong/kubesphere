/*
Copyright 2019 The KubeSphere Authors.

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

package cache

import (
	"fmt"

	"kubesphere.io/kubesphere/pkg/server/options"
)

type Options struct {
	Type    string                 `json:"type"`
	Options options.DynamicOptions `json:"options"`
}

// NewCacheOptions returns options points to nowhere,
// because redis is not required for some components
func NewCacheOptions() *Options {
	return &Options{
		Type:    "",
		Options: map[string]interface{}{},
	}
}

// Validate check options
func (r *Options) Validate() []error {
	errors := make([]error, 0)

	if r.Type == "" {
		errors = append(errors, fmt.Errorf("invalid cache type"))
	}

	return errors
}
