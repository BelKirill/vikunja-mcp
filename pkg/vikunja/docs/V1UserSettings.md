# V1UserSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultProjectId** | Pointer to **int32** | If a task is created without a specified project this value should be used. Applies to tasks made directly in API and from clients. | [optional] 
**DiscoverableByEmail** | Pointer to **bool** | If true, the user can be found when searching for their exact email. | [optional] 
**DiscoverableByName** | Pointer to **bool** | If true, this user can be found by their name or parts of it when searching for it. | [optional] 
**EmailRemindersEnabled** | Pointer to **bool** | If enabled, sends email reminders of tasks to the user. | [optional] 
**FrontendSettings** | Pointer to **map[string]interface{}** | Additional settings only used by the frontend | [optional] 
**Language** | Pointer to **string** | The user&#39;s language | [optional] 
**Name** | Pointer to **string** | The new name of the current user. | [optional] 
**OverdueTasksRemindersEnabled** | Pointer to **bool** | If enabled, the user will get an email for their overdue tasks each morning. | [optional] 
**OverdueTasksRemindersTime** | Pointer to **string** | The time when the daily summary of overdue tasks will be sent via email. | [optional] 
**Timezone** | Pointer to **string** | The user&#39;s time zone. Used to send task reminders in the time zone of the user. | [optional] 
**WeekStart** | Pointer to **int32** | The day when the week starts for this user. 0 &#x3D; sunday, 1 &#x3D; monday, etc. | [optional] 

## Methods

### NewV1UserSettings

`func NewV1UserSettings() *V1UserSettings`

NewV1UserSettings instantiates a new V1UserSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1UserSettingsWithDefaults

`func NewV1UserSettingsWithDefaults() *V1UserSettings`

NewV1UserSettingsWithDefaults instantiates a new V1UserSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultProjectId

`func (o *V1UserSettings) GetDefaultProjectId() int32`

GetDefaultProjectId returns the DefaultProjectId field if non-nil, zero value otherwise.

### GetDefaultProjectIdOk

`func (o *V1UserSettings) GetDefaultProjectIdOk() (*int32, bool)`

GetDefaultProjectIdOk returns a tuple with the DefaultProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProjectId

`func (o *V1UserSettings) SetDefaultProjectId(v int32)`

SetDefaultProjectId sets DefaultProjectId field to given value.

### HasDefaultProjectId

`func (o *V1UserSettings) HasDefaultProjectId() bool`

HasDefaultProjectId returns a boolean if a field has been set.

### GetDiscoverableByEmail

`func (o *V1UserSettings) GetDiscoverableByEmail() bool`

GetDiscoverableByEmail returns the DiscoverableByEmail field if non-nil, zero value otherwise.

### GetDiscoverableByEmailOk

`func (o *V1UserSettings) GetDiscoverableByEmailOk() (*bool, bool)`

GetDiscoverableByEmailOk returns a tuple with the DiscoverableByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverableByEmail

`func (o *V1UserSettings) SetDiscoverableByEmail(v bool)`

SetDiscoverableByEmail sets DiscoverableByEmail field to given value.

### HasDiscoverableByEmail

`func (o *V1UserSettings) HasDiscoverableByEmail() bool`

HasDiscoverableByEmail returns a boolean if a field has been set.

### GetDiscoverableByName

`func (o *V1UserSettings) GetDiscoverableByName() bool`

GetDiscoverableByName returns the DiscoverableByName field if non-nil, zero value otherwise.

### GetDiscoverableByNameOk

`func (o *V1UserSettings) GetDiscoverableByNameOk() (*bool, bool)`

GetDiscoverableByNameOk returns a tuple with the DiscoverableByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverableByName

`func (o *V1UserSettings) SetDiscoverableByName(v bool)`

SetDiscoverableByName sets DiscoverableByName field to given value.

### HasDiscoverableByName

`func (o *V1UserSettings) HasDiscoverableByName() bool`

HasDiscoverableByName returns a boolean if a field has been set.

### GetEmailRemindersEnabled

`func (o *V1UserSettings) GetEmailRemindersEnabled() bool`

