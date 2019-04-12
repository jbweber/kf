// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was generated with option-builder.go, DO NOT EDIT IT.

package secrets

type getConfig struct {
	// Namespace is the Kubernetes namespace to use
	Namespace string
}

// GetOption is a single option for configuring a getConfig
type GetOption func(*getConfig)

// GetOptions is a configuration set defining a getConfig
type GetOptions []GetOption

// toConfig applies all the options to a new getConfig and returns it.
func (opts GetOptions) toConfig() getConfig {
	cfg := getConfig{}

	for _, v := range opts {
		v(&cfg)
	}

	return cfg
}

// Extend creates a new GetOptions with the contents of other overriding
// the values set in this GetOptions.
func (opts GetOptions) Extend(other GetOptions) GetOptions {
	var out GetOptions
	out = append(out, opts...)
	out = append(out, other...)
	return out
}

// Namespace returns the last set value for Namespace or the empty value
// if not set.
func (opts GetOptions) Namespace() string {
	return opts.toConfig().Namespace
}

// WithGetNamespace creates an Option that sets the Kubernetes namespace to use
func WithGetNamespace(val string) GetOption {
	return func(cfg *getConfig) {
		cfg.Namespace = val
	}
}

// GetOptionDefaults gets the default values for Get.
func GetOptionDefaults() GetOptions {
	return GetOptions{
		WithGetNamespace("default"),
	}
}
