// Code generated by solo-kit. DO NOT EDIT.

package kubernetes

import (
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionServiceFunc func(original, desired *Service) (bool, error)

type ServiceReconciler interface {
	Reconcile(namespace string, desiredResources ServiceList, transition TransitionServiceFunc, opts clients.ListOpts) error
}

func servicesToResources(list ServiceList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, service := range list {
		resourceList = append(resourceList, service)
	}
	return resourceList
}

func NewServiceReconciler(client ServiceClient) ServiceReconciler {
	return &serviceReconciler{
		base: reconcile.NewReconciler(client.BaseClient()),
	}
}

type serviceReconciler struct {
	base reconcile.Reconciler
}

func (r *serviceReconciler) Reconcile(namespace string, desiredResources ServiceList, transition TransitionServiceFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "service_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*Service), desired.(*Service))
		}
	}
	return r.base.Reconcile(namespace, servicesToResources(desiredResources), transitionResources, opts)
}
