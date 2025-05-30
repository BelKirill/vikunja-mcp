# UserToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewUserToken

`func NewUserToken() *UserToken`

NewUserToken instantiates a new UserToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTokenWithDefaults

`func NewUserTokenWithDefaults() *UserToken`

NewUserTokenWithDefaults instantiates a new UserToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *UserToken) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UserToken) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UserToken) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *UserToken) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *UserToken) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserToken) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserToken) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetToken

`func (o *UserToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UserToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UserToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UserToken) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


