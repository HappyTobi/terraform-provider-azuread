package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId{}

// GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId is a struct representing the Resource ID for a Group Id Calendar Calendar View Id Exception Occurrence Id Attachment
type GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId struct {
	GroupId      string
	EventId      string
	EventId1     string
	AttachmentId string
}

// NewGroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentID returns a new GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId struct
func NewGroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentID(groupId string, eventId string, eventId1 string, attachmentId string) GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId {
	return GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId{
		GroupId:      groupId,
		EventId:      eventId,
		EventId1:     eventId1,
		AttachmentId: attachmentId,
	}
}

// ParseGroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentID parses 'input' into a GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId
func ParseGroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentID(input string) (*GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentIDInsensitively parses 'input' case-insensitively into a GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseGroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentIDInsensitively(input string) (*GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateGroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentID checks that 'input' can be parsed as a Group Id Calendar Calendar View Id Exception Occurrence Id Attachment ID
func ValidateGroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Calendar Calendar View Id Exception Occurrence Id Attachment ID
func (id GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId) ID() string {
	fmtString := "/groups/%s/calendar/calendarView/%s/exceptionOccurrences/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Calendar Calendar View Id Exception Occurrence Id Attachment ID
func (id GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this Group Id Calendar Calendar View Id Exception Occurrence Id Attachment ID
func (id GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Group Id Calendar Calendar View Id Exception Occurrence Id Attachment (%s)", strings.Join(components, "\n"))
}