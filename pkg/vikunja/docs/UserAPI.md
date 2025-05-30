# \UserAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserConfirmPost**](UserAPI.md#UserConfirmPost) | **Post** /user/confirm | Confirm the email of a new user
[**UserDeletionCancelPost**](UserAPI.md#UserDeletionCancelPost) | **Post** /user/deletion/cancel | Abort a user deletion request
[**UserDeletionConfirmPost**](UserAPI.md#UserDeletionConfirmPost) | **Post** /user/deletion/confirm | Confirm a user deletion request
[**UserDeletionRequestPost**](UserAPI.md#UserDeletionRequestPost) | **Post** /user/deletion/request | Request the deletion of the user
[**UserExportDownloadPost**](UserAPI.md#UserExportDownloadPost) | **Post** /user/export/download | Download a user data export.
[**UserExportRequestPost**](UserAPI.md#UserExportRequestPost) | **Post** /user/export/request | Request a user data export.
[**UserGet**](UserAPI.md#UserGet) | **Get** /user | Get user information
[**UserPasswordPost**](UserAPI.md#UserPasswordPost) | **Post** /user/password | Change password
[**UserPasswordResetPost**](UserAPI.md#UserPasswordResetPost) | **Post** /user/password/reset | Resets a password
[**UserPasswordTokenPost**](UserAPI.md#UserPasswordTokenPost) | **Post** /user/password/token | Request password reset token
[**UserSettingsAvatarGet**](UserAPI.md#UserSettingsAvatarGet) | **Get** /user/settings/avatar | Return user avatar setting
[**UserSettingsAvatarPost**](UserAPI.md#UserSettingsAvatarPost) | **Post** /user/settings/avatar | Set the user&#39;s avatar
[**UserSettingsAvatarUploadPut**](UserAPI.md#UserSettingsAvatarUploadPut) | **Put** /user/settings/avatar/upload | Upload a user avatar
[**UserSettingsEmailPost**](UserAPI.md#UserSettingsEmailPost) | **Post** /user/settings/email | Update email address
[**UserSettingsGeneralPost**](UserAPI.md#UserSettingsGeneralPost) | **Post** /user/settings/general | Change general user settings of the current user.
[**UserSettingsTokenCaldavGet**](UserAPI.md#UserSettingsTokenCaldavGet) | **Get** /user/settings/token/caldav | Returns the caldav tokens for the current user
[**UserSettingsTokenCaldavIdGet**](UserAPI.md#UserSettingsTokenCaldavIdGet) | **Get** /user/settings/token/caldav/{id} | Delete a caldav token by id
[**UserSettingsTokenCaldavPut**](UserAPI.md#UserSettingsTokenCaldavPut) | **Put** /user/settings/token/caldav | Generate a caldav token
[**UserSettingsTotpDisablePost**](UserAPI.md#UserSettingsTotpDisablePost) | **Post** /user/settings/totp/disable | Disable totp settings
[**UserSettingsTotpEnablePost**](UserAPI.md#UserSettingsTotpEnablePost) | **Post** /user/settings/totp/enable | Enable a previously enrolled totp setting.
[**UserSettingsTotpEnrollPost**](UserAPI.md#UserSettingsTotpEnrollPost) | **Post** /user/settings/totp/enroll | Enroll a user into totp
[**UserSettingsTotpGet**](UserAPI.md#UserSettingsTotpGet) | **Get** /user/settings/totp | Totp setting for the current user
[**UserSettingsTotpQrcodeGet**](UserAPI.md#UserSettingsTotpQrcodeGet) | **Get** /user/settings/totp/qrcode | Totp QR Code
[**UserTimezonesGet**](UserAPI.md#UserTimezonesGet) | **Get** /user/timezones | Get all available time zones on this vikunja instance
[**UserTokenPost**](UserAPI.md#UserTokenPost) | **Post** /user/token | Renew user token
[**UsernameAvatarGet**](UserAPI.md#UsernameAvatarGet) | **Get** /{username}/avatar | User Avatar
[**UsersGet**](UserAPI.md#UsersGet) | **Get** /users | Get users



## UserConfirmPost

> ModelsMessage UserConfirmPost(ctx).Credentials(credentials).Execute()

Confirm the email of a new user



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
	credentials := *openapiclient.NewUserEmailConfirm() // UserEmailConfirm | The token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserConfirmPost(context.Background()).Credentials(credentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserConfirmPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserConfirmPost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserConfirmPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserConfirmPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentials** | [**UserEmailConfirm**](UserEmailConfirm.md) | The token. | 

### Return type

[**ModelsMessage**](ModelsMessage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserDeletionCancelPost

> ModelsMessage UserDeletionCancelPost(ctx).Credentials(credentials).Execute()

Abort a user deletion request



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
	credentials := *openapiclient.NewV1UserPasswordConfirmation() // V1UserPasswordConfirmation | The user password to confirm.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserDeletionCancelPost(context.Background()).Credentials(credentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserDeletionCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserDeletionCancelPost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserDeletionCancelPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserDeletionCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentials** | [**V1UserPasswordConfirmation**](V1UserPasswordConfirmation.md) | The user password to confirm. | 

### Return type

[**ModelsMessage**](ModelsMessage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserDeletionConfirmPost

> ModelsMessage UserDeletionConfirmPost(ctx).Credentials(credentials).Execute()

Confirm a user deletion request



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
	credentials := *openapiclient.NewV1UserDeletionRequestConfirm() // V1UserDeletionRequestConfirm | The token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserDeletionConfirmPost(context.Background()).Credentials(credentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserDeletionConfirmPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserDeletionConfirmPost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserDeletionConfirmPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserDeletionConfirmPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentials** | [**V1UserDeletionRequestConfirm**](V1UserDeletionRequestConfirm.md) | The token. | 

### Return type

[**ModelsMessage**](ModelsMessage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserDeletionRequestPost

> ModelsMessage UserDeletionRequestPost(ctx).Credentials(credentials).Execute()

Request the deletion of the user



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
	credentials := *openapiclient.NewV1UserPasswordConfirmation() // V1UserPasswordConfirmation | The user password.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserDeletionRequestPost(context.Background()).Credentials(credentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserDeletionRequestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserDeletionRequestPost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserDeletionRequestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserDeletionRequestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentials** | [**V1UserPasswordConfirmation**](V1UserPasswordConfirmation.md) | The user password. | 

### Return type

[**ModelsMessage**](ModelsMessage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserExportDownloadPost

> ModelsMessage UserExportDownloadPost(ctx).Password(password).Execute()

Download a user data export.

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
	password := *openapiclient.NewV1UserPasswordConfirmation() // V1UserPasswordConfirmation | User password to confirm the download.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserExportDownloadPost(context.Background()).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserExportDownloadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserExportDownloadPost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserExportDownloadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserExportDownloadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **password** | [**V1UserPasswordConfirmation**](V1UserPasswordConfirmation.md) | User password to confirm the download. | 

### Return type

[**ModelsMessage**](ModelsMessage.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserExportRequestPost

> ModelsMessage UserExportRequestPost(ctx).Password(password).Execute()

Request a user data export.

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
	password := *openapiclient.NewV1UserPasswordConfirmation() // V1UserPasswordConfirmation | User password to confirm the data export request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserExportRequestPost(context.Background()).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserExportRequestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserExportRequestPost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserExportRequestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserExportRequestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **password** | [**V1UserPasswordConfirmation**](V1UserPasswordConfirmation.md) | User password to confirm the data export request. | 

### Return type

[**ModelsMessage**](ModelsMessage.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGet

> V1UserWithSettings UserGet(ctx).Execute()

Get user information



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
	resp, r, err := apiClient.UserAPI.UserGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserGet`: V1UserWithSettings
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetRequest struct via the builder pattern


### Return type

[**V1UserWithSettings**](V1UserWithSettings.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserPasswordPost

> ModelsMessage UserPasswordPost(ctx).UserPassword(userPassword).Execute()

Change password



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
	userPassword := *openapiclient.NewV1UserPassword() // V1UserPassword | The current and new password.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserPasswordPost(context.Background()).UserPassword(userPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserPasswordPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserPasswordPost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserPasswordPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserPasswordPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userPassword** | [**V1UserPassword**](V1UserPassword.md) | The current and new password. | 

### Return type

[**ModelsMessage**](ModelsMessage.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserPasswordResetPost

> ModelsMessage UserPasswordResetPost(ctx).Credentials(credentials).Execute()

Resets a password



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
	credentials := *openapiclient.NewUserPasswordReset() // UserPasswordReset | The token with the new password.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserPasswordResetPost(context.Background()).Credentials(credentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserPasswordResetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserPasswordResetPost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserPasswordResetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserPasswordResetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentials** | [**UserPasswordReset**](UserPasswordReset.md) | The token with the new password. | 

### Return type

[**ModelsMessage**](ModelsMessage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserPasswordTokenPost

> ModelsMessage UserPasswordTokenPost(ctx).Credentials(credentials).Execute()

Request password reset token



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
	credentials := *openapiclient.NewUserPasswordTokenRequest() // UserPasswordTokenRequest | The username of the user to request a token for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserPasswordTokenPost(context.Background()).Credentials(credentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserPasswordTokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserPasswordTokenPost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserPasswordTokenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserPasswordTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentials** | [**UserPasswordTokenRequest**](UserPasswordTokenRequest.md) | The username of the user to request a token for. | 

### Return type

[**ModelsMessage**](ModelsMessage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSettingsAvatarGet

> V1UserAvatarProvider UserSettingsAvatarGet(ctx).Execute()

Return user avatar setting



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
	resp, r, err := apiClient.UserAPI.UserSettingsAvatarGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserSettingsAvatarGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSettingsAvatarGet`: V1UserAvatarProvider
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserSettingsAvatarGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserSettingsAvatarGetRequest struct via the builder pattern


### Return type

[**V1UserAvatarProvider**](V1UserAvatarProvider.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSettingsAvatarPost

> ModelsMessage UserSettingsAvatarPost(ctx).Avatar(avatar).Execute()

Set the user's avatar



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
	avatar := *openapiclient.NewV1UserAvatarProvider() // V1UserAvatarProvider | The user's avatar setting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserSettingsAvatarPost(context.Background()).Avatar(avatar).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserSettingsAvatarPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSettingsAvatarPost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserSettingsAvatarPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserSettingsAvatarPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avatar** | [**V1UserAvatarProvider**](V1UserAvatarProvider.md) | The user&#39;s avatar setting | 

### Return type

[**ModelsMessage**](ModelsMessage.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSettingsAvatarUploadPut

> ModelsMessage UserSettingsAvatarUploadPut(ctx).Avatar(avatar).Execute()

Upload a user avatar



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
	avatar := "avatar_example" // string | The avatar as single file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserSettingsAvatarUploadPut(context.Background()).Avatar(avatar).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserSettingsAvatarUploadPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSettingsAvatarUploadPut`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserSettingsAvatarUploadPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserSettingsAvatarUploadPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avatar** | **string** | The avatar as single file. | 

### Return type

[**ModelsMessage**](ModelsMessage.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSettingsEmailPost

> ModelsMessage UserSettingsEmailPost(ctx).UserEmailUpdate(userEmailUpdate).Execute()

Update email address



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
	userEmailUpdate := *openapiclient.NewUserEmailUpdate() // UserEmailUpdate | The new email address and current password.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserSettingsEmailPost(context.Background()).UserEmailUpdate(userEmailUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserSettingsEmailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSettingsEmailPost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserSettingsEmailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserSettingsEmailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userEmailUpdate** | [**UserEmailUpdate**](UserEmailUpdate.md) | The new email address and current password. | 

### Return type

[**ModelsMessage**](ModelsMessage.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSettingsGeneralPost

> ModelsMessage UserSettingsGeneralPost(ctx).Avatar(avatar).Execute()

Change general user settings of the current user.

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
	avatar := *openapiclient.NewV1UserSettings() // V1UserSettings | The updated user settings

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserSettingsGeneralPost(context.Background()).Avatar(avatar).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserSettingsGeneralPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSettingsGeneralPost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserSettingsGeneralPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserSettingsGeneralPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avatar** | [**V1UserSettings**](V1UserSettings.md) | The updated user settings | 

### Return type

[**ModelsMessage**](ModelsMessage.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSettingsTokenCaldavGet

> []UserToken UserSettingsTokenCaldavGet(ctx).Execute()

Returns the caldav tokens for the current user



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
	resp, r, err := apiClient.UserAPI.UserSettingsTokenCaldavGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserSettingsTokenCaldavGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSettingsTokenCaldavGet`: []UserToken
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserSettingsTokenCaldavGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserSettingsTokenCaldavGetRequest struct via the builder pattern


### Return type

[**[]UserToken**](UserToken.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSettingsTokenCaldavIdGet

> ModelsMessage UserSettingsTokenCaldavIdGet(ctx, id).Execute()

Delete a caldav token by id

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
	id := int32(56) // int32 | Token ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserSettingsTokenCaldavIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserSettingsTokenCaldavIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSettingsTokenCaldavIdGet`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserSettingsTokenCaldavIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Token ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserSettingsTokenCaldavIdGetRequest struct via the builder pattern


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


## UserSettingsTokenCaldavPut

> UserToken UserSettingsTokenCaldavPut(ctx).Execute()

Generate a caldav token



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
	resp, r, err := apiClient.UserAPI.UserSettingsTokenCaldavPut(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserSettingsTokenCaldavPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSettingsTokenCaldavPut`: UserToken
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserSettingsTokenCaldavPut`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserSettingsTokenCaldavPutRequest struct via the builder pattern


### Return type

[**UserToken**](UserToken.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSettingsTotpDisablePost

> ModelsMessage UserSettingsTotpDisablePost(ctx).Totp(totp).Execute()

Disable totp settings



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
	totp := *openapiclient.NewUserLogin() // UserLogin | The current user's password (only password is enough).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserSettingsTotpDisablePost(context.Background()).Totp(totp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserSettingsTotpDisablePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSettingsTotpDisablePost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserSettingsTotpDisablePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserSettingsTotpDisablePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **totp** | [**UserLogin**](UserLogin.md) | The current user&#39;s password (only password is enough). | 

### Return type

[**ModelsMessage**](ModelsMessage.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSettingsTotpEnablePost

> ModelsMessage UserSettingsTotpEnablePost(ctx).Totp(totp).Execute()

Enable a previously enrolled totp setting.



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
	totp := *openapiclient.NewUserTOTPPasscode() // UserTOTPPasscode | The totp passcode.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserSettingsTotpEnablePost(context.Background()).Totp(totp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserSettingsTotpEnablePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSettingsTotpEnablePost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserSettingsTotpEnablePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserSettingsTotpEnablePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **totp** | [**UserTOTPPasscode**](UserTOTPPasscode.md) | The totp passcode. | 

### Return type

[**ModelsMessage**](ModelsMessage.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSettingsTotpEnrollPost

> UserTOTP UserSettingsTotpEnrollPost(ctx).Execute()

Enroll a user into totp



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
	resp, r, err := apiClient.UserAPI.UserSettingsTotpEnrollPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserSettingsTotpEnrollPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSettingsTotpEnrollPost`: UserTOTP
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserSettingsTotpEnrollPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserSettingsTotpEnrollPostRequest struct via the builder pattern


### Return type

[**UserTOTP**](UserTOTP.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSettingsTotpGet

> UserTOTP UserSettingsTotpGet(ctx).Execute()

Totp setting for the current user



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
	resp, r, err := apiClient.UserAPI.UserSettingsTotpGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserSettingsTotpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSettingsTotpGet`: UserTOTP
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserSettingsTotpGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserSettingsTotpGetRequest struct via the builder pattern


### Return type

[**UserTOTP**](UserTOTP.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSettingsTotpQrcodeGet

> *os.File UserSettingsTotpQrcodeGet(ctx).Execute()

Totp QR Code



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
	resp, r, err := apiClient.UserAPI.UserSettingsTotpQrcodeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserSettingsTotpQrcodeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSettingsTotpQrcodeGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserSettingsTotpQrcodeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserSettingsTotpQrcodeGetRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserTimezonesGet

> []string UserTimezonesGet(ctx).Execute()

Get all available time zones on this vikunja instance



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
	resp, r, err := apiClient.UserAPI.UserTimezonesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserTimezonesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserTimezonesGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserTimezonesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserTimezonesGetRequest struct via the builder pattern


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


## UserTokenPost

> AuthToken UserTokenPost(ctx).Execute()

Renew user token



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
	resp, r, err := apiClient.UserAPI.UserTokenPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserTokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserTokenPost`: AuthToken
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserTokenPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserTokenPostRequest struct via the builder pattern


### Return type

[**AuthToken**](AuthToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsernameAvatarGet

> *os.File UsernameAvatarGet(ctx, username).Size(size).Execute()

User Avatar



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
	username := "username_example" // string | The username of the user who's avatar you want to get
	size := int32(56) // int32 | The size of the avatar you want to get. If bigger than the max configured size this will be adjusted to the maximum size. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UsernameAvatarGet(context.Background(), username).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UsernameAvatarGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsernameAvatarGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UsernameAvatarGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The username of the user who&#39;s avatar you want to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsernameAvatarGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **int32** | The size of the avatar you want to get. If bigger than the max configured size this will be adjusted to the maximum size. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGet

> []UserUser UsersGet(ctx).S(s).Execute()

Get users



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
	s := "s_example" // string | The search criteria. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UsersGet(context.Background()).S(s).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersGet`: []UserUser
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UsersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **s** | **string** | The search criteria. | 

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

