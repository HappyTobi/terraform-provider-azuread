package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationFeatureSummary struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Total number of users accounts, excluding those that are blocked.
	TotalUserCount *int64 `json:"totalUserCount,omitempty"`

	// Number of users registered or capable for multifactor authentication, self-service password reset, and passwordless
	// authentication.
	UserRegistrationFeatureCounts *[]UserRegistrationFeatureCount `json:"userRegistrationFeatureCounts,omitempty"`

	// The role type of the user. Possible values are: all, privilegedAdmin, admin, user, unknownFutureValue.
	UserRoles *IncludedUserRoles `json:"userRoles,omitempty"`

	// User type. Possible values are: all, member, guest, unknownFutureValue.
	UserTypes *IncludedUserTypes `json:"userTypes,omitempty"`
}