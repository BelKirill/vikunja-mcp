# ModelsTaskCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **string** | The filter query to match tasks by. Check out https://vikunja.io/docs/filters for a full explanation. | [optional] 
**FilterIncludeNulls** | Pointer to **bool** | If set to true, the result will also include null values | [optional] 
**OrderBy** | Pointer to **[]string** | The query parameter to order the items by. This can be either asc or desc, with asc being the default. | [optional] 
**SortBy** | Pointer to **[]string** | The query parameter to sort by. This is for ex. done, priority, etc. | [optional] 

## Methods

### NewModelsTaskCollection

`func NewModelsTaskCollection() *ModelsTaskCollection`

NewModelsTaskCollection instantiates a new ModelsTaskCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTaskCollectionWithDefaults

`func NewModelsTaskCollectionWithDefaults() *ModelsTaskCollection`

NewModelsTaskCollectionWithDefaults instantiates a new ModelsTaskCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ModelsTaskCollection) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ModelsTaskCollection) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ModelsTaskCollection) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ModelsTaskCollection) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFilterIncludeNulls

`func (o *ModelsTaskCollection) GetFilterIncludeNulls() bool`

GetFilterIncludeNulls returns the FilterIncludeNulls field if non-nil, zero value otherwise.

### GetFilterIncludeNullsOk

`func (o *ModelsTaskCollection) GetFilterIncludeNullsOk() (*bool, bool)`

GetFilterIncludeNullsOk returns a tuple with the FilterIncludeNulls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterIncludeNulls

`func (o *ModelsTaskCollection) SetFilterIncludeNulls(v bool)`

SetFilterIncludeNulls sets FilterIncludeNulls field to given value.

### HasFilterIncludeNulls

`func (o *ModelsTaskCollection) HasFilterIncludeNulls() bool`

HasFilterIncludeNulls returns a boolean if a field has been set.

### GetOrderBy

`func (o *ModelsTaskCollection) GetOrderBy() []string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ModelsTaskCollection) GetOrderByOk() (*[]string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ModelsTaskCollection) SetOrderBy(v []string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *ModelsTaskCollection) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetSortBy

`func (o *ModelsTaskCollection) GetSortBy() []string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *ModelsTaskCollection) GetSortByOk() (*[]string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *ModelsTaskCollection) SetSortBy(v []string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *ModelsTaskCollection) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


