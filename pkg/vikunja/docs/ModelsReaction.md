# ModelsReaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | A timestamp when this reaction was created. You cannot change this value. | [optional] 
**User** | Pointer to [**UserUser**](UserUser.md) | The user who reacted | [optional] 
**Value** | Pointer to **string** | The actual reaction. This can be any valid utf character or text, up to a length of 20. | [optional] 

## Methods

### NewModelsReaction

`func NewModelsReaction() *ModelsReaction`

NewModelsReaction instantiates a new ModelsReaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsReactionWithDefaults

`func NewModelsReactionWithDefaults() *ModelsReaction`

NewModelsReactionWithDefaults instantiates a new ModelsReaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ModelsReaction) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsReaction) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsReaction) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsReaction) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUser

`func (o *ModelsReaction) GetUser() UserUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ModelsReaction) GetUserOk() (*UserUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ModelsReaction) SetUser(v UserUser)`

SetUser sets User field to given value.

### HasUser

`func (o *ModelsReaction) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetValue

`func (o *ModelsReaction) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ModelsReaction) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ModelsReaction) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ModelsReaction) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


