# \SubscriptionsAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationsGet**](SubscriptionsAPI.md#NotificationsGet) | **Get** /notifications | Get all notifications for the current user
[**NotificationsIdPost**](SubscriptionsAPI.md#NotificationsIdPost) | **Post** /notifications/{id} | Mark a notification as (un-)read
[**SubscriptionsEntityEntityIDDelete**](SubscriptionsAPI.md#SubscriptionsEntityEntityIDDelete) | **Delete** /subscriptions/{entity}/{entityID} | Unsubscribe the current user from an entity.
[**SubscriptionsEntityEntityIDPut**](SubscriptionsAPI.md#SubscriptionsEntityEntityIDPut) | **Put** /subscriptions/{entity}/{entityID} | Subscribes the current user to an entity.



## NotificationsGet

> []NotificationsDatabaseNotification NotificationsGet(ctx).Page(page).PerPage(perPage).Execute()

Get all notifications for the current user



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.NotificationsGet(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.NotificationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationsGet`: []NotificationsDatabaseNotification
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.NotificationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page number. Used for pagination. If not provided, the first page of results is returned. | 
 **perPage** | **int32** | The maximum number of items per page. Note this parameter is limited by the configured maximum of items per page. | 

### Return type

[**[]NotificationsDatabaseNotification**](NotificationsDatabaseNotification.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationsIdPost

> ModelsDatabaseNotifications NotificationsIdPost(ctx, id).Execute()

Mark a notification as (un-)read



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
	id := int32(56) // int32 | Notification ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.NotificationsIdPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.NotificationsIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationsIdPost`: ModelsDatabaseNotifications
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.NotificationsIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Notification ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsDatabaseNotifications**](ModelsDatabaseNotifications.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsEntityEntityIDDelete

> ModelsSubscription SubscriptionsEntityEntityIDDelete(ctx, entity, entityID).Execute()

Unsubscribe the current user from an entity.



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
	entity := "entity_example" // string | The entity the user subscribed to. Can be either `project` or `task`.
	entityID := "entityID_example" // string | The numeric id of the subscribed entity to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsEntityEntityIDDelete(context.Background(), entity, entityID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsEntityEntityIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsEntityEntityIDDelete`: ModelsSubscription
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsEntityEntityIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entity** | **string** | The entity the user subscribed to. Can be either &#x60;project&#x60; or &#x60;task&#x60;. | 
**entityID** | **string** | The numeric id of the subscribed entity to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsEntityEntityIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelsSubscription**](ModelsSubscription.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsEntityEntityIDPut

> ModelsSubscription SubscriptionsEntityEntityIDPut(ctx, entity, entityID).Execute()

Subscribes the current user to an entity.



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
	entity := "entity_example" // string | The entity the user subscribes to. Can be either `project` or `task`.
	entityID := "entityID_example" // string | The numeric id of the entity to subscribe to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsEntityEntityIDPut(context.Background(), entity, entityID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsEntityEntityIDPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsEntityEntityIDPut`: ModelsSubscription
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsEntityEntityIDPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entity** | **string** | The entity the user subscribes to. Can be either &#x60;project&#x60; or &#x60;task&#x60;. | 
**entityID** | **string** | The numeric id of the entity to subscribe to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsEntityEntityIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelsSubscription**](ModelsSubscription.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

