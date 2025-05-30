# \MigrationAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MigrationMicrosoftTodoAuthGet**](MigrationAPI.md#MigrationMicrosoftTodoAuthGet) | **Get** /migration/microsoft-todo/auth | Get the auth url from Microsoft Todo
[**MigrationMicrosoftTodoMigratePost**](MigrationAPI.md#MigrationMicrosoftTodoMigratePost) | **Post** /migration/microsoft-todo/migrate | Migrate all projects, tasks etc. from Microsoft Todo
[**MigrationMicrosoftTodoStatusGet**](MigrationAPI.md#MigrationMicrosoftTodoStatusGet) | **Get** /migration/microsoft-todo/status | Get migration status
[**MigrationTicktickMigratePost**](MigrationAPI.md#MigrationTicktickMigratePost) | **Post** /migration/ticktick/migrate | Import all projects, tasks etc. from a TickTick backup export
[**MigrationTicktickStatusGet**](MigrationAPI.md#MigrationTicktickStatusGet) | **Get** /migration/ticktick/status | Get migration status
[**MigrationTodoistAuthGet**](MigrationAPI.md#MigrationTodoistAuthGet) | **Get** /migration/todoist/auth | Get the auth url from todoist
[**MigrationTodoistMigratePost**](MigrationAPI.md#MigrationTodoistMigratePost) | **Post** /migration/todoist/migrate | Migrate all lists, tasks etc. from todoist
[**MigrationTodoistStatusGet**](MigrationAPI.md#MigrationTodoistStatusGet) | **Get** /migration/todoist/status | Get migration status
[**MigrationTrelloAuthGet**](MigrationAPI.md#MigrationTrelloAuthGet) | **Get** /migration/trello/auth | Get the auth url from trello
[**MigrationTrelloMigratePost**](MigrationAPI.md#MigrationTrelloMigratePost) | **Post** /migration/trello/migrate | Migrate all projects, tasks etc. from trello
[**MigrationTrelloStatusGet**](MigrationAPI.md#MigrationTrelloStatusGet) | **Get** /migration/trello/status | Get migration status
[**MigrationVikunjaFileMigratePost**](MigrationAPI.md#MigrationVikunjaFileMigratePost) | **Post** /migration/vikunja-file/migrate | Import all projects, tasks etc. from a Vikunja data export
[**MigrationVikunjaFileStatusGet**](MigrationAPI.md#MigrationVikunjaFileStatusGet) | **Get** /migration/vikunja-file/status | Get migration status



## MigrationMicrosoftTodoAuthGet

> HandlerAuthURL MigrationMicrosoftTodoAuthGet(ctx).Execute()

Get the auth url from Microsoft Todo



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
	resp, r, err := apiClient.MigrationAPI.MigrationMicrosoftTodoAuthGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationAPI.MigrationMicrosoftTodoAuthGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MigrationMicrosoftTodoAuthGet`: HandlerAuthURL
	fmt.Fprintf(os.Stdout, "Response from `MigrationAPI.MigrationMicrosoftTodoAuthGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationMicrosoftTodoAuthGetRequest struct via the builder pattern


### Return type

[**HandlerAuthURL**](HandlerAuthURL.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationMicrosoftTodoMigratePost

> ModelsMessage MigrationMicrosoftTodoMigratePost(ctx).MigrationCode(migrationCode).Execute()

Migrate all projects, tasks etc. from Microsoft Todo



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
	migrationCode := *openapiclient.NewMicrosofttodoMigration() // MicrosofttodoMigration | The auth token previously obtained from the auth url. See the docs for /migration/microsoft-todo/auth.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MigrationAPI.MigrationMicrosoftTodoMigratePost(context.Background()).MigrationCode(migrationCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationAPI.MigrationMicrosoftTodoMigratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MigrationMicrosoftTodoMigratePost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `MigrationAPI.MigrationMicrosoftTodoMigratePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMigrationMicrosoftTodoMigratePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **migrationCode** | [**MicrosofttodoMigration**](MicrosofttodoMigration.md) | The auth token previously obtained from the auth url. See the docs for /migration/microsoft-todo/auth. | 

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


## MigrationMicrosoftTodoStatusGet

> MigrationStatus MigrationMicrosoftTodoStatusGet(ctx).Execute()

Get migration status



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
	resp, r, err := apiClient.MigrationAPI.MigrationMicrosoftTodoStatusGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationAPI.MigrationMicrosoftTodoStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MigrationMicrosoftTodoStatusGet`: MigrationStatus
	fmt.Fprintf(os.Stdout, "Response from `MigrationAPI.MigrationMicrosoftTodoStatusGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationMicrosoftTodoStatusGetRequest struct via the builder pattern


### Return type

[**MigrationStatus**](MigrationStatus.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationTicktickMigratePost

> ModelsMessage MigrationTicktickMigratePost(ctx).Import_(import_).Execute()

Import all projects, tasks etc. from a TickTick backup export



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
	import_ := "import__example" // string | The TickTick backup csv file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MigrationAPI.MigrationTicktickMigratePost(context.Background()).Import_(import_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationAPI.MigrationTicktickMigratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MigrationTicktickMigratePost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `MigrationAPI.MigrationTicktickMigratePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMigrationTicktickMigratePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **import_** | **string** | The TickTick backup csv file. | 

### Return type

[**ModelsMessage**](ModelsMessage.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationTicktickStatusGet

> MigrationStatus MigrationTicktickStatusGet(ctx).Execute()

Get migration status



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
	resp, r, err := apiClient.MigrationAPI.MigrationTicktickStatusGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationAPI.MigrationTicktickStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MigrationTicktickStatusGet`: MigrationStatus
	fmt.Fprintf(os.Stdout, "Response from `MigrationAPI.MigrationTicktickStatusGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationTicktickStatusGetRequest struct via the builder pattern


### Return type

[**MigrationStatus**](MigrationStatus.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationTodoistAuthGet

> HandlerAuthURL MigrationTodoistAuthGet(ctx).Execute()

Get the auth url from todoist



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
	resp, r, err := apiClient.MigrationAPI.MigrationTodoistAuthGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationAPI.MigrationTodoistAuthGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MigrationTodoistAuthGet`: HandlerAuthURL
	fmt.Fprintf(os.Stdout, "Response from `MigrationAPI.MigrationTodoistAuthGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationTodoistAuthGetRequest struct via the builder pattern


### Return type

[**HandlerAuthURL**](HandlerAuthURL.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationTodoistMigratePost

> ModelsMessage MigrationTodoistMigratePost(ctx).MigrationCode(migrationCode).Execute()

Migrate all lists, tasks etc. from todoist



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
	migrationCode := *openapiclient.NewTodoistMigration() // TodoistMigration | The auth code previously obtained from the auth url. See the docs for /migration/todoist/auth.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MigrationAPI.MigrationTodoistMigratePost(context.Background()).MigrationCode(migrationCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationAPI.MigrationTodoistMigratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MigrationTodoistMigratePost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `MigrationAPI.MigrationTodoistMigratePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMigrationTodoistMigratePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **migrationCode** | [**TodoistMigration**](TodoistMigration.md) | The auth code previously obtained from the auth url. See the docs for /migration/todoist/auth. | 

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


## MigrationTodoistStatusGet

> MigrationStatus MigrationTodoistStatusGet(ctx).Execute()

Get migration status



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
	resp, r, err := apiClient.MigrationAPI.MigrationTodoistStatusGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationAPI.MigrationTodoistStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MigrationTodoistStatusGet`: MigrationStatus
	fmt.Fprintf(os.Stdout, "Response from `MigrationAPI.MigrationTodoistStatusGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationTodoistStatusGetRequest struct via the builder pattern


### Return type

[**MigrationStatus**](MigrationStatus.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationTrelloAuthGet

> HandlerAuthURL MigrationTrelloAuthGet(ctx).Execute()

Get the auth url from trello



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
	resp, r, err := apiClient.MigrationAPI.MigrationTrelloAuthGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationAPI.MigrationTrelloAuthGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MigrationTrelloAuthGet`: HandlerAuthURL
	fmt.Fprintf(os.Stdout, "Response from `MigrationAPI.MigrationTrelloAuthGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationTrelloAuthGetRequest struct via the builder pattern


### Return type

[**HandlerAuthURL**](HandlerAuthURL.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationTrelloMigratePost

> ModelsMessage MigrationTrelloMigratePost(ctx).MigrationCode(migrationCode).Execute()

Migrate all projects, tasks etc. from trello



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
	migrationCode := *openapiclient.NewTrelloMigration() // TrelloMigration | The auth token previously obtained from the auth url. See the docs for /migration/trello/auth.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MigrationAPI.MigrationTrelloMigratePost(context.Background()).MigrationCode(migrationCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationAPI.MigrationTrelloMigratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MigrationTrelloMigratePost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `MigrationAPI.MigrationTrelloMigratePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMigrationTrelloMigratePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **migrationCode** | [**TrelloMigration**](TrelloMigration.md) | The auth token previously obtained from the auth url. See the docs for /migration/trello/auth. | 

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


## MigrationTrelloStatusGet

> MigrationStatus MigrationTrelloStatusGet(ctx).Execute()

Get migration status



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
	resp, r, err := apiClient.MigrationAPI.MigrationTrelloStatusGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationAPI.MigrationTrelloStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MigrationTrelloStatusGet`: MigrationStatus
	fmt.Fprintf(os.Stdout, "Response from `MigrationAPI.MigrationTrelloStatusGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationTrelloStatusGetRequest struct via the builder pattern


### Return type

[**MigrationStatus**](MigrationStatus.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationVikunjaFileMigratePost

> ModelsMessage MigrationVikunjaFileMigratePost(ctx).Import_(import_).Execute()

Import all projects, tasks etc. from a Vikunja data export



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
	import_ := "import__example" // string | The Vikunja export zip file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MigrationAPI.MigrationVikunjaFileMigratePost(context.Background()).Import_(import_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationAPI.MigrationVikunjaFileMigratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MigrationVikunjaFileMigratePost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `MigrationAPI.MigrationVikunjaFileMigratePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMigrationVikunjaFileMigratePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **import_** | **string** | The Vikunja export zip file. | 

### Return type

[**ModelsMessage**](ModelsMessage.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationVikunjaFileStatusGet

> MigrationStatus MigrationVikunjaFileStatusGet(ctx).Execute()

Get migration status



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
	resp, r, err := apiClient.MigrationAPI.MigrationVikunjaFileStatusGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationAPI.MigrationVikunjaFileStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MigrationVikunjaFileStatusGet`: MigrationStatus
	fmt.Fprintf(os.Stdout, "Response from `MigrationAPI.MigrationVikunjaFileStatusGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationVikunjaFileStatusGetRequest struct via the builder pattern


### Return type

[**MigrationStatus**](MigrationStatus.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

