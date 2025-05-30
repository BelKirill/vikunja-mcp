# \ApiAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RoutesGet**](ApiAPI.md#RoutesGet) | **Get** /routes | Get a list of all token api routes
[**TokensGet**](ApiAPI.md#TokensGet) | **Get** /tokens | Get all api tokens of the current user
[**TokensPut**](ApiAPI.md#TokensPut) | **Put** /tokens | Create a new api token
[**TokensTokenIDDelete**](ApiAPI.md#TokensTokenIDDelete) | **Delete** /tokens/{tokenID} | Deletes an existing api token



## RoutesGet

> []map[string]ModelsRouteDetail RoutesGet(ctx).Execute()

Get a list of all token api routes



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
	resp, r, err := apiClient.ApiAPI.RoutesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.RoutesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RoutesGet`: []map[string]ModelsRouteDetail
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.RoutesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRoutesGetRequest struct via the builder pattern


### Return type

[**[]map[string]ModelsRouteDetail**](map.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokensGet

> []ModelsAPIToken TokensGet(ctx).Page(page).PerPage(perPage).S(s).Execute()

Get all api tokens of the current user



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
	page := int32(56) // int32 | The page number, used for pagination. If not provided, the first page of results is returned. (optional)
	perPage := int32(56) // int32 | The maximum number of tokens per page. This parameter is limited by the configured maximum of items per page. (optional)
	s := "s_example" // string | Search tokens by their title. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.TokensGet(context.Background()).Page(page).PerPage(perPage).S(s).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.TokensGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokensGet`: []ModelsAPIToken
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.TokensGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokensGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page number, used for pagination. If not provided, the first page of results is returned. | 
 **perPage** | **int32** | The maximum number of tokens per page. This parameter is limited by the configured maximum of items per page. | 
 **s** | **string** | Search tokens by their title. | 

### Return type

[**[]ModelsAPIToken**](ModelsAPIToken.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokensPut

> ModelsAPIToken TokensPut(ctx).Token(token).Execute()

Create a new api token



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
	token := *openapiclient.NewModelsAPIToken() // ModelsAPIToken | The token object with required fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.TokensPut(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.TokensPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokensPut`: ModelsAPIToken
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.TokensPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokensPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | [**ModelsAPIToken**](ModelsAPIToken.md) | The token object with required fields | 

### Return type

[**ModelsAPIToken**](ModelsAPIToken.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokensTokenIDDelete

> ModelsMessage TokensTokenIDDelete(ctx, tokenID).Execute()

Deletes an existing api token



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
	tokenID := int32(56) // int32 | Token ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.TokensTokenIDDelete(context.Background(), tokenID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.TokensTokenIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokensTokenIDDelete`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.TokensTokenIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenID** | **int32** | Token ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTokensTokenIDDeleteRequest struct via the builder pattern


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

