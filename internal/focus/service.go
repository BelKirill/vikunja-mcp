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

// --- Default stubs (override via setter) ------------------------------------

// --- Default stubs (override via setter) ------------------------------------

var (
	focusService FocusService = &stubFocusService{}
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
