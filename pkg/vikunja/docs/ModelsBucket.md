# ModelsBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The number of tasks currently in this bucket | [optional] 
**Created** | Pointer to **string** | A timestamp when this bucket was created. You cannot change this value. | [optional] 
**CreatedBy** | Pointer to [**UserUser**](UserUser.md) | The user who initially created the bucket. | [optional] 
**Id** | Pointer to **int32** | The unique, numeric id of this bucket. | [optional] 
**Limit** | Pointer to **int32** | How many tasks can be at the same time on this board max | [optional] 
**Position** | Pointer to **float32** | The position this bucket has when querying all buckets. See the tasks.position property on how to use this. | [optional] 
**ProjectViewId** | Pointer to **int32** | The project view this bucket belongs to. | [optional] 
**Tasks** | Pointer to [**[]ModelsTask**](ModelsTask.md) | All tasks which belong to this bucket. | [optional] 
**Title** | Pointer to **string** | The title of this bucket. | [optional] 
**Updated** | Pointer to **string** | A timestamp when this bucket was last updated. You cannot change this value. | [optional] 

## Methods

### NewModelsBucket

`func NewModelsBucket() *ModelsBucket`

NewModelsBucket instantiates a new ModelsBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsBucketWithDefaults

`func NewModelsBucketWithDefaults() *ModelsBucket`

NewModelsBucketWithDefaults instantiates a new ModelsBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ModelsBucket) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ModelsBucket) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ModelsBucket) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ModelsBucket) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreated

`func (o *ModelsBucket) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsBucket) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsBucket) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsBucket) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ModelsBucket) GetCreatedBy() UserUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ModelsBucket) GetCreatedByOk() (*UserUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ModelsBucket) SetCreatedBy(v UserUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ModelsBucket) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetId

`func (o *ModelsBucket) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsBucket) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsBucket) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsBucket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLimit

`func (o *ModelsBucket) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ModelsBucket) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ModelsBucket) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ModelsBucket) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetPosition

`func (o *ModelsBucket) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ModelsBucket) GetPositionOk() (*float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ModelsBucket) SetPosition(v float32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ModelsBucket) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetProjectViewId

`func (o *ModelsBucket) GetProjectViewId() int32`

GetProjectViewId returns the ProjectViewId field if non-nil, zero value otherwise.

### GetProjectViewIdOk

`func (o *ModelsBucket) GetProjectViewIdOk() (*int32, bool)`

GetProjectViewIdOk returns a tuple with the ProjectViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectViewId

`func (o *ModelsBucket) SetProjectViewId(v int32)`

SetProjectViewId sets ProjectViewId field to given value.

### HasProjectViewId

`func (o *ModelsBucket) HasProjectViewId() bool`

HasProjectViewId returns a boolean if a field has been set.

### GetTasks

`func (o *ModelsBucket) GetTasks() []ModelsTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ModelsBucket) GetTasksOk() (*[]ModelsTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ModelsBucket) SetTasks(v []ModelsTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ModelsBucket) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetTitle

`func (o *ModelsBucket) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModelsBucket) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModelsBucket) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ModelsBucket) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdated

`func (o *ModelsBucket) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ModelsBucket) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ModelsBucket) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ModelsBucket) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


