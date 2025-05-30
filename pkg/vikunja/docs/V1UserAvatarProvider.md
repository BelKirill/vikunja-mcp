# V1UserAvatarProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarProvider** | Pointer to **string** | The avatar provider. Valid types are &#x60;gravatar&#x60; (uses the user email), &#x60;upload&#x60;, &#x60;initials&#x60;, &#x60;marble&#x60; (generates a random avatar for each user), &#x60;default&#x60;. | [optional] 

## Methods

### NewV1UserAvatarProvider

`func NewV1UserAvatarProvider() *V1UserAvatarProvider`

NewV1UserAvatarProvider instantiates a new V1UserAvatarProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1UserAvatarProviderWithDefaults

`func NewV1UserAvatarProviderWithDefaults() *V1UserAvatarProvider`

NewV1UserAvatarProviderWithDefaults instantiates a new V1UserAvatarProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarProvider

`func (o *V1UserAvatarProvider) GetAvatarProvider() string`

GetAvatarProvider returns the AvatarProvider field if non-nil, zero value otherwise.

### GetAvatarProviderOk

`func (o *V1UserAvatarProvider) GetAvatarProviderOk() (*string, bool)`

GetAvatarProviderOk returns a tuple with the AvatarProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarProvider

`func (o *V1UserAvatarProvider) SetAvatarProvider(v string)`

SetAvatarProvider sets AvatarProvider field to given value.

### HasAvatarProvider

`func (o *V1UserAvatarProvider) HasAvatarProvider() bool`

HasAvatarProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


