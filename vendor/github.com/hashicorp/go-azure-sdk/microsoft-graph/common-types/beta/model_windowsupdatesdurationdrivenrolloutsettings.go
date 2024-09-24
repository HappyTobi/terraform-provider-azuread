package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdatesGradualRolloutSettings = WindowsUpdatesDurationDrivenRolloutSettings{}

type WindowsUpdatesDurationDrivenRolloutSettings struct {
	// The target duration of the rollout. Given durationBetweenOffers and durationUntilDeploymentEnd, the system will
	// automatically calculate how many devices are in each offering.
	DurationUntilDeploymentEnd *string `json:"durationUntilDeploymentEnd,omitempty"`

	// Fields inherited from WindowsUpdatesGradualRolloutSettings

	// The duration between each set of devices being offered the update. The value is represented in ISO 8601 format for
	// duration. Default value is P1D (one day).
	DurationBetweenOffers nullable.Type[string] `json:"durationBetweenOffers,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsUpdatesDurationDrivenRolloutSettings) WindowsUpdatesGradualRolloutSettings() BaseWindowsUpdatesGradualRolloutSettingsImpl {
	return BaseWindowsUpdatesGradualRolloutSettingsImpl{
		DurationBetweenOffers: s.DurationBetweenOffers,
		ODataId:               s.ODataId,
		ODataType:             s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesDurationDrivenRolloutSettings{}

func (s WindowsUpdatesDurationDrivenRolloutSettings) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesDurationDrivenRolloutSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesDurationDrivenRolloutSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesDurationDrivenRolloutSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.durationDrivenRolloutSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesDurationDrivenRolloutSettings: %+v", err)
	}

	return encoded, nil
}