# ModelsTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignees** | Pointer to [**[]UserUser**](UserUser.md) | An array of users who are assigned to this task | [optional] 
**Attachments** | Pointer to [**[]ModelsTaskAttachment**](ModelsTaskAttachment.md) | All attachments this task has. This property is read-onlym, you must use the separate endpoint to add attachments to a task. | [optional] 
**BucketId** | Pointer to **int32** | The bucket id. Will only be populated when the task is accessed via a view with buckets. Can be used to move a task between buckets. In that case, the new bucket must be in the same view as the old one. | [optional] 
**CoverImageAttachmentId** | Pointer to **int32** | If this task has a cover image, the field will return the id of the attachment that is the cover image. | [optional] 
**Created** | Pointer to **string** | A timestamp when this task was created. You cannot change this value. | [optional] 
**CreatedBy** | Pointer to [**UserUser**](UserUser.md) | The user who initially created the task. | [optional] 
**Description** | Pointer to **string** | The task description. | [optional] 
**Done** | Pointer to **bool** | Whether a task is done or not. | [optional] 
**DoneAt** | Pointer to **string** | The time when a task was marked as done. | [optional] 
**DueDate** | Pointer to **string** | The time when the task is due. | [optional] 
**EndDate** | Pointer to **string** | When this task ends. | [optional] 
**HexColor** | Pointer to **string** | The task color in hex | [optional] 
**Id** | Pointer to **int32** | The unique, numeric id of this task. | [optional] 
**Identifier** | Pointer to **string** | The task identifier, based on the project identifier and the task&#39;s index | [optional] 
**Index** | Pointer to **int32** | The task index, calculated per project | [optional] 
**IsFavorite** | Pointer to **bool** | True if a task is a favorite task. Favorite tasks show up in a separate \&quot;Important\&quot; project. This value depends on the user making the call to the api. | [optional] 
**Labels** | Pointer to [**[]ModelsLabel**](ModelsLabel.md) | An array of labels which are associated with this task. This property is read-only, you must use the separate endpoint to add labels to a task. | [optional] 
**PercentDone** | Pointer to **float32** | Determines how far a task is left from being done | [optional] 
**Position** | Pointer to **float32** | The position of the task - any task project can be sorted as usual by this parameter. When accessing tasks via views with buckets, this is primarily used to sort them based on a range. Positions are always saved per view. They will automatically be set if you request the tasks through a view endpoint, otherwise they will always be 0. To update them, take a look at the Task Position endpoint. | [optional] 
**Priority** | Pointer to **int32** | The task priority. Can be anything you want, it is possible to sort by this later. | [optional] 
**ProjectId** | Pointer to **int32** | The project this task belongs to. | [optional] 
**Reactions** | Pointer to [**map[string][]UserUser**](array.md) | Reactions on that task. | [optional] 
**RelatedTasks** | Pointer to [**map[string][]ModelsTask**](array.md) | All related tasks, grouped by their relation kind | [optional] 
**Reminders** | Pointer to [**[]ModelsTaskReminder**](ModelsTaskReminder.md) | An array of reminders that are associated with this task. | [optional] 
**RepeatAfter** | Pointer to **int32** | An amount in seconds this task repeats itself. If this is set, when marking the task as done, it will mark itself as \&quot;undone\&quot; and then increase all remindes and the due date by its amount. | [optional] 
**RepeatMode** | Pointer to [**ModelsTaskRepeatMode**](ModelsTaskRepeatMode.md) | Can have three possible values which will trigger when the task is marked as done: 0 &#x3D; repeats after the amount specified in repeat_after, 1 &#x3D; repeats all dates each months (ignoring repeat_after), 3 &#x3D; repeats from the current date rather than the last set date. | [optional] 
**StartDate** | Pointer to **string** | When this task starts. | [optional] 
**Subscription** | Pointer to [**ModelsSubscription**](ModelsSubscription.md) | The subscription status for the user reading this task. You can only read this property, use the subscription endpoints to modify it. Will only returned when retrieving one task. | [optional] 
**Title** | Pointer to **string** | The task text. This is what you&#39;ll see in the project. | [optional] 
**Updated** | Pointer to **string** | A timestamp when this task was last updated. You cannot change this value. | [optional] 

## Methods

### NewModelsTask

`func NewModelsTask() *ModelsTask`

NewModelsTask instantiates a new ModelsTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTaskWithDefaults

`func NewModelsTaskWithDefaults() *ModelsTask`

