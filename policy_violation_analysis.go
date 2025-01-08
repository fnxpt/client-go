package dtrack

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type ViolationPortfolioResponse []Violation

type Violation struct {
	Type            string          `json:"type"`
	Project         Project         `json:"project"`
	Component       Component       `json:"component"`
	PolicyCondition PolicyCondition `json:"policyCondition"`
	Timestamp       int64           `json:"timestamp"`
	UUID            string          `json:"uuid"`
}

type ViolationAnalysisState string

const (
	ViolationAnalysisStateNotSet   ViolationAnalysisState = "NOT_SET"
	ViolationAnalysisStateApproved ViolationAnalysisState = "APPROVED"
	ViolationAnalysisStateRejected ViolationAnalysisState = "REJECTED"
)

type ViolationAnalysis struct {
	Comments   []ViolationAnalysisComment `json:"analysisComments"`
	State      ViolationAnalysisState     `json:"analysisState"`
	Suppressed bool                       `json:"isSuppressed"`
}

type ViolationAnalysisComment struct {
	Comment   string `json:"comment"`
	Commenter string `json:"commenter"`
	Timestamp int    `json:"timestamp"`
}

type ViolationAnalysisRequest struct {
	Component       uuid.UUID              `json:"component"`
	PolicyViolation uuid.UUID              `json:"policyViolation"`
	Comment         string                 `json:"comment,omitempty"`
	State           ViolationAnalysisState `json:"analysisState,omitempty"`
	Suppressed      *bool                  `json:"isSuppressed,omitempty"`
}

type ViolationAnalysisService struct {
	client *Client
}

func (vas ViolationAnalysisService) Portfolio(ctx context.Context) (va ViolationPortfolioResponse, err error) {
	params := map[string]string{}

	// pageNumber=1
	// pageSize=100
	// offset=1
	// limit=1
	// sortName=1
	// sortOrder=asc%2C%20desc
	// suppressed=true
	// showInactive=true
	// violationState=1
	// riskType=1
	// policy=1
	// analysisState=1
	// occurredOnDateFrom=1
	// occurredOnDateTo=1
	// textSearchField=1
	// textSearchInput=1

	req, err := vas.client.newRequest(ctx, http.MethodGet, "/api/v1/violation", withParams(params))
	if err != nil {
		return
	}

	_, err = vas.client.doRequest(req, &va)
	return
}

func (vas ViolationAnalysisService) Get(ctx context.Context, componentUUID, policyViolationUUID uuid.UUID) (va ViolationAnalysis, err error) {
	params := map[string]string{
		"component":       componentUUID.String(),
		"policyViolation": policyViolationUUID.String(),
	}

	req, err := vas.client.newRequest(ctx, http.MethodGet, "/api/v1/violation/analysis", withParams(params))
	if err != nil {
		return
	}

	_, err = vas.client.doRequest(req, &va)
	return
}

func (vas ViolationAnalysisService) Update(ctx context.Context, analysisReq ViolationAnalysisRequest) (va ViolationAnalysis, err error) {
	req, err := vas.client.newRequest(ctx, http.MethodPut, "/api/v1/violation/analysis", withBody(analysisReq))
	if err != nil {
		return
	}

	_, err = vas.client.doRequest(req, &va)
	return
}
