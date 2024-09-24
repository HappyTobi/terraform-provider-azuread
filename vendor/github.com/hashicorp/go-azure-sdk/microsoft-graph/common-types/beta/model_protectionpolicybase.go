package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectionPolicyBase interface {
	Entity
	ProtectionPolicyBase() BaseProtectionPolicyBaseImpl
}

var _ ProtectionPolicyBase = BaseProtectionPolicyBaseImpl{}

type BaseProtectionPolicyBaseImpl struct {
	// The identity of person who created the policy.
	CreatedBy IdentitySet `json:"createdBy"`

	// The time of creation of the policy.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The name of the policy to be created.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The identity of the person who last modified the policy.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The timestamp of the last modification of the policy.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Contains the retention setting details for the policy.
	RetentionSettings *[]RetentionSetting `json:"retentionSettings,omitempty"`

	// The aggregated status of the protection units associated with the policy. The possible values are: inactive,
	// activeWithErrors, updating, active, unknownFutureValue.
	Status *ProtectionPolicyStatus `json:"status,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseProtectionPolicyBaseImpl) ProtectionPolicyBase() BaseProtectionPolicyBaseImpl {
	return s
}

func (s BaseProtectionPolicyBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ProtectionPolicyBase = RawProtectionPolicyBaseImpl{}

// RawProtectionPolicyBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawProtectionPolicyBaseImpl struct {
	protectionPolicyBase BaseProtectionPolicyBaseImpl
	Type                 string
	Values               map[string]interface{}
}

func (s RawProtectionPolicyBaseImpl) ProtectionPolicyBase() BaseProtectionPolicyBaseImpl {
	return s.protectionPolicyBase
}

func (s RawProtectionPolicyBaseImpl) Entity() BaseEntityImpl {
	return s.protectionPolicyBase.Entity()
}

var _ json.Marshaler = BaseProtectionPolicyBaseImpl{}

func (s BaseProtectionPolicyBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseProtectionPolicyBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseProtectionPolicyBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseProtectionPolicyBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.protectionPolicyBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseProtectionPolicyBaseImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseProtectionPolicyBaseImpl{}

func (s *BaseProtectionPolicyBaseImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime      nullable.Type[string]   `json:"createdDateTime,omitempty"`
		DisplayName          nullable.Type[string]   `json:"displayName,omitempty"`
		LastModifiedDateTime nullable.Type[string]   `json:"lastModifiedDateTime,omitempty"`
		RetentionSettings    *[]RetentionSetting     `json:"retentionSettings,omitempty"`
		Status               *ProtectionPolicyStatus `json:"status,omitempty"`
		Id                   *string                 `json:"id,omitempty"`
		ODataId              *string                 `json:"@odata.id,omitempty"`
		ODataType            *string                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.RetentionSettings = decoded.RetentionSettings
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseProtectionPolicyBaseImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'BaseProtectionPolicyBaseImpl': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BaseProtectionPolicyBaseImpl': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}

func UnmarshalProtectionPolicyBaseImplementation(input []byte) (ProtectionPolicyBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ProtectionPolicyBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.exchangeProtectionPolicy") {
		var out ExchangeProtectionPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExchangeProtectionPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.oneDriveForBusinessProtectionPolicy") {
		var out OneDriveForBusinessProtectionPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OneDriveForBusinessProtectionPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sharePointProtectionPolicy") {
		var out SharePointProtectionPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharePointProtectionPolicy: %+v", err)
		}
		return out, nil
	}

	var parent BaseProtectionPolicyBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseProtectionPolicyBaseImpl: %+v", err)
	}

	return RawProtectionPolicyBaseImpl{
		protectionPolicyBase: parent,
		Type:                 value,
		Values:               temp,
	}, nil

}