# ModelsTaskComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**UserUser**](UserUser.md) |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Reactions** | Pointer to [**map[string][]UserUser**](array.md) |  | [optional] 
**Updated** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsTaskComment

`func NewModelsTaskComment() *ModelsTaskComment`

NewModelsTaskComment instantiates a new ModelsTaskComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTaskCommentWithDefaults

`func NewModelsTaskCommentWithDefaults() *ModelsTaskComment`

NewModelsTaskCommentWithDefaults instantiates a new ModelsTaskComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *ModelsTaskComment) GetAuthor() UserUser`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ModelsTaskComment) GetAuthorOk() (*UserUser, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ModelsTaskComment) SetAuthor(v UserUser)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ModelsTaskComment) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetComment

`func (o *ModelsTaskComment) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ModelsTaskComment) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ModelsTaskComment) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ModelsTaskComment) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreated

`func (o *ModelsTaskComment) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsTaskComment) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsTaskComment) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsTaskComment) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *ModelsTaskComment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsTaskComment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsTaskComment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsTaskComment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReactions

`func (o *ModelsTaskComment) GetReactions() map[string][]UserUser`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *ModelsTaskComment) GetReactionsOk() (*map[string][]UserUser, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *ModelsTaskComment) SetReactions(v map[string][]UserUser)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *ModelsTaskComment) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetUpdated

`func (o *ModelsTaskComment) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ModelsTaskComment) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ModelsTaskComment) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ModelsTaskComment) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