NewModelsTaskWithDefaults instantiates a new ModelsTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignees

`func (o *ModelsTask) GetAssignees() []UserUser`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *ModelsTask) GetAssigneesOk() (*[]UserUser, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *ModelsTask) SetAssignees(v []UserUser)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *ModelsTask) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.

### GetAttachments

`func (o *ModelsTask) GetAttachments() []ModelsTaskAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ModelsTask) GetAttachmentsOk() (*[]ModelsTaskAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ModelsTask) SetAttachments(v []ModelsTaskAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ModelsTask) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetBucketId

`func (o *ModelsTask) GetBucketId() int32`

GetBucketId returns the BucketId field if non-nil, zero value otherwise.

### GetBucketIdOk

`func (o *ModelsTask) GetBucketIdOk() (*int32, bool)`

GetBucketIdOk returns a tuple with the BucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketId

`func (o *ModelsTask) SetBucketId(v int32)`

SetBucketId sets BucketId field to given value.

### HasBucketId

`func (o *ModelsTask) HasBucketId() bool`

HasBucketId returns a boolean if a field has been set.

### GetCoverImageAttachmentId

`func (o *ModelsTask) GetCoverImageAttachmentId() int32`

GetCoverImageAttachmentId returns the CoverImageAttachmentId field if non-nil, zero value otherwise.

### GetCoverImageAttachmentIdOk

`func (o *ModelsTask) GetCoverImageAttachmentIdOk() (*int32, bool)`

GetCoverImageAttachmentIdOk returns a tuple with the CoverImageAttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImageAttachmentId

`func (o *ModelsTask) SetCoverImageAttachmentId(v int32)`

SetCoverImageAttachmentId sets CoverImageAttachmentId field to given value.

### HasCoverImageAttachmentId

`func (o *ModelsTask) HasCoverImageAttachmentId() bool`

HasCoverImageAttachmentId returns a boolean if a field has been set.

### GetCreated

`func (o *ModelsTask) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsTask) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsTask) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsTask) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ModelsTask) GetCreatedBy() UserUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ModelsTask) GetCreatedByOk() (*UserUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ModelsTask) SetCreatedBy(v UserUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ModelsTask) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *ModelsTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelsTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDone

`func (o *ModelsTask) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *ModelsTask) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *ModelsTask) SetDone(v bool)`

SetDone sets Done field to given value.

### HasDone

`func (o *ModelsTask) HasDone() bool`

HasDone returns a boolean if a field has been set.

### GetDoneAt

`func (o *ModelsTask) GetDoneAt() string`

GetDoneAt returns the DoneAt field if non-nil, zero value otherwise.

### GetDoneAtOk

`func (o *ModelsTask) GetDoneAtOk() (*string, bool)`

GetDoneAtOk returns a tuple with the DoneAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoneAt

`func (o *ModelsTask) SetDoneAt(v string)`

SetDoneAt sets DoneAt field to given value.

### HasDoneAt

`func (o *ModelsTask) HasDoneAt() bool`

HasDoneAt returns a boolean if a field has been set.

### GetDueDate

`func (o *ModelsTask) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *ModelsTask) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *ModelsTask) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *ModelsTask) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ModelsTask) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ModelsTask) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ModelsTask) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ModelsTask) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetHexColor

`func (o *ModelsTask) GetHexColor() string`

GetHexColor returns the HexColor field if non-nil, zero value otherwise.

### GetHexColorOk

`func (o *ModelsTask) GetHexColorOk() (*string, bool)`

GetHexColorOk returns a tuple with the HexColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexColor

`func (o *ModelsTask) SetHexColor(v string)`

SetHexColor sets HexColor field to given value.

### HasHexColor

`func (o *ModelsTask) HasHexColor() bool`

HasHexColor returns a boolean if a field has been set.

### GetId

`func (o *ModelsTask) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsTask) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsTask) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *ModelsTask) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ModelsTask) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ModelsTask) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *ModelsTask) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetIndex

`func (o *ModelsTask) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ModelsTask) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ModelsTask) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ModelsTask) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetIsFavorite

`func (o *ModelsTask) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *ModelsTask) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *ModelsTask) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *ModelsTask) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetLabels

