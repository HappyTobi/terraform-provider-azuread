package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdInferenceClassificationOverrideId{}

// UserIdInferenceClassificationOverrideId is a struct representing the Resource ID for a User Id Inference Classification Override
type UserIdInferenceClassificationOverrideId struct {
	UserId                            string
	InferenceClassificationOverrideId string
}

// NewUserIdInferenceClassificationOverrideID returns a new UserIdInferenceClassificationOverrideId struct
func NewUserIdInferenceClassificationOverrideID(userId string, inferenceClassificationOverrideId string) UserIdInferenceClassificationOverrideId {
	return UserIdInferenceClassificationOverrideId{
		UserId:                            userId,
		InferenceClassificationOverrideId: inferenceClassificationOverrideId,
	}
}

// ParseUserIdInferenceClassificationOverrideID parses 'input' into a UserIdInferenceClassificationOverrideId
func ParseUserIdInferenceClassificationOverrideID(input string) (*UserIdInferenceClassificationOverrideId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInferenceClassificationOverrideId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInferenceClassificationOverrideId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdInferenceClassificationOverrideIDInsensitively parses 'input' case-insensitively into a UserIdInferenceClassificationOverrideId
// note: this method should only be used for API response data and not user input
func ParseUserIdInferenceClassificationOverrideIDInsensitively(input string) (*UserIdInferenceClassificationOverrideId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInferenceClassificationOverrideId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInferenceClassificationOverrideId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdInferenceClassificationOverrideId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.InferenceClassificationOverrideId, ok = input.Parsed["inferenceClassificationOverrideId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "inferenceClassificationOverrideId", input)
	}

	return nil
}

// ValidateUserIdInferenceClassificationOverrideID checks that 'input' can be parsed as a User Id Inference Classification Override ID
func ValidateUserIdInferenceClassificationOverrideID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdInferenceClassificationOverrideID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Inference Classification Override ID
func (id UserIdInferenceClassificationOverrideId) ID() string {
	fmtString := "/users/%s/inferenceClassification/overrides/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.InferenceClassificationOverrideId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Inference Classification Override ID
func (id UserIdInferenceClassificationOverrideId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("inferenceClassification", "inferenceClassification", "inferenceClassification"),
		resourceids.StaticSegment("overrides", "overrides", "overrides"),
		resourceids.UserSpecifiedSegment("inferenceClassificationOverrideId", "inferenceClassificationOverrideId"),
	}
}

// String returns a human-readable description of this User Id Inference Classification Override ID
func (id UserIdInferenceClassificationOverrideId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Inference Classification Override: %q", id.InferenceClassificationOverrideId),
	}
	return fmt.Sprintf("User Id Inference Classification Override (%s)", strings.Join(components, "\n"))
}