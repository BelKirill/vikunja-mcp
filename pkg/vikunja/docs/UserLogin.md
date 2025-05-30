# UserLogin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LongToken** | Pointer to **bool** | If true, the token returned will be valid a lot longer than default. Useful for \&quot;remember me\&quot; style logins. | [optional] 
**Password** | Pointer to **string** | The password for the user. | [optional] 
**TotpPasscode** | Pointer to **string** | The totp passcode of a user. Only needs to be provided when enabled. | [optional] 
**Username** | Pointer to **string** | The username used to log in. | [optional] 

## Methods

### NewUserLogin

`func NewUserLogin() *UserLogin`

NewUserLogin instantiates a new UserLogin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLoginWithDefaults

`func NewUserLoginWithDefaults() *UserLogin`

NewUserLoginWithDefaults instantiates a new UserLogin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLongToken

`func (o *UserLogin) GetLongToken() bool`

GetLongToken returns the LongToken field if non-nil, zero value otherwise.

### GetLongTokenOk

`func (o *UserLogin) GetLongTokenOk() (*bool, bool)`

GetLongTokenOk returns a tuple with the LongToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongToken

`func (o *UserLogin) SetLongToken(v bool)`

SetLongToken sets LongToken field to given value.

### HasLongToken

`func (o *UserLogin) HasLongToken() bool`

HasLongToken returns a boolean if a field has been set.

### GetPassword

`func (o *UserLogin) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserLogin) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserLogin) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserLogin) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetTotpPasscode

`func (o *UserLogin) GetTotpPasscode() string`

GetTotpPasscode returns the TotpPasscode field if non-nil, zero value otherwise.

### GetTotpPasscodeOk

`func (o *UserLogin) GetTotpPasscodeOk() (*string, bool)`

GetTotpPasscodeOk returns a tuple with the TotpPasscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpPasscode

`func (o *UserLogin) SetTotpPasscode(v string)`

SetTotpPasscode sets TotpPasscode field to given value.

### HasTotpPasscode

`func (o *UserLogin) HasTotpPasscode() bool`

HasTotpPasscode returns a boolean if a field has been set.

### GetUsername

`func (o *UserLogin) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserLogin) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserLogin) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserLogin) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


