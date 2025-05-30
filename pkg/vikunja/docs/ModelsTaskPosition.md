# ModelsTaskPosition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | Pointer to **float32** | The position of the task - any task project can be sorted as usual by this parameter. When accessing tasks via kanban buckets, this is primarily used to sort them based on a range We&#39;re using a float64 here to make it possible to put any task within any two other tasks (by changing the number). You would calculate the new position between two tasks with something like task3.position &#x3D; (task2.position - task1.position) / 2. A 64-Bit float leaves plenty of room to initially give tasks a position with 2^16 difference to the previous task which also leaves a lot of room for rearranging and sorting later. Positions are always saved per view. They will automatically be set if you request the tasks through a view endpoint, otherwise they will always be 0. To update them, take a look at the Task Position endpoint. | [optional] 
**ProjectViewId** | Pointer to **int32** | The project view this task is related to | [optional] 
**TaskId** | Pointer to **int32** | The ID of the task this position is for | [optional] 

## Methods

### NewModelsTaskPosition

`func NewModelsTaskPosition() *ModelsTaskPosition`

NewModelsTaskPosition instantiates a new ModelsTaskPosition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTaskPositionWithDefaults

`func NewModelsTaskPositionWithDefaults() *ModelsTaskPosition`

NewModelsTaskPositionWithDefaults instantiates a new ModelsTaskPosition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *ModelsTaskPosition) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ModelsTaskPosition) GetPositionOk() (*float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ModelsTaskPosition) SetPosition(v float32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ModelsTaskPosition) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetProjectViewId

`func (o *ModelsTaskPosition) GetProjectViewId() int32`

GetProjectViewId returns the ProjectViewId field if non-nil, zero value otherwise.

### GetProjectViewIdOk

`func (o *ModelsTaskPosition) GetProjectViewIdOk() (*int32, bool)`

GetProjectViewIdOk returns a tuple with the ProjectViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectViewId

`func (o *ModelsTaskPosition) SetProjectViewId(v int32)`

SetProjectViewId sets ProjectViewId field to given value.

### HasProjectViewId

`func (o *ModelsTaskPosition) HasProjectViewId() bool`

HasProjectViewId returns a boolean if a field has been set.

### GetTaskId

`func (o *ModelsTaskPosition) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ModelsTaskPosition) GetTaskIdOk() (*int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ModelsTaskPosition) SetTaskId(v int32)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *ModelsTaskPosition) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


