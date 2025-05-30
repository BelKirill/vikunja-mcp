# MigrationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinishedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**MigratorName** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewMigrationStatus

`func NewMigrationStatus() *MigrationStatus`

NewMigrationStatus instantiates a new MigrationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationStatusWithDefaults

`func NewMigrationStatusWithDefaults() *MigrationStatus`

NewMigrationStatusWithDefaults instantiates a new MigrationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinishedAt

`func (o *MigrationStatus) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *MigrationStatus) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *MigrationStatus) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *MigrationStatus) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetId

`func (o *MigrationStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MigrationStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MigrationStatus) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MigrationStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMigratorName

`func (o *MigrationStatus) GetMigratorName() string`

GetMigratorName returns the MigratorName field if non-nil, zero value otherwise.

### GetMigratorNameOk

`func (o *MigrationStatus) GetMigratorNameOk() (*string, bool)`

GetMigratorNameOk returns a tuple with the MigratorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratorName

`func (o *MigrationStatus) SetMigratorName(v string)`

SetMigratorName sets MigratorName field to given value.

### HasMigratorName

`func (o *MigrationStatus) HasMigratorName() bool`

HasMigratorName returns a boolean if a field has been set.

### GetStartedAt

`func (o *MigrationStatus) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *MigrationStatus) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *MigrationStatus) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *MigrationStatus) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


