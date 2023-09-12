// Copyright 2023 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package advisorynotifications provides access to the Advisory Notifications API.
//
// For product documentation, see: https://cloud.google.com/advisory-notifications
//
// # Creating a client
//
// Usage example:
//
//	import "google.golang.org/api/advisorynotifications/v1"
//	...
//	ctx := context.Background()
//	advisorynotificationsService, err := advisorynotifications.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// # Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//	advisorynotificationsService, err := advisorynotifications.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//	config := &oauth2.Config{...}
//	// ...
//	token, err := config.Exchange(ctx, ...)
//	advisorynotificationsService, err := advisorynotifications.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package advisorynotifications // import "google.golang.org/api/advisorynotifications/v1"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	internal "google.golang.org/api/internal"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint

const apiId = "advisorynotifications:v1"
const apiName = "advisorynotifications"
const apiVersion = "v1"
const basePath = "https://advisorynotifications.googleapis.com/"
const mtlsBasePath = "https://advisorynotifications.mtls.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// See, edit, configure, and delete your Google Cloud data and see the
	// email address for your Google Account.
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := internaloption.WithDefaultScopes(
		"https://www.googleapis.com/auth/cloud-platform",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Organizations = NewOrganizationsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Organizations *OrganizationsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewOrganizationsService(s *Service) *OrganizationsService {
	rs := &OrganizationsService{s: s}
	rs.Locations = NewOrganizationsLocationsService(s)
	return rs
}

type OrganizationsService struct {
	s *Service

	Locations *OrganizationsLocationsService
}

func NewOrganizationsLocationsService(s *Service) *OrganizationsLocationsService {
	rs := &OrganizationsLocationsService{s: s}
	rs.Notifications = NewOrganizationsLocationsNotificationsService(s)
	return rs
}

type OrganizationsLocationsService struct {
	s *Service

	Notifications *OrganizationsLocationsNotificationsService
}

func NewOrganizationsLocationsNotificationsService(s *Service) *OrganizationsLocationsNotificationsService {
	rs := &OrganizationsLocationsNotificationsService{s: s}
	return rs
}

type OrganizationsLocationsNotificationsService struct {
	s *Service
}

// GoogleCloudAdvisorynotificationsV1Attachment: Attachment with
// specific information about the issue.
type GoogleCloudAdvisorynotificationsV1Attachment struct {
	// Csv: A CSV file attachment. Max size is 10 MB.
	Csv *GoogleCloudAdvisorynotificationsV1Csv `json:"csv,omitempty"`

	// DisplayName: The title of the attachment.
	DisplayName string `json:"displayName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Csv") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Csv") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudAdvisorynotificationsV1Attachment) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudAdvisorynotificationsV1Attachment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudAdvisorynotificationsV1Csv: A representation of a CSV file
// attachment, as a list of column headers and a list of data rows.
type GoogleCloudAdvisorynotificationsV1Csv struct {
	// DataRows: The list of data rows in a CSV file, as string arrays
	// rather than as a single comma-separated string.
	DataRows []*GoogleCloudAdvisorynotificationsV1CsvCsvRow `json:"dataRows,omitempty"`

	// Headers: The list of headers for data columns in a CSV file.
	Headers []string `json:"headers,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DataRows") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DataRows") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudAdvisorynotificationsV1Csv) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudAdvisorynotificationsV1Csv
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudAdvisorynotificationsV1CsvCsvRow: A representation of a
// single data row in a CSV file.
type GoogleCloudAdvisorynotificationsV1CsvCsvRow struct {
	// Entries: The data entries in a CSV file row, as a string array rather
	// than a single comma-separated string.
	Entries []string `json:"entries,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Entries") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Entries") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudAdvisorynotificationsV1CsvCsvRow) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudAdvisorynotificationsV1CsvCsvRow
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudAdvisorynotificationsV1ListNotificationsResponse: Response
// of ListNotifications endpoint.
type GoogleCloudAdvisorynotificationsV1ListNotificationsResponse struct {
	// NextPageToken: A token, which can be sent as `page_token` to retrieve
	// the next page. If this field is omitted, there are no subsequent
	// pages.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Notifications: List of notifications under a given parent.
	Notifications []*GoogleCloudAdvisorynotificationsV1Notification `json:"notifications,omitempty"`

	// TotalSize: Estimation of a total number of notifications.
	TotalSize int64 `json:"totalSize,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudAdvisorynotificationsV1ListNotificationsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudAdvisorynotificationsV1ListNotificationsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudAdvisorynotificationsV1Message: A message which contains
// notification details.
type GoogleCloudAdvisorynotificationsV1Message struct {
	// Attachments: The attachments to download.
	Attachments []*GoogleCloudAdvisorynotificationsV1Attachment `json:"attachments,omitempty"`

	// Body: The message content.
	Body *GoogleCloudAdvisorynotificationsV1MessageBody `json:"body,omitempty"`

	// CreateTime: The Message creation timestamp.
	CreateTime string `json:"createTime,omitempty"`

	// LocalizationTime: Time when Message was localized
	LocalizationTime string `json:"localizationTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Attachments") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Attachments") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudAdvisorynotificationsV1Message) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudAdvisorynotificationsV1Message
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudAdvisorynotificationsV1MessageBody: A message body
// containing text.
type GoogleCloudAdvisorynotificationsV1MessageBody struct {
	// Text: The text content of the message body.
	Text *GoogleCloudAdvisorynotificationsV1Text `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Text") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Text") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudAdvisorynotificationsV1MessageBody) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudAdvisorynotificationsV1MessageBody
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudAdvisorynotificationsV1Notification: A notification object
// for notifying customers about security and privacy issues.
type GoogleCloudAdvisorynotificationsV1Notification struct {
	// CreateTime: Output only. Time the notification was created.
	CreateTime string `json:"createTime,omitempty"`

	// Messages: A list of messages in the notification.
	Messages []*GoogleCloudAdvisorynotificationsV1Message `json:"messages,omitempty"`

	// Name: The resource name of the notification. Format:
	// organizations/{organization}/locations/{location}/notifications/{notif
	// ication}.
	Name string `json:"name,omitempty"`

	// Subject: The subject line of the notification.
	Subject *GoogleCloudAdvisorynotificationsV1Subject `json:"subject,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudAdvisorynotificationsV1Notification) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudAdvisorynotificationsV1Notification
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudAdvisorynotificationsV1Subject: A subject line of a
// notification.
type GoogleCloudAdvisorynotificationsV1Subject struct {
	// Text: The text content.
	Text *GoogleCloudAdvisorynotificationsV1Text `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Text") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Text") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudAdvisorynotificationsV1Subject) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudAdvisorynotificationsV1Subject
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudAdvisorynotificationsV1Text: A text object containing the
// English text and its localized copies.
type GoogleCloudAdvisorynotificationsV1Text struct {
	// EnText: The English copy.
	EnText string `json:"enText,omitempty"`

	// LocalizationState: Status of the localization.
	//
	// Possible values:
	//   "LOCALIZATION_STATE_UNSPECIFIED" - Not used.
	//   "LOCALIZATION_STATE_NOT_APPLICABLE" - Localization is not
	// applicable for requested language. This can happen when: - The
	// requested language was not supported by Advisory Notifications at the
	// time of localization (including notifications created before the
	// localization feature was launched). - The requested language is
	// English, so only the English text is returned.
	//   "LOCALIZATION_STATE_PENDING" - Localization for requested language
	// is in progress, and not ready yet.
	//   "LOCALIZATION_STATE_COMPLETED" - Localization for requested
	// language is completed.
	LocalizationState string `json:"localizationState,omitempty"`

	// LocalizedText: The requested localized copy (if applicable).
	LocalizedText string `json:"localizedText,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EnText") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EnText") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudAdvisorynotificationsV1Text) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudAdvisorynotificationsV1Text
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "advisorynotifications.organizations.locations.notifications.get":

type OrganizationsLocationsNotificationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a notification.
//
//   - name: A name of the notification to retrieve. Format:
//     organizations/{organization}/locations/{location}/notifications/{not
//     ification}.
func (r *OrganizationsLocationsNotificationsService) Get(name string) *OrganizationsLocationsNotificationsGetCall {
	c := &OrganizationsLocationsNotificationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// LanguageCode sets the optional parameter "languageCode": ISO code for
// requested localization language. If unset, will be interpereted as
// "en". If the requested language is valid, but not supported for this
// notification, English will be returned with an "Not applicable"
// LocalizationState. If the ISO code is invalid (i.e. not a real
// language), this RPC will throw an error.
func (c *OrganizationsLocationsNotificationsGetCall) LanguageCode(languageCode string) *OrganizationsLocationsNotificationsGetCall {
	c.urlParams_.Set("languageCode", languageCode)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsLocationsNotificationsGetCall) Fields(s ...googleapi.Field) *OrganizationsLocationsNotificationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OrganizationsLocationsNotificationsGetCall) IfNoneMatch(entityTag string) *OrganizationsLocationsNotificationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsLocationsNotificationsGetCall) Context(ctx context.Context) *OrganizationsLocationsNotificationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrganizationsLocationsNotificationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsLocationsNotificationsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/"+internal.Version)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "advisorynotifications.organizations.locations.notifications.get" call.
// Exactly one of *GoogleCloudAdvisorynotificationsV1Notification or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudAdvisorynotificationsV1Notification.ServerResponse.Header
// or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *OrganizationsLocationsNotificationsGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudAdvisorynotificationsV1Notification, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &GoogleCloudAdvisorynotificationsV1Notification{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a notification.",
	//   "flatPath": "v1/organizations/{organizationsId}/locations/{locationsId}/notifications/{notificationsId}",
	//   "httpMethod": "GET",
	//   "id": "advisorynotifications.organizations.locations.notifications.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "languageCode": {
	//       "description": "ISO code for requested localization language. If unset, will be interpereted as \"en\". If the requested language is valid, but not supported for this notification, English will be returned with an \"Not applicable\" LocalizationState. If the ISO code is invalid (i.e. not a real language), this RPC will throw an error.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "Required. A name of the notification to retrieve. Format: organizations/{organization}/locations/{location}/notifications/{notification}.",
	//       "location": "path",
	//       "pattern": "^organizations/[^/]+/locations/[^/]+/notifications/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudAdvisorynotificationsV1Notification"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "advisorynotifications.organizations.locations.notifications.list":

type OrganizationsLocationsNotificationsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists notifications under a given parent.
//
//   - parent: The parent, which owns this collection of notifications.
//     Must be of the form
//     "organizations/{organization}/locations/{location}".
func (r *OrganizationsLocationsNotificationsService) List(parent string) *OrganizationsLocationsNotificationsListCall {
	c := &OrganizationsLocationsNotificationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// LanguageCode sets the optional parameter "languageCode": ISO code for
// requested localization language. If unset, will be interpereted as
// "en". If the requested language is valid, but not supported for this
// notification, English will be returned with an "Not applicable"
// LocalizationState. If the ISO code is invalid (i.e. not a real
// language), this RPC will throw an error.
func (c *OrganizationsLocationsNotificationsListCall) LanguageCode(languageCode string) *OrganizationsLocationsNotificationsListCall {
	c.urlParams_.Set("languageCode", languageCode)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of notifications to return. The service may return fewer than this
// value. If unspecified or equal to 0, at most 50 notifications will be
// returned. The maximum value is 50; values above 50 will be coerced to
// 50.
func (c *OrganizationsLocationsNotificationsListCall) PageSize(pageSize int64) *OrganizationsLocationsNotificationsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A page token
// returned from a previous request. When paginating, all other
// parameters provided in the request must match the call that returned
// the page token.
func (c *OrganizationsLocationsNotificationsListCall) PageToken(pageToken string) *OrganizationsLocationsNotificationsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// View sets the optional parameter "view": Specifies which parts of the
// notification resource should be returned in the response.
//
// Possible values:
//
//	"NOTIFICATION_VIEW_UNSPECIFIED" - Not specified, equivalent to
//
// BASIC.
//
//	"BASIC" - Server responses only include title, creation time and
//
// Notification ID. Note: for internal use responses also include the
// last update time, the latest message text and whether notification
// has attachments.
//
//	"FULL" - Include everything.
func (c *OrganizationsLocationsNotificationsListCall) View(view string) *OrganizationsLocationsNotificationsListCall {
	c.urlParams_.Set("view", view)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsLocationsNotificationsListCall) Fields(s ...googleapi.Field) *OrganizationsLocationsNotificationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OrganizationsLocationsNotificationsListCall) IfNoneMatch(entityTag string) *OrganizationsLocationsNotificationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsLocationsNotificationsListCall) Context(ctx context.Context) *OrganizationsLocationsNotificationsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrganizationsLocationsNotificationsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsLocationsNotificationsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/"+internal.Version)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/notifications")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "advisorynotifications.organizations.locations.notifications.list" call.
// Exactly one of
// *GoogleCloudAdvisorynotificationsV1ListNotificationsResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudAdvisorynotificationsV1ListNotificationsResponse.ServerRes
// ponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *OrganizationsLocationsNotificationsListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudAdvisorynotificationsV1ListNotificationsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &GoogleCloudAdvisorynotificationsV1ListNotificationsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists notifications under a given parent.",
	//   "flatPath": "v1/organizations/{organizationsId}/locations/{locationsId}/notifications",
	//   "httpMethod": "GET",
	//   "id": "advisorynotifications.organizations.locations.notifications.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "languageCode": {
	//       "description": "ISO code for requested localization language. If unset, will be interpereted as \"en\". If the requested language is valid, but not supported for this notification, English will be returned with an \"Not applicable\" LocalizationState. If the ISO code is invalid (i.e. not a real language), this RPC will throw an error.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of notifications to return. The service may return fewer than this value. If unspecified or equal to 0, at most 50 notifications will be returned. The maximum value is 50; values above 50 will be coerced to 50.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A page token returned from a previous request. When paginating, all other parameters provided in the request must match the call that returned the page token.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The parent, which owns this collection of notifications. Must be of the form \"organizations/{organization}/locations/{location}\".",
	//       "location": "path",
	//       "pattern": "^organizations/[^/]+/locations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Specifies which parts of the notification resource should be returned in the response.",
	//       "enum": [
	//         "NOTIFICATION_VIEW_UNSPECIFIED",
	//         "BASIC",
	//         "FULL"
	//       ],
	//       "enumDescriptions": [
	//         "Not specified, equivalent to BASIC.",
	//         "Server responses only include title, creation time and Notification ID. Note: for internal use responses also include the last update time, the latest message text and whether notification has attachments.",
	//         "Include everything."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/notifications",
	//   "response": {
	//     "$ref": "GoogleCloudAdvisorynotificationsV1ListNotificationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *OrganizationsLocationsNotificationsListCall) Pages(ctx context.Context, f func(*GoogleCloudAdvisorynotificationsV1ListNotificationsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}
