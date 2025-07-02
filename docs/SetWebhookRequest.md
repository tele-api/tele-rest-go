# SetWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | HTTPS URL to send updates to. Use an empty string to remove webhook integration | 
**Certificate** | Pointer to **interface{}** |  | [optional] 
**IpAddress** | Pointer to **string** | The fixed IP address which will be used to send webhook requests instead of the IP address resolved through DNS | [optional] 
**MaxConnections** | Pointer to **int32** | The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to *40*. Use lower values to limit the load on your bot&#39;s server, and higher values to increase your bot&#39;s throughput. | [optional] [default to 40]
**AllowedUpdates** | Pointer to **[]string** | A JSON-serialized list of the update types you want your bot to receive. For example, specify &#x60;[\&quot;message\&quot;, \&quot;edited_channel_post\&quot;, \&quot;callback_query\&quot;]&#x60; to only receive updates of these types. See [Update](https://core.telegram.org/bots/api/#update) for a complete list of available update types. Specify an empty list to receive all update types except *chat\\_member*, *message\\_reaction*, and *message\\_reaction\\_count* (default). If not specified, the previous setting will be used.   Please note that this parameter doesn&#39;t affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time. | [optional] 
**DropPendingUpdates** | Pointer to **bool** | Pass *True* to drop all pending updates | [optional] 
**SecretToken** | Pointer to **string** | A secret token to be sent in a header “X-Telegram-Bot-Api-Secret-Token” in every webhook request, 1-256 characters. Only characters &#x60;A-Z&#x60;, &#x60;a-z&#x60;, &#x60;0-9&#x60;, &#x60;_&#x60; and &#x60;-&#x60; are allowed. The header is useful to ensure that the request comes from a webhook set by you. | [optional] 

## Methods

### NewSetWebhookRequest

`func NewSetWebhookRequest(url string, ) *SetWebhookRequest`

NewSetWebhookRequest instantiates a new SetWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetWebhookRequestWithDefaults

`func NewSetWebhookRequestWithDefaults() *SetWebhookRequest`

NewSetWebhookRequestWithDefaults instantiates a new SetWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *SetWebhookRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SetWebhookRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SetWebhookRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCertificate

`func (o *SetWebhookRequest) GetCertificate() interface{}`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SetWebhookRequest) GetCertificateOk() (*interface{}, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SetWebhookRequest) SetCertificate(v interface{})`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *SetWebhookRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *SetWebhookRequest) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *SetWebhookRequest) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetIpAddress

`func (o *SetWebhookRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *SetWebhookRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *SetWebhookRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *SetWebhookRequest) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMaxConnections

`func (o *SetWebhookRequest) GetMaxConnections() int32`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *SetWebhookRequest) GetMaxConnectionsOk() (*int32, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *SetWebhookRequest) SetMaxConnections(v int32)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *SetWebhookRequest) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetAllowedUpdates

`func (o *SetWebhookRequest) GetAllowedUpdates() []string`

GetAllowedUpdates returns the AllowedUpdates field if non-nil, zero value otherwise.

### GetAllowedUpdatesOk

`func (o *SetWebhookRequest) GetAllowedUpdatesOk() (*[]string, bool)`

GetAllowedUpdatesOk returns a tuple with the AllowedUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUpdates

`func (o *SetWebhookRequest) SetAllowedUpdates(v []string)`

SetAllowedUpdates sets AllowedUpdates field to given value.

### HasAllowedUpdates

`func (o *SetWebhookRequest) HasAllowedUpdates() bool`

HasAllowedUpdates returns a boolean if a field has been set.

### GetDropPendingUpdates

`func (o *SetWebhookRequest) GetDropPendingUpdates() bool`

GetDropPendingUpdates returns the DropPendingUpdates field if non-nil, zero value otherwise.

### GetDropPendingUpdatesOk

`func (o *SetWebhookRequest) GetDropPendingUpdatesOk() (*bool, bool)`

GetDropPendingUpdatesOk returns a tuple with the DropPendingUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropPendingUpdates

`func (o *SetWebhookRequest) SetDropPendingUpdates(v bool)`

SetDropPendingUpdates sets DropPendingUpdates field to given value.

### HasDropPendingUpdates

`func (o *SetWebhookRequest) HasDropPendingUpdates() bool`

HasDropPendingUpdates returns a boolean if a field has been set.

### GetSecretToken

`func (o *SetWebhookRequest) GetSecretToken() string`

GetSecretToken returns the SecretToken field if non-nil, zero value otherwise.

### GetSecretTokenOk

`func (o *SetWebhookRequest) GetSecretTokenOk() (*string, bool)`

GetSecretTokenOk returns a tuple with the SecretToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretToken

`func (o *SetWebhookRequest) SetSecretToken(v string)`

SetSecretToken sets SecretToken field to given value.

### HasSecretToken

`func (o *SetWebhookRequest) HasSecretToken() bool`

HasSecretToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


