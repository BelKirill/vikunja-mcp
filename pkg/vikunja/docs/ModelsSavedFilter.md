# ModelsSavedFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | A timestamp when this filter was created. You cannot change this value. | [optional] 
**Description** | Pointer to **string** | The description of the filter | [optional] 
**Filters** | Pointer to [**ModelsTaskCollection**](ModelsTaskCollection.md) | The actual filters this filter contains | [optional] 
**Id** | Pointer to **int32** | The unique numeric id of this saved filter | [optional] 
**IsFavorite** | Pointer to **bool** | True if the filter is a favorite. Favorite filters show up in a separate parent project together with favorite projects. | [optional] 
**Owner** | Pointer to [**UserUser**](UserUser.md) | The user who owns this filter | [optional] 
**Title** | Pointer to **string** | The title of the filter. | [optional] 
**Updated** | Pointer to **string** | A timestamp when this filter was last updated. You cannot change this value. | [optional] 

## Methods

### NewModelsSavedFilter

`func NewModelsSavedFilter() *ModelsSavedFilter`

NewModelsSavedFilter instantiates a new ModelsSavedFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSavedFilterWithDefaults

`func NewModelsSavedFilterWithDefaults() *ModelsSavedFilter`

NewModelsSavedFilterWithDefaults instantiates a new ModelsSavedFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ModelsSavedFilter) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsSavedFilter) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsSavedFilter) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsSavedFilter) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *ModelsSavedFilter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsSavedFilter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsSavedFilter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelsSavedFilter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilters

`func (o *ModelsSavedFilter) GetFilters() ModelsTaskCollection`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ModelsSavedFilter) GetFiltersOk() (*ModelsTaskCollection, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ModelsSavedFilter) SetFilters(v ModelsTaskCollection)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ModelsSavedFilter) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetId

`func (o *ModelsSavedFilter) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsSavedFilter) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsSavedFilter) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsSavedFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsFavorite

`func (o *ModelsSavedFilter) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *ModelsSavedFilter) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *ModelsSavedFilter) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *ModelsSavedFilter) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetOwner

`func (o *ModelsSavedFilter) GetOwner() UserUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ModelsSavedFilter) GetOwnerOk() (*UserUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ModelsSavedFilter) SetOwner(v UserUser)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ModelsSavedFilter) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTitle

`func (o *ModelsSavedFilter) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModelsSavedFilter) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModelsSavedFilter) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ModelsSavedFilter) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdated

`func (o *ModelsSavedFilter) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ModelsSavedFilter) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ModelsSavedFilter) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ModelsSavedFilter) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


