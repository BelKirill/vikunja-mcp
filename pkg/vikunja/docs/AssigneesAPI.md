# \AssigneesAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TasksTaskIDAssigneesBulkPost**](AssigneesAPI.md#TasksTaskIDAssigneesBulkPost) | **Post** /tasks/{taskID}/assignees/bulk | Add multiple new assignees to a task
[**TasksTaskIDAssigneesGet**](AssigneesAPI.md#TasksTaskIDAssigneesGet) | **Get** /tasks/{taskID}/assignees | Get all assignees for a task
[**TasksTaskIDAssigneesPut**](AssigneesAPI.md#TasksTaskIDAssigneesPut) | **Put** /tasks/{taskID}/assignees | Add a new assignee to a task
[**TasksTaskIDAssigneesUserIDDelete**](AssigneesAPI.md#TasksTaskIDAssigneesUserIDDelete) | **Delete** /tasks/{taskID}/assignees/{userID} | Delete an assignee



## TasksTaskIDAssigneesBulkPost

> ModelsTaskAssginee TasksTaskIDAssigneesBulkPost(ctx, taskID).Assignee(assignee).Execute()

Add multiple new assignees to a task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	taskID := int32(56) // int32 | Task ID
	assignee := *openapiclient.NewModelsBulkAssignees() // ModelsBulkAssignees | The array of assignees

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssigneesAPI.TasksTaskIDAssigneesBulkPost(context.Background(), taskID).Assignee(assignee).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssigneesAPI.TasksTaskIDAssigneesBulkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIDAssigneesBulkPost`: ModelsTaskAssginee
	fmt.Fprintf(os.Stdout, "Response from `AssigneesAPI.TasksTaskIDAssigneesBulkPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskID** | **int32** | Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIDAssigneesBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignee** | [**ModelsBulkAssignees**](ModelsBulkAssignees.md) | The array of assignees | 

### Return type

[**ModelsTaskAssginee**](ModelsTaskAssginee.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskIDAssigneesGet

> []UserUser TasksTaskIDAssigneesGet(ctx, taskID).Page(page).PerPage(perPage).S(s).Execute()

Get all assignees for a task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	taskID := int32(56) // int32 | Task ID
	page := int32(56) // int32 | The page number. Used for pagination. If not provided, the first page of results is returned. (optional)
	perPage := int32(56) // int32 | The maximum number of items per page. Note this parameter is limited by the configured maximum of items per page. (optional)
	s := "s_example" // string | Search assignees by their username. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssigneesAPI.TasksTaskIDAssigneesGet(context.Background(), taskID).Page(page).PerPage(perPage).S(s).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssigneesAPI.TasksTaskIDAssigneesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIDAssigneesGet`: []UserUser
	fmt.Fprintf(os.Stdout, "Response from `AssigneesAPI.TasksTaskIDAssigneesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskID** | **int32** | Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIDAssigneesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page number. Used for pagination. If not provided, the first page of results is returned. | 
 **perPage** | **int32** | The maximum number of items per page. Note this parameter is limited by the configured maximum of items per page. | 
 **s** | **string** | Search assignees by their username. | 

### Return type

[**[]UserUser**](UserUser.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskIDAssigneesPut

> ModelsTaskAssginee TasksTaskIDAssigneesPut(ctx, taskID).Assignee(assignee).Execute()

Add a new assignee to a task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	taskID := int32(56) // int32 | Task ID
	assignee := *openapiclient.NewModelsTaskAssginee() // ModelsTaskAssginee | The assingee object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssigneesAPI.TasksTaskIDAssigneesPut(context.Background(), taskID).Assignee(assignee).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssigneesAPI.TasksTaskIDAssigneesPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIDAssigneesPut`: ModelsTaskAssginee
	fmt.Fprintf(os.Stdout, "Response from `AssigneesAPI.TasksTaskIDAssigneesPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskID** | **int32** | Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIDAssigneesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignee** | [**ModelsTaskAssginee**](ModelsTaskAssginee.md) | The assingee object | 

### Return type

[**ModelsTaskAssginee**](ModelsTaskAssginee.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskIDAssigneesUserIDDelete

> ModelsMessage TasksTaskIDAssigneesUserIDDelete(ctx, taskID, userID).Execute()

Delete an assignee



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	taskID := int32(56) // int32 | Task ID
	userID := int32(56) // int32 | Assignee user ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssigneesAPI.TasksTaskIDAssigneesUserIDDelete(context.Background(), taskID, userID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssigneesAPI.TasksTaskIDAssigneesUserIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIDAssigneesUserIDDelete`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `AssigneesAPI.TasksTaskIDAssigneesUserIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskID** | **int32** | Task ID | 
**userID** | **int32** | Assignee user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIDAssigneesUserIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelsMessage**](ModelsMessage.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

