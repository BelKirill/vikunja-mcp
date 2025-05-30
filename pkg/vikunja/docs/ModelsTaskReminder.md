# ModelsTaskReminder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelativePeriod** | Pointer to **int32** | A period in seconds relative to another date argument. Negative values mean the reminder triggers before the date. Default: 0, tiggers when RelativeTo is due. | [optional] 
**RelativeTo** | Pointer to [**ModelsReminderRelation**](ModelsReminderRelation.md) | The name of the date field to which the relative period refers to. | [optional] 
**Reminder** | Pointer to **string** | The absolute time when the user wants to be reminded of the task. | [optional] 

## Methods

### NewModelsTaskReminder

`func NewModelsTaskReminder() *ModelsTaskReminder`

NewModelsTaskReminder instantiates a new ModelsTaskReminder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTaskReminderWithDefaults

`func NewModelsTaskReminderWithDefaults() *ModelsTaskReminder`

NewModelsTaskReminderWithDefaults instantiates a new ModelsTaskReminder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelativePeriod

`func (o *ModelsTaskReminder) GetRelativePeriod() int32`

GetRelativePeriod returns the RelativePeriod field if non-nil, zero value otherwise.

### GetRelativePeriodOk

`func (o *ModelsTaskReminder) GetRelativePeriodOk() (*int32, bool)`

GetRelativePeriodOk returns a tuple with the RelativePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePeriod

`func (o *ModelsTaskReminder) SetRelativePeriod(v int32)`

SetRelativePeriod sets RelativePeriod field to given value.

### HasRelativePeriod

`func (o *ModelsTaskReminder) HasRelativePeriod() bool`

HasRelativePeriod returns a boolean if a field has been set.

### GetRelativeTo

`func (o *ModelsTaskReminder) GetRelativeTo() ModelsReminderRelation`

GetRelativeTo returns the RelativeTo field if non-nil, zero value otherwise.

### GetRelativeToOk

`func (o *ModelsTaskReminder) GetRelativeToOk() (*ModelsReminderRelation, bool)`

GetRelativeToOk returns a tuple with the RelativeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeTo

`func (o *ModelsTaskReminder) SetRelativeTo(v ModelsReminderRelation)`

SetRelativeTo sets RelativeTo field to given value.

### HasRelativeTo

`func (o *ModelsTaskReminder) HasRelativeTo() bool`

HasRelativeTo returns a boolean if a field has been set.

### GetReminder

`func (o *ModelsTaskReminder) GetReminder() string`

GetReminder returns the Reminder field if non-nil, zero value otherwise.

### GetReminderOk

`func (o *ModelsTaskReminder) GetReminderOk() (*string, bool)`

GetReminderOk returns a tuple with the Reminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminder

`func (o *ModelsTaskReminder) SetReminder(v string)`

SetReminder sets Reminder field to given value.

### HasReminder

`func (o *ModelsTaskReminder) HasReminder() bool`

HasReminder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


