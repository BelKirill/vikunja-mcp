# \WebhooksAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsIdWebhooksGet**](WebhooksAPI.md#ProjectsIdWebhooksGet) | **Get** /projects/{id}/webhooks | Get all api webhook targets for the specified project
[**ProjectsIdWebhooksPut**](WebhooksAPI.md#ProjectsIdWebhooksPut) | **Put** /projects/{id}/webhooks | Create a webhook target
[**ProjectsIdWebhooksWebhookIDDelete**](WebhooksAPI.md#ProjectsIdWebhooksWebhookIDDelete) | **Delete** /projects/{id}/webhooks/{webhookID} | Deletes an existing webhook target
[**ProjectsIdWebhooksWebhookIDPost**](WebhooksAPI.md#ProjectsIdWebhooksWebhookIDPost) | **Post** /projects/{id}/webhooks/{webhookID} | Change a webhook target&#39;s events.
[**WebhooksEventsGet**](WebhooksAPI.md#WebhooksEventsGet) | **Get** /webhooks/events | Get all possible webhook events



## ProjectsIdWebhooksGet

> []ModelsWebhook ProjectsIdWebhooksGet(ctx, id).Page(page).PerPage(perPage).Execute()

Get all api webhook targets for the specified project



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
	id := int32(56) // int32 | Project ID
	page := int32(56) // int32 | The page number. Used for pagination. If not provided, the first page of results is returned. (optional)
	perPage := int32(56) // int32 | The maximum number of items per bucket per page. This parameter is limited by the configured maximum of items per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.ProjectsIdWebhooksGet(context.Background(), id).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ProjectsIdWebhooksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsIdWebhooksGet`: []ModelsWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.ProjectsIdWebhooksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsIdWebhooksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page number. Used for pagination. If not provided, the first page of results is returned. | 
 **perPage** | **int32** | The maximum number of items per bucket per page. This parameter is limited by the configured maximum of items per page. | 

### Return type

[**[]ModelsWebhook**](ModelsWebhook.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsIdWebhooksPut

> ModelsWebhook ProjectsIdWebhooksPut(ctx, id).Webhook(webhook).Execute()

Create a webhook target



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
	id := int32(56) // int32 | Project ID
	webhook := *openapiclient.NewModelsWebhook() // ModelsWebhook | The webhook target object with required fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.ProjectsIdWebhooksPut(context.Background(), id).Webhook(webhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ProjectsIdWebhooksPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsIdWebhooksPut`: ModelsWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.ProjectsIdWebhooksPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsIdWebhooksPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhook** | [**ModelsWebhook**](ModelsWebhook.md) | The webhook target object with required fields | 

### Return type

[**ModelsWebhook**](ModelsWebhook.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsIdWebhooksWebhookIDDelete

> ModelsMessage ProjectsIdWebhooksWebhookIDDelete(ctx, id, webhookID).Execute()

Deletes an existing webhook target



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
	id := int32(56) // int32 | Project ID
	webhookID := int32(56) // int32 | Webhook ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.ProjectsIdWebhooksWebhookIDDelete(context.Background(), id, webhookID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ProjectsIdWebhooksWebhookIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsIdWebhooksWebhookIDDelete`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.ProjectsIdWebhooksWebhookIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Project ID | 
**webhookID** | **int32** | Webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsIdWebhooksWebhookIDDeleteRequest struct via the builder pattern


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


## ProjectsIdWebhooksWebhookIDPost

> ModelsWebhook ProjectsIdWebhooksWebhookIDPost(ctx, id, webhookID).Execute()

Change a webhook target's events.



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
	id := int32(56) // int32 | Project ID
	webhookID := int32(56) // int32 | Webhook ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.ProjectsIdWebhooksWebhookIDPost(context.Background(), id, webhookID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ProjectsIdWebhooksWebhookIDPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsIdWebhooksWebhookIDPost`: ModelsWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.ProjectsIdWebhooksWebhookIDPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Project ID | 
**webhookID** | **int32** | Webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsIdWebhooksWebhookIDPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelsWebhook**](ModelsWebhook.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksEventsGet

> []string WebhooksEventsGet(ctx).Execute()

Get all possible webhook events



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.WebhooksEventsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksEventsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhooksEventsGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.WebhooksEventsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksEventsGetRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

