package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomCalloutExtension interface {
	Entity
	CustomCalloutExtension() BaseCustomCalloutExtensionImpl
}

var _ CustomCalloutExtension = BaseCustomCalloutExtensionImpl{}

type BaseCustomCalloutExtensionImpl struct {
	// Configuration for securing the API call to the logic app. For example, using OAuth client credentials flow.
	AuthenticationConfiguration CustomExtensionAuthenticationConfiguration `json:"authenticationConfiguration"`

	// HTTP connection settings that define how long Microsoft Entra ID can wait for a connection to a logic app, how many
	// times you can retry a timed-out connection and the exception scenarios when retries are allowed.
	ClientConfiguration *CustomExtensionClientConfiguration `json:"clientConfiguration,omitempty"`

	// Description for the customCalloutExtension object.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display name for the customCalloutExtension object.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The type and details for configuring the endpoint to call the logic app's workflow.
	EndpointConfiguration CustomExtensionEndpointConfiguration `json:"endpointConfiguration"`

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

func (s BaseCustomCalloutExtensionImpl) CustomCalloutExtension() BaseCustomCalloutExtensionImpl {
	return s
}

func (s BaseCustomCalloutExtensionImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ CustomCalloutExtension = RawCustomCalloutExtensionImpl{}

// RawCustomCalloutExtensionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCustomCalloutExtensionImpl struct {
	customCalloutExtension BaseCustomCalloutExtensionImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawCustomCalloutExtensionImpl) CustomCalloutExtension() BaseCustomCalloutExtensionImpl {
	return s.customCalloutExtension
}

func (s RawCustomCalloutExtensionImpl) Entity() BaseEntityImpl {
	return s.customCalloutExtension.Entity()
}

var _ json.Marshaler = BaseCustomCalloutExtensionImpl{}

func (s BaseCustomCalloutExtensionImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseCustomCalloutExtensionImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseCustomCalloutExtensionImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseCustomCalloutExtensionImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.customCalloutExtension"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseCustomCalloutExtensionImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseCustomCalloutExtensionImpl{}

func (s *BaseCustomCalloutExtensionImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ClientConfiguration *CustomExtensionClientConfiguration `json:"clientConfiguration,omitempty"`
		Description         nullable.Type[string]               `json:"description,omitempty"`
		DisplayName         nullable.Type[string]               `json:"displayName,omitempty"`
		Id                  *string                             `json:"id,omitempty"`
		ODataId             *string                             `json:"@odata.id,omitempty"`
		ODataType           *string                             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ClientConfiguration = decoded.ClientConfiguration
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseCustomCalloutExtensionImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authenticationConfiguration"]; ok {
		impl, err := UnmarshalCustomExtensionAuthenticationConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AuthenticationConfiguration' for 'BaseCustomCalloutExtensionImpl': %+v", err)
		}
		s.AuthenticationConfiguration = impl
	}

	if v, ok := temp["endpointConfiguration"]; ok {
		impl, err := UnmarshalCustomExtensionEndpointConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'EndpointConfiguration' for 'BaseCustomCalloutExtensionImpl': %+v", err)
		}
		s.EndpointConfiguration = impl
	}

	return nil
}

func UnmarshalCustomCalloutExtensionImplementation(input []byte) (CustomCalloutExtension, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomCalloutExtension into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageAssignmentRequestWorkflowExtension") {
		var out AccessPackageAssignmentRequestWorkflowExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageAssignmentRequestWorkflowExtension: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageAssignmentWorkflowExtension") {
		var out AccessPackageAssignmentWorkflowExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageAssignmentWorkflowExtension: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customAuthenticationExtension") {
		var out CustomAuthenticationExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomAuthenticationExtension: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.customTaskExtension") {
		var out IdentityGovernanceCustomTaskExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceCustomTaskExtension: %+v", err)
		}
		return out, nil
	}

	var parent BaseCustomCalloutExtensionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCustomCalloutExtensionImpl: %+v", err)
	}

	return RawCustomCalloutExtensionImpl{
		customCalloutExtension: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}