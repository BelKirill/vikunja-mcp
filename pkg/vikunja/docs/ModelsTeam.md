# ModelsTeam

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
**Updated** | Pointer to **string** | A timestamp when this relation was last updated. You cannot change this value. | [optional] 

## Methods

### NewModelsTeam

`func NewModelsTeam() *ModelsTeam`

NewModelsTeam instantiates a new ModelsTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTeamWithDefaults

`func NewModelsTeamWithDefaults() *ModelsTeam`

NewModelsTeamWithDefaults instantiates a new ModelsTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ModelsTeam) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsTeam) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsTeam) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsTeam) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ModelsTeam) GetCreatedBy() UserUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ModelsTeam) GetCreatedByOk() (*UserUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ModelsTeam) SetCreatedBy(v UserUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ModelsTeam) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *ModelsTeam) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsTeam) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsTeam) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelsTeam) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ModelsTeam) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsTeam) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsTeam) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsTeam) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncludePublic

`func (o *ModelsTeam) GetIncludePublic() bool`

GetIncludePublic returns the IncludePublic field if non-nil, zero value otherwise.

### GetIncludePublicOk

`func (o *ModelsTeam) GetIncludePublicOk() (*bool, bool)`

GetIncludePublicOk returns a tuple with the IncludePublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePublic

`func (o *ModelsTeam) SetIncludePublic(v bool)`

SetIncludePublic sets IncludePublic field to given value.

### HasIncludePublic

`func (o *ModelsTeam) HasIncludePublic() bool`

HasIncludePublic returns a boolean if a field has been set.

### GetIsPublic

`func (o *ModelsTeam) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *ModelsTeam) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *ModelsTeam) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *ModelsTeam) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetMembers

`func (o *ModelsTeam) GetMembers() []ModelsTeamUser`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ModelsTeam) GetMembersOk() (*[]ModelsTeamUser, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ModelsTeam) SetMembers(v []ModelsTeamUser)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ModelsTeam) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetName

`func (o *ModelsTeam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsTeam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsTeam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsTeam) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOidcId

`func (o *ModelsTeam) GetOidcId() string`

GetOidcId returns the OidcId field if non-nil, zero value otherwise.

### GetOidcIdOk

`func (o *ModelsTeam) GetOidcIdOk() (*string, bool)`

GetOidcIdOk returns a tuple with the OidcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcId

`func (o *ModelsTeam) SetOidcId(v string)`

SetOidcId sets OidcId field to given value.

### HasOidcId

`func (o *ModelsTeam) HasOidcId() bool`

HasOidcId returns a boolean if a field has been set.

### GetUpdated

`func (o *ModelsTeam) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ModelsTeam) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ModelsTeam) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ModelsTeam) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


