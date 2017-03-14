package blobauditingpolicies

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
)

// BlobAuditingPolicyState enumerates the values for blob auditing policy
// state.
type BlobAuditingPolicyState string

const (
    // Disabled specifies the disabled state for blob auditing policy state.
    Disabled BlobAuditingPolicyState = "Disabled"
    // Enabled specifies the enabled state for blob auditing policy state.
    Enabled BlobAuditingPolicyState = "Enabled"
)

// DatabaseBlobAuditingPolicy is contains information about a database Blob
// Auditing policy.
type DatabaseBlobAuditingPolicy struct {
    autorest.Response `json:"-"`
    ID *string `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Type *string `json:"type,omitempty"`
    Location *string `json:"location,omitempty"`
    Kind *string `json:"kind,omitempty"`
    *DatabaseBlobAuditingPolicyProperties `json:"properties,omitempty"`
}

// DatabaseBlobAuditingPolicyProperties is properties for a database Blob
// Auditing policy.
type DatabaseBlobAuditingPolicyProperties struct {
    State BlobAuditingPolicyState `json:"state,omitempty"`
    StorageEndpoint *string `json:"storageEndpoint,omitempty"`
    StorageAccountAccessKey *string `json:"storageAccountAccessKey,omitempty"`
    RetentionDays *int32 `json:"retentionDays,omitempty"`
    AuditActionsAndGroups *[]string `json:"auditActionsAndGroups,omitempty"`
    StorageAccountSubscriptionID *string `json:"storageAccountSubscriptionId,omitempty"`
    IsStorageSecondaryKeyInUse *bool `json:"isStorageSecondaryKeyInUse,omitempty"`
}

// ProxyResource is aRM proxy resource.
type ProxyResource struct {
    ID *string `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Type *string `json:"type,omitempty"`
}

// Resource is aRM resource.
type Resource struct {
    ID *string `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Type *string `json:"type,omitempty"`
}

// TrackedResource is aRM tracked top level resource.
type TrackedResource struct {
    ID *string `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Type *string `json:"type,omitempty"`
    Tags *map[string]*string `json:"tags,omitempty"`
    Location *string `json:"location,omitempty"`
}
