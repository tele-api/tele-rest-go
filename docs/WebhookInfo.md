# WebhookInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Webhook URL, may be empty if webhook is not set up | 
**HasCustomCertificate** | **bool** | *True*, if a custom certificate was provided for webhook certificate checks | 
**PendingUpdateCount** | **int32** | Number of updates awaiting delivery | 
**IpAddress** | Pointer to **string** | *Optional*. Currently used webhook IP address | [optional] 
**LastErrorDate** | Pointer to **int32** | *Optional*. Unix time for the most recent error that happened when trying to deliver an update via webhook | [optional] 
**LastErrorMessage** | Pointer to **string** | *Optional*. Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook | [optional] 
**LastSynchronizationErrorDate** | Pointer to **int32** | *Optional*. Unix time of the most recent error that happened when trying to synchronize available updates with Telegram datacenters | [optional] 
**MaxConnections** | Pointer to **int32** | *Optional*. The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery | [optional] 
**AllowedUpdates** | Pointer to **[]string** | *Optional*. A list of update types the bot is subscribed to. Defaults to all update types except *chat\\_member* | [optional] 

## Methods

### NewWebhookInfo

`func NewWebhookInfo(url string, hasCustomCertificate bool, pendingUpdateCount int32, ) *WebhookInfo`

NewWebhookInfo instantiates a new WebhookInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookInfoWithDefaults

`func NewWebhookInfoWithDefaults() *WebhookInfo`

NewWebhookInfoWithDefaults instantiates a new WebhookInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebhookInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookInfo) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHasCustomCertificate

`func (o *WebhookInfo) GetHasCustomCertificate() bool`

GetHasCustomCertificate returns the HasCustomCertificate field if non-nil, zero value otherwise.

### GetHasCustomCertificateOk

`func (o *WebhookInfo) GetHasCustomCertificateOk() (*bool, bool)`

GetHasCustomCertificateOk returns a tuple with the HasCustomCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCustomCertificate

`func (o *WebhookInfo) SetHasCustomCertificate(v bool)`

SetHasCustomCertificate sets HasCustomCertificate field to given value.


### GetPendingUpdateCount

`func (o *WebhookInfo) GetPendingUpdateCount() int32`

GetPendingUpdateCount returns the PendingUpdateCount field if non-nil, zero value otherwise.

### GetPendingUpdateCountOk

`func (o *WebhookInfo) GetPendingUpdateCountOk() (*int32, bool)`

GetPendingUpdateCountOk returns a tuple with the PendingUpdateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUpdateCount

`func (o *WebhookInfo) SetPendingUpdateCount(v int32)`

SetPendingUpdateCount sets PendingUpdateCount field to given value.


### GetIpAddress

`func (o *WebhookInfo) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *WebhookInfo) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *WebhookInfo) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *WebhookInfo) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLastErrorDate

`func (o *WebhookInfo) GetLastErrorDate() int32`

GetLastErrorDate returns the LastErrorDate field if non-nil, zero value otherwise.

### GetLastErrorDateOk

`func (o *WebhookInfo) GetLastErrorDateOk() (*int32, bool)`

GetLastErrorDateOk returns a tuple with the LastErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDate

`func (o *WebhookInfo) SetLastErrorDate(v int32)`

SetLastErrorDate sets LastErrorDate field to given value.

### HasLastErrorDate

`func (o *WebhookInfo) HasLastErrorDate() bool`

HasLastErrorDate returns a boolean if a field has been set.

### GetLastErrorMessage

`func (o *WebhookInfo) GetLastErrorMessage() string`

GetLastErrorMessage returns the LastErrorMessage field if non-nil, zero value otherwise.

### GetLastErrorMessageOk

`func (o *WebhookInfo) GetLastErrorMessageOk() (*string, bool)`

GetLastErrorMessageOk returns a tuple with the LastErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorMessage

`func (o *WebhookInfo) SetLastErrorMessage(v string)`

SetLastErrorMessage sets LastErrorMessage field to given value.

### HasLastErrorMessage

`func (o *WebhookInfo) HasLastErrorMessage() bool`

HasLastErrorMessage returns a boolean if a field has been set.

### GetLastSynchronizationErrorDate

`func (o *WebhookInfo) GetLastSynchronizationErrorDate() int32`

GetLastSynchronizationErrorDate returns the LastSynchronizationErrorDate field if non-nil, zero value otherwise.

### GetLastSynchronizationErrorDateOk

`func (o *WebhookInfo) GetLastSynchronizationErrorDateOk() (*int32, bool)`

GetLastSynchronizationErrorDateOk returns a tuple with the LastSynchronizationErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSynchronizationErrorDate

`func (o *WebhookInfo) SetLastSynchronizationErrorDate(v int32)`

SetLastSynchronizationErrorDate sets LastSynchronizationErrorDate field to given value.

### HasLastSynchronizationErrorDate

`func (o *WebhookInfo) HasLastSynchronizationErrorDate() bool`

HasLastSynchronizationErrorDate returns a boolean if a field has been set.

### GetMaxConnections

`func (o *WebhookInfo) GetMaxConnections() int32`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *WebhookInfo) GetMaxConnectionsOk() (*int32, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *WebhookInfo) SetMaxConnections(v int32)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *WebhookInfo) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetAllowedUpdates

`func (o *WebhookInfo) GetAllowedUpdates() []string`

GetAllowedUpdates returns the AllowedUpdates field if non-nil, zero value otherwise.

### GetAllowedUpdatesOk

`func (o *WebhookInfo) GetAllowedUpdatesOk() (*[]string, bool)`

GetAllowedUpdatesOk returns a tuple with the AllowedUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUpdates

`func (o *WebhookInfo) SetAllowedUpdates(v []string)`

SetAllowedUpdates sets AllowedUpdates field to given value.

### HasAllowedUpdates

`func (o *WebhookInfo) HasAllowedUpdates() bool`

HasAllowedUpdates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


