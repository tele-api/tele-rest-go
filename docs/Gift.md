# Gift

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the gift | 
**Sticker** | [**Sticker**](Sticker.md) |  | 
**StarCount** | **int32** | The number of Telegram Stars that must be paid to send the sticker | 
**UpgradeStarCount** | Pointer to **int32** | *Optional*. The number of Telegram Stars that must be paid to upgrade the gift to a unique one | [optional] 
**IsPremium** | Pointer to **bool** | *Optional*. *True*, if the gift can only be purchased by Telegram Premium subscribers | [optional] [default to true]
**HasColors** | Pointer to **bool** | *Optional*. *True*, if the gift can be used (after being upgraded) to customize a user&#39;s appearance | [optional] [default to true]
**TotalCount** | Pointer to **int32** | *Optional*. The total number of gifts of this type that can be sent by all users; for limited gifts only | [optional] 
**RemainingCount** | Pointer to **int32** | *Optional*. The number of remaining gifts of this type that can be sent by all users; for limited gifts only | [optional] 
**PersonalTotalCount** | Pointer to **int32** | *Optional*. The total number of gifts of this type that can be sent by the bot; for limited gifts only | [optional] 
**PersonalRemainingCount** | Pointer to **int32** | *Optional*. The number of remaining gifts of this type that can be sent by the bot; for limited gifts only | [optional] 
**Background** | Pointer to [**GiftBackground**](GiftBackground.md) |  | [optional] 
**UniqueGiftVariantCount** | Pointer to **int32** | *Optional*. The total number of different unique gifts that can be obtained by upgrading the gift | [optional] 
**PublisherChat** | Pointer to [**Chat**](Chat.md) |  | [optional] 

## Methods

### NewGift

`func NewGift(id string, sticker Sticker, starCount int32, ) *Gift`

NewGift instantiates a new Gift object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftWithDefaults

`func NewGiftWithDefaults() *Gift`

NewGiftWithDefaults instantiates a new Gift object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Gift) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Gift) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Gift) SetId(v string)`

SetId sets Id field to given value.


### GetSticker

`func (o *Gift) GetSticker() Sticker`

GetSticker returns the Sticker field if non-nil, zero value otherwise.

### GetStickerOk

`func (o *Gift) GetStickerOk() (*Sticker, bool)`

GetStickerOk returns a tuple with the Sticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticker

`func (o *Gift) SetSticker(v Sticker)`

SetSticker sets Sticker field to given value.


### GetStarCount

`func (o *Gift) GetStarCount() int32`

GetStarCount returns the StarCount field if non-nil, zero value otherwise.

### GetStarCountOk

`func (o *Gift) GetStarCountOk() (*int32, bool)`

GetStarCountOk returns a tuple with the StarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarCount

`func (o *Gift) SetStarCount(v int32)`

SetStarCount sets StarCount field to given value.


### GetUpgradeStarCount

`func (o *Gift) GetUpgradeStarCount() int32`

GetUpgradeStarCount returns the UpgradeStarCount field if non-nil, zero value otherwise.

### GetUpgradeStarCountOk

`func (o *Gift) GetUpgradeStarCountOk() (*int32, bool)`

GetUpgradeStarCountOk returns a tuple with the UpgradeStarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStarCount

`func (o *Gift) SetUpgradeStarCount(v int32)`

SetUpgradeStarCount sets UpgradeStarCount field to given value.

### HasUpgradeStarCount

`func (o *Gift) HasUpgradeStarCount() bool`

HasUpgradeStarCount returns a boolean if a field has been set.

### GetIsPremium

`func (o *Gift) GetIsPremium() bool`

GetIsPremium returns the IsPremium field if non-nil, zero value otherwise.

### GetIsPremiumOk

`func (o *Gift) GetIsPremiumOk() (*bool, bool)`

GetIsPremiumOk returns a tuple with the IsPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPremium

`func (o *Gift) SetIsPremium(v bool)`

SetIsPremium sets IsPremium field to given value.

### HasIsPremium

`func (o *Gift) HasIsPremium() bool`

HasIsPremium returns a boolean if a field has been set.

### GetHasColors

`func (o *Gift) GetHasColors() bool`

GetHasColors returns the HasColors field if non-nil, zero value otherwise.

### GetHasColorsOk

`func (o *Gift) GetHasColorsOk() (*bool, bool)`

GetHasColorsOk returns a tuple with the HasColors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasColors

`func (o *Gift) SetHasColors(v bool)`

SetHasColors sets HasColors field to given value.

### HasHasColors

`func (o *Gift) HasHasColors() bool`

HasHasColors returns a boolean if a field has been set.

### GetTotalCount

`func (o *Gift) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *Gift) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *Gift) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *Gift) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetRemainingCount

