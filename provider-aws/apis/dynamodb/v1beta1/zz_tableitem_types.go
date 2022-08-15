/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TableItemObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TableItemParameters struct {

	// Hash key to use for lookups and identification of the item
	// +kubebuilder:validation:Required
	HashKey *string `json:"hashKey" tf:"hash_key,omitempty"`

	// JSON representation of a map of attribute name/value pairs, one for each attribute. Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
	// +kubebuilder:validation:Required
	Item *string `json:"item" tf:"item,omitempty"`

	// Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
	// +kubebuilder:validation:Optional
	RangeKey *string `json:"rangeKey,omitempty" tf:"range_key,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Name of the table to contain the item.
	// +crossplane:generate:reference:type=Table
	// +kubebuilder:validation:Optional
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`

	// +kubebuilder:validation:Optional
	TableNameRef *v1.Reference `json:"tableNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TableNameSelector *v1.Selector `json:"tableNameSelector,omitempty" tf:"-"`
}

// TableItemSpec defines the desired state of TableItem
type TableItemSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableItemParameters `json:"forProvider"`
}

// TableItemStatus defines the observed state of TableItem.
type TableItemStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableItemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TableItem is the Schema for the TableItems API. Provides a DynamoDB table item resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TableItem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TableItemSpec   `json:"spec"`
	Status            TableItemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableItemList contains a list of TableItems
type TableItemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TableItem `json:"items"`
}

// Repository type metadata.
var (
	TableItem_Kind             = "TableItem"
	TableItem_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TableItem_Kind}.String()
	TableItem_KindAPIVersion   = TableItem_Kind + "." + CRDGroupVersion.String()
	TableItem_GroupVersionKind = CRDGroupVersion.WithKind(TableItem_Kind)
)

func init() {
	SchemeBuilder.Register(&TableItem{}, &TableItemList{})
}
