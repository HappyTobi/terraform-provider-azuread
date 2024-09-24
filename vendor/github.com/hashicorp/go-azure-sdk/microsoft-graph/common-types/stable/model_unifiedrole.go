package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRole struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The unified role definition ID of the directory role. Refer to unifiedRoleDefinition resource.
	RoleDefinitionId *string `json:"roleDefinitionId,omitempty"`
}