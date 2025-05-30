# \TaskAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KindIdReactionsDeletePost**](TaskAPI.md#KindIdReactionsDeletePost) | **Post** /{kind}/{id}/reactions/delete | Removes the user&#39;s reaction
[**KindIdReactionsGet**](TaskAPI.md#KindIdReactionsGet) | **Get** /{kind}/{id}/reactions | Get all reactions for an entity
[**KindIdReactionsPut**](TaskAPI.md#KindIdReactionsPut) | **Put** /{kind}/{id}/reactions | Add a reaction to an entity
[**ProjectsIdTasksPut**](TaskAPI.md#ProjectsIdTasksPut) | **Put** /projects/{id}/tasks | Create a task
[**ProjectsIdViewsViewTasksGet**](TaskAPI.md#ProjectsIdViewsViewTasksGet) | **Get** /projects/{id}/views/{view}/tasks | Get tasks in a project
[**ProjectsProjectViewsViewBucketsBucketTasksPost**](TaskAPI.md#ProjectsProjectViewsViewBucketsBucketTasksPost) | **Post** /projects/{project}/views/{view}/buckets/{bucket}/tasks | Update a task bucket
[**TasksAllGet**](TaskAPI.md#TasksAllGet) | **Get** /tasks/all | Get tasks
[**TasksBulkPost**](TaskAPI.md#TasksBulkPost) | **Post** /tasks/bulk | Update a bunch of tasks at once
[**TasksIdAttachmentsAttachmentIDDelete**](TaskAPI.md#TasksIdAttachmentsAttachmentIDDelete) | **Delete** /tasks/{id}/attachments/{attachmentID} | Delete an attachment
[**TasksIdAttachmentsAttachmentIDGet**](TaskAPI.md#TasksIdAttachmentsAttachmentIDGet) | **Get** /tasks/{id}/attachments/{attachmentID} | Get one attachment.
[**TasksIdAttachmentsGet**](TaskAPI.md#TasksIdAttachmentsGet) | **Get** /tasks/{id}/attachments | Get  all attachments for one task.
[**TasksIdAttachmentsPut**](TaskAPI.md#TasksIdAttachmentsPut) | **Put** /tasks/{id}/attachments | Upload a task attachment
[**TasksIdDelete**](TaskAPI.md#TasksIdDelete) | **Delete** /tasks/{id} | Delete a task
[**TasksIdGet**](TaskAPI.md#TasksIdGet) | **Get** /tasks/{id} | Get one task
[**TasksIdPositionPost**](TaskAPI.md#TasksIdPositionPost) | **Post** /tasks/{id}/position | Updates a task position
[**TasksIdPost**](TaskAPI.md#TasksIdPost) | **Post** /tasks/{id} | Update a task
[**TasksTaskIDCommentsCommentIDDelete**](TaskAPI.md#TasksTaskIDCommentsCommentIDDelete) | **Delete** /tasks/{taskID}/comments/{commentID} | Remove a task comment
[**TasksTaskIDCommentsCommentIDGet**](TaskAPI.md#TasksTaskIDCommentsCommentIDGet) | **Get** /tasks/{taskID}/comments/{commentID} | Remove a task comment
[**TasksTaskIDCommentsCommentIDPost**](TaskAPI.md#TasksTaskIDCommentsCommentIDPost) | **Post** /tasks/{taskID}/comments/{commentID} | Update an existing task comment
[**TasksTaskIDCommentsGet**](TaskAPI.md#TasksTaskIDCommentsGet) | **Get** /tasks/{taskID}/comments | Get all task comments
[**TasksTaskIDCommentsPut**](TaskAPI.md#TasksTaskIDCommentsPut) | **Put** /tasks/{taskID}/comments | Create a new task comment
[**TasksTaskIDRelationsPut**](TaskAPI.md#TasksTaskIDRelationsPut) | **Put** /tasks/{taskID}/relations | Create a new relation between two tasks
[**TasksTaskIDRelationsRelationKindOtherTaskIDDelete**](TaskAPI.md#TasksTaskIDRelationsRelationKindOtherTaskIDDelete) | **Delete** /tasks/{taskID}/relations/{relationKind}/{otherTaskID} | Remove a task relation



## KindIdReactionsDeletePost

> ModelsMessage KindIdReactionsDeletePost(ctx, id, kind).Project(project).Execute()

Removes the user's reaction



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
	id := int32(56) // int32 | Entity ID
	kind := int32(56) // int32 | The kind of the entity. Can be either `tasks` or `comments` for task comments
	project := *openapiclient.NewModelsReaction() // ModelsReaction | The reaction you want to add to the entity.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.KindIdReactionsDeletePost(context.Background(), id, kind).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.KindIdReactionsDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KindIdReactionsDeletePost`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.KindIdReactionsDeletePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Entity ID | 
**kind** | **int32** | The kind of the entity. Can be either &#x60;tasks&#x60; or &#x60;comments&#x60; for task comments | 

### Other Parameters

Other parameters are passed through a pointer to a apiKindIdReactionsDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **project** | [**ModelsReaction**](ModelsReaction.md) | The reaction you want to add to the entity. | 

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


## KindIdReactionsGet

> []map[string][]UserUser KindIdReactionsGet(ctx, id, kind).Execute()

Get all reactions for an entity



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
	id := int32(56) // int32 | Entity ID
	kind := int32(56) // int32 | The kind of the entity. Can be either `tasks` or `comments` for task comments

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.KindIdReactionsGet(context.Background(), id, kind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.KindIdReactionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KindIdReactionsGet`: []map[string][]UserUser
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.KindIdReactionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Entity ID | 
**kind** | **int32** | The kind of the entity. Can be either &#x60;tasks&#x60; or &#x60;comments&#x60; for task comments | 

### Other Parameters

Other parameters are passed through a pointer to a apiKindIdReactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]map[string][]UserUser**](map.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KindIdReactionsPut

> ModelsReaction KindIdReactionsPut(ctx, id, kind).Project(project).Execute()

Add a reaction to an entity



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
	id := int32(56) // int32 | Entity ID
	kind := int32(56) // int32 | The kind of the entity. Can be either `tasks` or `comments` for task comments
	project := *openapiclient.NewModelsReaction() // ModelsReaction | The reaction you want to add to the entity.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.KindIdReactionsPut(context.Background(), id, kind).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.KindIdReactionsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KindIdReactionsPut`: ModelsReaction
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.KindIdReactionsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Entity ID | 
**kind** | **int32** | The kind of the entity. Can be either &#x60;tasks&#x60; or &#x60;comments&#x60; for task comments | 

### Other Parameters

Other parameters are passed through a pointer to a apiKindIdReactionsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **project** | [**ModelsReaction**](ModelsReaction.md) | The reaction you want to add to the entity. | 

### Return type

[**ModelsReaction**](ModelsReaction.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsIdTasksPut

> ModelsTask ProjectsIdTasksPut(ctx, id).Task(task).Execute()

Create a task



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
	task := *openapiclient.NewModelsTask() // ModelsTask | The task object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.ProjectsIdTasksPut(context.Background(), id).Task(task).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.ProjectsIdTasksPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsIdTasksPut`: ModelsTask
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.ProjectsIdTasksPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsIdTasksPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **task** | [**ModelsTask**](ModelsTask.md) | The task object | 

### Return type

[**ModelsTask**](ModelsTask.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsIdViewsViewTasksGet

> []ModelsTask ProjectsIdViewsViewTasksGet(ctx, id, view).Page(page).PerPage(perPage).S(s).SortBy(sortBy).OrderBy(orderBy).Filter(filter).FilterTimezone(filterTimezone).FilterIncludeNulls(filterIncludeNulls).Expand(expand).Execute()

Get tasks in a project



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
	id := int32(56) // int32 | The project ID.
	view := int32(56) // int32 | The project view ID.
	page := int32(56) // int32 | The page number. Used for pagination. If not provided, the first page of results is returned. (optional)
	perPage := int32(56) // int32 | The maximum number of items per page. Note this parameter is limited by the configured maximum of items per page. (optional)
	s := "s_example" // string | Search tasks by task text. (optional)
	sortBy := "sortBy_example" // string | The sorting parameter. You can pass this multiple times to get the tasks ordered by multiple different parametes, along with `order_by`. Possible values to sort by are `id`, `title`, `description`, `done`, `done_at`, `due_date`, `created_by_id`, `project_id`, `repeat_after`, `priority`, `start_date`, `end_date`, `hex_color`, `percent_done`, `uid`, `created`, `updated`. Default is `id`. (optional)
	orderBy := "orderBy_example" // string | The ordering parameter. Possible values to order by are `asc` or `desc`. Default is `asc`. (optional)
	filter := "filter_example" // string | The filter query to match tasks by. Check out https://vikunja.io/docs/filters for a full explanation of the feature. (optional)
	filterTimezone := "filterTimezone_example" // string | The time zone which should be used for date match (statements like  (optional)
	filterIncludeNulls := "filterIncludeNulls_example" // string | If set to true the result will include filtered fields whose value is set to `null`. Available values are `true` or `false`. Defaults to `false`. (optional)
	expand := "expand_example" // string | If set to `subtasks`, Vikunja will fetch only tasks which do not have subtasks and then in a second step, will fetch all of these subtasks. This may result in more tasks than the pagination limit being returned, but all subtasks will be present in the response. You can only set this to `subtasks`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.ProjectsIdViewsViewTasksGet(context.Background(), id, view).Page(page).PerPage(perPage).S(s).SortBy(sortBy).OrderBy(orderBy).Filter(filter).FilterTimezone(filterTimezone).FilterIncludeNulls(filterIncludeNulls).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.ProjectsIdViewsViewTasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsIdViewsViewTasksGet`: []ModelsTask
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.ProjectsIdViewsViewTasksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The project ID. | 
**view** | **int32** | The project view ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsIdViewsViewTasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | The page number. Used for pagination. If not provided, the first page of results is returned. | 
 **perPage** | **int32** | The maximum number of items per page. Note this parameter is limited by the configured maximum of items per page. | 
 **s** | **string** | Search tasks by task text. | 
 **sortBy** | **string** | The sorting parameter. You can pass this multiple times to get the tasks ordered by multiple different parametes, along with &#x60;order_by&#x60;. Possible values to sort by are &#x60;id&#x60;, &#x60;title&#x60;, &#x60;description&#x60;, &#x60;done&#x60;, &#x60;done_at&#x60;, &#x60;due_date&#x60;, &#x60;created_by_id&#x60;, &#x60;project_id&#x60;, &#x60;repeat_after&#x60;, &#x60;priority&#x60;, &#x60;start_date&#x60;, &#x60;end_date&#x60;, &#x60;hex_color&#x60;, &#x60;percent_done&#x60;, &#x60;uid&#x60;, &#x60;created&#x60;, &#x60;updated&#x60;. Default is &#x60;id&#x60;. | 
 **orderBy** | **string** | The ordering parameter. Possible values to order by are &#x60;asc&#x60; or &#x60;desc&#x60;. Default is &#x60;asc&#x60;. | 
 **filter** | **string** | The filter query to match tasks by. Check out https://vikunja.io/docs/filters for a full explanation of the feature. | 
 **filterTimezone** | **string** | The time zone which should be used for date match (statements like  | 
 **filterIncludeNulls** | **string** | If set to true the result will include filtered fields whose value is set to &#x60;null&#x60;. Available values are &#x60;true&#x60; or &#x60;false&#x60;. Defaults to &#x60;false&#x60;. | 
 **expand** | **string** | If set to &#x60;subtasks&#x60;, Vikunja will fetch only tasks which do not have subtasks and then in a second step, will fetch all of these subtasks. This may result in more tasks than the pagination limit being returned, but all subtasks will be present in the response. You can only set this to &#x60;subtasks&#x60;. | 

### Return type

[**[]ModelsTask**](ModelsTask.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectViewsViewBucketsBucketTasksPost

> ModelsTaskBucket ProjectsProjectViewsViewBucketsBucketTasksPost(ctx, project, view, bucket).TaskBucket(taskBucket).Execute()

Update a task bucket



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
	view := int32(56) // int32 | Project View ID
	bucket := int32(56) // int32 | Bucket ID
	taskBucket := *openapiclient.NewModelsTaskBucket() // ModelsTaskBucket | The id of the task you want to move into the bucket.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.ProjectsProjectViewsViewBucketsBucketTasksPost(context.Background(), project, view, bucket).TaskBucket(taskBucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.ProjectsProjectViewsViewBucketsBucketTasksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectViewsViewBucketsBucketTasksPost`: ModelsTaskBucket
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.ProjectsProjectViewsViewBucketsBucketTasksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **int32** | Project ID | 
**view** | **int32** | Project View ID | 
**bucket** | **int32** | Bucket ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectViewsViewBucketsBucketTasksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **taskBucket** | [**ModelsTaskBucket**](ModelsTaskBucket.md) | The id of the task you want to move into the bucket. | 

### Return type

[**ModelsTaskBucket**](ModelsTaskBucket.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksAllGet

> []ModelsTask TasksAllGet(ctx).Page(page).PerPage(perPage).S(s).SortBy(sortBy).OrderBy(orderBy).Filter(filter).FilterTimezone(filterTimezone).FilterIncludeNulls(filterIncludeNulls).Expand(expand).Execute()

Get tasks



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
	s := "s_example" // string | Search tasks by task text. (optional)
	sortBy := "sortBy_example" // string | The sorting parameter. You can pass this multiple times to get the tasks ordered by multiple different parametes, along with `order_by`. Possible values to sort by are `id`, `title`, `description`, `done`, `done_at`, `due_date`, `created_by_id`, `project_id`, `repeat_after`, `priority`, `start_date`, `end_date`, `hex_color`, `percent_done`, `uid`, `created`, `updated`. Default is `id`. (optional)
	orderBy := "orderBy_example" // string | The ordering parameter. Possible values to order by are `asc` or `desc`. Default is `asc`. (optional)
	filter := "filter_example" // string | The filter query to match tasks by. Check out https://vikunja.io/docs/filters for a full explanation of the feature. (optional)
	filterTimezone := "filterTimezone_example" // string | The time zone which should be used for date match (statements like  (optional)
	filterIncludeNulls := "filterIncludeNulls_example" // string | If set to true the result will include filtered fields whose value is set to `null`. Available values are `true` or `false`. Defaults to `false`. (optional)
	expand := "expand_example" // string | If set to `subtasks`, Vikunja will fetch only tasks which do not have subtasks and then in a second step, will fetch all of these subtasks. This may result in more tasks than the pagination limit being returned, but all subtasks will be present in the response. You can only set this to `subtasks`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.TasksAllGet(context.Background()).Page(page).PerPage(perPage).S(s).SortBy(sortBy).OrderBy(orderBy).Filter(filter).FilterTimezone(filterTimezone).FilterIncludeNulls(filterIncludeNulls).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksAllGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksAllGet`: []ModelsTask
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.TasksAllGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTasksAllGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page number. Used for pagination. If not provided, the first page of results is returned. | 
 **perPage** | **int32** | The maximum number of items per page. Note this parameter is limited by the configured maximum of items per page. | 
 **s** | **string** | Search tasks by task text. | 
 **sortBy** | **string** | The sorting parameter. You can pass this multiple times to get the tasks ordered by multiple different parametes, along with &#x60;order_by&#x60;. Possible values to sort by are &#x60;id&#x60;, &#x60;title&#x60;, &#x60;description&#x60;, &#x60;done&#x60;, &#x60;done_at&#x60;, &#x60;due_date&#x60;, &#x60;created_by_id&#x60;, &#x60;project_id&#x60;, &#x60;repeat_after&#x60;, &#x60;priority&#x60;, &#x60;start_date&#x60;, &#x60;end_date&#x60;, &#x60;hex_color&#x60;, &#x60;percent_done&#x60;, &#x60;uid&#x60;, &#x60;created&#x60;, &#x60;updated&#x60;. Default is &#x60;id&#x60;. | 
 **orderBy** | **string** | The ordering parameter. Possible values to order by are &#x60;asc&#x60; or &#x60;desc&#x60;. Default is &#x60;asc&#x60;. | 
 **filter** | **string** | The filter query to match tasks by. Check out https://vikunja.io/docs/filters for a full explanation of the feature. | 
 **filterTimezone** | **string** | The time zone which should be used for date match (statements like  | 
 **filterIncludeNulls** | **string** | If set to true the result will include filtered fields whose value is set to &#x60;null&#x60;. Available values are &#x60;true&#x60; or &#x60;false&#x60;. Defaults to &#x60;false&#x60;. | 
 **expand** | **string** | If set to &#x60;subtasks&#x60;, Vikunja will fetch only tasks which do not have subtasks and then in a second step, will fetch all of these subtasks. This may result in more tasks than the pagination limit being returned, but all subtasks will be present in the response. You can only set this to &#x60;subtasks&#x60;. | 

### Return type

[**[]ModelsTask**](ModelsTask.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksBulkPost

> ModelsTask TasksBulkPost(ctx).Task(task).Execute()

Update a bunch of tasks at once



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
	task := *openapiclient.NewModelsBulkTask() // ModelsBulkTask | The task object. Looks like a normal task, the only difference is it uses an array of project_ids to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.TasksBulkPost(context.Background()).Task(task).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksBulkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksBulkPost`: ModelsTask
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.TasksBulkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTasksBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **task** | [**ModelsBulkTask**](ModelsBulkTask.md) | The task object. Looks like a normal task, the only difference is it uses an array of project_ids to update. | 

### Return type

[**ModelsTask**](ModelsTask.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksIdAttachmentsAttachmentIDDelete

> ModelsMessage TasksIdAttachmentsAttachmentIDDelete(ctx, id, attachmentID).Execute()

Delete an attachment



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
	id := int32(56) // int32 | Task ID
	attachmentID := int32(56) // int32 | Attachment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.TasksIdAttachmentsAttachmentIDDelete(context.Background(), id, attachmentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksIdAttachmentsAttachmentIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksIdAttachmentsAttachmentIDDelete`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.TasksIdAttachmentsAttachmentIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Task ID | 
**attachmentID** | **int32** | Attachment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksIdAttachmentsAttachmentIDDeleteRequest struct via the builder pattern


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


## TasksIdAttachmentsAttachmentIDGet

> *os.File TasksIdAttachmentsAttachmentIDGet(ctx, id, attachmentID).PreviewSize(previewSize).Execute()

Get one attachment.



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
	id := int32(56) // int32 | Task ID
	attachmentID := int32(56) // int32 | Attachment ID
	previewSize := "previewSize_example" // string | The size of the preview image. Can be sm = 100px, md = 200px, lg = 400px or xl = 800px. If provided, a preview image will be returned if the attachment is an image. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.TasksIdAttachmentsAttachmentIDGet(context.Background(), id, attachmentID).PreviewSize(previewSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksIdAttachmentsAttachmentIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksIdAttachmentsAttachmentIDGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.TasksIdAttachmentsAttachmentIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Task ID | 
**attachmentID** | **int32** | Attachment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksIdAttachmentsAttachmentIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **previewSize** | **string** | The size of the preview image. Can be sm &#x3D; 100px, md &#x3D; 200px, lg &#x3D; 400px or xl &#x3D; 800px. If provided, a preview image will be returned if the attachment is an image. | 

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


## TasksIdAttachmentsGet

> []ModelsTaskAttachment TasksIdAttachmentsGet(ctx, id).Page(page).PerPage(perPage).Execute()

Get  all attachments for one task.



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
	id := int32(56) // int32 | Task ID
	page := int32(56) // int32 | The page number. Used for pagination. If not provided, the first page of results is returned. (optional)
	perPage := int32(56) // int32 | The maximum number of items per page. Note this parameter is limited by the configured maximum of items per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.TasksIdAttachmentsGet(context.Background(), id).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksIdAttachmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksIdAttachmentsGet`: []ModelsTaskAttachment
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.TasksIdAttachmentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksIdAttachmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page number. Used for pagination. If not provided, the first page of results is returned. | 
 **perPage** | **int32** | The maximum number of items per page. Note this parameter is limited by the configured maximum of items per page. | 

### Return type

[**[]ModelsTaskAttachment**](ModelsTaskAttachment.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksIdAttachmentsPut

> ModelsMessage TasksIdAttachmentsPut(ctx, id).Files(files).Execute()

Upload a task attachment



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
	id := int32(56) // int32 | Task ID
	files := "files_example" // string | The file, as multipart form file. You can pass multiple.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.TasksIdAttachmentsPut(context.Background(), id).Files(files).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksIdAttachmentsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksIdAttachmentsPut`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.TasksIdAttachmentsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksIdAttachmentsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **files** | **string** | The file, as multipart form file. You can pass multiple. | 

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


## TasksIdDelete

> ModelsMessage TasksIdDelete(ctx, id).Execute()

Delete a task



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
	id := int32(56) // int32 | Task ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.TasksIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksIdDelete`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.TasksIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksIdDeleteRequest struct via the builder pattern


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


## TasksIdGet

> ModelsTask TasksIdGet(ctx, id).Execute()

Get one task



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
	id := int32(56) // int32 | The task ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.TasksIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksIdGet`: ModelsTask
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.TasksIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsTask**](ModelsTask.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksIdPositionPost

> ModelsTaskPosition TasksIdPositionPost(ctx, id).View(view).Execute()

Updates a task position



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
	id := int32(56) // int32 | Task ID
	view := *openapiclient.NewModelsTaskPosition() // ModelsTaskPosition | The task position with updated values you want to change.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.TasksIdPositionPost(context.Background(), id).View(view).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksIdPositionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksIdPositionPost`: ModelsTaskPosition
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.TasksIdPositionPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksIdPositionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | [**ModelsTaskPosition**](ModelsTaskPosition.md) | The task position with updated values you want to change. | 

### Return type

[**ModelsTaskPosition**](ModelsTaskPosition.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksIdPost

> ModelsTask TasksIdPost(ctx, id).Task(task).Execute()

Update a task



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
	id := int32(56) // int32 | The Task ID
	task := *openapiclient.NewModelsTask() // ModelsTask | The task object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.TasksIdPost(context.Background(), id).Task(task).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksIdPost`: ModelsTask
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.TasksIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **task** | [**ModelsTask**](ModelsTask.md) | The task object | 

### Return type

[**ModelsTask**](ModelsTask.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskIDCommentsCommentIDDelete

> ModelsMessage TasksTaskIDCommentsCommentIDDelete(ctx, taskID, commentID).Execute()

Remove a task comment



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
	commentID := int32(56) // int32 | Comment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.TasksTaskIDCommentsCommentIDDelete(context.Background(), taskID, commentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksTaskIDCommentsCommentIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIDCommentsCommentIDDelete`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.TasksTaskIDCommentsCommentIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskID** | **int32** | Task ID | 
**commentID** | **int32** | Comment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIDCommentsCommentIDDeleteRequest struct via the builder pattern


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


## TasksTaskIDCommentsCommentIDGet

> ModelsTaskComment TasksTaskIDCommentsCommentIDGet(ctx, taskID, commentID).Execute()

Remove a task comment



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
	commentID := int32(56) // int32 | Comment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.TasksTaskIDCommentsCommentIDGet(context.Background(), taskID, commentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksTaskIDCommentsCommentIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIDCommentsCommentIDGet`: ModelsTaskComment
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.TasksTaskIDCommentsCommentIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskID** | **int32** | Task ID | 
**commentID** | **int32** | Comment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIDCommentsCommentIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelsTaskComment**](ModelsTaskComment.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskIDCommentsCommentIDPost

> ModelsTaskComment TasksTaskIDCommentsCommentIDPost(ctx, taskID, commentID).Execute()

Update an existing task comment



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
	commentID := int32(56) // int32 | Comment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.TasksTaskIDCommentsCommentIDPost(context.Background(), taskID, commentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksTaskIDCommentsCommentIDPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIDCommentsCommentIDPost`: ModelsTaskComment
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.TasksTaskIDCommentsCommentIDPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskID** | **int32** | Task ID | 
**commentID** | **int32** | Comment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIDCommentsCommentIDPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelsTaskComment**](ModelsTaskComment.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskIDCommentsGet

> []ModelsTaskComment TasksTaskIDCommentsGet(ctx, taskID).Execute()

Get all task comments



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.TasksTaskIDCommentsGet(context.Background(), taskID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksTaskIDCommentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIDCommentsGet`: []ModelsTaskComment
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.TasksTaskIDCommentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskID** | **int32** | Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIDCommentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ModelsTaskComment**](ModelsTaskComment.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskIDCommentsPut

> ModelsTaskComment TasksTaskIDCommentsPut(ctx, taskID).Relation(relation).Execute()

Create a new task comment



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
	relation := *openapiclient.NewModelsTaskComment() // ModelsTaskComment | The task comment object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.TasksTaskIDCommentsPut(context.Background(), taskID).Relation(relation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksTaskIDCommentsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIDCommentsPut`: ModelsTaskComment
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.TasksTaskIDCommentsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskID** | **int32** | Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIDCommentsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **relation** | [**ModelsTaskComment**](ModelsTaskComment.md) | The task comment object | 

### Return type

[**ModelsTaskComment**](ModelsTaskComment.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskIDRelationsPut

> ModelsTaskRelation TasksTaskIDRelationsPut(ctx, taskID).Relation(relation).Execute()

Create a new relation between two tasks



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
	relation := *openapiclient.NewModelsTaskRelation() // ModelsTaskRelation | The relation object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.TasksTaskIDRelationsPut(context.Background(), taskID).Relation(relation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksTaskIDRelationsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIDRelationsPut`: ModelsTaskRelation
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.TasksTaskIDRelationsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskID** | **int32** | Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIDRelationsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **relation** | [**ModelsTaskRelation**](ModelsTaskRelation.md) | The relation object | 

### Return type

[**ModelsTaskRelation**](ModelsTaskRelation.md)

### Authorization

[JWTKeyAuth](../README.md#JWTKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskIDRelationsRelationKindOtherTaskIDDelete

> ModelsMessage TasksTaskIDRelationsRelationKindOtherTaskIDDelete(ctx, taskID, relationKind, otherTaskID).Relation(relation).Execute()

Remove a task relation

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
	relationKind := "relationKind_example" // string | The kind of the relation. See the TaskRelation type for more info.
	otherTaskID := int32(56) // int32 | The id of the other task.
	relation := *openapiclient.NewModelsTaskRelation() // ModelsTaskRelation | The relation object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.TasksTaskIDRelationsRelationKindOtherTaskIDDelete(context.Background(), taskID, relationKind, otherTaskID).Relation(relation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksTaskIDRelationsRelationKindOtherTaskIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIDRelationsRelationKindOtherTaskIDDelete`: ModelsMessage
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.TasksTaskIDRelationsRelationKindOtherTaskIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskID** | **int32** | Task ID | 
**relationKind** | **string** | The kind of the relation. See the TaskRelation type for more info. | 
**otherTaskID** | **int32** | The id of the other task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIDRelationsRelationKindOtherTaskIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **relation** | [**ModelsTaskRelation**](ModelsTaskRelation.md) | The relation object | 

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

