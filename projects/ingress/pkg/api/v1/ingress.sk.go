// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"sort"

	"github.com/gogo/protobuf/proto"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/hashutils"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TODO: modify as needed to populate additional fields
func NewIngress(namespace, name string) *Ingress {
	return &Ingress{
		Metadata: core.Metadata{
			Name:      name,
			Namespace: namespace,
		},
	}
}

func (r *Ingress) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *Ingress) Hash() uint64 {
	metaCopy := r.GetMetadata()
	metaCopy.ResourceVersion = ""
	return hashutils.HashAll(
		metaCopy,
		r.KubeIngressSpec,
	)
}

type IngressList []*Ingress
type IngressesByNamespace map[string]IngressList

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list IngressList) Find(namespace, name string) (*Ingress, error) {
	for _, ingress := range list {
		if ingress.Metadata.Name == name {
			if namespace == "" || ingress.Metadata.Namespace == namespace {
				return ingress, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find ingress %v.%v", namespace, name)
}

func (list IngressList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, ingress := range list {
		ress = append(ress, ingress)
	}
	return ress
}

func (list IngressList) Names() []string {
	var names []string
	for _, ingress := range list {
		names = append(names, ingress.Metadata.Name)
	}
	return names
}

func (list IngressList) NamespacesDotNames() []string {
	var names []string
	for _, ingress := range list {
		names = append(names, ingress.Metadata.Namespace+"."+ingress.Metadata.Name)
	}
	return names
}

func (list IngressList) Sort() IngressList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Metadata.Less(list[j].Metadata)
	})
	return list
}

func (list IngressList) Clone() IngressList {
	var ingressList IngressList
	for _, ingress := range list {
		ingressList = append(ingressList, proto.Clone(ingress).(*Ingress))
	}
	return ingressList
}

func (list IngressList) Each(f func(element *Ingress)) {
	for _, ingress := range list {
		f(ingress)
	}
}

func (list IngressList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *Ingress) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

func (list IngressList) ByNamespace() IngressesByNamespace {
	byNamespace := make(IngressesByNamespace)
	for _, ingress := range list {
		byNamespace.Add(ingress)
	}
	return byNamespace
}

func (byNamespace IngressesByNamespace) Add(ingress ...*Ingress) {
	for _, item := range ingress {
		byNamespace[item.Metadata.Namespace] = append(byNamespace[item.Metadata.Namespace], item)
	}
}

func (byNamespace IngressesByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace IngressesByNamespace) List() IngressList {
	var list IngressList
	for _, ingressList := range byNamespace {
		list = append(list, ingressList...)
	}
	return list.Sort()
}

func (byNamespace IngressesByNamespace) Clone() IngressesByNamespace {
	return byNamespace.List().Clone().ByNamespace()
}

var _ resources.Resource = &Ingress{}

// Kubernetes Adapter for Ingress

func (o *Ingress) GetObjectKind() schema.ObjectKind {
	t := IngressCrd.TypeMeta()
	return &t
}

func (o *Ingress) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*Ingress)
}

var IngressCrd = crd.NewCrd("ingress.solo.io",
	"ingresses",
	"ingress.solo.io",
	"v1",
	"Ingress",
	"ig",
	&Ingress{})
