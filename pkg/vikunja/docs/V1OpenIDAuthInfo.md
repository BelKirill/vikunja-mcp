# V1OpenIDAuthInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Providers** | Pointer to [**[]OpenidProvider**](OpenidProvider.md) |  | [optional] 

## Methods

### NewV1OpenIDAuthInfo

`func NewV1OpenIDAuthInfo() *V1OpenIDAuthInfo`

NewV1OpenIDAuthInfo instantiates a new V1OpenIDAuthInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1OpenIDAuthInfoWithDefaults

`func NewV1OpenIDAuthInfoWithDefaults() *V1OpenIDAuthInfo`

NewV1OpenIDAuthInfoWithDefaults instantiates a new V1OpenIDAuthInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *V1OpenIDAuthInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *V1OpenIDAuthInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *V1OpenIDAuthInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *V1OpenIDAuthInfo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProviders

`func (o *V1OpenIDAuthInfo) GetProviders() []OpenidProvider`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *V1OpenIDAuthInfo) GetProvidersOk() (*[]OpenidProvider, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *V1OpenIDAuthInfo) SetProviders(v []OpenidProvider)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *V1OpenIDAuthInfo) HasProviders() bool`

HasProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


