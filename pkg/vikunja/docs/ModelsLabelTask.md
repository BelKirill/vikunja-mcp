# ModelsLabelTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | A timestamp when this task was created. You cannot change this value. | [optional] 
**LabelId** | Pointer to **int32** | The label id you want to associate with a task. | [optional] 

## Methods

### NewModelsLabelTask

`func NewModelsLabelTask() *ModelsLabelTask`

NewModelsLabelTask instantiates a new ModelsLabelTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsLabelTaskWithDefaults

`func NewModelsLabelTaskWithDefaults() *ModelsLabelTask`

NewModelsLabelTaskWithDefaults instantiates a new ModelsLabelTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ModelsLabelTask) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsLabelTask) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsLabelTask) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsLabelTask) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLabelId

`func (o *ModelsLabelTask) GetLabelId() int32`

GetLabelId returns the LabelId field if non-nil, zero value otherwise.

### GetLabelIdOk

`func (o *ModelsLabelTask) GetLabelIdOk() (*int32, bool)`

GetLabelIdOk returns a tuple with the LabelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelId

`func (o *ModelsLabelTask) SetLabelId(v int32)`

SetLabelId sets LabelId field to given value.

### HasLabelId

`func (o *ModelsLabelTask) HasLabelId() bool`

HasLabelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


