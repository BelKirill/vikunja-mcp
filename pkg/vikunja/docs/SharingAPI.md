# \SharingAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationsPost**](SharingAPI.md#NotificationsPost) | **Post** /notifications | Mark all notifications of a user as read
[**ProjectsIdTeamsGet**](SharingAPI.md#ProjectsIdTeamsGet) | **Get** /projects/{id}/teams | Get teams on a project
[**ProjectsIdTeamsPut**](SharingAPI.md#ProjectsIdTeamsPut) | **Put** /projects/{id}/teams | Add a team to a project
[**ProjectsIdUsersGet**](SharingAPI.md#ProjectsIdUsersGet) | **Get** /projects/{id}/users | Get users on a project
[**ProjectsIdUsersPut**](SharingAPI.md#ProjectsIdUsersPut) | **Put** /projects/{id}/users | Add a user to a project
[**ProjectsProjectIDTeamsTeamIDDelete**](SharingAPI.md#ProjectsProjectIDTeamsTeamIDDelete) | **Delete** /projects/{projectID}/teams/{teamID} | Delete a team from a project
[**ProjectsProjectIDTeamsTeamIDPost**](SharingAPI.md#ProjectsProjectIDTeamsTeamIDPost) | **Post** /projects/{projectID}/teams/{teamID} | Update a team &lt;-&gt; project relation
[**ProjectsProjectIDUsersUserIDDelete**](SharingAPI.md#ProjectsProjectIDUsersUserIDDelete) | **Delete** /projects/{projectID}/users/{userID} | Delete a user from a project
[**ProjectsProjectIDUsersUserIDPost**](SharingAPI.md#ProjectsProjectIDUsersUserIDPost) | **Post** /projects/{projectID}/users/{userID} | Update a user &lt;-&gt; project relation
[**ProjectsProjectSharesGet**](SharingAPI.md#ProjectsProjectSharesGet) | **Get** /projects/{project}/shares | Get all link shares for a project
[**ProjectsProjectSharesPut**](SharingAPI.md#ProjectsProjectSharesPut) | **Put** /projects/{project}/shares | Share a project via link
[**ProjectsProjectSharesShareDelete**](SharingAPI.md#ProjectsProjectSharesShareDelete) | **Delete** /projects/{project}/shares/{share} | Remove a link share
[**ProjectsProjectSharesShareGet**](SharingAPI.md#ProjectsProjectSharesShareGet) | **Get** /projects/{project}/shares/{share} | Get one link shares for a project
[**SharesShareAuthPost**](SharingAPI.md#SharesShareAuthPost) | **Post** /shares/{share}/auth | Get an auth token for a share



## NotificationsPost

> ModelsMessage NotificationsPost(ctx).Execute()

Mark all notifications of a user as read

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
	resp, r, err := apiClient.SharingAPI.NotificationsPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharingAPI.NotificationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationsPost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `SharingAPI.NotificationsPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsPostRequest struct via the builder pattern


### Return type

[**ModelsMessage**](ModelsMessage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsIdTeamsGet

> []ModelsTeamWithRight ProjectsIdTeamsGet(ctx, id).Page(page).PerPage(perPage).S(s).Execute()

Get teams on a project



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
	perPage := int32(56) // int32 | The maximum number of items per page. Note this parameter is limited by the configured maximum of items per page. (optional)
	s := "s_example" // string | Search teams by its name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharingAPI.ProjectsIdTeamsGet(context.Background(), id).Page(page).PerPage(perPage).S(s).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharingAPI.ProjectsIdTeamsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsIdTeamsGet`: []ModelsTeamWithRight
	fmt.Fprintf(os.Stdout, "Response from `SharingAPI.ProjectsIdTeamsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsIdTeamsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page number. Used for pagination. If not provided, the first page of results is returned. | 
 **perPage** | **int32** | The maximum number of items per page. Note this parameter is limited by the configured maximum of items per page. | 
 **s** | **string** | Search teams by its name. | 

### Return type

[**[]ModelsTeamWithRight**](ModelsTeamWithRight.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsIdTeamsPut

> ModelsTeamProject ProjectsIdTeamsPut(ctx, id).Project(project).Execute()

Add a team to a project



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
	project := *openapiclient.NewModelsTeamProject() // ModelsTeamProject | The team you want to add to the project.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharingAPI.ProjectsIdTeamsPut(context.Background(), id).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharingAPI.ProjectsIdTeamsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsIdTeamsPut`: ModelsTeamProject
	fmt.Fprintf(os.Stdout, "Response from `SharingAPI.ProjectsIdTeamsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsIdTeamsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | [**ModelsTeamProject**](ModelsTeamProject.md) | The team you want to add to the project. | 

### Return type

[**ModelsTeamProject**](ModelsTeamProject.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsIdUsersGet

> []ModelsUserWithRight ProjectsIdUsersGet(ctx, id).Page(page).PerPage(perPage).S(s).Execute()

Get users on a project



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
	perPage := int32(56) // int32 | The maximum number of items per page. Note this parameter is limited by the configured maximum of items per page. (optional)
	s := "s_example" // string | Search users by its name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharingAPI.ProjectsIdUsersGet(context.Background(), id).Page(page).PerPage(perPage).S(s).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharingAPI.ProjectsIdUsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsIdUsersGet`: []ModelsUserWithRight
	fmt.Fprintf(os.Stdout, "Response from `SharingAPI.ProjectsIdUsersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsIdUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page number. Used for pagination. If not provided, the first page of results is returned. | 
 **perPage** | **int32** | The maximum number of items per page. Note this parameter is limited by the configured maximum of items per page. | 
 **s** | **string** | Search users by its name. | 

### Return type

[**[]ModelsUserWithRight**](ModelsUserWithRight.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsIdUsersPut

> ModelsProjectUser ProjectsIdUsersPut(ctx, id).Project(project).Execute()

Add a user to a project



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
	project := *openapiclient.NewModelsProjectUser() // ModelsProjectUser | The user you want to add to the project.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharingAPI.ProjectsIdUsersPut(context.Background(), id).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharingAPI.ProjectsIdUsersPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsIdUsersPut`: ModelsProjectUser
	fmt.Fprintf(os.Stdout, "Response from `SharingAPI.ProjectsIdUsersPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsIdUsersPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | [**ModelsProjectUser**](ModelsProjectUser.md) | The user you want to add to the project. | 

### Return type

[**ModelsProjectUser**](ModelsProjectUser.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectIDTeamsTeamIDDelete

> ModelsMessage ProjectsProjectIDTeamsTeamIDDelete(ctx, projectID, teamID).Execute()

Delete a team from a project



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
	projectID := int32(56) // int32 | Project ID
	teamID := int32(56) // int32 | Team ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharingAPI.ProjectsProjectIDTeamsTeamIDDelete(context.Background(), projectID, teamID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharingAPI.ProjectsProjectIDTeamsTeamIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectIDTeamsTeamIDDelete`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `SharingAPI.ProjectsProjectIDTeamsTeamIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectID** | **int32** | Project ID | 
**teamID** | **int32** | Team ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectIDTeamsTeamIDDeleteRequest struct via the builder pattern


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


## ProjectsProjectIDTeamsTeamIDPost

> ModelsTeamProject ProjectsProjectIDTeamsTeamIDPost(ctx, projectID, teamID).Project(project).Execute()

Update a team <-> project relation



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
	projectID := int32(56) // int32 | Project ID
	teamID := int32(56) // int32 | Team ID
	project := *openapiclient.NewModelsTeamProject() // ModelsTeamProject | The team you want to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharingAPI.ProjectsProjectIDTeamsTeamIDPost(context.Background(), projectID, teamID).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharingAPI.ProjectsProjectIDTeamsTeamIDPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectIDTeamsTeamIDPost`: ModelsTeamProject
	fmt.Fprintf(os.Stdout, "Response from `SharingAPI.ProjectsProjectIDTeamsTeamIDPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectID** | **int32** | Project ID | 
**teamID** | **int32** | Team ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectIDTeamsTeamIDPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **project** | [**ModelsTeamProject**](ModelsTeamProject.md) | The team you want to update. | 

### Return type

[**ModelsTeamProject**](ModelsTeamProject.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectIDUsersUserIDDelete

> ModelsMessage ProjectsProjectIDUsersUserIDDelete(ctx, projectID, userID).Execute()

Delete a user from a project



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
	projectID := int32(56) // int32 | Project ID
	userID := int32(56) // int32 | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharingAPI.ProjectsProjectIDUsersUserIDDelete(context.Background(), projectID, userID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharingAPI.ProjectsProjectIDUsersUserIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectIDUsersUserIDDelete`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `SharingAPI.ProjectsProjectIDUsersUserIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectID** | **int32** | Project ID | 
**userID** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectIDUsersUserIDDeleteRequest struct via the builder pattern


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


## ProjectsProjectIDUsersUserIDPost

> ModelsProjectUser ProjectsProjectIDUsersUserIDPost(ctx, projectID, userID).Project(project).Execute()

Update a user <-> project relation



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
	projectID := int32(56) // int32 | Project ID
	userID := int32(56) // int32 | User ID
	project := *openapiclient.NewModelsProjectUser() // ModelsProjectUser | The user you want to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharingAPI.ProjectsProjectIDUsersUserIDPost(context.Background(), projectID, userID).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharingAPI.ProjectsProjectIDUsersUserIDPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectIDUsersUserIDPost`: ModelsProjectUser
	fmt.Fprintf(os.Stdout, "Response from `SharingAPI.ProjectsProjectIDUsersUserIDPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectID** | **int32** | Project ID | 
**userID** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectIDUsersUserIDPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **project** | [**ModelsProjectUser**](ModelsProjectUser.md) | The user you want to update. | 

### Return type

[**ModelsProjectUser**](ModelsProjectUser.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectSharesGet

> []ModelsLinkSharing ProjectsProjectSharesGet(ctx, project).Page(page).PerPage(perPage).S(s).Execute()

Get all link shares for a project



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
	project := int32(56) // int32 | Project ID
	page := int32(56) // int32 | The page number. Used for pagination. If not provided, the first page of results is returned. (optional)
	perPage := int32(56) // int32 | The maximum number of items per page. Note this parameter is limited by the configured maximum of items per page. (optional)
	s := "s_example" // string | Search shares by hash. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharingAPI.ProjectsProjectSharesGet(context.Background(), project).Page(page).PerPage(perPage).S(s).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharingAPI.ProjectsProjectSharesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectSharesGet`: []ModelsLinkSharing
	fmt.Fprintf(os.Stdout, "Response from `SharingAPI.ProjectsProjectSharesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectSharesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page number. Used for pagination. If not provided, the first page of results is returned. | 
 **perPage** | **int32** | The maximum number of items per page. Note this parameter is limited by the configured maximum of items per page. | 
 **s** | **string** | Search shares by hash. | 

### Return type

[**[]ModelsLinkSharing**](ModelsLinkSharing.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectSharesPut

> ModelsLinkSharing ProjectsProjectSharesPut(ctx, project).Label(label).Execute()

Share a project via link



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
	project := int32(56) // int32 | Project ID
	label := *openapiclient.NewModelsLinkSharing() // ModelsLinkSharing | The new link share object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharingAPI.ProjectsProjectSharesPut(context.Background(), project).Label(label).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharingAPI.ProjectsProjectSharesPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectSharesPut`: ModelsLinkSharing
	fmt.Fprintf(os.Stdout, "Response from `SharingAPI.ProjectsProjectSharesPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectSharesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **label** | [**ModelsLinkSharing**](ModelsLinkSharing.md) | The new link share object | 

### Return type

[**ModelsLinkSharing**](ModelsLinkSharing.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectSharesShareDelete

> ModelsMessage ProjectsProjectSharesShareDelete(ctx, project, share).Execute()

Remove a link share



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
	project := int32(56) // int32 | Project ID
	share := int32(56) // int32 | Share Link ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharingAPI.ProjectsProjectSharesShareDelete(context.Background(), project, share).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharingAPI.ProjectsProjectSharesShareDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectSharesShareDelete`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `SharingAPI.ProjectsProjectSharesShareDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **int32** | Project ID | 
**share** | **int32** | Share Link ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectSharesShareDeleteRequest struct via the builder pattern


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


## ProjectsProjectSharesShareGet

> ModelsLinkSharing ProjectsProjectSharesShareGet(ctx, project, share).Execute()

Get one link shares for a project



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
	project := int32(56) // int32 | Project ID
	share := int32(56) // int32 | Share ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharingAPI.ProjectsProjectSharesShareGet(context.Background(), project, share).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharingAPI.ProjectsProjectSharesShareGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectSharesShareGet`: ModelsLinkSharing
	fmt.Fprintf(os.Stdout, "Response from `SharingAPI.ProjectsProjectSharesShareGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **int32** | Project ID | 
**share** | **int32** | Share ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectSharesShareGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelsLinkSharing**](ModelsLinkSharing.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesShareAuthPost

> AuthToken SharesShareAuthPost(ctx, share).Password(password).Execute()

Get an auth token for a share



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
	share := "share_example" // string | The share hash
	password := *openapiclient.NewV1LinkShareAuth() // V1LinkShareAuth | The password for link shares which require one.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharingAPI.SharesShareAuthPost(context.Background(), share).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharingAPI.SharesShareAuthPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharesShareAuthPost`: AuthToken
	fmt.Fprintf(os.Stdout, "Response from `SharingAPI.SharesShareAuthPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**share** | **string** | The share hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesShareAuthPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **password** | [**V1LinkShareAuth**](V1LinkShareAuth.md) | The password for link shares which require one. | 

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

