# \ProjectAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackgroundsUnsplashImageImageGet**](ProjectAPI.md#BackgroundsUnsplashImageImageGet) | **Get** /backgrounds/unsplash/image/{image} | Get an unsplash image
[**BackgroundsUnsplashImageImageThumbGet**](ProjectAPI.md#BackgroundsUnsplashImageImageThumbGet) | **Get** /backgrounds/unsplash/image/{image}/thumb | Get an unsplash thumbnail image
[**BackgroundsUnsplashSearchGet**](ProjectAPI.md#BackgroundsUnsplashSearchGet) | **Get** /backgrounds/unsplash/search | Search for a background from unsplash
[**ProjectsGet**](ProjectAPI.md#ProjectsGet) | **Get** /projects | Get all projects a user has access to
[**ProjectsIdBackgroundDelete**](ProjectAPI.md#ProjectsIdBackgroundDelete) | **Delete** /projects/{id}/background | Remove a project background
[**ProjectsIdBackgroundGet**](ProjectAPI.md#ProjectsIdBackgroundGet) | **Get** /projects/{id}/background | Get the project background
[**ProjectsIdBackgroundsUnsplashPost**](ProjectAPI.md#ProjectsIdBackgroundsUnsplashPost) | **Post** /projects/{id}/backgrounds/unsplash | Set an unsplash photo as project background
[**ProjectsIdBackgroundsUploadPut**](ProjectAPI.md#ProjectsIdBackgroundsUploadPut) | **Put** /projects/{id}/backgrounds/upload | Upload a project background
[**ProjectsIdDelete**](ProjectAPI.md#ProjectsIdDelete) | **Delete** /projects/{id} | Deletes a project
[**ProjectsIdGet**](ProjectAPI.md#ProjectsIdGet) | **Get** /projects/{id} | Gets one project
[**ProjectsIdPost**](ProjectAPI.md#ProjectsIdPost) | **Post** /projects/{id} | Updates a project
[**ProjectsIdProjectusersGet**](ProjectAPI.md#ProjectsIdProjectusersGet) | **Get** /projects/{id}/projectusers | Get users
[**ProjectsIdViewsViewBucketsGet**](ProjectAPI.md#ProjectsIdViewsViewBucketsGet) | **Get** /projects/{id}/views/{view}/buckets | Get all kanban buckets of a project
[**ProjectsIdViewsViewBucketsPut**](ProjectAPI.md#ProjectsIdViewsViewBucketsPut) | **Put** /projects/{id}/views/{view}/buckets | Create a new bucket
[**ProjectsProjectIDDuplicatePut**](ProjectAPI.md#ProjectsProjectIDDuplicatePut) | **Put** /projects/{projectID}/duplicate | Duplicate an existing project
[**ProjectsProjectIDViewsViewBucketsBucketIDDelete**](ProjectAPI.md#ProjectsProjectIDViewsViewBucketsBucketIDDelete) | **Delete** /projects/{projectID}/views/{view}/buckets/{bucketID} | Deletes an existing bucket
[**ProjectsProjectIDViewsViewBucketsBucketIDPost**](ProjectAPI.md#ProjectsProjectIDViewsViewBucketsBucketIDPost) | **Post** /projects/{projectID}/views/{view}/buckets/{bucketID} | Update an existing bucket
[**ProjectsProjectViewsGet**](ProjectAPI.md#ProjectsProjectViewsGet) | **Get** /projects/{project}/views | Get all project views for a project
[**ProjectsProjectViewsIdDelete**](ProjectAPI.md#ProjectsProjectViewsIdDelete) | **Delete** /projects/{project}/views/{id} | Delete a project view
[**ProjectsProjectViewsIdGet**](ProjectAPI.md#ProjectsProjectViewsIdGet) | **Get** /projects/{project}/views/{id} | Get one project view
[**ProjectsProjectViewsIdPost**](ProjectAPI.md#ProjectsProjectViewsIdPost) | **Post** /projects/{project}/views/{id} | Updates a project view
[**ProjectsProjectViewsPut**](ProjectAPI.md#ProjectsProjectViewsPut) | **Put** /projects/{project}/views | Create a project view
[**ProjectsPut**](ProjectAPI.md#ProjectsPut) | **Put** /projects | Creates a new project



## BackgroundsUnsplashImageImageGet

> *os.File BackgroundsUnsplashImageImageGet(ctx, image).Execute()

Get an unsplash image



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
	image := int32(56) // int32 | Unsplash Image ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.BackgroundsUnsplashImageImageGet(context.Background(), image).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.BackgroundsUnsplashImageImageGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackgroundsUnsplashImageImageGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.BackgroundsUnsplashImageImageGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**image** | **int32** | Unsplash Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackgroundsUnsplashImageImageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackgroundsUnsplashImageImageThumbGet

> *os.File BackgroundsUnsplashImageImageThumbGet(ctx, image).Execute()

Get an unsplash thumbnail image



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
	image := int32(56) // int32 | Unsplash Image ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.BackgroundsUnsplashImageImageThumbGet(context.Background(), image).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.BackgroundsUnsplashImageImageThumbGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackgroundsUnsplashImageImageThumbGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.BackgroundsUnsplashImageImageThumbGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**image** | **int32** | Unsplash Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackgroundsUnsplashImageImageThumbGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackgroundsUnsplashSearchGet

> []BackgroundImage BackgroundsUnsplashSearchGet(ctx).S(s).P(p).Execute()

Search for a background from unsplash



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
	s := "s_example" // string | Search backgrounds from unsplash with this search term. (optional)
	p := int32(56) // int32 | The page number. Used for pagination. If not provided, the first page of results is returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.BackgroundsUnsplashSearchGet(context.Background()).S(s).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.BackgroundsUnsplashSearchGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackgroundsUnsplashSearchGet`: []BackgroundImage
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.BackgroundsUnsplashSearchGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackgroundsUnsplashSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **s** | **string** | Search backgrounds from unsplash with this search term. | 
 **p** | **int32** | The page number. Used for pagination. If not provided, the first page of results is returned. | 

### Return type

[**[]BackgroundImage**](BackgroundImage.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsGet

> []ModelsProject ProjectsGet(ctx).Page(page).PerPage(perPage).S(s).IsArchived(isArchived).Execute()

Get all projects a user has access to



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
	s := "s_example" // string | Search projects by title. (optional)
	isArchived := true // bool | If true, also returns all archived projects. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectsGet(context.Background()).Page(page).PerPage(perPage).S(s).IsArchived(isArchived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsGet`: []ModelsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page number. Used for pagination. If not provided, the first page of results is returned. | 
 **perPage** | **int32** | The maximum number of items per page. Note this parameter is limited by the configured maximum of items per page. | 
 **s** | **string** | Search projects by title. | 
 **isArchived** | **bool** | If true, also returns all archived projects. | 

### Return type

[**[]ModelsProject**](ModelsProject.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsIdBackgroundDelete

> ModelsProject ProjectsIdBackgroundDelete(ctx, id).Execute()

Remove a project background



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectsIdBackgroundDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectsIdBackgroundDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsIdBackgroundDelete`: ModelsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectsIdBackgroundDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsIdBackgroundDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsProject**](ModelsProject.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsIdBackgroundGet

> *os.File ProjectsIdBackgroundGet(ctx, id).Execute()

Get the project background



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectsIdBackgroundGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectsIdBackgroundGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsIdBackgroundGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectsIdBackgroundGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsIdBackgroundGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsIdBackgroundsUnsplashPost

> ModelsProject ProjectsIdBackgroundsUnsplashPost(ctx, id).Project(project).Execute()

Set an unsplash photo as project background



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
	project := *openapiclient.NewBackgroundImage() // BackgroundImage | The image you want to set as background

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectsIdBackgroundsUnsplashPost(context.Background(), id).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectsIdBackgroundsUnsplashPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsIdBackgroundsUnsplashPost`: ModelsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectsIdBackgroundsUnsplashPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsIdBackgroundsUnsplashPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | [**BackgroundImage**](BackgroundImage.md) | The image you want to set as background | 

### Return type

[**ModelsProject**](ModelsProject.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsIdBackgroundsUploadPut

> ModelsMessage ProjectsIdBackgroundsUploadPut(ctx, id).Background(background).Execute()

Upload a project background



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
	background := "background_example" // string | The file as single file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectsIdBackgroundsUploadPut(context.Background(), id).Background(background).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectsIdBackgroundsUploadPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsIdBackgroundsUploadPut`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectsIdBackgroundsUploadPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsIdBackgroundsUploadPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **background** | **string** | The file as single file. | 

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


## ProjectsIdDelete

> ModelsMessage ProjectsIdDelete(ctx, id).Execute()

Deletes a project



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsIdDelete`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsIdDeleteRequest struct via the builder pattern


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


## ProjectsIdGet

> ModelsProject ProjectsIdGet(ctx, id).Execute()

Gets one project



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsIdGet`: ModelsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsProject**](ModelsProject.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsIdPost

> ModelsProject ProjectsIdPost(ctx, id).Project(project).Execute()

Updates a project



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
	project := *openapiclient.NewModelsProject() // ModelsProject | The project with updated values you want to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectsIdPost(context.Background(), id).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectsIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsIdPost`: ModelsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectsIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | [**ModelsProject**](ModelsProject.md) | The project with updated values you want to update. | 

### Return type

[**ModelsProject**](ModelsProject.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsIdProjectusersGet

> []UserUser ProjectsIdProjectusersGet(ctx, id).S(s).Execute()

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
	id := int32(56) // int32 | Project ID
	s := "s_example" // string | Search for a user by its name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectsIdProjectusersGet(context.Background(), id).S(s).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectsIdProjectusersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsIdProjectusersGet`: []UserUser
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectsIdProjectusersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsIdProjectusersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **s** | **string** | Search for a user by its name. | 

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


## ProjectsIdViewsViewBucketsGet

> []ModelsBucket ProjectsIdViewsViewBucketsGet(ctx, id, view).Execute()

Get all kanban buckets of a project



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
	view := int32(56) // int32 | Project view ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectsIdViewsViewBucketsGet(context.Background(), id, view).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectsIdViewsViewBucketsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsIdViewsViewBucketsGet`: []ModelsBucket
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectsIdViewsViewBucketsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Project ID | 
**view** | **int32** | Project view ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsIdViewsViewBucketsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ModelsBucket**](ModelsBucket.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsIdViewsViewBucketsPut

> ModelsBucket ProjectsIdViewsViewBucketsPut(ctx, id, view).Bucket(bucket).Execute()

Create a new bucket



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
	id := int32(56) // int32 | Project Id
	view := int32(56) // int32 | Project view ID
	bucket := *openapiclient.NewModelsBucket() // ModelsBucket | The bucket object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectsIdViewsViewBucketsPut(context.Background(), id, view).Bucket(bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectsIdViewsViewBucketsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsIdViewsViewBucketsPut`: ModelsBucket
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectsIdViewsViewBucketsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Project Id | 
**view** | **int32** | Project view ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsIdViewsViewBucketsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bucket** | [**ModelsBucket**](ModelsBucket.md) | The bucket object | 

### Return type

[**ModelsBucket**](ModelsBucket.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectIDDuplicatePut

> ModelsProjectDuplicate ProjectsProjectIDDuplicatePut(ctx, projectID).Project(project).Execute()

Duplicate an existing project



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
	projectID := int32(56) // int32 | The project ID to duplicate
	project := *openapiclient.NewModelsProjectDuplicate() // ModelsProjectDuplicate | The target parent project which should hold the copied project.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectsProjectIDDuplicatePut(context.Background(), projectID).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectsProjectIDDuplicatePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectIDDuplicatePut`: ModelsProjectDuplicate
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectsProjectIDDuplicatePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectID** | **int32** | The project ID to duplicate | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectIDDuplicatePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | [**ModelsProjectDuplicate**](ModelsProjectDuplicate.md) | The target parent project which should hold the copied project. | 

### Return type

[**ModelsProjectDuplicate**](ModelsProjectDuplicate.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectIDViewsViewBucketsBucketIDDelete

> ModelsMessage ProjectsProjectIDViewsViewBucketsBucketIDDelete(ctx, projectID, bucketID, view).Execute()

Deletes an existing bucket



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
	projectID := int32(56) // int32 | Project Id
	bucketID := int32(56) // int32 | Bucket Id
	view := int32(56) // int32 | Project view ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectsProjectIDViewsViewBucketsBucketIDDelete(context.Background(), projectID, bucketID, view).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectsProjectIDViewsViewBucketsBucketIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectIDViewsViewBucketsBucketIDDelete`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectsProjectIDViewsViewBucketsBucketIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectID** | **int32** | Project Id | 
**bucketID** | **int32** | Bucket Id | 
**view** | **int32** | Project view ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectIDViewsViewBucketsBucketIDDeleteRequest struct via the builder pattern


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


## ProjectsProjectIDViewsViewBucketsBucketIDPost

> ModelsBucket ProjectsProjectIDViewsViewBucketsBucketIDPost(ctx, projectID, bucketID, view).Bucket(bucket).Execute()

Update an existing bucket



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
	projectID := int32(56) // int32 | Project Id
	bucketID := int32(56) // int32 | Bucket Id
	view := int32(56) // int32 | Project view ID
	bucket := *openapiclient.NewModelsBucket() // ModelsBucket | The bucket object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectsProjectIDViewsViewBucketsBucketIDPost(context.Background(), projectID, bucketID, view).Bucket(bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectsProjectIDViewsViewBucketsBucketIDPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectIDViewsViewBucketsBucketIDPost`: ModelsBucket
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectsProjectIDViewsViewBucketsBucketIDPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectID** | **int32** | Project Id | 
**bucketID** | **int32** | Bucket Id | 
**view** | **int32** | Project view ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectIDViewsViewBucketsBucketIDPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bucket** | [**ModelsBucket**](ModelsBucket.md) | The bucket object | 

### Return type

[**ModelsBucket**](ModelsBucket.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectViewsGet

> []ModelsProjectView ProjectsProjectViewsGet(ctx, project).Execute()

Get all project views for a project



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectsProjectViewsGet(context.Background(), project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectsProjectViewsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectViewsGet`: []ModelsProjectView
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectsProjectViewsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectViewsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ModelsProjectView**](ModelsProjectView.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectViewsIdDelete

> ModelsMessage ProjectsProjectViewsIdDelete(ctx, project, id).Execute()

Delete a project view



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
	id := int32(56) // int32 | Project View ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectsProjectViewsIdDelete(context.Background(), project, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectsProjectViewsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectViewsIdDelete`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectsProjectViewsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **int32** | Project ID | 
**id** | **int32** | Project View ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectViewsIdDeleteRequest struct via the builder pattern


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


## ProjectsProjectViewsIdGet

> ModelsProjectView ProjectsProjectViewsIdGet(ctx, project, id).Execute()

Get one project view



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
	id := int32(56) // int32 | Project View ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectsProjectViewsIdGet(context.Background(), project, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectsProjectViewsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectViewsIdGet`: ModelsProjectView
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectsProjectViewsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **int32** | Project ID | 
**id** | **int32** | Project View ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectViewsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelsProjectView**](ModelsProjectView.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectViewsIdPost

> ModelsProjectView ProjectsProjectViewsIdPost(ctx, project, id).View(view).Execute()

Updates a project view



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
	id := int32(56) // int32 | Project View ID
	view := *openapiclient.NewModelsProjectView() // ModelsProjectView | The project view with updated values you want to change.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectsProjectViewsIdPost(context.Background(), project, id).View(view).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectsProjectViewsIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectViewsIdPost`: ModelsProjectView
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectsProjectViewsIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **int32** | Project ID | 
**id** | **int32** | Project View ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectViewsIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **view** | [**ModelsProjectView**](ModelsProjectView.md) | The project view with updated values you want to change. | 

### Return type

[**ModelsProjectView**](ModelsProjectView.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectViewsPut

> ModelsProjectView ProjectsProjectViewsPut(ctx, project).View(view).Execute()

Create a project view



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
	view := *openapiclient.NewModelsProjectView() // ModelsProjectView | The project view you want to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectsProjectViewsPut(context.Background(), project).View(view).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectsProjectViewsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectViewsPut`: ModelsProjectView
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectsProjectViewsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectViewsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | [**ModelsProjectView**](ModelsProjectView.md) | The project view you want to create. | 

### Return type

[**ModelsProjectView**](ModelsProjectView.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsPut

> ModelsProject ProjectsPut(ctx).Project(project).Execute()

Creates a new project



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
	project := *openapiclient.NewModelsProject() // ModelsProject | The project you want to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectsPut(context.Background()).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsPut`: ModelsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | [**ModelsProject**](ModelsProject.md) | The project you want to create. | 

### Return type

[**ModelsProject**](ModelsProject.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