`func (o *ModelsTask) GetLabels() []ModelsLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ModelsTask) GetLabelsOk() (*[]ModelsLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ModelsTask) SetLabels(v []ModelsLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ModelsTask) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetPercentDone

`func (o *ModelsTask) GetPercentDone() float32`

GetPercentDone returns the PercentDone field if non-nil, zero value otherwise.

### GetPercentDoneOk

`func (o *ModelsTask) GetPercentDoneOk() (*float32, bool)`

GetPercentDoneOk returns a tuple with the PercentDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentDone

`func (o *ModelsTask) SetPercentDone(v float32)`

SetPercentDone sets PercentDone field to given value.

### HasPercentDone

`func (o *ModelsTask) HasPercentDone() bool`

HasPercentDone returns a boolean if a field has been set.

### GetPosition

`func (o *ModelsTask) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ModelsTask) GetPositionOk() (*float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ModelsTask) SetPosition(v float32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ModelsTask) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPriority

`func (o *ModelsTask) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ModelsTask) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ModelsTask) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ModelsTask) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProjectId

`func (o *ModelsTask) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModelsTask) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModelsTask) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ModelsTask) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReactions

`func (o *ModelsTask) GetReactions() map[string][]UserUser`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *ModelsTask) GetReactionsOk() (*map[string][]UserUser, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *ModelsTask) SetReactions(v map[string][]UserUser)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *ModelsTask) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetRelatedTasks

`func (o *ModelsTask) GetRelatedTasks() map[string][]ModelsTask`

GetRelatedTasks returns the RelatedTasks field if non-nil, zero value otherwise.

### GetRelatedTasksOk

`func (o *ModelsTask) GetRelatedTasksOk() (*map[string][]ModelsTask, bool)`

GetRelatedTasksOk returns a tuple with the RelatedTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedTasks

`func (o *ModelsTask) SetRelatedTasks(v map[string][]ModelsTask)`

SetRelatedTasks sets RelatedTasks field to given value.

### HasRelatedTasks

`func (o *ModelsTask) HasRelatedTasks() bool`

HasRelatedTasks returns a boolean if a field has been set.

### GetReminders

`func (o *ModelsTask) GetReminders() []ModelsTaskReminder`

GetReminders returns the Reminders field if non-nil, zero value otherwise.

### GetRemindersOk

`func (o *ModelsTask) GetRemindersOk() (*[]ModelsTaskReminder, bool)`

GetRemindersOk returns a tuple with the Reminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminders

`func (o *ModelsTask) SetReminders(v []ModelsTaskReminder)`

SetReminders sets Reminders field to given value.

### HasReminders

`func (o *ModelsTask) HasReminders() bool`

HasReminders returns a boolean if a field has been set.

### GetRepeatAfter

`func (o *ModelsTask) GetRepeatAfter() int32`

GetRepeatAfter returns the RepeatAfter field if non-nil, zero value otherwise.

### GetRepeatAfterOk

`func (o *ModelsTask) GetRepeatAfterOk() (*int32, bool)`

GetRepeatAfterOk returns a tuple with the RepeatAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatAfter

`func (o *ModelsTask) SetRepeatAfter(v int32)`

SetRepeatAfter sets RepeatAfter field to given value.

### HasRepeatAfter

`func (o *ModelsTask) HasRepeatAfter() bool`

HasRepeatAfter returns a boolean if a field has been set.

### GetRepeatMode

`func (o *ModelsTask) GetRepeatMode() ModelsTaskRepeatMode`

GetRepeatMode returns the RepeatMode field if non-nil, zero value otherwise.

### GetRepeatModeOk

`func (o *ModelsTask) GetRepeatModeOk() (*ModelsTaskRepeatMode, bool)`

GetRepeatModeOk returns a tuple with the RepeatMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatMode

`func (o *ModelsTask) SetRepeatMode(v ModelsTaskRepeatMode)`

SetRepeatMode sets RepeatMode field to given value.

### HasRepeatMode

`func (o *ModelsTask) HasRepeatMode() bool`

HasRepeatMode returns a boolean if a field has been set.

### GetStartDate

`func (o *ModelsTask) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ModelsTask) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ModelsTask) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ModelsTask) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetSubscription

`func (o *ModelsTask) GetSubscription() ModelsSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *ModelsTask) GetSubscriptionOk() (*ModelsSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *ModelsTask) SetSubscription(v ModelsSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *ModelsTask) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetTitle

`func (o *ModelsTask) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModelsTask) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModelsTask) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ModelsTask) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdated

`func (o *ModelsTask) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ModelsTask) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ModelsTask) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ModelsTask) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


