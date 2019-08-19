// TODO: Docs
package client

import (
	api "github.com/weaveworks/gitops-toolkit/cmd/sample-app/apis/sample"
	meta "github.com/weaveworks/gitops-toolkit/pkg/apis/meta/v1alpha1"
	"github.com/weaveworks/gitops-toolkit/pkg/client"
	"github.com/weaveworks/gitops-toolkit/pkg/storage"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TODO: Autogenerate this!

// NewClient creates a client for the specified storage
func NewClient(s storage.Storage) *Client {
	return &Client{
		SampleInternalClient: NewSampleInternalClient(s),
	}
}

// Client is a struct providing high-level access to objects in a storage
// The resource-specific client interfaces are automatically generated based
// off client_resource_template.go. The auto-generation can be done with hack/client.sh
// At the moment SampleInternalClient is the default client. If more than this client
// is created in the future, the SampleInternalClient will be accessible under
// Client.SampleInternal() instead.
type Client struct {
	*SampleInternalClient
}

func NewSampleInternalClient(s storage.Storage) *SampleInternalClient {
	return &SampleInternalClient{
		storage:        s,
		dynamicClients: map[schema.GroupVersionKind]client.DynamicClient{},
		gv:             api.SchemeGroupVersion,
	}
}

type SampleInternalClient struct {
	storage          storage.Storage
	gv               schema.GroupVersion
	carClient        CarClient
	motorcycleClient MotorcycleClient
	dynamicClients   map[schema.GroupVersionKind]client.DynamicClient
}

// Dynamic returns the DynamicClient for the Client instance, for the specific kind
func (c *SampleInternalClient) Dynamic(kind meta.Kind) (dc client.DynamicClient) {
	var ok bool
	gvk := c.gv.WithKind(kind.Title())
	if dc, ok = c.dynamicClients[gvk]; !ok {
		dc = client.NewDynamicClient(c.storage, gvk)
		c.dynamicClients[gvk] = dc
	}

	return
}