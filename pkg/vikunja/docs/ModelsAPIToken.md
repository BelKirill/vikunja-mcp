# ModelsAPIToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | A timestamp when this api key was created. You cannot change this value. | [optional] 
**ExpiresAt** | Pointer to **string** | The date when this key expires. | [optional] 
**Id** | Pointer to **int32** | The unique, numeric id of this api key. | [optional] 
**Permissions** | Pointer to **map[string][]string** | The permissions this token has. Possible values are available via the /routes endpoint and consist of the keys of the list from that endpoint. For example, if the token should be able to read all tasks as well as update existing tasks, you should add &#x60;{\&quot;tasks\&quot;:[\&quot;read_all\&quot;,\&quot;update\&quot;]}&#x60;. | [optional] 
**Title** | Pointer to **string** | A human-readable name for this token | [optional] 
**Token** | Pointer to **string** | The actual api key. Only visible after creation. | [optional] 

## Methods

### NewModelsAPIToken

`func NewModelsAPIToken() *ModelsAPIToken`

NewModelsAPIToken instantiates a new ModelsAPIToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsAPITokenWithDefaults

`func NewModelsAPITokenWithDefaults() *ModelsAPIToken`

NewModelsAPITokenWithDefaults instantiates a new ModelsAPIToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ModelsAPIToken) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsAPIToken) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsAPIToken) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsAPIToken) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ModelsAPIToken) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ModelsAPIToken) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ModelsAPIToken) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ModelsAPIToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *ModelsAPIToken) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsAPIToken) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsAPIToken) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsAPIToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPermissions

`func (o *ModelsAPIToken) GetPermissions() map[string][]string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ModelsAPIToken) GetPermissionsOk() (*map[string][]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ModelsAPIToken) SetPermissions(v map[string][]string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ModelsAPIToken) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetTitle

`func (o *ModelsAPIToken) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModelsAPIToken) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModelsAPIToken) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ModelsAPIToken) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetToken

`func (o *ModelsAPIToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ModelsAPIToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ModelsAPIToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ModelsAPIToken) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


