# ModelsWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | A timestamp when this webhook target was created. You cannot change this value. | [optional] 
**CreatedBy** | Pointer to [**UserUser**](UserUser.md) | The user who initially created the webhook target. | [optional] 
**Events** | Pointer to **[]string** | The webhook events which should fire this webhook target | [optional] 
**Id** | Pointer to **int32** | The generated ID of this webhook target | [optional] 
**ProjectId** | Pointer to **int32** | The project ID of the project this webhook target belongs to | [optional] 
**Secret** | Pointer to **string** | If provided, webhook requests will be signed using HMAC. Check out the docs about how to use this: https://vikunja.io/docs/webhooks/#signing | [optional] 
**TargetUrl** | Pointer to **string** | The target URL where the POST request with the webhook payload will be made | [optional] 
**Updated** | Pointer to **string** | A timestamp when this webhook target was last updated. You cannot change this value. | [optional] 

## Methods

### NewModelsWebhook

`func NewModelsWebhook() *ModelsWebhook`

NewModelsWebhook instantiates a new ModelsWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsWebhookWithDefaults

`func NewModelsWebhookWithDefaults() *ModelsWebhook`

NewModelsWebhookWithDefaults instantiates a new ModelsWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ModelsWebhook) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsWebhook) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsWebhook) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsWebhook) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ModelsWebhook) GetCreatedBy() UserUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ModelsWebhook) GetCreatedByOk() (*UserUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ModelsWebhook) SetCreatedBy(v UserUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ModelsWebhook) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetEvents

`func (o *ModelsWebhook) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ModelsWebhook) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ModelsWebhook) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ModelsWebhook) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetId

`func (o *ModelsWebhook) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsWebhook) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsWebhook) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsWebhook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectId

`func (o *ModelsWebhook) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModelsWebhook) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModelsWebhook) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ModelsWebhook) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetSecret

`func (o *ModelsWebhook) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ModelsWebhook) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ModelsWebhook) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ModelsWebhook) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetTargetUrl

`func (o *ModelsWebhook) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *ModelsWebhook) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *ModelsWebhook) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *ModelsWebhook) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### GetUpdated

`func (o *ModelsWebhook) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ModelsWebhook) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ModelsWebhook) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ModelsWebhook) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


