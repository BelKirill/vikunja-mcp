# ModelsProjectUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | A timestamp when this relation was created. You cannot change this value. | [optional] 
**Id** | Pointer to **int32** | The unique, numeric id of this project &lt;-&gt; user relation. | [optional] 
**Right** | Pointer to [**ModelsRight**](ModelsRight.md) | The right this user has. 0 &#x3D; Read only, 1 &#x3D; Read &amp; Write, 2 &#x3D; Admin. See the docs for more details. | [optional] 
**Updated** | Pointer to **string** | A timestamp when this relation was last updated. You cannot change this value. | [optional] 
**UserId** | Pointer to **string** | The username. | [optional] 

## Methods

### NewModelsProjectUser

`func NewModelsProjectUser() *ModelsProjectUser`

NewModelsProjectUser instantiates a new ModelsProjectUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsProjectUserWithDefaults

`func NewModelsProjectUserWithDefaults() *ModelsProjectUser`

NewModelsProjectUserWithDefaults instantiates a new ModelsProjectUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ModelsProjectUser) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsProjectUser) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsProjectUser) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsProjectUser) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *ModelsProjectUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsProjectUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsProjectUser) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsProjectUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRight

`func (o *ModelsProjectUser) GetRight() ModelsRight`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *ModelsProjectUser) GetRightOk() (*ModelsRight, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *ModelsProjectUser) SetRight(v ModelsRight)`

SetRight sets Right field to given value.

### HasRight

`func (o *ModelsProjectUser) HasRight() bool`

HasRight returns a boolean if a field has been set.

### GetUpdated

`func (o *ModelsProjectUser) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ModelsProjectUser) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ModelsProjectUser) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ModelsProjectUser) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUserId

`func (o *ModelsProjectUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelsProjectUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelsProjectUser) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelsProjectUser) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


