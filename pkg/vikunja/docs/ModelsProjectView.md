# ModelsProjectView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketConfiguration** | Pointer to [**[]ModelsProjectViewBucketConfiguration**](ModelsProjectViewBucketConfiguration.md) | When the bucket configuration mode is not &#x60;manual&#x60;, this field holds the options of that configuration. | [optional] 
**BucketConfigurationMode** | Pointer to [**ModelsBucketConfigurationModeKind**](ModelsBucketConfigurationModeKind.md) | The bucket configuration mode. Can be &#x60;none&#x60;, &#x60;manual&#x60; or &#x60;filter&#x60;. &#x60;manual&#x60; allows to move tasks between buckets as you normally would. &#x60;filter&#x60; creates buckets based on a filter for each bucket. | [optional] 
**Created** | Pointer to **string** | A timestamp when this reaction was created. You cannot change this value. | [optional] 
**DefaultBucketId** | Pointer to **int32** | The ID of the bucket where new tasks without a bucket are added to. By default, this is the leftmost bucket in a view. | [optional] 
**DoneBucketId** | Pointer to **int32** | If tasks are moved to the done bucket, they are marked as done. If they are marked as done individually, they are moved into the done bucket. | [optional] 
**Filter** | Pointer to **string** | The filter query to match tasks by. Check out https://vikunja.io/docs/filters for a full explanation. | [optional] 
**Id** | Pointer to **int32** | The unique numeric id of this view | [optional] 
**Position** | Pointer to **float32** | The position of this view in the list. The list of all views will be sorted by this parameter. | [optional] 
**ProjectId** | Pointer to **int32** | The project this view belongs to | [optional] 
**Title** | Pointer to **string** | The title of this view | [optional] 
**Updated** | Pointer to **string** | A timestamp when this view was updated. You cannot change this value. | [optional] 
**ViewKind** | Pointer to [**ModelsProjectViewKind**](ModelsProjectViewKind.md) | The kind of this view. Can be &#x60;list&#x60;, &#x60;gantt&#x60;, &#x60;table&#x60; or &#x60;kanban&#x60;. | [optional] 

## Methods

### NewModelsProjectView

`func NewModelsProjectView() *ModelsProjectView`

NewModelsProjectView instantiates a new ModelsProjectView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsProjectViewWithDefaults

`func NewModelsProjectViewWithDefaults() *ModelsProjectView`

NewModelsProjectViewWithDefaults instantiates a new ModelsProjectView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketConfiguration

`func (o *ModelsProjectView) GetBucketConfiguration() []ModelsProjectViewBucketConfiguration`

GetBucketConfiguration returns the BucketConfiguration field if non-nil, zero value otherwise.

### GetBucketConfigurationOk

`func (o *ModelsProjectView) GetBucketConfigurationOk() (*[]ModelsProjectViewBucketConfiguration, bool)`

GetBucketConfigurationOk returns a tuple with the BucketConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketConfiguration

`func (o *ModelsProjectView) SetBucketConfiguration(v []ModelsProjectViewBucketConfiguration)`

SetBucketConfiguration sets BucketConfiguration field to given value.

### HasBucketConfiguration

`func (o *ModelsProjectView) HasBucketConfiguration() bool`

HasBucketConfiguration returns a boolean if a field has been set.

### GetBucketConfigurationMode

`func (o *ModelsProjectView) GetBucketConfigurationMode() ModelsBucketConfigurationModeKind`

GetBucketConfigurationMode returns the BucketConfigurationMode field if non-nil, zero value otherwise.

### GetBucketConfigurationModeOk

`func (o *ModelsProjectView) GetBucketConfigurationModeOk() (*ModelsBucketConfigurationModeKind, bool)`

GetBucketConfigurationModeOk returns a tuple with the BucketConfigurationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketConfigurationMode

`func (o *ModelsProjectView) SetBucketConfigurationMode(v ModelsBucketConfigurationModeKind)`

SetBucketConfigurationMode sets BucketConfigurationMode field to given value.

### HasBucketConfigurationMode

`func (o *ModelsProjectView) HasBucketConfigurationMode() bool`

HasBucketConfigurationMode returns a boolean if a field has been set.

### GetCreated

`func (o *ModelsProjectView) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsProjectView) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsProjectView) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsProjectView) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDefaultBucketId

`func (o *ModelsProjectView) GetDefaultBucketId() int32`

GetDefaultBucketId returns the DefaultBucketId field if non-nil, zero value otherwise.

### GetDefaultBucketIdOk

`func (o *ModelsProjectView) GetDefaultBucketIdOk() (*int32, bool)`

GetDefaultBucketIdOk returns a tuple with the DefaultBucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBucketId

`func (o *ModelsProjectView) SetDefaultBucketId(v int32)`

SetDefaultBucketId sets DefaultBucketId field to given value.

### HasDefaultBucketId

`func (o *ModelsProjectView) HasDefaultBucketId() bool`

HasDefaultBucketId returns a boolean if a field has been set.

### GetDoneBucketId

`func (o *ModelsProjectView) GetDoneBucketId() int32`

GetDoneBucketId returns the DoneBucketId field if non-nil, zero value otherwise.

### GetDoneBucketIdOk

`func (o *ModelsProjectView) GetDoneBucketIdOk() (*int32, bool)`

GetDoneBucketIdOk returns a tuple with the DoneBucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoneBucketId

`func (o *ModelsProjectView) SetDoneBucketId(v int32)`

SetDoneBucketId sets DoneBucketId field to given value.

### HasDoneBucketId

`func (o *ModelsProjectView) HasDoneBucketId() bool`

HasDoneBucketId returns a boolean if a field has been set.

### GetFilter

`func (o *ModelsProjectView) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ModelsProjectView) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ModelsProjectView) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ModelsProjectView) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetId

`func (o *ModelsProjectView) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsProjectView) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsProjectView) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsProjectView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPosition

`func (o *ModelsProjectView) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ModelsProjectView) GetPositionOk() (*float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ModelsProjectView) SetPosition(v float32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ModelsProjectView) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetProjectId

`func (o *ModelsProjectView) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModelsProjectView) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModelsProjectView) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ModelsProjectView) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetTitle

`func (o *ModelsProjectView) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModelsProjectView) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModelsProjectView) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ModelsProjectView) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdated

`func (o *ModelsProjectView) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ModelsProjectView) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ModelsProjectView) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ModelsProjectView) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetViewKind

`func (o *ModelsProjectView) GetViewKind() ModelsProjectViewKind`

GetViewKind returns the ViewKind field if non-nil, zero value otherwise.

### GetViewKindOk

`func (o *ModelsProjectView) GetViewKindOk() (*ModelsProjectViewKind, bool)`

GetViewKindOk returns a tuple with the ViewKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewKind

`func (o *ModelsProjectView) SetViewKind(v ModelsProjectViewKind)`

SetViewKind sets ViewKind field to given value.

### HasViewKind

`func (o *ModelsProjectView) HasViewKind() bool`

HasViewKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


