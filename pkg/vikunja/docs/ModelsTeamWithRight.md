# ModelsTeamWithRight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | A timestamp when this relation was created. You cannot change this value. | [optional] 
**CreatedBy** | Pointer to [**UserUser**](UserUser.md) | The user who created this team. | [optional] 
**Description** | Pointer to **string** | The team&#39;s description. | [optional] 
**Id** | Pointer to **int32** | The unique, numeric id of this team. | [optional] 
**IncludePublic** | Pointer to **bool** | Query parameter controlling whether to include public projects or not | [optional] 
**IsPublic** | Pointer to **bool** | Defines wether the team should be publicly discoverable when sharing a project | [optional] 
**Members** | Pointer to [**[]ModelsTeamUser**](ModelsTeamUser.md) | An array of all members in this team. | [optional] 
**Name** | Pointer to **string** | The name of this team. | [optional] 
**OidcId** | Pointer to **string** | The team&#39;s oidc id delivered by the oidc provider | [optional] 
**Right** | Pointer to [**ModelsRight**](ModelsRight.md) |  | [optional] 
**Updated** | Pointer to **string** | A timestamp when this relation was last updated. You cannot change this value. | [optional] 

## Methods

### NewModelsTeamWithRight

`func NewModelsTeamWithRight() *ModelsTeamWithRight`

NewModelsTeamWithRight instantiates a new ModelsTeamWithRight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTeamWithRightWithDefaults

`func NewModelsTeamWithRightWithDefaults() *ModelsTeamWithRight`

NewModelsTeamWithRightWithDefaults instantiates a new ModelsTeamWithRight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ModelsTeamWithRight) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsTeamWithRight) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsTeamWithRight) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsTeamWithRight) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ModelsTeamWithRight) GetCreatedBy() UserUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ModelsTeamWithRight) GetCreatedByOk() (*UserUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ModelsTeamWithRight) SetCreatedBy(v UserUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ModelsTeamWithRight) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *ModelsTeamWithRight) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsTeamWithRight) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsTeamWithRight) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelsTeamWithRight) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ModelsTeamWithRight) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsTeamWithRight) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsTeamWithRight) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsTeamWithRight) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncludePublic

`func (o *ModelsTeamWithRight) GetIncludePublic() bool`

GetIncludePublic returns the IncludePublic field if non-nil, zero value otherwise.

### GetIncludePublicOk

`func (o *ModelsTeamWithRight) GetIncludePublicOk() (*bool, bool)`

GetIncludePublicOk returns a tuple with the IncludePublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePublic

`func (o *ModelsTeamWithRight) SetIncludePublic(v bool)`

SetIncludePublic sets IncludePublic field to given value.

### HasIncludePublic

`func (o *ModelsTeamWithRight) HasIncludePublic() bool`

HasIncludePublic returns a boolean if a field has been set.

### GetIsPublic

`func (o *ModelsTeamWithRight) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *ModelsTeamWithRight) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *ModelsTeamWithRight) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *ModelsTeamWithRight) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetMembers

`func (o *ModelsTeamWithRight) GetMembers() []ModelsTeamUser`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ModelsTeamWithRight) GetMembersOk() (*[]ModelsTeamUser, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ModelsTeamWithRight) SetMembers(v []ModelsTeamUser)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ModelsTeamWithRight) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetName

`func (o *ModelsTeamWithRight) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsTeamWithRight) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsTeamWithRight) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsTeamWithRight) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOidcId

`func (o *ModelsTeamWithRight) GetOidcId() string`

GetOidcId returns the OidcId field if non-nil, zero value otherwise.

### GetOidcIdOk

`func (o *ModelsTeamWithRight) GetOidcIdOk() (*string, bool)`

GetOidcIdOk returns a tuple with the OidcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcId

`func (o *ModelsTeamWithRight) SetOidcId(v string)`

SetOidcId sets OidcId field to given value.

### HasOidcId

`func (o *ModelsTeamWithRight) HasOidcId() bool`

HasOidcId returns a boolean if a field has been set.

### GetRight

`func (o *ModelsTeamWithRight) GetRight() ModelsRight`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *ModelsTeamWithRight) GetRightOk() (*ModelsRight, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *ModelsTeamWithRight) SetRight(v ModelsRight)`

SetRight sets Right field to given value.

### HasRight

`func (o *ModelsTeamWithRight) HasRight() bool`

HasRight returns a boolean if a field has been set.

### GetUpdated

`func (o *ModelsTeamWithRight) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ModelsTeamWithRight) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ModelsTeamWithRight) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ModelsTeamWithRight) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


