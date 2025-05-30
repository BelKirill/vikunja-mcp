# V1VikunjaInfos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to [**V1AuthInfo**](V1AuthInfo.md) |  | [optional] 
**AvailableMigrators** | Pointer to **[]string** |  | [optional] 
**CaldavEnabled** | Pointer to **bool** |  | [optional] 
**DemoModeEnabled** | Pointer to **bool** |  | [optional] 
**EmailRemindersEnabled** | Pointer to **bool** |  | [optional] 
**EnabledBackgroundProviders** | Pointer to **[]string** |  | [optional] 
**FrontendUrl** | Pointer to **string** |  | [optional] 
**Legal** | Pointer to [**V1LegalInfo**](V1LegalInfo.md) |  | [optional] 
**LinkSharingEnabled** | Pointer to **bool** |  | [optional] 
**MaxFileSize** | Pointer to **string** |  | [optional] 
**Motd** | Pointer to **string** |  | [optional] 
**PublicTeamsEnabled** | Pointer to **bool** |  | [optional] 
**RegistrationEnabled** | Pointer to **bool** |  | [optional] 
**TaskAttachmentsEnabled** | Pointer to **bool** |  | [optional] 
**TaskCommentsEnabled** | Pointer to **bool** |  | [optional] 
**TotpEnabled** | Pointer to **bool** |  | [optional] 
**UserDeletionEnabled** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**WebhooksEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewV1VikunjaInfos

`func NewV1VikunjaInfos() *V1VikunjaInfos`

NewV1VikunjaInfos instantiates a new V1VikunjaInfos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1VikunjaInfosWithDefaults

`func NewV1VikunjaInfosWithDefaults() *V1VikunjaInfos`

