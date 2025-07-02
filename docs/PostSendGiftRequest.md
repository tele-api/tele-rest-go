# PostSendGiftRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int32** | Required if *chat\\_id* is not specified. Unique identifier of the target user who will receive the gift. | [optional] 
**ChatId** | Pointer to [**PostSendGiftRequestChatId**](PostSendGiftRequestChatId.md) |  | [optional] 
**GiftId** | **string** | Identifier of the gift | 
**PayForUpgrade** | Pointer to **bool** | Pass *True* to pay for the gift upgrade from the bot&#39;s balance, thereby making the upgrade free for the receiver | [optional] 
**Text** | Pointer to **string** | Text that will be shown along with the gift; 0-128 characters | [optional] 
**TextParseMode** | Pointer to **string** | Mode for parsing entities in the text. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. Entities other than “bold”, “italic”, “underline”, “strikethrough”, “spoiler”, and “custom\\_emoji” are ignored. | [optional] 
**TextEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | A JSON-serialized list of special entities that appear in the gift text. It can be specified instead of *text\\_parse\\_mode*. Entities other than “bold”, “italic”, “underline”, “strikethrough”, “spoiler”, and “custom\\_emoji” are ignored. | [optional] 

## Methods

### NewPostSendGiftRequest

`func NewPostSendGiftRequest(giftId string, ) *PostSendGiftRequest`

NewPostSendGiftRequest instantiates a new PostSendGiftRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSendGiftRequestWithDefaults

`func NewPostSendGiftRequestWithDefaults() *PostSendGiftRequest`

NewPostSendGiftRequestWithDefaults instantiates a new PostSendGiftRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *PostSendGiftRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostSendGiftRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostSendGiftRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PostSendGiftRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetChatId

`func (o *PostSendGiftRequest) GetChatId() PostSendGiftRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *PostSendGiftRequest) GetChatIdOk() (*PostSendGiftRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *PostSendGiftRequest) SetChatId(v PostSendGiftRequestChatId)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *PostSendGiftRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetGiftId

`func (o *PostSendGiftRequest) GetGiftId() string`

GetGiftId returns the GiftId field if non-nil, zero value otherwise.

### GetGiftIdOk

`func (o *PostSendGiftRequest) GetGiftIdOk() (*string, bool)`

GetGiftIdOk returns a tuple with the GiftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftId

`func (o *PostSendGiftRequest) SetGiftId(v string)`

SetGiftId sets GiftId field to given value.


### GetPayForUpgrade

`func (o *PostSendGiftRequest) GetPayForUpgrade() bool`

GetPayForUpgrade returns the PayForUpgrade field if non-nil, zero value otherwise.

### GetPayForUpgradeOk

`func (o *PostSendGiftRequest) GetPayForUpgradeOk() (*bool, bool)`

GetPayForUpgradeOk returns a tuple with the PayForUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayForUpgrade

`func (o *PostSendGiftRequest) SetPayForUpgrade(v bool)`

SetPayForUpgrade sets PayForUpgrade field to given value.

### HasPayForUpgrade

`func (o *PostSendGiftRequest) HasPayForUpgrade() bool`

HasPayForUpgrade returns a boolean if a field has been set.

### GetText

`func (o *PostSendGiftRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostSendGiftRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostSendGiftRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PostSendGiftRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTextParseMode

`func (o *PostSendGiftRequest) GetTextParseMode() string`

GetTextParseMode returns the TextParseMode field if non-nil, zero value otherwise.

### GetTextParseModeOk

`func (o *PostSendGiftRequest) GetTextParseModeOk() (*string, bool)`

GetTextParseModeOk returns a tuple with the TextParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextParseMode

`func (o *PostSendGiftRequest) SetTextParseMode(v string)`

SetTextParseMode sets TextParseMode field to given value.

### HasTextParseMode

`func (o *PostSendGiftRequest) HasTextParseMode() bool`

HasTextParseMode returns a boolean if a field has been set.

### GetTextEntities

`func (o *PostSendGiftRequest) GetTextEntities() []MessageEntity`

GetTextEntities returns the TextEntities field if non-nil, zero value otherwise.

### GetTextEntitiesOk

`func (o *PostSendGiftRequest) GetTextEntitiesOk() (*[]MessageEntity, bool)`

GetTextEntitiesOk returns a tuple with the TextEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextEntities

`func (o *PostSendGiftRequest) SetTextEntities(v []MessageEntity)`

SetTextEntities sets TextEntities field to given value.

### HasTextEntities

`func (o *PostSendGiftRequest) HasTextEntities() bool`

HasTextEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


