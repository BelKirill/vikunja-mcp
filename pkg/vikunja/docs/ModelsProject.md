# ModelsProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackgroundBlurHash** | Pointer to **string** | Contains a very small version of the project background to use as a blurry preview until the actual background is loaded. Check out https://blurha.sh/ to learn how it works. | [optional] 
**BackgroundInformation** | Pointer to **map[string]interface{}** | Holds extra information about the background set since some background providers require attribution or similar. If not null, the background can be accessed at /projects/{projectID}/background | [optional] 
**Created** | Pointer to **string** | A timestamp when this project was created. You cannot change this value. | [optional] 
**Description** | Pointer to **string** | The description of the project. | [optional] 
**HexColor** | Pointer to **string** | The hex color of this project | [optional] 
**Id** | Pointer to **int32** | The unique, numeric id of this project. | [optional] 
**Identifier** | Pointer to **string** | The unique project short identifier. Used to build task identifiers. | [optional] 
**IsArchived** | Pointer to **bool** | Whether a project is archived. | [optional] 
**IsFavorite** | Pointer to **bool** | True if a project is a favorite. Favorite projects show up in a separate parent project. This value depends on the user making the call to the api. | [optional] 
**Owner** | Pointer to [**UserUser**](UserUser.md) | The user who created this project. | [optional] 
**ParentProjectId** | Pointer to **int32** |  | [optional] 
**Position** | Pointer to **float32** | The position this project has when querying all projects. See the tasks.position property on how to use this. | [optional] 
**Subscription** | Pointer to [**ModelsSubscription**](ModelsSubscription.md) | The subscription status for the user reading this project. You can only read this property, use the subscription endpoints to modify it. Will only returned when retreiving one project. | [optional] 
**Title** | Pointer to **string** | The title of the project. You&#39;ll see this in the overview. | [optional] 
**Updated** | Pointer to **string** | A timestamp when this project was last updated. You cannot change this value. | [optional] 
**Views** | Pointer to [**[]ModelsProjectView**](ModelsProjectView.md) |  | [optional] 

## Methods

### NewModelsProject

`func NewModelsProject() *ModelsProject`

NewModelsProject instantiates a new ModelsProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsProjectWithDefaults

`func NewModelsProjectWithDefaults() *ModelsProject`

NewModelsProjectWithDefaults instantiates a new ModelsProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackgroundBlurHash

`func (o *ModelsProject) GetBackgroundBlurHash() string`

GetBackgroundBlurHash returns the BackgroundBlurHash field if non-nil, zero value otherwise.

### GetBackgroundBlurHashOk

`func (o *ModelsProject) GetBackgroundBlurHashOk() (*string, bool)`

GetBackgroundBlurHashOk returns a tuple with the BackgroundBlurHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundBlurHash

`func (o *ModelsProject) SetBackgroundBlurHash(v string)`

SetBackgroundBlurHash sets BackgroundBlurHash field to given value.

### HasBackgroundBlurHash

`func (o *ModelsProject) HasBackgroundBlurHash() bool`

HasBackgroundBlurHash returns a boolean if a field has been set.

### GetBackgroundInformation

`func (o *ModelsProject) GetBackgroundInformation() map[string]interface{}`

GetBackgroundInformation returns the BackgroundInformation field if non-nil, zero value otherwise.

### GetBackgroundInformationOk

`func (o *ModelsProject) GetBackgroundInformationOk() (*map[string]interface{}, bool)`

GetBackgroundInformationOk returns a tuple with the BackgroundInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundInformation

`func (o *ModelsProject) SetBackgroundInformation(v map[string]interface{})`

SetBackgroundInformation sets BackgroundInformation field to given value.

### HasBackgroundInformation

`func (o *ModelsProject) HasBackgroundInformation() bool`

HasBackgroundInformation returns a boolean if a field has been set.

### GetCreated

`func (o *ModelsProject) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsProject) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsProject) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsProject) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *ModelsProject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsProject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsProject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelsProject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHexColor

`func (o *ModelsProject) GetHexColor() string`

GetHexColor returns the HexColor field if non-nil, zero value otherwise.

### GetHexColorOk

`func (o *ModelsProject) GetHexColorOk() (*string, bool)`

GetHexColorOk returns a tuple with the HexColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexColor

`func (o *ModelsProject) SetHexColor(v string)`

SetHexColor sets HexColor field to given value.

### HasHexColor

`func (o *ModelsProject) HasHexColor() bool`

HasHexColor returns a boolean if a field has been set.

### GetId

`func (o *ModelsProject) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsProject) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsProject) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsProject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *ModelsProject) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ModelsProject) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ModelsProject) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *ModelsProject) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetIsArchived

`func (o *ModelsProject) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *ModelsProject) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *ModelsProject) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *ModelsProject) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetIsFavorite

`func (o *ModelsProject) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *ModelsProject) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *ModelsProject) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *ModelsProject) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetOwner

`func (o *ModelsProject) GetOwner() UserUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ModelsProject) GetOwnerOk() (*UserUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ModelsProject) SetOwner(v UserUser)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ModelsProject) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetParentProjectId

`func (o *ModelsProject) GetParentProjectId() int32`

GetParentProjectId returns the ParentProjectId field if non-nil, zero value otherwise.

### GetParentProjectIdOk

`func (o *ModelsProject) GetParentProjectIdOk() (*int32, bool)`

GetParentProjectIdOk returns a tuple with the ParentProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentProjectId

`func (o *ModelsProject) SetParentProjectId(v int32)`

SetParentProjectId sets ParentProjectId field to given value.

### HasParentProjectId

`func (o *ModelsProject) HasParentProjectId() bool`

HasParentProjectId returns a boolean if a field has been set.

### GetPosition

`func (o *ModelsProject) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ModelsProject) GetPositionOk() (*float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ModelsProject) SetPosition(v float32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ModelsProject) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetSubscription

`func (o *ModelsProject) GetSubscription() ModelsSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *ModelsProject) GetSubscriptionOk() (*ModelsSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *ModelsProject) SetSubscription(v ModelsSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *ModelsProject) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetTitle

`func (o *ModelsProject) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModelsProject) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModelsProject) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ModelsProject) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdated

`func (o *ModelsProject) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ModelsProject) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ModelsProject) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ModelsProject) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetViews

`func (o *ModelsProject) GetViews() []ModelsProjectView`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *ModelsProject) GetViewsOk() (*[]ModelsProjectView, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *ModelsProject) SetViews(v []ModelsProjectView)`

SetViews sets Views field to given value.

### HasViews

`func (o *ModelsProject) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


