/*
Copyright 2022 k0s authors

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

package v1beta2

import (
	"context"
	"time"

	v1beta2 "github.com/k0sproject/k0s/pkg/apis/autopilot.k0sproject.io/v1beta2"
	scheme "github.com/k0sproject/k0s/pkg/apis/autopilot.k0sproject.io/v1beta2/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// UpdateConfigsGetter has a method to return a UpdateConfigInterface.
// A group's client should implement this interface.
type UpdateConfigsGetter interface {
	UpdateConfigs() UpdateConfigInterface
}

// UpdateConfigInterface has methods to work with UpdateConfig resources.
type UpdateConfigInterface interface {
	Create(ctx context.Context, updateConfig *v1beta2.UpdateConfig, opts v1.CreateOptions) (*v1beta2.UpdateConfig, error)
	Update(ctx context.Context, updateConfig *v1beta2.UpdateConfig, opts v1.UpdateOptions) (*v1beta2.UpdateConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta2.UpdateConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta2.UpdateConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	UpdateConfigExpansion
}

// updateConfigs implements UpdateConfigInterface
type updateConfigs struct {
	client rest.Interface
}

// newUpdateConfigs returns a UpdateConfigs
func newUpdateConfigs(c *AutopilotV1beta2Client) *updateConfigs {
	return &updateConfigs{
		client: c.RESTClient(),
	}
}

// Get takes name of the updateConfig, and returns the corresponding updateConfig object, and an error if there is any.
func (c *updateConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.UpdateConfig, err error) {
	result = &v1beta2.UpdateConfig{}
	err = c.client.Get().
		Resource("updateconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of UpdateConfigs that match those selectors.
func (c *updateConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.UpdateConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta2.UpdateConfigList{}
	err = c.client.Get().
		Resource("updateconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested updateConfigs.
func (c *updateConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("updateconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a updateConfig and creates it.  Returns the server's representation of the updateConfig, and an error, if there is any.
func (c *updateConfigs) Create(ctx context.Context, updateConfig *v1beta2.UpdateConfig, opts v1.CreateOptions) (result *v1beta2.UpdateConfig, err error) {
	result = &v1beta2.UpdateConfig{}
	err = c.client.Post().
		Resource("updateconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(updateConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a updateConfig and updates it. Returns the server's representation of the updateConfig, and an error, if there is any.
func (c *updateConfigs) Update(ctx context.Context, updateConfig *v1beta2.UpdateConfig, opts v1.UpdateOptions) (result *v1beta2.UpdateConfig, err error) {
	result = &v1beta2.UpdateConfig{}
	err = c.client.Put().
		Resource("updateconfigs").
		Name(updateConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(updateConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the updateConfig and deletes it. Returns an error if one occurs.
func (c *updateConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("updateconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}
