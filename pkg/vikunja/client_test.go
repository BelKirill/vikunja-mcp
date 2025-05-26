package vikunja

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// type testHTTPClient struct {
// 	doFunc func(req *http.Request) (*http.Response, error)
// }

// func (c *testHTTPClient) Do(req *http.Request) (*http.Response, error) {
// 	return c.doFunc(req)
// }

func TestNewClient_EnvVars(t *testing.T) {
	os.Setenv("VIKUNJA_URL", "http://localhost:3456")
	os.Setenv("VIKUNJA_TOKEN", "testtoken")
	client, err := NewClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "testtoken", client.Token)
	assert.Equal(t, "http://localhost:3456", client.BaseURL.String())
}

func TestNewClient_MissingEnv(t *testing.T) {
	os.Unsetenv("VIKUNJA_URL")
	os.Unsetenv("VIKUNJA_TOKEN")
	_, err := NewClient()
	assert.Error(t, err)
}

func TestClient_GetTask_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": map[string]interface{}{
				"id": "1", "project_id": "p1", "due_date": "2025-01-01", "priority": 2, "label_ids": []int{}, "estimate": 30,
			},
		})
	}))
	defer server.Close()

	u, _ := url.Parse(server.URL)
	client := &Client{BaseURL: u, HTTPClient: server.Client(), Token: "tok"}
	task, err := client.GetTask("1")
	assert.NoError(t, err)
	assert.Equal(t, "1", task.ID)
	assert.Equal(t, "p1", task.ProjectID)
}

func TestClient_GetTask_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	u, _ := url.Parse(server.URL)
	client := &Client{BaseURL: u, HTTPClient: server.Client(), Token: "tok"}
	_, err := client.GetTask("notfound")
	assert.Error(t, err)
}

func TestClient_GetProject_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": map[string]interface{}{
				"id": "p1", "title": "Project 1",
			},
		})
	}))
	defer server.Close()

	u, _ := url.Parse(server.URL)
	client := &Client{BaseURL: u, HTTPClient: server.Client(), Token: "tok"}
	proj, err := client.GetProject("p1")
	assert.NoError(t, err)
	assert.Equal(t, "p1", proj.ID)
	assert.Equal(t, "Project 1", proj.Name)
}

func TestClient_SearchTasks_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": []map[string]interface{}{
				{"id": "1", "project_id": "p1", "due_date": "2025-01-01", "priority": 2, "label_ids": []int{}, "estimate": 30},
			},
		})
	}))
	defer server.Close()

	u, _ := url.Parse(server.URL)
	client := &Client{BaseURL: u, HTTPClient: server.Client(), Token: "tok"}
	tasks, err := client.SearchTasks(map[string]string{"project_id": "p1"})
	assert.NoError(t, err)
	assert.Len(t, tasks, 1)
	assert.Equal(t, "1", tasks[0].ID)
}

func TestClient_SearchTasks_ErrorStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	u, _ := url.Parse(server.URL)
	client := &Client{BaseURL: u, HTTPClient: server.Client(), Token: "tok"}
	_, err := client.SearchTasks(map[string]string{"project_id": "p1"})
	assert.Error(t, err)
}
