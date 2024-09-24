package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileLobApp = Win32LobApp{}

type Win32LobApp struct {
	// Contains properties for Windows architecture.
	ApplicableArchitectures *WindowsArchitecture `json:"applicableArchitectures,omitempty"`

	// The command line to install this app
	InstallCommandLine nullable.Type[string] `json:"installCommandLine,omitempty"`

	// The install experience for this app.
	InstallExperience *Win32LobAppInstallExperience `json:"installExperience,omitempty"`

	// The value for the minimum CPU speed which is required to install this app.
	MinimumCpuSpeedInMHz nullable.Type[int64] `json:"minimumCpuSpeedInMHz,omitempty"`

	// The value for the minimum free disk space which is required to install this app.
	MinimumFreeDiskSpaceInMB nullable.Type[int64] `json:"minimumFreeDiskSpaceInMB,omitempty"`

	// The value for the minimum physical memory which is required to install this app.
	MinimumMemoryInMB nullable.Type[int64] `json:"minimumMemoryInMB,omitempty"`

	// The value for the minimum number of processors which is required to install this app.
	MinimumNumberOfProcessors nullable.Type[int64] `json:"minimumNumberOfProcessors,omitempty"`

	// The value for the minimum supported windows release.
	MinimumSupportedWindowsRelease nullable.Type[string] `json:"minimumSupportedWindowsRelease,omitempty"`

	// The MSI details if this Win32 app is an MSI app.
	MsiInformation *Win32LobAppMsiInformation `json:"msiInformation,omitempty"`

	// The return codes for post installation behavior.
	ReturnCodes *[]Win32LobAppReturnCode `json:"returnCodes,omitempty"`

	// The detection and requirement rules for this app.
	Rules *[]Win32LobAppRule `json:"rules,omitempty"`

	// The relative path of the setup file in the encrypted Win32LobApp package.
	SetupFilePath nullable.Type[string] `json:"setupFilePath,omitempty"`

	// The command line to uninstall this app
	UninstallCommandLine nullable.Type[string] `json:"uninstallCommandLine,omitempty"`

	// Fields inherited from MobileLobApp

	// The internal committed content version.
	CommittedContentVersion nullable.Type[string] `json:"committedContentVersion,omitempty"`

	// The list of content versions for this app.
	ContentVersions *[]MobileAppContent `json:"contentVersions,omitempty"`

	// The name of the main Lob application file.
	FileName nullable.Type[string] `json:"fileName,omitempty"`

	// The total size, including all uploaded files.
	Size *int64 `json:"size,omitempty"`

	// Fields inherited from MobileApp

	// The list of group assignments for this mobile app.
	Assignments *[]MobileAppAssignment `json:"assignments,omitempty"`

	// The list of categories for this app.
	Categories *[]MobileAppCategory `json:"categories,omitempty"`

	// The date and time the app was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The description of the app.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The developer of the app.
	Developer nullable.Type[string] `json:"developer,omitempty"`

	// The admin provided or imported title of the app.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The more information Url.
	InformationUrl nullable.Type[string] `json:"informationUrl,omitempty"`

	// The value indicating whether the app is marked as featured by the admin.
	IsFeatured *bool `json:"isFeatured,omitempty"`

	// The large icon, to be displayed in the app details and used for upload of the icon.
	LargeIcon *MimeContent `json:"largeIcon,omitempty"`

	// The date and time the app was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Notes for the app.
	Notes nullable.Type[string] `json:"notes,omitempty"`

	// The owner of the app.
	Owner nullable.Type[string] `json:"owner,omitempty"`

	// The privacy statement Url.
	PrivacyInformationUrl nullable.Type[string] `json:"privacyInformationUrl,omitempty"`

	// The publisher of the app.
	Publisher nullable.Type[string] `json:"publisher,omitempty"`

	// Indicates the publishing state of an app.
	PublishingState *MobileAppPublishingState `json:"publishingState,omitempty"`

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

func (s Win32LobApp) MobileLobApp() BaseMobileLobAppImpl {
	return BaseMobileLobAppImpl{
		CommittedContentVersion: s.CommittedContentVersion,
		ContentVersions:         s.ContentVersions,
		FileName:                s.FileName,
		Size:                    s.Size,
		Assignments:             s.Assignments,
		Categories:              s.Categories,
		CreatedDateTime:         s.CreatedDateTime,
		Description:             s.Description,
		Developer:               s.Developer,
		DisplayName:             s.DisplayName,
		InformationUrl:          s.InformationUrl,
		IsFeatured:              s.IsFeatured,
		LargeIcon:               s.LargeIcon,
		LastModifiedDateTime:    s.LastModifiedDateTime,
		Notes:                   s.Notes,
		Owner:                   s.Owner,
		PrivacyInformationUrl:   s.PrivacyInformationUrl,
		Publisher:               s.Publisher,
		PublishingState:         s.PublishingState,
		Id:                      s.Id,
		ODataId:                 s.ODataId,
		ODataType:               s.ODataType,
	}
}

func (s Win32LobApp) MobileApp() BaseMobileAppImpl {
	return BaseMobileAppImpl{
		Assignments:           s.Assignments,
		Categories:            s.Categories,
		CreatedDateTime:       s.CreatedDateTime,
		Description:           s.Description,
		Developer:             s.Developer,
		DisplayName:           s.DisplayName,
		InformationUrl:        s.InformationUrl,
		IsFeatured:            s.IsFeatured,
		LargeIcon:             s.LargeIcon,
		LastModifiedDateTime:  s.LastModifiedDateTime,
		Notes:                 s.Notes,
		Owner:                 s.Owner,
		PrivacyInformationUrl: s.PrivacyInformationUrl,
		Publisher:             s.Publisher,
		PublishingState:       s.PublishingState,
		Id:                    s.Id,
		ODataId:               s.ODataId,
		ODataType:             s.ODataType,
	}
}

func (s Win32LobApp) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Win32LobApp{}

func (s Win32LobApp) MarshalJSON() ([]byte, error) {
	type wrapper Win32LobApp
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Win32LobApp: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Win32LobApp: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.win32LobApp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Win32LobApp: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Win32LobApp{}

func (s *Win32LobApp) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ApplicableArchitectures        *WindowsArchitecture          `json:"applicableArchitectures,omitempty"`
		InstallCommandLine             nullable.Type[string]         `json:"installCommandLine,omitempty"`
		InstallExperience              *Win32LobAppInstallExperience `json:"installExperience,omitempty"`
		MinimumCpuSpeedInMHz           nullable.Type[int64]          `json:"minimumCpuSpeedInMHz,omitempty"`
		MinimumFreeDiskSpaceInMB       nullable.Type[int64]          `json:"minimumFreeDiskSpaceInMB,omitempty"`
		MinimumMemoryInMB              nullable.Type[int64]          `json:"minimumMemoryInMB,omitempty"`
		MinimumNumberOfProcessors      nullable.Type[int64]          `json:"minimumNumberOfProcessors,omitempty"`
		MinimumSupportedWindowsRelease nullable.Type[string]         `json:"minimumSupportedWindowsRelease,omitempty"`
		MsiInformation                 *Win32LobAppMsiInformation    `json:"msiInformation,omitempty"`
		ReturnCodes                    *[]Win32LobAppReturnCode      `json:"returnCodes,omitempty"`
		SetupFilePath                  nullable.Type[string]         `json:"setupFilePath,omitempty"`
		UninstallCommandLine           nullable.Type[string]         `json:"uninstallCommandLine,omitempty"`
		CommittedContentVersion        nullable.Type[string]         `json:"committedContentVersion,omitempty"`
		ContentVersions                *[]MobileAppContent           `json:"contentVersions,omitempty"`
		FileName                       nullable.Type[string]         `json:"fileName,omitempty"`
		Size                           *int64                        `json:"size,omitempty"`
		Assignments                    *[]MobileAppAssignment        `json:"assignments,omitempty"`
		Categories                     *[]MobileAppCategory          `json:"categories,omitempty"`
		CreatedDateTime                *string                       `json:"createdDateTime,omitempty"`
		Description                    nullable.Type[string]         `json:"description,omitempty"`
		Developer                      nullable.Type[string]         `json:"developer,omitempty"`
		DisplayName                    nullable.Type[string]         `json:"displayName,omitempty"`
		InformationUrl                 nullable.Type[string]         `json:"informationUrl,omitempty"`
		IsFeatured                     *bool                         `json:"isFeatured,omitempty"`
		LargeIcon                      *MimeContent                  `json:"largeIcon,omitempty"`
		LastModifiedDateTime           *string                       `json:"lastModifiedDateTime,omitempty"`
		Notes                          nullable.Type[string]         `json:"notes,omitempty"`
		Owner                          nullable.Type[string]         `json:"owner,omitempty"`
		PrivacyInformationUrl          nullable.Type[string]         `json:"privacyInformationUrl,omitempty"`
		Publisher                      nullable.Type[string]         `json:"publisher,omitempty"`
		PublishingState                *MobileAppPublishingState     `json:"publishingState,omitempty"`
		Id                             *string                       `json:"id,omitempty"`
		ODataId                        *string                       `json:"@odata.id,omitempty"`
		ODataType                      *string                       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ApplicableArchitectures = decoded.ApplicableArchitectures
	s.InstallCommandLine = decoded.InstallCommandLine
	s.InstallExperience = decoded.InstallExperience
	s.MinimumCpuSpeedInMHz = decoded.MinimumCpuSpeedInMHz
	s.MinimumFreeDiskSpaceInMB = decoded.MinimumFreeDiskSpaceInMB
	s.MinimumMemoryInMB = decoded.MinimumMemoryInMB
	s.MinimumNumberOfProcessors = decoded.MinimumNumberOfProcessors
	s.MinimumSupportedWindowsRelease = decoded.MinimumSupportedWindowsRelease
	s.MsiInformation = decoded.MsiInformation
	s.ReturnCodes = decoded.ReturnCodes
	s.SetupFilePath = decoded.SetupFilePath
	s.UninstallCommandLine = decoded.UninstallCommandLine
	s.Assignments = decoded.Assignments
	s.Categories = decoded.Categories
	s.CommittedContentVersion = decoded.CommittedContentVersion
	s.ContentVersions = decoded.ContentVersions
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.Developer = decoded.Developer
	s.DisplayName = decoded.DisplayName
	s.FileName = decoded.FileName
	s.Id = decoded.Id
	s.InformationUrl = decoded.InformationUrl
	s.IsFeatured = decoded.IsFeatured
	s.LargeIcon = decoded.LargeIcon
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Notes = decoded.Notes
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Owner = decoded.Owner
	s.PrivacyInformationUrl = decoded.PrivacyInformationUrl
	s.Publisher = decoded.Publisher
	s.PublishingState = decoded.PublishingState
	s.Size = decoded.Size

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Win32LobApp into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["rules"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Rules into list []json.RawMessage: %+v", err)
		}

		output := make([]Win32LobAppRule, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWin32LobAppRuleImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Rules' for 'Win32LobApp': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Rules = &output
	}

	return nil
}