`func (o *Gift) GetRemainingCount() int32`

GetRemainingCount returns the RemainingCount field if non-nil, zero value otherwise.

### GetRemainingCountOk

`func (o *Gift) GetRemainingCountOk() (*int32, bool)`

GetRemainingCountOk returns a tuple with the RemainingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCount

`func (o *Gift) SetRemainingCount(v int32)`

SetRemainingCount sets RemainingCount field to given value.

### HasRemainingCount

`func (o *Gift) HasRemainingCount() bool`

HasRemainingCount returns a boolean if a field has been set.

### GetPersonalTotalCount

`func (o *Gift) GetPersonalTotalCount() int32`

GetPersonalTotalCount returns the PersonalTotalCount field if non-nil, zero value otherwise.

### GetPersonalTotalCountOk

`func (o *Gift) GetPersonalTotalCountOk() (*int32, bool)`

GetPersonalTotalCountOk returns a tuple with the PersonalTotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalTotalCount

`func (o *Gift) SetPersonalTotalCount(v int32)`

SetPersonalTotalCount sets PersonalTotalCount field to given value.

### HasPersonalTotalCount

`func (o *Gift) HasPersonalTotalCount() bool`

HasPersonalTotalCount returns a boolean if a field has been set.

### GetPersonalRemainingCount

`func (o *Gift) GetPersonalRemainingCount() int32`

GetPersonalRemainingCount returns the PersonalRemainingCount field if non-nil, zero value otherwise.

### GetPersonalRemainingCountOk

`func (o *Gift) GetPersonalRemainingCountOk() (*int32, bool)`

GetPersonalRemainingCountOk returns a tuple with the PersonalRemainingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalRemainingCount

`func (o *Gift) SetPersonalRemainingCount(v int32)`

SetPersonalRemainingCount sets PersonalRemainingCount field to given value.

### HasPersonalRemainingCount

`func (o *Gift) HasPersonalRemainingCount() bool`

HasPersonalRemainingCount returns a boolean if a field has been set.

### GetBackground

`func (o *Gift) GetBackground() GiftBackground`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *Gift) GetBackgroundOk() (*GiftBackground, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *Gift) SetBackground(v GiftBackground)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *Gift) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### GetUniqueGiftVariantCount

`func (o *Gift) GetUniqueGiftVariantCount() int32`

GetUniqueGiftVariantCount returns the UniqueGiftVariantCount field if non-nil, zero value otherwise.

### GetUniqueGiftVariantCountOk

`func (o *Gift) GetUniqueGiftVariantCountOk() (*int32, bool)`

GetUniqueGiftVariantCountOk returns a tuple with the UniqueGiftVariantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueGiftVariantCount

`func (o *Gift) SetUniqueGiftVariantCount(v int32)`

SetUniqueGiftVariantCount sets UniqueGiftVariantCount field to given value.

### HasUniqueGiftVariantCount

`func (o *Gift) HasUniqueGiftVariantCount() bool`

HasUniqueGiftVariantCount returns a boolean if a field has been set.

### GetPublisherChat

`func (o *Gift) GetPublisherChat() Chat`

GetPublisherChat returns the PublisherChat field if non-nil, zero value otherwise.

### GetPublisherChatOk

`func (o *Gift) GetPublisherChatOk() (*Chat, bool)`

GetPublisherChatOk returns a tuple with the PublisherChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherChat

`func (o *Gift) SetPublisherChat(v Chat)`

SetPublisherChat sets PublisherChat field to given value.

### HasPublisherChat

`func (o *Gift) HasPublisherChat() bool`

HasPublisherChat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


