# ModelsTaskRelation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | A timestamp when this label was created. You cannot change this value. | [optional] 
**CreatedBy** | Pointer to [**UserUser**](UserUser.md) | The user who created this relation | [optional] 
**OtherTaskId** | Pointer to **int32** | The ID of the other task, the task which is being related. | [optional] 
**RelationKind** | Pointer to [**ModelsRelationKind**](ModelsRelationKind.md) | The kind of the relation. | [optional] 
**TaskId** | Pointer to **int32** | The ID of the \&quot;base\&quot; task, the task which has a relation to another. | [optional] 

## Methods

### NewModelsTaskRelation

`func NewModelsTaskRelation() *ModelsTaskRelation`

NewModelsTaskRelation instantiates a new ModelsTaskRelation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTaskRelationWithDefaults

`func NewModelsTaskRelationWithDefaults() *ModelsTaskRelation`

NewModelsTaskRelationWithDefaults instantiates a new ModelsTaskRelation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ModelsTaskRelation) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsTaskRelation) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsTaskRelation) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsTaskRelation) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ModelsTaskRelation) GetCreatedBy() UserUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ModelsTaskRelation) GetCreatedByOk() (*UserUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ModelsTaskRelation) SetCreatedBy(v UserUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ModelsTaskRelation) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetOtherTaskId

`func (o *ModelsTaskRelation) GetOtherTaskId() int32`

GetOtherTaskId returns the OtherTaskId field if non-nil, zero value otherwise.

### GetOtherTaskIdOk

`func (o *ModelsTaskRelation) GetOtherTaskIdOk() (*int32, bool)`

GetOtherTaskIdOk returns a tuple with the OtherTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherTaskId

`func (o *ModelsTaskRelation) SetOtherTaskId(v int32)`

SetOtherTaskId sets OtherTaskId field to given value.

### HasOtherTaskId

`func (o *ModelsTaskRelation) HasOtherTaskId() bool`

HasOtherTaskId returns a boolean if a field has been set.

### GetRelationKind

`func (o *ModelsTaskRelation) GetRelationKind() ModelsRelationKind`

GetRelationKind returns the RelationKind field if non-nil, zero value otherwise.

### GetRelationKindOk

`func (o *ModelsTaskRelation) GetRelationKindOk() (*ModelsRelationKind, bool)`

GetRelationKindOk returns a tuple with the RelationKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationKind

`func (o *ModelsTaskRelation) SetRelationKind(v ModelsRelationKind)`

SetRelationKind sets RelationKind field to given value.

### HasRelationKind

`func (o *ModelsTaskRelation) HasRelationKind() bool`

HasRelationKind returns a boolean if a field has been set.

### GetTaskId

`func (o *ModelsTaskRelation) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ModelsTaskRelation) GetTaskIdOk() (*int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ModelsTaskRelation) SetTaskId(v int32)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *ModelsTaskRelation) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


