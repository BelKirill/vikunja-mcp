# ModelsTeamUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | Pointer to **bool** | Whether the member is an admin of the team. See the docs for more about what a team admin can do | [optional] 
**Created** | Pointer to **string** | A timestamp when this task was created. You cannot change this value. | [optional] 
**Email** | Pointer to **string** | The user&#39;s email address. | [optional] 
**Id** | Pointer to **int32** | The unique, numeric id of this user. | [optional] 
**Name** | Pointer to **string** | The full name of the user. | [optional] 
**Updated** | Pointer to **string** | A timestamp when this task was last updated. You cannot change this value. | [optional] 
**Username** | Pointer to **string** | The username of the user. Is always unique. | [optional] 

## Methods

### NewModelsTeamUser

`func NewModelsTeamUser() *ModelsTeamUser`

NewModelsTeamUser instantiates a new ModelsTeamUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTeamUserWithDefaults

`func NewModelsTeamUserWithDefaults() *ModelsTeamUser`

NewModelsTeamUserWithDefaults instantiates a new ModelsTeamUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *ModelsTeamUser) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *ModelsTeamUser) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *ModelsTeamUser) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *ModelsTeamUser) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetCreated

`func (o *ModelsTeamUser) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsTeamUser) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsTeamUser) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsTeamUser) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEmail

`func (o *ModelsTeamUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ModelsTeamUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ModelsTeamUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ModelsTeamUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *ModelsTeamUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsTeamUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsTeamUser) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsTeamUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ModelsTeamUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsTeamUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsTeamUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsTeamUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdated

`func (o *ModelsTeamUser) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ModelsTeamUser) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ModelsTeamUser) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ModelsTeamUser) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUsername

`func (o *ModelsTeamUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ModelsTeamUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ModelsTeamUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ModelsTeamUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


