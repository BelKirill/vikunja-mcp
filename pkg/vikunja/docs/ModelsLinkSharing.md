# ModelsLinkSharing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | A timestamp when this project was shared. You cannot change this value. | [optional] 
**Hash** | Pointer to **string** | The public id to get this shared project | [optional] 
**Id** | Pointer to **int32** | The ID of the shared thing | [optional] 
**Name** | Pointer to **string** | The name of this link share. All actions someone takes while being authenticated with that link will appear with that name. | [optional] 
**Password** | Pointer to **string** | The password of this link share. You can only set it, not retrieve it after the link share has been created. | [optional] 
**Right** | Pointer to [**ModelsRight**](ModelsRight.md) | The right this project is shared with. 0 &#x3D; Read only, 1 &#x3D; Read &amp; Write, 2 &#x3D; Admin. See the docs for more details. | [optional] 
**SharedBy** | Pointer to [**UserUser**](UserUser.md) | The user who shared this project | [optional] 
**SharingType** | Pointer to [**ModelsSharingType**](ModelsSharingType.md) | The kind of this link. 0 &#x3D; undefined, 1 &#x3D; without password, 2 &#x3D; with password. | [optional] 
**Updated** | Pointer to **string** | A timestamp when this share was last updated. You cannot change this value. | [optional] 

## Methods

### NewModelsLinkSharing

`func NewModelsLinkSharing() *ModelsLinkSharing`

NewModelsLinkSharing instantiates a new ModelsLinkSharing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsLinkSharingWithDefaults

`func NewModelsLinkSharingWithDefaults() *ModelsLinkSharing`

NewModelsLinkSharingWithDefaults instantiates a new ModelsLinkSharing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ModelsLinkSharing) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsLinkSharing) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsLinkSharing) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsLinkSharing) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetHash

`func (o *ModelsLinkSharing) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ModelsLinkSharing) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ModelsLinkSharing) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ModelsLinkSharing) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetId

`func (o *ModelsLinkSharing) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsLinkSharing) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsLinkSharing) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsLinkSharing) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ModelsLinkSharing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsLinkSharing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsLinkSharing) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsLinkSharing) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *ModelsLinkSharing) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModelsLinkSharing) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModelsLinkSharing) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ModelsLinkSharing) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRight

`func (o *ModelsLinkSharing) GetRight() ModelsRight`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *ModelsLinkSharing) GetRightOk() (*ModelsRight, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *ModelsLinkSharing) SetRight(v ModelsRight)`

SetRight sets Right field to given value.

### HasRight

`func (o *ModelsLinkSharing) HasRight() bool`

HasRight returns a boolean if a field has been set.

### GetSharedBy

`func (o *ModelsLinkSharing) GetSharedBy() UserUser`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *ModelsLinkSharing) GetSharedByOk() (*UserUser, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBy

`func (o *ModelsLinkSharing) SetSharedBy(v UserUser)`

SetSharedBy sets SharedBy field to given value.

### HasSharedBy

`func (o *ModelsLinkSharing) HasSharedBy() bool`

HasSharedBy returns a boolean if a field has been set.

### GetSharingType

`func (o *ModelsLinkSharing) GetSharingType() ModelsSharingType`

GetSharingType returns the SharingType field if non-nil, zero value otherwise.

### GetSharingTypeOk

`func (o *ModelsLinkSharing) GetSharingTypeOk() (*ModelsSharingType, bool)`

GetSharingTypeOk returns a tuple with the SharingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingType

`func (o *ModelsLinkSharing) SetSharingType(v ModelsSharingType)`

SetSharingType sets SharingType field to given value.

### HasSharingType

`func (o *ModelsLinkSharing) HasSharingType() bool`

HasSharingType returns a boolean if a field has been set.

### GetUpdated

`func (o *ModelsLinkSharing) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ModelsLinkSharing) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ModelsLinkSharing) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ModelsLinkSharing) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


