package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventMessageDetail = MembersLeftEventMessageDetail{}

type MembersLeftEventMessageDetail struct {
	// Initiator of the event.
	Initiator IdentitySet `json:"initiator"`

	// List of members who left the chat.
	Members *[]TeamworkUserIdentity `json:"members,omitempty"`

	// Fields inherited from EventMessageDetail

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s MembersLeftEventMessageDetail) EventMessageDetail() BaseEventMessageDetailImpl {
	return BaseEventMessageDetailImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MembersLeftEventMessageDetail{}

func (s MembersLeftEventMessageDetail) MarshalJSON() ([]byte, error) {
	type wrapper MembersLeftEventMessageDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MembersLeftEventMessageDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MembersLeftEventMessageDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.membersLeftEventMessageDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MembersLeftEventMessageDetail: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MembersLeftEventMessageDetail{}

func (s *MembersLeftEventMessageDetail) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Members   *[]TeamworkUserIdentity `json:"members,omitempty"`
		ODataId   *string                 `json:"@odata.id,omitempty"`
		ODataType *string                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Members = decoded.Members
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MembersLeftEventMessageDetail into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["initiator"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Initiator' for 'MembersLeftEventMessageDetail': %+v", err)
		}
		s.Initiator = impl
	}

	return nil
}