NewV1VikunjaInfosWithDefaults instantiates a new V1VikunjaInfos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *V1VikunjaInfos) GetAuth() V1AuthInfo`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *V1VikunjaInfos) GetAuthOk() (*V1AuthInfo, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *V1VikunjaInfos) SetAuth(v V1AuthInfo)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *V1VikunjaInfos) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetAvailableMigrators

`func (o *V1VikunjaInfos) GetAvailableMigrators() []string`

GetAvailableMigrators returns the AvailableMigrators field if non-nil, zero value otherwise.

### GetAvailableMigratorsOk

`func (o *V1VikunjaInfos) GetAvailableMigratorsOk() (*[]string, bool)`

GetAvailableMigratorsOk returns a tuple with the AvailableMigrators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMigrators

`func (o *V1VikunjaInfos) SetAvailableMigrators(v []string)`

SetAvailableMigrators sets AvailableMigrators field to given value.

### HasAvailableMigrators

`func (o *V1VikunjaInfos) HasAvailableMigrators() bool`

HasAvailableMigrators returns a boolean if a field has been set.

### GetCaldavEnabled

`func (o *V1VikunjaInfos) GetCaldavEnabled() bool`

GetCaldavEnabled returns the CaldavEnabled field if non-nil, zero value otherwise.

### GetCaldavEnabledOk

`func (o *V1VikunjaInfos) GetCaldavEnabledOk() (*bool, bool)`

GetCaldavEnabledOk returns a tuple with the CaldavEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaldavEnabled

`func (o *V1VikunjaInfos) SetCaldavEnabled(v bool)`

SetCaldavEnabled sets CaldavEnabled field to given value.

### HasCaldavEnabled

`func (o *V1VikunjaInfos) HasCaldavEnabled() bool`

HasCaldavEnabled returns a boolean if a field has been set.

### GetDemoModeEnabled

`func (o *V1VikunjaInfos) GetDemoModeEnabled() bool`

GetDemoModeEnabled returns the DemoModeEnabled field if non-nil, zero value otherwise.

### GetDemoModeEnabledOk

`func (o *V1VikunjaInfos) GetDemoModeEnabledOk() (*bool, bool)`

GetDemoModeEnabledOk returns a tuple with the DemoModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemoModeEnabled

`func (o *V1VikunjaInfos) SetDemoModeEnabled(v bool)`

SetDemoModeEnabled sets DemoModeEnabled field to given value.

### HasDemoModeEnabled

`func (o *V1VikunjaInfos) HasDemoModeEnabled() bool`

HasDemoModeEnabled returns a boolean if a field has been set.

### GetEmailRemindersEnabled

`func (o *V1VikunjaInfos) GetEmailRemindersEnabled() bool`

GetEmailRemindersEnabled returns the EmailRemindersEnabled field if non-nil, zero value otherwise.

### GetEmailRemindersEnabledOk

`func (o *V1VikunjaInfos) GetEmailRemindersEnabledOk() (*bool, bool)`

GetEmailRemindersEnabledOk returns a tuple with the EmailRemindersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailRemindersEnabled

`func (o *V1VikunjaInfos) SetEmailRemindersEnabled(v bool)`

SetEmailRemindersEnabled sets EmailRemindersEnabled field to given value.

### HasEmailRemindersEnabled

`func (o *V1VikunjaInfos) HasEmailRemindersEnabled() bool`

HasEmailRemindersEnabled returns a boolean if a field has been set.

### GetEnabledBackgroundProviders

`func (o *V1VikunjaInfos) GetEnabledBackgroundProviders() []string`

GetEnabledBackgroundProviders returns the EnabledBackgroundProviders field if non-nil, zero value otherwise.

### GetEnabledBackgroundProvidersOk

`func (o *V1VikunjaInfos) GetEnabledBackgroundProvidersOk() (*[]string, bool)`

GetEnabledBackgroundProvidersOk returns a tuple with the EnabledBackgroundProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledBackgroundProviders

`func (o *V1VikunjaInfos) SetEnabledBackgroundProviders(v []string)`

SetEnabledBackgroundProviders sets EnabledBackgroundProviders field to given value.

### HasEnabledBackgroundProviders

`func (o *V1VikunjaInfos) HasEnabledBackgroundProviders() bool`

HasEnabledBackgroundProviders returns a boolean if a field has been set.

### GetFrontendUrl

`func (o *V1VikunjaInfos) GetFrontendUrl() string`

GetFrontendUrl returns the FrontendUrl field if non-nil, zero value otherwise.

### GetFrontendUrlOk

`func (o *V1VikunjaInfos) GetFrontendUrlOk() (*string, bool)`

GetFrontendUrlOk returns a tuple with the FrontendUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontendUrl

`func (o *V1VikunjaInfos) SetFrontendUrl(v string)`

SetFrontendUrl sets FrontendUrl field to given value.

### HasFrontendUrl

`func (o *V1VikunjaInfos) HasFrontendUrl() bool`

HasFrontendUrl returns a boolean if a field has been set.

### GetLegal

`func (o *V1VikunjaInfos) GetLegal() V1LegalInfo`

GetLegal returns the Legal field if non-nil, zero value otherwise.

### GetLegalOk

`func (o *V1VikunjaInfos) GetLegalOk() (*V1LegalInfo, bool)`

GetLegalOk returns a tuple with the Legal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegal

`func (o *V1VikunjaInfos) SetLegal(v V1LegalInfo)`

SetLegal sets Legal field to given value.

### HasLegal

`func (o *V1VikunjaInfos) HasLegal() bool`

HasLegal returns a boolean if a field has been set.

### GetLinkSharingEnabled

`func (o *V1VikunjaInfos) GetLinkSharingEnabled() bool`

GetLinkSharingEnabled returns the LinkSharingEnabled field if non-nil, zero value otherwise.

### GetLinkSharingEnabledOk

`func (o *V1VikunjaInfos) GetLinkSharingEnabledOk() (*bool, bool)`

GetLinkSharingEnabledOk returns a tuple with the LinkSharingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSharingEnabled

`func (o *V1VikunjaInfos) SetLinkSharingEnabled(v bool)`

SetLinkSharingEnabled sets LinkSharingEnabled field to given value.

### HasLinkSharingEnabled

`func (o *V1VikunjaInfos) HasLinkSharingEnabled() bool`

HasLinkSharingEnabled returns a boolean if a field has been set.

### GetMaxFileSize

`func (o *V1VikunjaInfos) GetMaxFileSize() string`

GetMaxFileSize returns the MaxFileSize field if non-nil, zero value otherwise.

### GetMaxFileSizeOk

`func (o *V1VikunjaInfos) GetMaxFileSizeOk() (*string, bool)`

GetMaxFileSizeOk returns a tuple with the MaxFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFileSize

`func (o *V1VikunjaInfos) SetMaxFileSize(v string)`

SetMaxFileSize sets MaxFileSize field to given value.

### HasMaxFileSize

`func (o *V1VikunjaInfos) HasMaxFileSize() bool`

HasMaxFileSize returns a boolean if a field has been set.

### GetMotd

`func (o *V1VikunjaInfos) GetMotd() string`

GetMotd returns the Motd field if non-nil, zero value otherwise.

### GetMotdOk

`func (o *V1VikunjaInfos) GetMotdOk() (*string, bool)`

GetMotdOk returns a tuple with the Motd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotd

`func (o *V1VikunjaInfos) SetMotd(v string)`

SetMotd sets Motd field to given value.

### HasMotd

`func (o *V1VikunjaInfos) HasMotd() bool`

HasMotd returns a boolean if a field has been set.

### GetPublicTeamsEnabled

`func (o *V1VikunjaInfos) GetPublicTeamsEnabled() bool`

GetPublicTeamsEnabled returns the PublicTeamsEnabled field if non-nil, zero value otherwise.

### GetPublicTeamsEnabledOk

`func (o *V1VikunjaInfos) GetPublicTeamsEnabledOk() (*bool, bool)`

GetPublicTeamsEnabledOk returns a tuple with the PublicTeamsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicTeamsEnabled

`func (o *V1VikunjaInfos) SetPublicTeamsEnabled(v bool)`

SetPublicTeamsEnabled sets PublicTeamsEnabled field to given value.

### HasPublicTeamsEnabled

`func (o *V1VikunjaInfos) HasPublicTeamsEnabled() bool`

HasPublicTeamsEnabled returns a boolean if a field has been set.

### GetRegistrationEnabled

`func (o *V1VikunjaInfos) GetRegistrationEnabled() bool`

GetRegistrationEnabled returns the RegistrationEnabled field if non-nil, zero value otherwise.

### GetRegistrationEnabledOk

`func (o *V1VikunjaInfos) GetRegistrationEnabledOk() (*bool, bool)`

GetRegistrationEnabledOk returns a tuple with the RegistrationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationEnabled

`func (o *V1VikunjaInfos) SetRegistrationEnabled(v bool)`

SetRegistrationEnabled sets RegistrationEnabled field to given value.

### HasRegistrationEnabled

`func (o *V1VikunjaInfos) HasRegistrationEnabled() bool`

HasRegistrationEnabled returns a boolean if a field has been set.

### GetTaskAttachmentsEnabled

`func (o *V1VikunjaInfos) GetTaskAttachmentsEnabled() bool`

GetTaskAttachmentsEnabled returns the TaskAttachmentsEnabled field if non-nil, zero value otherwise.

### GetTaskAttachmentsEnabledOk

`func (o *V1VikunjaInfos) GetTaskAttachmentsEnabledOk() (*bool, bool)`

GetTaskAttachmentsEnabledOk returns a tuple with the TaskAttachmentsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskAttachmentsEnabled

`func (o *V1VikunjaInfos) SetTaskAttachmentsEnabled(v bool)`

SetTaskAttachmentsEnabled sets TaskAttachmentsEnabled field to given value.

### HasTaskAttachmentsEnabled

`func (o *V1VikunjaInfos) HasTaskAttachmentsEnabled() bool`

HasTaskAttachmentsEnabled returns a boolean if a field has been set.

### GetTaskCommentsEnabled

`func (o *V1VikunjaInfos) GetTaskCommentsEnabled() bool`

GetTaskCommentsEnabled returns the TaskCommentsEnabled field if non-nil, zero value otherwise.

### GetTaskCommentsEnabledOk

`func (o *V1VikunjaInfos) GetTaskCommentsEnabledOk() (*bool, bool)`

GetTaskCommentsEnabledOk returns a tuple with the TaskCommentsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCommentsEnabled

`func (o *V1VikunjaInfos) SetTaskCommentsEnabled(v bool)`

SetTaskCommentsEnabled sets TaskCommentsEnabled field to given value.

### HasTaskCommentsEnabled

`func (o *V1VikunjaInfos) HasTaskCommentsEnabled() bool`

HasTaskCommentsEnabled returns a boolean if a field has been set.

### GetTotpEnabled

`func (o *V1VikunjaInfos) GetTotpEnabled() bool`

GetTotpEnabled returns the TotpEnabled field if non-nil, zero value otherwise.

### GetTotpEnabledOk

`func (o *V1VikunjaInfos) GetTotpEnabledOk() (*bool, bool)`

GetTotpEnabledOk returns a tuple with the TotpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpEnabled

`func (o *V1VikunjaInfos) SetTotpEnabled(v bool)`

SetTotpEnabled sets TotpEnabled field to given value.

### HasTotpEnabled

`func (o *V1VikunjaInfos) HasTotpEnabled() bool`

HasTotpEnabled returns a boolean if a field has been set.

### GetUserDeletionEnabled

`func (o *V1VikunjaInfos) GetUserDeletionEnabled() bool`

GetUserDeletionEnabled returns the UserDeletionEnabled field if non-nil, zero value otherwise.

### GetUserDeletionEnabledOk

`func (o *V1VikunjaInfos) GetUserDeletionEnabledOk() (*bool, bool)`

GetUserDeletionEnabledOk returns a tuple with the UserDeletionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDeletionEnabled

`func (o *V1VikunjaInfos) SetUserDeletionEnabled(v bool)`

SetUserDeletionEnabled sets UserDeletionEnabled field to given value.

### HasUserDeletionEnabled

`func (o *V1VikunjaInfos) HasUserDeletionEnabled() bool`

HasUserDeletionEnabled returns a boolean if a field has been set.

### GetVersion

`func (o *V1VikunjaInfos) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V1VikunjaInfos) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V1VikunjaInfos) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V1VikunjaInfos) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWebhooksEnabled

`func (o *V1VikunjaInfos) GetWebhooksEnabled() bool`

GetWebhooksEnabled returns the WebhooksEnabled field if non-nil, zero value otherwise.

### GetWebhooksEnabledOk

`func (o *V1VikunjaInfos) GetWebhooksEnabledOk() (*bool, bool)`

GetWebhooksEnabledOk returns a tuple with the WebhooksEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooksEnabled

`func (o *V1VikunjaInfos) SetWebhooksEnabled(v bool)`

SetWebhooksEnabled sets WebhooksEnabled field to given value.

### HasWebhooksEnabled

`func (o *V1VikunjaInfos) HasWebhooksEnabled() bool`

HasWebhooksEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


