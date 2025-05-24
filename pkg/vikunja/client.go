package vikunja

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "os"
    "time"
)

// Client encapsulates HTTP settings for Vikunja.
type Client struct {
    BaseURL    *url.URL
    HTTPClient *http.Client
    Token      string
}

// NewClient reads VIKUNJA_URL and VIKUNJA_TOKEN from env and returns a configured client.
func NewClient() (*Client, error) {
    rawURL := os.Getenv("VIKUNJA_URL")
    if rawURL == "" {
        return nil, fmt.Errorf("VIKUNJA_URL is not set")
    }
    parsed, err := url.Parse(rawURL)
    if err != nil {
        return nil, fmt.Errorf("invalid VIKUNJA_URL: %w", err)
    }
    token := os.Getenv("VIKUNJA_TOKEN")
    if token == "" {
        return nil, fmt.Errorf("VIKUNJA_TOKEN is not set")
    }

    return &Client{
        BaseURL:    parsed,
        HTTPClient: &http.Client{Timeout: 10 * time.Second},
        Token:      token,
    }, nil
}

// RawTask matches our internal RawTask struct JSON from Vikunja.
type RawTask struct {
    ID         string    `json:"id"`
    ProjectID  string    `json:"project_id"`
    DueDate    string    `json:"due_date"`   // ISO 8601
    Priority   int       `json:"priority"`
    LabelIDs   []int     `json:"label_ids"`
    EstimateMin int      `json:"estimate"`    // minutes
}


// GetTask fetches a single task by ID.
func (c *Client) GetTask(taskID string) (*RawTask, error) {
    rel := &url.URL{Path: fmt.Sprintf("/v1/tasks/%s", taskID)}
    u := c.BaseURL.ResolveReference(rel)

    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Authorization", "Bearer "+c.Token)

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected status %d fetching task %s", resp.StatusCode, taskID)
    }

    var wrapper struct {
        Data RawTask `json:"data"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
        return nil, err
    }
    return &wrapper.Data, nil
}

// GetProject fetches a single project by ID.
func (c *Client) GetProject(projID string) (*struct {
    ID   string `json:"id"`
    Name string `json:"title"`
}, error) {
    rel := &url.URL{Path: fmt.Sprintf("/v1/projects/%s", projID)}
    u := c.BaseURL.ResolveReference(rel)

    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Authorization", "Bearer "+c.Token)

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected status %d fetching project %s", resp.StatusCode, projID)
    }

    var wrapper struct {
        Data struct {
            ID    string `json:"id"`
            Name string `json:"title"`
        } `json:"data"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
        return nil, err
    }
    return &wrapper.Data, nil
}

// SearchTasks fetches tasks by query parameters. Example: filter by project IDs, priority, labels.
func (c *Client) SearchTasks(params map[string]string) ([]RawTask, error) {
    rel := &url.URL{Path: "/v1/tasks"}
    u := c.BaseURL.ResolveReference(rel)

    q := u.Query()
    for k, v := range params {
        q.Set(k, v)
    }
    u.RawQuery = q.Encode()

    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Authorization", "Bearer "+c.Token)

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected status %d searching tasks", resp.StatusCode)
    }

    var wrapper struct {
        Data []RawTask `json:"data"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
        return nil, err
    }
    return wrapper.Data, nil
}