GetEmailRemindersEnabled returns the EmailRemindersEnabled field if non-nil, zero value otherwise.

### GetEmailRemindersEnabledOk

`func (o *V1UserSettings) GetEmailRemindersEnabledOk() (*bool, bool)`

GetEmailRemindersEnabledOk returns a tuple with the EmailRemindersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailRemindersEnabled

`func (o *V1UserSettings) SetEmailRemindersEnabled(v bool)`

SetEmailRemindersEnabled sets EmailRemindersEnabled field to given value.

### HasEmailRemindersEnabled

`func (o *V1UserSettings) HasEmailRemindersEnabled() bool`

HasEmailRemindersEnabled returns a boolean if a field has been set.

### GetFrontendSettings

`func (o *V1UserSettings) GetFrontendSettings() map[string]interface{}`

GetFrontendSettings returns the FrontendSettings field if non-nil, zero value otherwise.

### GetFrontendSettingsOk

`func (o *V1UserSettings) GetFrontendSettingsOk() (*map[string]interface{}, bool)`

GetFrontendSettingsOk returns a tuple with the FrontendSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontendSettings

`func (o *V1UserSettings) SetFrontendSettings(v map[string]interface{})`

SetFrontendSettings sets FrontendSettings field to given value.

### HasFrontendSettings

`func (o *V1UserSettings) HasFrontendSettings() bool`

HasFrontendSettings returns a boolean if a field has been set.

### GetLanguage

`func (o *V1UserSettings) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *V1UserSettings) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *V1UserSettings) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *V1UserSettings) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetName

`func (o *V1UserSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1UserSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1UserSettings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1UserSettings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverdueTasksRemindersEnabled

`func (o *V1UserSettings) GetOverdueTasksRemindersEnabled() bool`

GetOverdueTasksRemindersEnabled returns the OverdueTasksRemindersEnabled field if non-nil, zero value otherwise.

### GetOverdueTasksRemindersEnabledOk

`func (o *V1UserSettings) GetOverdueTasksRemindersEnabledOk() (*bool, bool)`

GetOverdueTasksRemindersEnabledOk returns a tuple with the OverdueTasksRemindersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdueTasksRemindersEnabled

`func (o *V1UserSettings) SetOverdueTasksRemindersEnabled(v bool)`

SetOverdueTasksRemindersEnabled sets OverdueTasksRemindersEnabled field to given value.

### HasOverdueTasksRemindersEnabled

`func (o *V1UserSettings) HasOverdueTasksRemindersEnabled() bool`

HasOverdueTasksRemindersEnabled returns a boolean if a field has been set.

### GetOverdueTasksRemindersTime

`func (o *V1UserSettings) GetOverdueTasksRemindersTime() string`

GetOverdueTasksRemindersTime returns the OverdueTasksRemindersTime field if non-nil, zero value otherwise.

### GetOverdueTasksRemindersTimeOk

`func (o *V1UserSettings) GetOverdueTasksRemindersTimeOk() (*string, bool)`

GetOverdueTasksRemindersTimeOk returns a tuple with the OverdueTasksRemindersTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdueTasksRemindersTime

`func (o *V1UserSettings) SetOverdueTasksRemindersTime(v string)`

SetOverdueTasksRemindersTime sets OverdueTasksRemindersTime field to given value.

### HasOverdueTasksRemindersTime

`func (o *V1UserSettings) HasOverdueTasksRemindersTime() bool`

HasOverdueTasksRemindersTime returns a boolean if a field has been set.

### GetTimezone

`func (o *V1UserSettings) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *V1UserSettings) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *V1UserSettings) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *V1UserSettings) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetWeekStart

`func (o *V1UserSettings) GetWeekStart() int32`

GetWeekStart returns the WeekStart field if non-nil, zero value otherwise.

### GetWeekStartOk

`func (o *V1UserSettings) GetWeekStartOk() (*int32, bool)`

GetWeekStartOk returns a tuple with the WeekStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekStart

`func (o *V1UserSettings) SetWeekStart(v int32)`

SetWeekStart sets WeekStart field to given value.

### HasWeekStart

`func (o *V1UserSettings) HasWeekStart() bool`

HasWeekStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


