# UserTOTP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | The totp entry will only be enabled after the user verified they have a working totp setup. | [optional] 
**Secret** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | The totp url used to be able to enroll the user later | [optional] 

## Methods

### NewUserTOTP

`func NewUserTOTP() *UserTOTP`

NewUserTOTP instantiates a new UserTOTP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTOTPWithDefaults

`func NewUserTOTPWithDefaults() *UserTOTP`

NewUserTOTPWithDefaults instantiates a new UserTOTP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UserTOTP) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserTOTP) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserTOTP) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UserTOTP) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSecret

`func (o *UserTOTP) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *UserTOTP) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *UserTOTP) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *UserTOTP) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetUrl

`func (o *UserTOTP) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UserTOTP) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UserTOTP) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UserTOTP) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


