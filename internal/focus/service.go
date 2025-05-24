package focus

// --- Service Interfaces ------------------------------------------------------

type FocusService interface {
    // Focus takes a date and hours, returns a prioritized worklist
    Focus(date string, hours float64) ([]FocusItem, error)
}

// --- Data Models -------------------------------------------------------------

type FocusItem struct {
    TaskID   string  // t
    Project  string  // p
    Estimate float64 // est
}

// --- Request/Response DTOs ---------------------------------------------------

// req: {"date": "2025-05-26", "hours": 5}
type FocusRequest struct {
    Date  string  `json:"date,omitempty"`  // optional, default = tomorrow
    Hours float64 `json:"hours,omitempty"` // optional, derive from calendar if zero
}

// FocusResponseItem is slim: {t,p,est}
type FocusResponseItem struct {
    T   string  `json:"t"`
    P   string  `json:"p"`
    Est float64 `json:"est"`
}

// --- Default stubs (override via setter) ------------------------------------

// --- Default stubs (override via setter) ------------------------------------

var (
    focusService  FocusService  = &stubFocusService{}
)

func SetFocusService(svc FocusService) {
    focusService = svc
}

// stubFocusService returns empty slice
type stubFocusService struct{}

func (s *stubFocusService) Focus(date string, hours float64) ([]FocusItem, error) {
    // stub: return no items
    return []FocusItem{}, nil
}