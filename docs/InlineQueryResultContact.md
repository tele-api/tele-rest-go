# InlineQueryResultContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the result, must be *contact* | [default to "contact"]
**Id** | **string** | Unique identifier for this result, 1-64 Bytes | 
**PhoneNumber** | **string** | Contact&#39;s phone number | 
**FirstName** | **string** | Contact&#39;s first name | 
**LastName** | Pointer to **string** | *Optional*. Contact&#39;s last name | [optional] 
**Vcard** | Pointer to **string** | *Optional*. Additional data about the contact in the form of a [vCard](https://en.wikipedia.org/wiki/VCard), 0-2048 bytes | [optional] 
**ReplyMarkup** | Pointer to [**InlineKeyboardMarkup**](InlineKeyboardMarkup.md) |  | [optional] 
**InputMessageContent** | Pointer to [**InputMessageContent**](InputMessageContent.md) |  | [optional] 
**ThumbnailUrl** | Pointer to **string** | *Optional*. Url of the thumbnail for the result | [optional] 
**ThumbnailWidth** | Pointer to **int32** | *Optional*. Thumbnail width | [optional] 
**ThumbnailHeight** | Pointer to **int32** | *Optional*. Thumbnail height | [optional] 

## Methods

### NewInlineQueryResultContact

`func NewInlineQueryResultContact(type_ string, id string, phoneNumber string, firstName string, ) *InlineQueryResultContact`

NewInlineQueryResultContact instantiates a new InlineQueryResultContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineQueryResultContactWithDefaults

`func NewInlineQueryResultContactWithDefaults() *InlineQueryResultContact`

NewInlineQueryResultContactWithDefaults instantiates a new InlineQueryResultContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineQueryResultContact) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineQueryResultContact) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineQueryResultContact) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *InlineQueryResultContact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineQueryResultContact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineQueryResultContact) SetId(v string)`

SetId sets Id field to given value.


### GetPhoneNumber

`func (o *InlineQueryResultContact) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *InlineQueryResultContact) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *InlineQueryResultContact) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetFirstName

`func (o *InlineQueryResultContact) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *InlineQueryResultContact) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *InlineQueryResultContact) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *InlineQueryResultContact) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *InlineQueryResultContact) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *InlineQueryResultContact) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *InlineQueryResultContact) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetVcard

`func (o *InlineQueryResultContact) GetVcard() string`

GetVcard returns the Vcard field if non-nil, zero value otherwise.

### GetVcardOk

`func (o *InlineQueryResultContact) GetVcardOk() (*string, bool)`

GetVcardOk returns a tuple with the Vcard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcard

`func (o *InlineQueryResultContact) SetVcard(v string)`

SetVcard sets Vcard field to given value.

### HasVcard

`func (o *InlineQueryResultContact) HasVcard() bool`

HasVcard returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *InlineQueryResultContact) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *InlineQueryResultContact) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *InlineQueryResultContact) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *InlineQueryResultContact) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.

### GetInputMessageContent

`func (o *InlineQueryResultContact) GetInputMessageContent() InputMessageContent`

GetInputMessageContent returns the InputMessageContent field if non-nil, zero value otherwise.

### GetInputMessageContentOk

`func (o *InlineQueryResultContact) GetInputMessageContentOk() (*InputMessageContent, bool)`

GetInputMessageContentOk returns a tuple with the InputMessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputMessageContent

`func (o *InlineQueryResultContact) SetInputMessageContent(v InputMessageContent)`

SetInputMessageContent sets InputMessageContent field to given value.

### HasInputMessageContent

`func (o *InlineQueryResultContact) HasInputMessageContent() bool`

HasInputMessageContent returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *InlineQueryResultContact) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *InlineQueryResultContact) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *InlineQueryResultContact) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *InlineQueryResultContact) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetThumbnailWidth

`func (o *InlineQueryResultContact) GetThumbnailWidth() int32`

GetThumbnailWidth returns the ThumbnailWidth field if non-nil, zero value otherwise.

### GetThumbnailWidthOk

`func (o *InlineQueryResultContact) GetThumbnailWidthOk() (*int32, bool)`

GetThumbnailWidthOk returns a tuple with the ThumbnailWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailWidth

`func (o *InlineQueryResultContact) SetThumbnailWidth(v int32)`

SetThumbnailWidth sets ThumbnailWidth field to given value.

### HasThumbnailWidth

`func (o *InlineQueryResultContact) HasThumbnailWidth() bool`

HasThumbnailWidth returns a boolean if a field has been set.

### GetThumbnailHeight

`func (o *InlineQueryResultContact) GetThumbnailHeight() int32`

GetThumbnailHeight returns the ThumbnailHeight field if non-nil, zero value otherwise.

### GetThumbnailHeightOk

`func (o *InlineQueryResultContact) GetThumbnailHeightOk() (*int32, bool)`

GetThumbnailHeightOk returns a tuple with the ThumbnailHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailHeight

`func (o *InlineQueryResultContact) SetThumbnailHeight(v int32)`

SetThumbnailHeight sets ThumbnailHeight field to given value.

### HasThumbnailHeight

`func (o *InlineQueryResultContact) HasThumbnailHeight() bool`

HasThumbnailHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


