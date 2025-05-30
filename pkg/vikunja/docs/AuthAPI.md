# \AuthAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTokenOpenid**](AuthAPI.md#GetTokenOpenid) | **Post** /auth/openid/{provider}/callback | Authenticate a user with OpenID Connect
[**LoginPost**](AuthAPI.md#LoginPost) | **Post** /login | Login
[**RegisterPost**](AuthAPI.md#RegisterPost) | **Post** /register | Register



## GetTokenOpenid

> AuthToken GetTokenOpenid(ctx, provider).Callback(callback).Execute()

Authenticate a user with OpenID Connect



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
	provider := int32(56) // int32 | The OpenID Connect provider key as returned by the /info endpoint
	callback := *openapiclient.NewOpenidCallback() // OpenidCallback | The openid callback

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.GetTokenOpenid(context.Background(), provider).Callback(callback).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.GetTokenOpenid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenOpenid`: AuthToken
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.GetTokenOpenid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **int32** | The OpenID Connect provider key as returned by the /info endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenOpenidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **callback** | [**OpenidCallback**](OpenidCallback.md) | The openid callback | 

### Return type

[**AuthToken**](AuthToken.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginPost

> AuthToken LoginPost(ctx).Credentials(credentials).Execute()

Login



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
	credentials := *openapiclient.NewUserLogin() // UserLogin | The login credentials

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.LoginPost(context.Background()).Credentials(credentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.LoginPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoginPost`: AuthToken
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.LoginPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentials** | [**UserLogin**](UserLogin.md) | The login credentials | 

### Return type

[**AuthToken**](AuthToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterPost

> UserUser RegisterPost(ctx).Credentials(credentials).Execute()

Register



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
	credentials := *openapiclient.NewUserAPIUserPassword() // UserAPIUserPassword | The user credentials

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.RegisterPost(context.Background()).Credentials(credentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.RegisterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterPost`: UserUser
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.RegisterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentials** | [**UserAPIUserPassword**](UserAPIUserPassword.md) | The user credentials | 

### Return type

[**UserUser**](UserUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

