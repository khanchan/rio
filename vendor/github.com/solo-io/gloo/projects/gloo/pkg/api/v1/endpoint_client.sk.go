// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type EndpointWatcher interface {
	// watch namespace-scoped Endpoints
	Watch(namespace string, opts clients.WatchOpts) (<-chan EndpointList, <-chan error, error)
}

type EndpointClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*Endpoint, error)
	Write(resource *Endpoint, opts clients.WriteOpts) (*Endpoint, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (EndpointList, error)
	EndpointWatcher
}

type endpointClient struct {
	rc clients.ResourceClient
}

func NewEndpointClient(rcFactory factory.ResourceClientFactory) (EndpointClient, error) {
	return NewEndpointClientWithToken(rcFactory, "")
}

func NewEndpointClientWithToken(rcFactory factory.ResourceClientFactory, token string) (EndpointClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &Endpoint{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base Endpoint resource client")
	}
	return NewEndpointClientWithBase(rc), nil
}

func NewEndpointClientWithBase(rc clients.ResourceClient) EndpointClient {
	return &endpointClient{
		rc: rc,
	}
}

func (client *endpointClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *endpointClient) Register() error {
	return client.rc.Register()
}

func (client *endpointClient) Read(namespace, name string, opts clients.ReadOpts) (*Endpoint, error) {
	opts = opts.WithDefaults()

	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Endpoint), nil
}

func (client *endpointClient) Write(endpoint *Endpoint, opts clients.WriteOpts) (*Endpoint, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(endpoint, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Endpoint), nil
}

func (client *endpointClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()

	return client.rc.Delete(namespace, name, opts)
}

func (client *endpointClient) List(namespace string, opts clients.ListOpts) (EndpointList, error) {
	opts = opts.WithDefaults()

	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToEndpoint(resourceList), nil
}

func (client *endpointClient) Watch(namespace string, opts clients.WatchOpts) (<-chan EndpointList, <-chan error, error) {
	opts = opts.WithDefaults()

	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	endpointsChan := make(chan EndpointList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				endpointsChan <- convertToEndpoint(resourceList)
			case <-opts.Ctx.Done():
				close(endpointsChan)
				return
			}
		}
	}()
	return endpointsChan, errs, nil
}

func convertToEndpoint(resources resources.ResourceList) EndpointList {
	var endpointList EndpointList
	for _, resource := range resources {
		endpointList = append(endpointList, resource.(*Endpoint))
	}
	return endpointList
}
