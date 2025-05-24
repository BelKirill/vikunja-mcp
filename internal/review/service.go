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
