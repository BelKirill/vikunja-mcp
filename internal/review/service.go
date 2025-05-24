package review

import (
	"time"
)

// --- Service Interfaces ------------------------------------------------------

type ReviewService interface {
    // Review takes a list of project IDs and returns for each the nearest important task
    Review(projectIDs []string) ([]ReviewItem, error)
}
// --- Data Models -------------------------------------------------------------

type ReviewItem struct {
    ProjectID string    // p
    TaskID    string    // t
    DueDate   time.Time // d
}

// --- Request/Response DTOs ---------------------------------------------------

// ReviewRequest mirrors the `360-review` contract
// req: {"project_ids": ["proj-1", "proj-2"]}
type ReviewRequest struct {
    ProjectIDs []string `json:"project_ids"`
}

// ReviewResponseItem is slim: {p,t,d}
type ReviewResponseItem struct {
    P string `json:"p"`
    T string `json:"t"`
    D string `json:"d"` // date only YYYY-MM-DD
}

// --- Default stubs (override via setter) ------------------------------------

var (
    reviewService ReviewService = &stubReviewService{}
)

func SetReviewService(svc ReviewService) {
    reviewService = svc
}


// stubReviewService returns an empty slice for now
type stubReviewService struct{}

func (s *stubReviewService) Review(projectIDs []string) ([]ReviewItem, error) {
    // stub: return no items
    return []ReviewItem{}, nil
}
