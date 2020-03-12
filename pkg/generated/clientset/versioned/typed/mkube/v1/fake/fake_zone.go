// This file is part of MinIO Kubernetes Cloud
// Copyright (c) 2019 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	mkubev1 "github.com/minio/m3/pkg/apis/mkube/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeZones implements ZoneInterface
type FakeZones struct {
	Fake *FakeMkubeV1
	ns   string
}

var zonesResource = schema.GroupVersionResource{Group: "mkube.min.io", Version: "v1", Resource: "zones"}

var zonesKind = schema.GroupVersionKind{Group: "mkube.min.io", Version: "v1", Kind: "Zone"}

// Get takes name of the zone, and returns the corresponding zone object, and an error if there is any.
func (c *FakeZones) Get(name string, options v1.GetOptions) (result *mkubev1.Zone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(zonesResource, c.ns, name), &mkubev1.Zone{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mkubev1.Zone), err
}

// List takes label and field selectors, and returns the list of Zones that match those selectors.
func (c *FakeZones) List(opts v1.ListOptions) (result *mkubev1.ZoneList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(zonesResource, zonesKind, c.ns, opts), &mkubev1.ZoneList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &mkubev1.ZoneList{ListMeta: obj.(*mkubev1.ZoneList).ListMeta}
	for _, item := range obj.(*mkubev1.ZoneList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested zones.
func (c *FakeZones) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(zonesResource, c.ns, opts))

}

// Create takes the representation of a zone and creates it.  Returns the server's representation of the zone, and an error, if there is any.
func (c *FakeZones) Create(zone *mkubev1.Zone) (result *mkubev1.Zone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(zonesResource, c.ns, zone), &mkubev1.Zone{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mkubev1.Zone), err
}

// Update takes the representation of a zone and updates it. Returns the server's representation of the zone, and an error, if there is any.
func (c *FakeZones) Update(zone *mkubev1.Zone) (result *mkubev1.Zone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(zonesResource, c.ns, zone), &mkubev1.Zone{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mkubev1.Zone), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeZones) UpdateStatus(zone *mkubev1.Zone) (*mkubev1.Zone, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(zonesResource, "status", c.ns, zone), &mkubev1.Zone{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mkubev1.Zone), err
}

// Delete takes name of the zone and deletes it. Returns an error if one occurs.
func (c *FakeZones) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(zonesResource, c.ns, name), &mkubev1.Zone{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeZones) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(zonesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &mkubev1.ZoneList{})
	return err
}

// Patch applies the patch and returns the patched zone.
func (c *FakeZones) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mkubev1.Zone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(zonesResource, c.ns, name, pt, data, subresources...), &mkubev1.Zone{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mkubev1.Zone), err
}
