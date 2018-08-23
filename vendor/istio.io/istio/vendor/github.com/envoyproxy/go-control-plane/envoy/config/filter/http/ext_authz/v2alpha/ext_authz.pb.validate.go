// Code generated by protoc-gen-validate
// source: envoy/config/filter/http/ext_authz/v2alpha/ext_authz.proto
// DO NOT EDIT!!!

package v2alpha

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// Validate checks the field values on ExtAuthz with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ExtAuthz) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for FailureModeAllow

	switch m.Services.(type) {

	case *ExtAuthz_GrpcService:

		if v, ok := interface{}(m.GetGrpcService()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExtAuthzValidationError{
					Field:  "GrpcService",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *ExtAuthz_HttpService:

		if v, ok := interface{}(m.GetHttpService()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExtAuthzValidationError{
					Field:  "HttpService",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// ExtAuthzValidationError is the validation error returned by
// ExtAuthz.Validate if the designated constraints aren't met.
type ExtAuthzValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ExtAuthzValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExtAuthz.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ExtAuthzValidationError{}

// Validate checks the field values on HttpService with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HttpService) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetServerUri()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpServiceValidationError{
				Field:  "ServerUri",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for PathPrefix

	return nil
}

// HttpServiceValidationError is the validation error returned by
// HttpService.Validate if the designated constraints aren't met.
type HttpServiceValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e HttpServiceValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpService.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = HttpServiceValidationError{}
