# PostGiftPremiumSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | Unique identifier of the target user who will receive a Telegram Premium subscription | 
**MonthCount** | **int32** | Number of months the Telegram Premium subscription will be active for the user; must be one of 3, 6, or 12 | 
**StarCount** | **int32** | Number of Telegram Stars to pay for the Telegram Premium subscription; must be 1000 for 3 months, 1500 for 6 months, and 2500 for 12 months | 
**Text** | Pointer to **string** | Text that will be shown along with the service message about the subscription; 0-128 characters | [optional] 
**TextParseMode** | Pointer to **string** | Mode for parsing entities in the text. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. Entities other than “bold”, “italic”, “underline”, “strikethrough”, “spoiler”, and “custom\\_emoji” are ignored. | [optional] 
**TextEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | A JSON-serialized list of special entities that appear in the gift text. It can be specified instead of *text\\_parse\\_mode*. Entities other than “bold”, “italic”, “underline”, “strikethrough”, “spoiler”, and “custom\\_emoji” are ignored. | [optional] 

## Methods

### NewPostGiftPremiumSubscriptionRequest

`func NewPostGiftPremiumSubscriptionRequest(userId int32, monthCount int32, starCount int32, ) *PostGiftPremiumSubscriptionRequest`

NewPostGiftPremiumSubscriptionRequest instantiates a new PostGiftPremiumSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostGiftPremiumSubscriptionRequestWithDefaults

`func NewPostGiftPremiumSubscriptionRequestWithDefaults() *PostGiftPremiumSubscriptionRequest`

NewPostGiftPremiumSubscriptionRequestWithDefaults instantiates a new PostGiftPremiumSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *PostGiftPremiumSubscriptionRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostGiftPremiumSubscriptionRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostGiftPremiumSubscriptionRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetMonthCount

`func (o *PostGiftPremiumSubscriptionRequest) GetMonthCount() int32`

GetMonthCount returns the MonthCount field if non-nil, zero value otherwise.

### GetMonthCountOk

`func (o *PostGiftPremiumSubscriptionRequest) GetMonthCountOk() (*int32, bool)`

GetMonthCountOk returns a tuple with the MonthCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthCount

`func (o *PostGiftPremiumSubscriptionRequest) SetMonthCount(v int32)`

SetMonthCount sets MonthCount field to given value.


### GetStarCount

`func (o *PostGiftPremiumSubscriptionRequest) GetStarCount() int32`

GetStarCount returns the StarCount field if non-nil, zero value otherwise.

### GetStarCountOk

`func (o *PostGiftPremiumSubscriptionRequest) GetStarCountOk() (*int32, bool)`

GetStarCountOk returns a tuple with the StarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarCount

`func (o *PostGiftPremiumSubscriptionRequest) SetStarCount(v int32)`

SetStarCount sets StarCount field to given value.


### GetText

`func (o *PostGiftPremiumSubscriptionRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostGiftPremiumSubscriptionRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostGiftPremiumSubscriptionRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PostGiftPremiumSubscriptionRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTextParseMode

`func (o *PostGiftPremiumSubscriptionRequest) GetTextParseMode() string`

GetTextParseMode returns the TextParseMode field if non-nil, zero value otherwise.

### GetTextParseModeOk

`func (o *PostGiftPremiumSubscriptionRequest) GetTextParseModeOk() (*string, bool)`

GetTextParseModeOk returns a tuple with the TextParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextParseMode

`func (o *PostGiftPremiumSubscriptionRequest) SetTextParseMode(v string)`

SetTextParseMode sets TextParseMode field to given value.

### HasTextParseMode

`func (o *PostGiftPremiumSubscriptionRequest) HasTextParseMode() bool`

HasTextParseMode returns a boolean if a field has been set.

### GetTextEntities

`func (o *PostGiftPremiumSubscriptionRequest) GetTextEntities() []MessageEntity`

GetTextEntities returns the TextEntities field if non-nil, zero value otherwise.

### GetTextEntitiesOk

`func (o *PostGiftPremiumSubscriptionRequest) GetTextEntitiesOk() (*[]MessageEntity, bool)`

GetTextEntitiesOk returns a tuple with the TextEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextEntities

`func (o *PostGiftPremiumSubscriptionRequest) SetTextEntities(v []MessageEntity)`

SetTextEntities sets TextEntities field to given value.

### HasTextEntities

`func (o *PostGiftPremiumSubscriptionRequest) HasTextEntities() bool`

HasTextEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


