# \LabelsAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LabelsGet**](LabelsAPI.md#LabelsGet) | **Get** /labels | Get all labels a user has access to
[**LabelsIdDelete**](LabelsAPI.md#LabelsIdDelete) | **Delete** /labels/{id} | Delete a label
[**LabelsIdGet**](LabelsAPI.md#LabelsIdGet) | **Get** /labels/{id} | Gets one label
[**LabelsIdPut**](LabelsAPI.md#LabelsIdPut) | **Put** /labels/{id} | Update a label
[**LabelsPut**](LabelsAPI.md#LabelsPut) | **Put** /labels | Create a label
[**TasksTaskIDLabelsBulkPost**](LabelsAPI.md#TasksTaskIDLabelsBulkPost) | **Post** /tasks/{taskID}/labels/bulk | Update all labels on a task.
[**TasksTaskLabelsGet**](LabelsAPI.md#TasksTaskLabelsGet) | **Get** /tasks/{task}/labels | Get all labels on a task
[**TasksTaskLabelsLabelDelete**](LabelsAPI.md#TasksTaskLabelsLabelDelete) | **Delete** /tasks/{task}/labels/{label} | Remove a label from a task
[**TasksTaskLabelsPut**](LabelsAPI.md#TasksTaskLabelsPut) | **Put** /tasks/{task}/labels | Add a label to a task



## LabelsGet

> []ModelsLabel LabelsGet(ctx).Page(page).PerPage(perPage).S(s).Execute()

Get all labels a user has access to



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
	page := int32(56) // int32 | The page number. Used for pagination. If not provided, the first page of results is returned. (optional)
	perPage := int32(56) // int32 | The maximum number of items per page. Note this parameter is limited by the configured maximum of items per page. (optional)
	s := "s_example" // string | Search labels by label text. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelsAPI.LabelsGet(context.Background()).Page(page).PerPage(perPage).S(s).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.LabelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LabelsGet`: []ModelsLabel
	fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.LabelsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLabelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page number. Used for pagination. If not provided, the first page of results is returned. | 
 **perPage** | **int32** | The maximum number of items per page. Note this parameter is limited by the configured maximum of items per page. | 
 **s** | **string** | Search labels by label text. | 

### Return type

[**[]ModelsLabel**](ModelsLabel.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelsIdDelete

> ModelsLabel LabelsIdDelete(ctx, id).Execute()

Delete a label



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
	id := int32(56) // int32 | Label ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelsAPI.LabelsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.LabelsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LabelsIdDelete`: ModelsLabel
	fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.LabelsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Label ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsLabel**](ModelsLabel.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelsIdGet

> ModelsLabel LabelsIdGet(ctx, id).Execute()

Gets one label



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
	id := int32(56) // int32 | Label ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelsAPI.LabelsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.LabelsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LabelsIdGet`: ModelsLabel
	fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.LabelsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Label ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsLabel**](ModelsLabel.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelsIdPut

> ModelsLabel LabelsIdPut(ctx, id).Label(label).Execute()

Update a label



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
	id := int32(56) // int32 | Label ID
	label := *openapiclient.NewModelsLabel() // ModelsLabel | The label object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelsAPI.LabelsIdPut(context.Background(), id).Label(label).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.LabelsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LabelsIdPut`: ModelsLabel
	fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.LabelsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Label ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **label** | [**ModelsLabel**](ModelsLabel.md) | The label object | 

### Return type

[**ModelsLabel**](ModelsLabel.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelsPut

> ModelsLabel LabelsPut(ctx).Label(label).Execute()

Create a label



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
	label := *openapiclient.NewModelsLabel() // ModelsLabel | The label object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelsAPI.LabelsPut(context.Background()).Label(label).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.LabelsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LabelsPut`: ModelsLabel
	fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.LabelsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLabelsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **label** | [**ModelsLabel**](ModelsLabel.md) | The label object | 

### Return type

[**ModelsLabel**](ModelsLabel.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskIDLabelsBulkPost

> ModelsLabelTaskBulk TasksTaskIDLabelsBulkPost(ctx, taskID).Label(label).Execute()

Update all labels on a task.



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
	label := *openapiclient.NewModelsLabelTaskBulk() // ModelsLabelTaskBulk | The array of labels

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelsAPI.TasksTaskIDLabelsBulkPost(context.Background(), taskID).Label(label).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.TasksTaskIDLabelsBulkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIDLabelsBulkPost`: ModelsLabelTaskBulk
	fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.TasksTaskIDLabelsBulkPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskID** | **int32** | Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIDLabelsBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **label** | [**ModelsLabelTaskBulk**](ModelsLabelTaskBulk.md) | The array of labels | 

### Return type

[**ModelsLabelTaskBulk**](ModelsLabelTaskBulk.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskLabelsGet

> []ModelsLabel TasksTaskLabelsGet(ctx, task).Page(page).PerPage(perPage).S(s).Execute()

Get all labels on a task



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
	task := int32(56) // int32 | Task ID
	page := int32(56) // int32 | The page number. Used for pagination. If not provided, the first page of results is returned. (optional)
	perPage := int32(56) // int32 | The maximum number of items per page. Note this parameter is limited by the configured maximum of items per page. (optional)
	s := "s_example" // string | Search labels by label text. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelsAPI.TasksTaskLabelsGet(context.Background(), task).Page(page).PerPage(perPage).S(s).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.TasksTaskLabelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskLabelsGet`: []ModelsLabel
	fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.TasksTaskLabelsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**task** | **int32** | Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskLabelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page number. Used for pagination. If not provided, the first page of results is returned. | 
 **perPage** | **int32** | The maximum number of items per page. Note this parameter is limited by the configured maximum of items per page. | 
 **s** | **string** | Search labels by label text. | 

### Return type

[**[]ModelsLabel**](ModelsLabel.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskLabelsLabelDelete

> ModelsMessage TasksTaskLabelsLabelDelete(ctx, task, label).Execute()

Remove a label from a task



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
	task := int32(56) // int32 | Task ID
	label := int32(56) // int32 | Label ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelsAPI.TasksTaskLabelsLabelDelete(context.Background(), task, label).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.TasksTaskLabelsLabelDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskLabelsLabelDelete`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.TasksTaskLabelsLabelDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**task** | **int32** | Task ID | 
**label** | **int32** | Label ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskLabelsLabelDeleteRequest struct via the builder pattern


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


## TasksTaskLabelsPut

> ModelsLabelTask TasksTaskLabelsPut(ctx, task).Label(label).Execute()

Add a label to a task



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
	task := int32(56) // int32 | Task ID
	label := *openapiclient.NewModelsLabelTask() // ModelsLabelTask | The label object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelsAPI.TasksTaskLabelsPut(context.Background(), task).Label(label).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.TasksTaskLabelsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskLabelsPut`: ModelsLabelTask
	fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.TasksTaskLabelsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**task** | **int32** | Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskLabelsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **label** | [**ModelsLabelTask**](ModelsLabelTask.md) | The label object | 

### Return type

[**ModelsLabelTask**](ModelsLabelTask.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

