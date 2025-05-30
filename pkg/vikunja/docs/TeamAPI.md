# \TeamAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsGet**](TeamAPI.md#TeamsGet) | **Get** /teams | Get teams
[**TeamsIdDelete**](TeamAPI.md#TeamsIdDelete) | **Delete** /teams/{id} | Deletes a team
[**TeamsIdGet**](TeamAPI.md#TeamsIdGet) | **Get** /teams/{id} | Gets one team
[**TeamsIdMembersPut**](TeamAPI.md#TeamsIdMembersPut) | **Put** /teams/{id}/members | Add a user to a team
[**TeamsIdMembersUserIDAdminPost**](TeamAPI.md#TeamsIdMembersUserIDAdminPost) | **Post** /teams/{id}/members/{userID}/admin | Toggle a team member&#39;s admin status
[**TeamsIdMembersUserIDDelete**](TeamAPI.md#TeamsIdMembersUserIDDelete) | **Delete** /teams/{id}/members/{userID} | Remove a user from a team
[**TeamsIdPost**](TeamAPI.md#TeamsIdPost) | **Post** /teams/{id} | Updates a team
[**TeamsPut**](TeamAPI.md#TeamsPut) | **Put** /teams | Creates a new team



## TeamsGet

> []ModelsTeam TeamsGet(ctx).Page(page).PerPage(perPage).S(s).Execute()

Get teams



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
	s := "s_example" // string | Search teams by its name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.TeamsGet(context.Background()).Page(page).PerPage(perPage).S(s).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.TeamsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsGet`: []ModelsTeam
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.TeamsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page number. Used for pagination. If not provided, the first page of results is returned. | 
 **perPage** | **int32** | The maximum number of items per page. Note this parameter is limited by the configured maximum of items per page. | 
 **s** | **string** | Search teams by its name. | 

### Return type

[**[]ModelsTeam**](ModelsTeam.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsIdDelete

> ModelsMessage TeamsIdDelete(ctx, id).Execute()

Deletes a team



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
	id := int32(56) // int32 | Team ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.TeamsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.TeamsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsIdDelete`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.TeamsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Team ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsIdDeleteRequest struct via the builder pattern


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


## TeamsIdGet

> ModelsTeam TeamsIdGet(ctx, id).Execute()

Gets one team



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
	id := int32(56) // int32 | Team ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.TeamsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.TeamsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsIdGet`: ModelsTeam
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.TeamsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Team ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsTeam**](ModelsTeam.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsIdMembersPut

> ModelsTeamMember TeamsIdMembersPut(ctx, id).Team(team).Execute()

Add a user to a team



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
	id := int32(56) // int32 | Team ID
	team := *openapiclient.NewModelsTeamMember() // ModelsTeamMember | The user to be added to a team.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.TeamsIdMembersPut(context.Background(), id).Team(team).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.TeamsIdMembersPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsIdMembersPut`: ModelsTeamMember
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.TeamsIdMembersPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Team ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsIdMembersPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **team** | [**ModelsTeamMember**](ModelsTeamMember.md) | The user to be added to a team. | 

### Return type

[**ModelsTeamMember**](ModelsTeamMember.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsIdMembersUserIDAdminPost

> ModelsMessage TeamsIdMembersUserIDAdminPost(ctx, id, userID).Execute()

Toggle a team member's admin status



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
	id := int32(56) // int32 | Team ID
	userID := int32(56) // int32 | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.TeamsIdMembersUserIDAdminPost(context.Background(), id, userID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.TeamsIdMembersUserIDAdminPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsIdMembersUserIDAdminPost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.TeamsIdMembersUserIDAdminPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Team ID | 
**userID** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsIdMembersUserIDAdminPostRequest struct via the builder pattern


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


## TeamsIdMembersUserIDDelete

> ModelsMessage TeamsIdMembersUserIDDelete(ctx, id, userID).Execute()

Remove a user from a team



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
	id := int32(56) // int32 | Team ID
	userID := int32(56) // int32 | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.TeamsIdMembersUserIDDelete(context.Background(), id, userID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.TeamsIdMembersUserIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsIdMembersUserIDDelete`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.TeamsIdMembersUserIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Team ID | 
**userID** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsIdMembersUserIDDeleteRequest struct via the builder pattern


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


## TeamsIdPost

> ModelsTeam TeamsIdPost(ctx, id).Team(team).Execute()

Updates a team



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
	id := int32(56) // int32 | Team ID
	team := *openapiclient.NewModelsTeam() // ModelsTeam | The team with updated values you want to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.TeamsIdPost(context.Background(), id).Team(team).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.TeamsIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsIdPost`: ModelsTeam
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.TeamsIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Team ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **team** | [**ModelsTeam**](ModelsTeam.md) | The team with updated values you want to update. | 

### Return type

[**ModelsTeam**](ModelsTeam.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPut

> ModelsTeam TeamsPut(ctx).Team(team).Execute()

Creates a new team



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
	team := *openapiclient.NewModelsTeam() // ModelsTeam | The team you want to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.TeamsPut(context.Background()).Team(team).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.TeamsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsPut`: ModelsTeam
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.TeamsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **team** | [**ModelsTeam**](ModelsTeam.md) | The team you want to create. | 

### Return type

[**ModelsTeam**](ModelsTeam.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

