# UserUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | A timestamp when this task was created. You cannot change this value. | [optional] 
**Email** | Pointer to **string** | The user&#39;s email address. | [optional] 
**Id** | Pointer to **int32** | The unique, numeric id of this user. | [optional] 
**Name** | Pointer to **string** | The full name of the user. | [optional] 
**Updated** | Pointer to **string** | A timestamp when this task was last updated. You cannot change this value. | [optional] 
**Username** | Pointer to **string** | The username of the user. Is always unique. | [optional] 

## Methods

### NewUserUser

`func NewUserUser() *UserUser`

NewUserUser instantiates a new UserUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUserWithDefaults

`func NewUserUserWithDefaults() *UserUser`

NewUserUserWithDefaults instantiates a new UserUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *UserUser) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UserUser) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UserUser) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *UserUser) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEmail

`func (o *UserUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *UserUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserUser) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UserUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdated

`func (o *UserUser) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *UserUser) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *UserUser) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *UserUser) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUsername

`func (o *UserUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


