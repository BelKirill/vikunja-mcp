# ModelsSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | A timestamp when this subscription was created. You cannot change this value. | [optional] 
**Entity** | Pointer to **int32** |  | [optional] 
**EntityId** | Pointer to **int32** | The id of the entity to subscribe to. | [optional] 
**Id** | Pointer to **int32** | The numeric ID of the subscription | [optional] 

## Methods

### NewModelsSubscription

`func NewModelsSubscription() *ModelsSubscription`

NewModelsSubscription instantiates a new ModelsSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSubscriptionWithDefaults

`func NewModelsSubscriptionWithDefaults() *ModelsSubscription`

NewModelsSubscriptionWithDefaults instantiates a new ModelsSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ModelsSubscription) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsSubscription) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsSubscription) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsSubscription) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEntity

`func (o *ModelsSubscription) GetEntity() int32`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *ModelsSubscription) GetEntityOk() (*int32, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *ModelsSubscription) SetEntity(v int32)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *ModelsSubscription) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetEntityId

`func (o *ModelsSubscription) GetEntityId() int32`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ModelsSubscription) GetEntityIdOk() (*int32, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ModelsSubscription) SetEntityId(v int32)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *ModelsSubscription) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetId

`func (o *ModelsSubscription) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsSubscription) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsSubscription) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsSubscription) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


