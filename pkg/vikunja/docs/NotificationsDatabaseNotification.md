# NotificationsDatabaseNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | A timestamp when this notification was created. You cannot change this value. | [optional] 
**Id** | Pointer to **int32** | The unique, numeric id of this notification. | [optional] 
**Name** | Pointer to **string** | The name of the notification | [optional] 
**Notification** | Pointer to **map[string]interface{}** | The actual content of the notification. | [optional] 
**ReadAt** | Pointer to **string** | When this notification is marked as read, this will be updated with the current timestamp. | [optional] 

## Methods

### NewNotificationsDatabaseNotification

`func NewNotificationsDatabaseNotification() *NotificationsDatabaseNotification`

NewNotificationsDatabaseNotification instantiates a new NotificationsDatabaseNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsDatabaseNotificationWithDefaults

`func NewNotificationsDatabaseNotificationWithDefaults() *NotificationsDatabaseNotification`

NewNotificationsDatabaseNotificationWithDefaults instantiates a new NotificationsDatabaseNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *NotificationsDatabaseNotification) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NotificationsDatabaseNotification) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NotificationsDatabaseNotification) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *NotificationsDatabaseNotification) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *NotificationsDatabaseNotification) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationsDatabaseNotification) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationsDatabaseNotification) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationsDatabaseNotification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NotificationsDatabaseNotification) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationsDatabaseNotification) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationsDatabaseNotification) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationsDatabaseNotification) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotification

`func (o *NotificationsDatabaseNotification) GetNotification() map[string]interface{}`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *NotificationsDatabaseNotification) GetNotificationOk() (*map[string]interface{}, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *NotificationsDatabaseNotification) SetNotification(v map[string]interface{})`

SetNotification sets Notification field to given value.

### HasNotification

`func (o *NotificationsDatabaseNotification) HasNotification() bool`

HasNotification returns a boolean if a field has been set.

### GetReadAt

`func (o *NotificationsDatabaseNotification) GetReadAt() string`

GetReadAt returns the ReadAt field if non-nil, zero value otherwise.

### GetReadAtOk

`func (o *NotificationsDatabaseNotification) GetReadAtOk() (*string, bool)`

GetReadAtOk returns a tuple with the ReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAt

`func (o *NotificationsDatabaseNotification) SetReadAt(v string)`

SetReadAt sets ReadAt field to given value.

### HasReadAt

`func (o *NotificationsDatabaseNotification) HasReadAt() bool`

HasReadAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


