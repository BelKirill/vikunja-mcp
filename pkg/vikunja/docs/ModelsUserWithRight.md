# ModelsUserWithRight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | A timestamp when this task was created. You cannot change this value. | [optional] 
**Email** | Pointer to **string** | The user&#39;s email address. | [optional] 
**Id** | Pointer to **int32** | The unique, numeric id of this user. | [optional] 
**Name** | Pointer to **string** | The full name of the user. | [optional] 
**Right** | Pointer to [**ModelsRight**](ModelsRight.md) |  | [optional] 
**Updated** | Pointer to **string** | A timestamp when this task was last updated. You cannot change this value. | [optional] 
**Username** | Pointer to **string** | The username of the user. Is always unique. | [optional] 

## Methods

### NewModelsUserWithRight

`func NewModelsUserWithRight() *ModelsUserWithRight`

NewModelsUserWithRight instantiates a new ModelsUserWithRight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsUserWithRightWithDefaults

`func NewModelsUserWithRightWithDefaults() *ModelsUserWithRight`

NewModelsUserWithRightWithDefaults instantiates a new ModelsUserWithRight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ModelsUserWithRight) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsUserWithRight) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsUserWithRight) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsUserWithRight) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEmail

`func (o *ModelsUserWithRight) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ModelsUserWithRight) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ModelsUserWithRight) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ModelsUserWithRight) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *ModelsUserWithRight) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsUserWithRight) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsUserWithRight) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsUserWithRight) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ModelsUserWithRight) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsUserWithRight) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsUserWithRight) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsUserWithRight) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRight

`func (o *ModelsUserWithRight) GetRight() ModelsRight`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *ModelsUserWithRight) GetRightOk() (*ModelsRight, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *ModelsUserWithRight) SetRight(v ModelsRight)`

SetRight sets Right field to given value.

### HasRight

`func (o *ModelsUserWithRight) HasRight() bool`

HasRight returns a boolean if a field has been set.

### GetUpdated

`func (o *ModelsUserWithRight) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ModelsUserWithRight) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ModelsUserWithRight) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ModelsUserWithRight) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUsername

`func (o *ModelsUserWithRight) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ModelsUserWithRight) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ModelsUserWithRight) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ModelsUserWithRight) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


