# UserAPIUserPassword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The user&#39;s email address | [optional] 
**Id** | Pointer to **int32** | The unique, numeric id of this user. | [optional] 
**Password** | Pointer to **string** | The user&#39;s password in clear text. Only used when registering the user. The maximum limi is 72 bytes, which may be less than 72 characters. This is due to the limit in the bcrypt hashing algorithm used to store passwords in Vikunja. | [optional] 
**Username** | Pointer to **string** | The user&#39;s username. Cannot contain anything that looks like an url or whitespaces. | [optional] 

## Methods

### NewUserAPIUserPassword

`func NewUserAPIUserPassword() *UserAPIUserPassword`

NewUserAPIUserPassword instantiates a new UserAPIUserPassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAPIUserPasswordWithDefaults

`func NewUserAPIUserPasswordWithDefaults() *UserAPIUserPassword`

NewUserAPIUserPasswordWithDefaults instantiates a new UserAPIUserPassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserAPIUserPassword) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserAPIUserPassword) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserAPIUserPassword) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserAPIUserPassword) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *UserAPIUserPassword) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserAPIUserPassword) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserAPIUserPassword) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserAPIUserPassword) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPassword

`func (o *UserAPIUserPassword) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserAPIUserPassword) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserAPIUserPassword) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserAPIUserPassword) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *UserAPIUserPassword) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserAPIUserPassword) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserAPIUserPassword) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserAPIUserPassword) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


