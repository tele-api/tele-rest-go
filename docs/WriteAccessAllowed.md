# WriteAccessAllowed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromRequest** | Pointer to **bool** | *Optional*. *True*, if the access was granted after the user accepted an explicit request from a Web App sent by the method [requestWriteAccess](https://core.telegram.org/bots/webapps#initializing-mini-apps) | [optional] 
**WebAppName** | Pointer to **string** | *Optional*. Name of the Web App, if the access was granted when the Web App was launched from a link | [optional] 
**FromAttachmentMenu** | Pointer to **bool** | *Optional*. *True*, if the access was granted when the bot was added to the attachment or side menu | [optional] 

## Methods

### NewWriteAccessAllowed

`func NewWriteAccessAllowed() *WriteAccessAllowed`

NewWriteAccessAllowed instantiates a new WriteAccessAllowed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteAccessAllowedWithDefaults

`func NewWriteAccessAllowedWithDefaults() *WriteAccessAllowed`

NewWriteAccessAllowedWithDefaults instantiates a new WriteAccessAllowed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromRequest

`func (o *WriteAccessAllowed) GetFromRequest() bool`

GetFromRequest returns the FromRequest field if non-nil, zero value otherwise.

### GetFromRequestOk

`func (o *WriteAccessAllowed) GetFromRequestOk() (*bool, bool)`

GetFromRequestOk returns a tuple with the FromRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromRequest

`func (o *WriteAccessAllowed) SetFromRequest(v bool)`

SetFromRequest sets FromRequest field to given value.

### HasFromRequest

`func (o *WriteAccessAllowed) HasFromRequest() bool`

HasFromRequest returns a boolean if a field has been set.

### GetWebAppName

`func (o *WriteAccessAllowed) GetWebAppName() string`

GetWebAppName returns the WebAppName field if non-nil, zero value otherwise.

### GetWebAppNameOk

`func (o *WriteAccessAllowed) GetWebAppNameOk() (*string, bool)`

GetWebAppNameOk returns a tuple with the WebAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAppName

`func (o *WriteAccessAllowed) SetWebAppName(v string)`

SetWebAppName sets WebAppName field to given value.

### HasWebAppName

`func (o *WriteAccessAllowed) HasWebAppName() bool`

HasWebAppName returns a boolean if a field has been set.

### GetFromAttachmentMenu

`func (o *WriteAccessAllowed) GetFromAttachmentMenu() bool`

GetFromAttachmentMenu returns the FromAttachmentMenu field if non-nil, zero value otherwise.

### GetFromAttachmentMenuOk

`func (o *WriteAccessAllowed) GetFromAttachmentMenuOk() (*bool, bool)`

GetFromAttachmentMenuOk returns a tuple with the FromAttachmentMenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAttachmentMenu

`func (o *WriteAccessAllowed) SetFromAttachmentMenu(v bool)`

SetFromAttachmentMenu sets FromAttachmentMenu field to given value.

### HasFromAttachmentMenu

`func (o *WriteAccessAllowed) HasFromAttachmentMenu() bool`

HasFromAttachmentMenu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


