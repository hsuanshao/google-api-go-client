// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package commentanalyzer provides access to the Perspective Comment Analyzer API.
//
// For product documentation, see: https://github.com/conversationai/perspectiveapi/blob/master/README.md
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/commentanalyzer/v1alpha1"
//   ...
//   ctx := context.Background()
//   commentanalyzerService, err := commentanalyzer.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   commentanalyzerService, err := commentanalyzer.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   commentanalyzerService, err := commentanalyzer.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package commentanalyzer // import "google.golang.org/api/commentanalyzer/v1alpha1"

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

const apiId = "commentanalyzer:v1alpha1"
const apiName = "commentanalyzer"
const apiVersion = "v1alpha1"
const basePath = "https://commentanalyzer.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View your email address
	UserinfoEmailScope = "https://www.googleapis.com/auth/userinfo.email"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := option.WithScopes(
		"https://www.googleapis.com/auth/userinfo.email",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
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
	s.Comments = NewCommentsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Comments *CommentsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewCommentsService(s *Service) *CommentsService {
	rs := &CommentsService{s: s}
	return rs
}

type CommentsService struct {
	s *Service
}

// AnalyzeCommentRequest: The comment analysis request
// message.
// LINT.IfChange
type AnalyzeCommentRequest struct {
	// ClientToken: Opaque token that is echoed from the request to the
	// response.
	ClientToken string `json:"clientToken,omitempty"`

	// Comment: The comment to analyze.
	Comment *TextEntry `json:"comment,omitempty"`

	// CommunityId: Optional identifier associating this
	// AnalyzeCommentRequest with a
	// particular client's community. Different communities may have
	// different
	// norms and rules. Specifying this value enables us to explore
	// building
	// community-specific models for clients.
	CommunityId string `json:"communityId,omitempty"`

	// Context: The context of the comment.
	Context *Context `json:"context,omitempty"`

	// DoNotStore: Do not store the comment or context sent in this request.
	// By default, the
	// service may store comments/context for debugging purposes.
	DoNotStore bool `json:"doNotStore,omitempty"`

	// Languages: The language(s) of the comment and context. If none are
	// specified, we
	// attempt to automatically detect the language. Specifying multiple
	// languages
	// means the text contains multiple lanugages. Both ISO and BCP-47
	// language
	// codes are accepted.
	//
	// The server returns an error if no language was specified and
	// language
	// detection fails. The server also returns an error if the languages
	// (either
	// specified by the caller, or auto-detected) are not *all* supported by
	// the
	// service.
	Languages []string `json:"languages,omitempty"`

	// RequestedAttributes: Specification of requested attributes. The
	// AttributeParameters serve as
	// configuration for each associated attribute. The map keys are
	// attribute
	// names. The available attributes may be different on each RFE
	// installation,
	// and can be seen by calling ListAttributes (see above).
	// For the prod installation, known as Perspective API,
	// at
	// blade:commentanalyzer-esf and commentanalyzer.googleapis.com,
	// see
	// go/checker-models (internal)
	// and
	// https://github.com/conversationai/perspectiveapi/blob/master/2-api
	// /models.md#all-attribute-types.
	RequestedAttributes map[string]AttributeParameters `json:"requestedAttributes,omitempty"`

	// SessionId: Session ID. Used to join related RPCs into a single
	// session. For example,
	// an interactive tool that calls both the AnalyzeComment
	// and
	// SuggestCommentScore RPCs should set all invocations of both RPCs to
	// the
	// same Session ID, typically a random 64-bit integer.
	SessionId string `json:"sessionId,omitempty"`

	// SpanAnnotations: An advisory parameter that will return span
	// annotations if the model
	// is capable of providing scores with sub-comment resolution. This
	// will
	// likely increase the size of the returned message.
	SpanAnnotations bool `json:"spanAnnotations,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ClientToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ClientToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AnalyzeCommentRequest) MarshalJSON() ([]byte, error) {
	type NoMethod AnalyzeCommentRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AnalyzeCommentResponse: The comment analysis response message.
type AnalyzeCommentResponse struct {
	// AttributeScores: Scores for the requested attributes. The map keys
	// are attribute names (same
	// as the requested_attribute field in AnalyzeCommentRequest, for
	// example
	// "ATTACK_ON_AUTHOR", "INFLAMMATORY", etc).
	AttributeScores map[string]AttributeScores `json:"attributeScores,omitempty"`

	// ClientToken: Same token from the original AnalyzeCommentRequest.
	ClientToken string `json:"clientToken,omitempty"`

	// DetectedLanguages: Contains the languages detected from the text
	// content, sorted in order of
	// likelihood.
	DetectedLanguages []string `json:"detectedLanguages,omitempty"`

	// Languages: The language(s) used by CommentAnalyzer service to choose
	// which Model to
	// use when analyzing the comment. Might better be
	// called
	// "effective_languages". The logic used to make the choice is as
	// follows:
	//   if !Request.languages.empty()
	//     effective_languages = Request.languages
	//   else
	//     effective_languages = detected_languages[0]
	Languages []string `json:"languages,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AttributeScores") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AttributeScores") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *AnalyzeCommentResponse) MarshalJSON() ([]byte, error) {
	type NoMethod AnalyzeCommentResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ArticleAndParentComment: A type of context specific to a comment left
// on a single-threaded comment
// message board, where comments are either a top level comment or the
// child of
// a top level comment.
type ArticleAndParentComment struct {
	// Article: The source content about which the comment was made (article
	// text, article
	// summary, video transcript, etc).
	Article *TextEntry `json:"article,omitempty"`

	// ParentComment: Refers to text that is a direct parent of the source
	// comment, such as in a
	// one-deep threaded message board. This field will only be present
	// for
	// comments that are replies to other comments and will not be populated
	// for
	// direct comments on the article_text.
	ParentComment *TextEntry `json:"parentComment,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Article") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Article") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ArticleAndParentComment) MarshalJSON() ([]byte, error) {
	type NoMethod ArticleAndParentComment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AttributeParameters: Configurable parameters for attribute scoring.
type AttributeParameters struct {
	// ScoreThreshold: Don't return scores for this attribute that are below
	// this threshold. If
	// unset, a default threshold will be applied. A FloatValue wrapper is
	// used to
	// distinguish between 0 vs. default/unset.
	ScoreThreshold float64 `json:"scoreThreshold,omitempty"`

	// ScoreType: What type of scores to return. If unset, defaults to
	// probability scores.
	//
	// Possible values:
	//   "SCORE_TYPE_UNSPECIFIED" - Unspecified. Defaults to PROBABILITY
	// scores if available, and otherwise
	// RAW. Every model has a RAW score.
	//   "PROBABILITY" - Probability scores are in the range [0, 1] and
	// indicate level of confidence
	// in the attribute label.
	//   "STD_DEV_SCORE" - Standard deviation scores are in the range (-inf,
	// +inf).
	//   "PERCENTILE" - Percentile scores are in the range [0, 1] and
	// indicate the percentile of
	// the raw score, normalized with a test dataset. This is not
	// generally
	// recommended, as the normalization is dependent on the dataset used,
	// which
	// may not match other usecases.
	//   "RAW" - Raw scores are the raw values from the model, and may take
	// any value. This
	// is primarily for debugging/testing, and not generally recommended.
	ScoreType string `json:"scoreType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ScoreThreshold") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ScoreThreshold") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *AttributeParameters) MarshalJSON() ([]byte, error) {
	type NoMethod AttributeParameters
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *AttributeParameters) UnmarshalJSON(data []byte) error {
	type NoMethod AttributeParameters
	var s1 struct {
		ScoreThreshold gensupport.JSONFloat64 `json:"scoreThreshold"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.ScoreThreshold = float64(s1.ScoreThreshold)
	return nil
}

// AttributeScores: This holds score values for a single attribute. It
// contains both per-span
// scores as well as an overall summary score..
type AttributeScores struct {
	// SpanScores: Per-span scores.
	SpanScores []*SpanScore `json:"spanScores,omitempty"`

	// SummaryScore: Overall score for comment as a whole.
	SummaryScore *Score `json:"summaryScore,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SpanScores") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "SpanScores") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AttributeScores) MarshalJSON() ([]byte, error) {
	type NoMethod AttributeScores
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Context: Context is typically something that a Comment is referencing
// or replying to
// (such as an article, or previous comment).
// Note: Populate only ONE OF the following fields. The oneof syntax
// cannot be
// used because that would require nesting entries inside another
// message and
// breaking backwards compatibility. The server will return an error if
// more
// than one of the following fields is present.
type Context struct {
	// ArticleAndParentComment: Information about the source for which the
	// original comment was made, and
	// any parent comment info.
	ArticleAndParentComment *ArticleAndParentComment `json:"articleAndParentComment,omitempty"`

	// Entries: A list of messages. For example, a linear comments section
	// or forum thread.
	Entries []*TextEntry `json:"entries,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "ArticleAndParentComment") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ArticleAndParentComment")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Context) MarshalJSON() ([]byte, error) {
	type NoMethod Context
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Score: Analysis scores are described by a value and a ScoreType.
type Score struct {
	// Type: The type of the above value.
	//
	// Possible values:
	//   "SCORE_TYPE_UNSPECIFIED" - Unspecified. Defaults to PROBABILITY
	// scores if available, and otherwise
	// RAW. Every model has a RAW score.
	//   "PROBABILITY" - Probability scores are in the range [0, 1] and
	// indicate level of confidence
	// in the attribute label.
	//   "STD_DEV_SCORE" - Standard deviation scores are in the range (-inf,
	// +inf).
	//   "PERCENTILE" - Percentile scores are in the range [0, 1] and
	// indicate the percentile of
	// the raw score, normalized with a test dataset. This is not
	// generally
	// recommended, as the normalization is dependent on the dataset used,
	// which
	// may not match other usecases.
	//   "RAW" - Raw scores are the raw values from the model, and may take
	// any value. This
	// is primarily for debugging/testing, and not generally recommended.
	Type string `json:"type,omitempty"`

	// Value: Score value. Semantics described by type below.
	Value float64 `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Type") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Type") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Score) MarshalJSON() ([]byte, error) {
	type NoMethod Score
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *Score) UnmarshalJSON(data []byte) error {
	type NoMethod Score
	var s1 struct {
		Value gensupport.JSONFloat64 `json:"value"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Value = float64(s1.Value)
	return nil
}

// SpanScore: This is a single score for a given span of text.
type SpanScore struct {
	// Begin: "begin" and "end" describe the span of the original text that
	// the attribute
	// score applies to. The values are the UTF-16 codepoint range. "end"
	// is
	// exclusive. For example, with the text "Hi there", the begin/end pair
	// (0,2)
	// describes the text "Hi".
	//
	// If "begin" and "end" are unset, the score applies to the full text.
	Begin int64 `json:"begin,omitempty"`

	End int64 `json:"end,omitempty"`

	// Score: The score value.
	Score *Score `json:"score,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Begin") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Begin") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SpanScore) MarshalJSON() ([]byte, error) {
	type NoMethod SpanScore
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SuggestCommentScoreRequest: The comment score suggestion request
// message.
type SuggestCommentScoreRequest struct {
	// AttributeScores: Attribute scores for the comment. The map keys are
	// attribute names, same as
	// the requested_attribute field in AnalyzeCommentRequest (for
	// example
	// "ATTACK_ON_AUTHOR", "INFLAMMATORY", etc.). This field has the same
	// type as
	// the `attribute_scores` field in AnalyzeCommentResponse.
	//
	// To specify an overall attribute score for the entire comment as a
	// whole,
	// use the `summary_score` field of the mapped AttributeScores object.
	// To
	// specify scores on specific subparts of the comment, use the
	// `span_scores`
	// field. All SpanScore objects must have begin and end fields set.
	//
	// All Score objects must be explicitly set (for binary classification,
	// use
	// the score values 0 and 1). If Score objects don't include a
	// ScoreType,
	// `PROBABILITY` is assumed.
	//
	// `attribute_scores` must not be empty. The mapped AttributeScores
	// objects
	// also must not be empty. An `INVALID_ARGUMENT` error is returned for
	// all
	// malformed requests.
	AttributeScores map[string]AttributeScores `json:"attributeScores,omitempty"`

	// ClientToken: Opaque token that is echoed from the request to the
	// response.
	ClientToken string `json:"clientToken,omitempty"`

	// Comment: The comment being scored.
	Comment *TextEntry `json:"comment,omitempty"`

	// CommunityId: Optional identifier associating this comment score
	// suggestion with a
	// particular sub-community. Different communities may have different
	// norms
	// and rules. Specifying this value enables training
	// community-specific
	// models.
	CommunityId string `json:"communityId,omitempty"`

	// Context: The context of the comment.
	Context *Context `json:"context,omitempty"`

	// Languages: The language(s) of the comment and context. If none are
	// specified, we
	// attempt to automatically detect the language. Both ISO and BCP-47
	// language
	// codes are accepted.
	Languages []string `json:"languages,omitempty"`

	// SessionId: Session ID. Used to join related RPCs into a single
	// session. For example,
	// an interactive tool that calls both the AnalyzeComment
	// and
	// SuggestCommentScore RPCs should set all invocations of both RPCs to
	// the
	// same Session ID, typically a random 64-bit integer.
	SessionId string `json:"sessionId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AttributeScores") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AttributeScores") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *SuggestCommentScoreRequest) MarshalJSON() ([]byte, error) {
	type NoMethod SuggestCommentScoreRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SuggestCommentScoreResponse: The comment score suggestion response
// message.
type SuggestCommentScoreResponse struct {
	// ClientToken: Same token from the original SuggestCommentScoreRequest.
	ClientToken string `json:"clientToken,omitempty"`

	// DetectedLanguages: The list of languages detected from the comment
	// text.
	DetectedLanguages []string `json:"detectedLanguages,omitempty"`

	// RequestedLanguages: The list of languages provided in the request.
	RequestedLanguages []string `json:"requestedLanguages,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ClientToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ClientToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SuggestCommentScoreResponse) MarshalJSON() ([]byte, error) {
	type NoMethod SuggestCommentScoreResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TextEntry: Represents a body of text.
type TextEntry struct {
	// Text: UTF-8 encoded text.
	Text string `json:"text,omitempty"`

	// Type: Type of the text field.
	//
	// Possible values:
	//   "TEXT_TYPE_UNSPECIFIED" - The content type is not specified. Text
	// will be interpreted as plain text
	// by default.
	//   "PLAIN_TEXT" - Plain text.
	//   "HTML" - HTML.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Text") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Text") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TextEntry) MarshalJSON() ([]byte, error) {
	type NoMethod TextEntry
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "commentanalyzer.comments.analyze":

type CommentsAnalyzeCall struct {
	s                     *Service
	analyzecommentrequest *AnalyzeCommentRequest
	urlParams_            gensupport.URLParams
	ctx_                  context.Context
	header_               http.Header
}

// Analyze: Analyzes the provided text and returns scores for requested
// attributes.
func (r *CommentsService) Analyze(analyzecommentrequest *AnalyzeCommentRequest) *CommentsAnalyzeCall {
	c := &CommentsAnalyzeCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.analyzecommentrequest = analyzecommentrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CommentsAnalyzeCall) Fields(s ...googleapi.Field) *CommentsAnalyzeCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CommentsAnalyzeCall) Context(ctx context.Context) *CommentsAnalyzeCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CommentsAnalyzeCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CommentsAnalyzeCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20200316")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.analyzecommentrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/comments:analyze")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "commentanalyzer.comments.analyze" call.
// Exactly one of *AnalyzeCommentResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *AnalyzeCommentResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *CommentsAnalyzeCall) Do(opts ...googleapi.CallOption) (*AnalyzeCommentResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AnalyzeCommentResponse{
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
	//   "description": "Analyzes the provided text and returns scores for requested attributes.",
	//   "flatPath": "v1alpha1/comments:analyze",
	//   "httpMethod": "POST",
	//   "id": "commentanalyzer.comments.analyze",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1alpha1/comments:analyze",
	//   "request": {
	//     "$ref": "AnalyzeCommentRequest"
	//   },
	//   "response": {
	//     "$ref": "AnalyzeCommentResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "commentanalyzer.comments.suggestscore":

type CommentsSuggestscoreCall struct {
	s                          *Service
	suggestcommentscorerequest *SuggestCommentScoreRequest
	urlParams_                 gensupport.URLParams
	ctx_                       context.Context
	header_                    http.Header
}

// Suggestscore: Suggest comment scores as training data.
func (r *CommentsService) Suggestscore(suggestcommentscorerequest *SuggestCommentScoreRequest) *CommentsSuggestscoreCall {
	c := &CommentsSuggestscoreCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.suggestcommentscorerequest = suggestcommentscorerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CommentsSuggestscoreCall) Fields(s ...googleapi.Field) *CommentsSuggestscoreCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CommentsSuggestscoreCall) Context(ctx context.Context) *CommentsSuggestscoreCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CommentsSuggestscoreCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CommentsSuggestscoreCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20200316")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.suggestcommentscorerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/comments:suggestscore")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "commentanalyzer.comments.suggestscore" call.
// Exactly one of *SuggestCommentScoreResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *SuggestCommentScoreResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *CommentsSuggestscoreCall) Do(opts ...googleapi.CallOption) (*SuggestCommentScoreResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &SuggestCommentScoreResponse{
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
	//   "description": "Suggest comment scores as training data.",
	//   "flatPath": "v1alpha1/comments:suggestscore",
	//   "httpMethod": "POST",
	//   "id": "commentanalyzer.comments.suggestscore",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1alpha1/comments:suggestscore",
	//   "request": {
	//     "$ref": "SuggestCommentScoreRequest"
	//   },
	//   "response": {
	//     "$ref": "SuggestCommentScoreResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}
