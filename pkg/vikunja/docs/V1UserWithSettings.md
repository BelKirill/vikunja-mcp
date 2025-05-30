# V1UserWithSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | A timestamp when this task was created. You cannot change this value. | [optional] 
**DeletionScheduledAt** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** | The user&#39;s email address. | [optional] 
**Id** | Pointer to **int32** | The unique, numeric id of this user. | [optional] 
**IsLocalUser** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** | The full name of the user. | [optional] 
**Settings** | Pointer to [**V1UserSettings**](V1UserSettings.md) |  | [optional] 
**Updated** | Pointer to **string** | A timestamp when this task was last updated. You cannot change this value. | [optional] 
**Username** | Pointer to **string** | The username of the user. Is always unique. | [optional] 

## Methods

### NewV1UserWithSettings

`func NewV1UserWithSettings() *V1UserWithSettings`

NewV1UserWithSettings instantiates a new V1UserWithSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1UserWithSettingsWithDefaults

`func NewV1UserWithSettingsWithDefaults() *V1UserWithSettings`

NewV1UserWithSettingsWithDefaults instantiates a new V1UserWithSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *V1UserWithSettings) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *V1UserWithSettings) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *V1UserWithSettings) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *V1UserWithSettings) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDeletionScheduledAt

`func (o *V1UserWithSettings) GetDeletionScheduledAt() string`

GetDeletionScheduledAt returns the DeletionScheduledAt field if non-nil, zero value otherwise.

### GetDeletionScheduledAtOk

`func (o *V1UserWithSettings) GetDeletionScheduledAtOk() (*string, bool)`

GetDeletionScheduledAtOk returns a tuple with the DeletionScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionScheduledAt

`func (o *V1UserWithSettings) SetDeletionScheduledAt(v string)`

SetDeletionScheduledAt sets DeletionScheduledAt field to given value.

### HasDeletionScheduledAt

`func (o *V1UserWithSettings) HasDeletionScheduledAt() bool`

HasDeletionScheduledAt returns a boolean if a field has been set.

### GetEmail

`func (o *V1UserWithSettings) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *V1UserWithSettings) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *V1UserWithSettings) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *V1UserWithSettings) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *V1UserWithSettings) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1UserWithSettings) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1UserWithSettings) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V1UserWithSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsLocalUser

`func (o *V1UserWithSettings) GetIsLocalUser() bool`

GetIsLocalUser returns the IsLocalUser field if non-nil, zero value otherwise.

### GetIsLocalUserOk

`func (o *V1UserWithSettings) GetIsLocalUserOk() (*bool, bool)`

GetIsLocalUserOk returns a tuple with the IsLocalUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocalUser

`func (o *V1UserWithSettings) SetIsLocalUser(v bool)`

SetIsLocalUser sets IsLocalUser field to given value.

### HasIsLocalUser

`func (o *V1UserWithSettings) HasIsLocalUser() bool`

HasIsLocalUser returns a boolean if a field has been set.

### GetName

`func (o *V1UserWithSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1UserWithSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1UserWithSettings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1UserWithSettings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *V1UserWithSettings) GetSettings() V1UserSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *V1UserWithSettings) GetSettingsOk() (*V1UserSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *V1UserWithSettings) SetSettings(v V1UserSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *V1UserWithSettings) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetUpdated

`func (o *V1UserWithSettings) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *V1UserWithSettings) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *V1UserWithSettings) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *V1UserWithSettings) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUsername

`func (o *V1UserWithSettings) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *V1UserWithSettings) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *V1UserWithSettings) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *V1UserWithSettings) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


