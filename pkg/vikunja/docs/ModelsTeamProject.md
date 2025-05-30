# ModelsTeamProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | A timestamp when this relation was created. You cannot change this value. | [optional] 
**Id** | Pointer to **int32** | The unique, numeric id of this project &lt;-&gt; team relation. | [optional] 
**Right** | Pointer to [**ModelsRight**](ModelsRight.md) | The right this team has. 0 &#x3D; Read only, 1 &#x3D; Read &amp; Write, 2 &#x3D; Admin. See the docs for more details. | [optional] 
**TeamId** | Pointer to **int32** | The team id. | [optional] 
**Updated** | Pointer to **string** | A timestamp when this relation was last updated. You cannot change this value. | [optional] 

## Methods

### NewModelsTeamProject

`func NewModelsTeamProject() *ModelsTeamProject`

NewModelsTeamProject instantiates a new ModelsTeamProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTeamProjectWithDefaults

`func NewModelsTeamProjectWithDefaults() *ModelsTeamProject`

NewModelsTeamProjectWithDefaults instantiates a new ModelsTeamProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ModelsTeamProject) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsTeamProject) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsTeamProject) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsTeamProject) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *ModelsTeamProject) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsTeamProject) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsTeamProject) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsTeamProject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRight

`func (o *ModelsTeamProject) GetRight() ModelsRight`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *ModelsTeamProject) GetRightOk() (*ModelsRight, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *ModelsTeamProject) SetRight(v ModelsRight)`

SetRight sets Right field to given value.

### HasRight

`func (o *ModelsTeamProject) HasRight() bool`

HasRight returns a boolean if a field has been set.

### GetTeamId

`func (o *ModelsTeamProject) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *ModelsTeamProject) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *ModelsTeamProject) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *ModelsTeamProject) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetUpdated

`func (o *ModelsTeamProject) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ModelsTeamProject) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ModelsTeamProject) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ModelsTeamProject) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


