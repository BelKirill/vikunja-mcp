# \TestingAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestTablePatch**](TestingAPI.md#TestTablePatch) | **Patch** /test/{table} | Reset the db to a defined state



## TestTablePatch

> []UserUser TestTablePatch(ctx, table).Execute()

Reset the db to a defined state



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
	table := "table_example" // string | The table to reset

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestingAPI.TestTablePatch(context.Background(), table).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestingAPI.TestTablePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestTablePatch`: []UserUser
	fmt.Fprintf(os.Stdout, "Response from `TestingAPI.TestTablePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**table** | **string** | The table to reset | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestTablePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UserUser**](UserUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

