# ModelsTeamMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | Pointer to **bool** | Whether or not the member is an admin of the team. See the docs for more about what a team admin can do | [optional] 
**Created** | Pointer to **string** | A timestamp when this relation was created. You cannot change this value. | [optional] 
**Id** | Pointer to **int32** | The unique, numeric id of this team member relation. | [optional] 
**Username** | Pointer to **string** | The username of the member. We use this to prevent automated user id entering. | [optional] 

## Methods

### NewModelsTeamMember

`func NewModelsTeamMember() *ModelsTeamMember`

NewModelsTeamMember instantiates a new ModelsTeamMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTeamMemberWithDefaults

`func NewModelsTeamMemberWithDefaults() *ModelsTeamMember`

NewModelsTeamMemberWithDefaults instantiates a new ModelsTeamMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *ModelsTeamMember) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *ModelsTeamMember) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *ModelsTeamMember) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *ModelsTeamMember) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetCreated

`func (o *ModelsTeamMember) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsTeamMember) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsTeamMember) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsTeamMember) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *ModelsTeamMember) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsTeamMember) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsTeamMember) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsTeamMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *ModelsTeamMember) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ModelsTeamMember) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ModelsTeamMember) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ModelsTeamMember) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


