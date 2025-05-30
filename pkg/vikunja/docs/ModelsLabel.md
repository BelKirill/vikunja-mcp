# ModelsLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | A timestamp when this label was created. You cannot change this value. | [optional] 
**CreatedBy** | Pointer to [**UserUser**](UserUser.md) | The user who created this label | [optional] 
**Description** | Pointer to **string** | The label description. | [optional] 
**HexColor** | Pointer to **string** | The color this label has in hex format. | [optional] 
**Id** | Pointer to **int32** | The unique, numeric id of this label. | [optional] 
**Title** | Pointer to **string** | The title of the lable. You&#39;ll see this one on tasks associated with it. | [optional] 
**Updated** | Pointer to **string** | A timestamp when this label was last updated. You cannot change this value. | [optional] 

## Methods

### NewModelsLabel

`func NewModelsLabel() *ModelsLabel`

NewModelsLabel instantiates a new ModelsLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsLabelWithDefaults

`func NewModelsLabelWithDefaults() *ModelsLabel`

NewModelsLabelWithDefaults instantiates a new ModelsLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ModelsLabel) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsLabel) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsLabel) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsLabel) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ModelsLabel) GetCreatedBy() UserUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ModelsLabel) GetCreatedByOk() (*UserUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ModelsLabel) SetCreatedBy(v UserUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ModelsLabel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *ModelsLabel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsLabel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsLabel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelsLabel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHexColor

`func (o *ModelsLabel) GetHexColor() string`

GetHexColor returns the HexColor field if non-nil, zero value otherwise.

### GetHexColorOk

`func (o *ModelsLabel) GetHexColorOk() (*string, bool)`

GetHexColorOk returns a tuple with the HexColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexColor

`func (o *ModelsLabel) SetHexColor(v string)`

SetHexColor sets HexColor field to given value.

### HasHexColor

`func (o *ModelsLabel) HasHexColor() bool`

HasHexColor returns a boolean if a field has been set.

### GetId

`func (o *ModelsLabel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsLabel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsLabel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsLabel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ModelsLabel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModelsLabel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModelsLabel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ModelsLabel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdated

`func (o *ModelsLabel) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ModelsLabel) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ModelsLabel) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ModelsLabel) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


