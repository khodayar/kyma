package publish

import (
	"regexp"
	"time"
)

var (
	isValidEventID = regexp.MustCompile(AllowedEventIDChars).MatchString

	// fully-qualified topic name components
	isValidEventType        = regexp.MustCompile(AllowedEventTypeChars).MatchString
	isValidEventTypeVersion = regexp.MustCompile(AllowedEventTypeVersionChars).MatchString

	isValidSourceId = regexp.MustCompile(AllowedSourceIdChars).MatchString
)

//ValidatePublish validates a publish POST request
func ValidatePublish(r *PublishRequest) *Error {
	if len(r.SourceID) == 0 {
		return ErrorResponseMissingFieldSourceId()
	}
	if len(r.EventType) == 0 {
		return ErrorResponseMissingFieldEventType()
	}
	if len(r.EventTypeVersion) == 0 {
		return ErrorResponseMissingFieldEventTypeVersion()
	}
	if len(r.EventTime) == 0 {
		return ErrorResponseMissingFieldEventTime()
	}
	if r.Data == nil {
		return ErrorResponseMissingFieldData()
	} else if d, ok := (r.Data).(string); ok && d == "" {
		return ErrorResponseMissingFieldData()
	}

	// validate the fully-qualified topic name components
	if !isValidSourceId(r.SourceID){
		return ErrorResponseWrongSourceId(r.SourceIdFromHeader)
	}
	if !isValidEventType(r.EventType) {
		return ErrorResponseWrongEventType()
	}
	if !isValidEventTypeVersion(r.EventTypeVersion) {
		return ErrorResponseWrongEventTypeVersion()
	}

	if _, err := time.Parse(time.RFC3339, r.EventTime); err != nil {
		return ErrorResponseWrongEventTime(err)
	}
	if len(r.EventID) > 0 && !isValidEventID(r.EventID) {
		return ErrorResponseWrongEventId()
	}
	return nil
}
