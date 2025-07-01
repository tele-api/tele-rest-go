# GiveawayWinners

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chat** | [**Chat**](Chat.md) |  | 
**GiveawayMessageId** | **int32** | Identifier of the message with the giveaway in the chat | 
**WinnersSelectionDate** | **int32** | Point in time (Unix timestamp) when winners of the giveaway were selected | 
**WinnerCount** | **int32** | Total number of winners in the giveaway | 
**Winners** | [**[]User**](User.md) | List of up to 100 winners of the giveaway | 
**AdditionalChatCount** | Pointer to **int32** | *Optional*. The number of other chats the user had to join in order to be eligible for the giveaway | [optional] 
**PrizeStarCount** | Pointer to **int32** | *Optional*. The number of Telegram Stars that were split between giveaway winners; for Telegram Star giveaways only | [optional] 
**PremiumSubscriptionMonthCount** | Pointer to **int32** | *Optional*. The number of months the Telegram Premium subscription won from the giveaway will be active for; for Telegram Premium giveaways only | [optional] 
**UnclaimedPrizeCount** | Pointer to **int32** | *Optional*. Number of undistributed prizes | [optional] 
**OnlyNewMembers** | Pointer to **bool** | *Optional*. *True*, if only users who had joined the chats after the giveaway started were eligible to win | [optional] [default to true]
**WasRefunded** | Pointer to **bool** | *Optional*. *True*, if the giveaway was canceled because the payment for it was refunded | [optional] [default to true]
**PrizeDescription** | Pointer to **string** | *Optional*. Description of additional giveaway prize | [optional] 

## Methods

### NewGiveawayWinners

`func NewGiveawayWinners(chat Chat, giveawayMessageId int32, winnersSelectionDate int32, winnerCount int32, winners []User, ) *GiveawayWinners`

NewGiveawayWinners instantiates a new GiveawayWinners object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiveawayWinnersWithDefaults

`func NewGiveawayWinnersWithDefaults() *GiveawayWinners`

NewGiveawayWinnersWithDefaults instantiates a new GiveawayWinners object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChat

`func (o *GiveawayWinners) GetChat() Chat`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *GiveawayWinners) GetChatOk() (*Chat, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *GiveawayWinners) SetChat(v Chat)`

SetChat sets Chat field to given value.


### GetGiveawayMessageId

`func (o *GiveawayWinners) GetGiveawayMessageId() int32`

GetGiveawayMessageId returns the GiveawayMessageId field if non-nil, zero value otherwise.

### GetGiveawayMessageIdOk

`func (o *GiveawayWinners) GetGiveawayMessageIdOk() (*int32, bool)`

GetGiveawayMessageIdOk returns a tuple with the GiveawayMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiveawayMessageId

`func (o *GiveawayWinners) SetGiveawayMessageId(v int32)`

SetGiveawayMessageId sets GiveawayMessageId field to given value.


### GetWinnersSelectionDate

`func (o *GiveawayWinners) GetWinnersSelectionDate() int32`

GetWinnersSelectionDate returns the WinnersSelectionDate field if non-nil, zero value otherwise.

### GetWinnersSelectionDateOk

`func (o *GiveawayWinners) GetWinnersSelectionDateOk() (*int32, bool)`

GetWinnersSelectionDateOk returns a tuple with the WinnersSelectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnersSelectionDate

`func (o *GiveawayWinners) SetWinnersSelectionDate(v int32)`

SetWinnersSelectionDate sets WinnersSelectionDate field to given value.


### GetWinnerCount

`func (o *GiveawayWinners) GetWinnerCount() int32`

GetWinnerCount returns the WinnerCount field if non-nil, zero value otherwise.

### GetWinnerCountOk

`func (o *GiveawayWinners) GetWinnerCountOk() (*int32, bool)`

GetWinnerCountOk returns a tuple with the WinnerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnerCount

`func (o *GiveawayWinners) SetWinnerCount(v int32)`

SetWinnerCount sets WinnerCount field to given value.


### GetWinners

`func (o *GiveawayWinners) GetWinners() []User`

GetWinners returns the Winners field if non-nil, zero value otherwise.

### GetWinnersOk

`func (o *GiveawayWinners) GetWinnersOk() (*[]User, bool)`

GetWinnersOk returns a tuple with the Winners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinners

`func (o *GiveawayWinners) SetWinners(v []User)`

SetWinners sets Winners field to given value.


### GetAdditionalChatCount

`func (o *GiveawayWinners) GetAdditionalChatCount() int32`

GetAdditionalChatCount returns the AdditionalChatCount field if non-nil, zero value otherwise.

### GetAdditionalChatCountOk

`func (o *GiveawayWinners) GetAdditionalChatCountOk() (*int32, bool)`

GetAdditionalChatCountOk returns a tuple with the AdditionalChatCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalChatCount

`func (o *GiveawayWinners) SetAdditionalChatCount(v int32)`

SetAdditionalChatCount sets AdditionalChatCount field to given value.

### HasAdditionalChatCount

`func (o *GiveawayWinners) HasAdditionalChatCount() bool`

HasAdditionalChatCount returns a boolean if a field has been set.

### GetPrizeStarCount

`func (o *GiveawayWinners) GetPrizeStarCount() int32`

GetPrizeStarCount returns the PrizeStarCount field if non-nil, zero value otherwise.

### GetPrizeStarCountOk

`func (o *GiveawayWinners) GetPrizeStarCountOk() (*int32, bool)`

GetPrizeStarCountOk returns a tuple with the PrizeStarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrizeStarCount

`func (o *GiveawayWinners) SetPrizeStarCount(v int32)`

SetPrizeStarCount sets PrizeStarCount field to given value.

### HasPrizeStarCount

`func (o *GiveawayWinners) HasPrizeStarCount() bool`

HasPrizeStarCount returns a boolean if a field has been set.

### GetPremiumSubscriptionMonthCount

`func (o *GiveawayWinners) GetPremiumSubscriptionMonthCount() int32`

GetPremiumSubscriptionMonthCount returns the PremiumSubscriptionMonthCount field if non-nil, zero value otherwise.

### GetPremiumSubscriptionMonthCountOk

`func (o *GiveawayWinners) GetPremiumSubscriptionMonthCountOk() (*int32, bool)`

GetPremiumSubscriptionMonthCountOk returns a tuple with the PremiumSubscriptionMonthCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumSubscriptionMonthCount

`func (o *GiveawayWinners) SetPremiumSubscriptionMonthCount(v int32)`

SetPremiumSubscriptionMonthCount sets PremiumSubscriptionMonthCount field to given value.

### HasPremiumSubscriptionMonthCount

`func (o *GiveawayWinners) HasPremiumSubscriptionMonthCount() bool`

HasPremiumSubscriptionMonthCount returns a boolean if a field has been set.

### GetUnclaimedPrizeCount

`func (o *GiveawayWinners) GetUnclaimedPrizeCount() int32`

GetUnclaimedPrizeCount returns the UnclaimedPrizeCount field if non-nil, zero value otherwise.

### GetUnclaimedPrizeCountOk

`func (o *GiveawayWinners) GetUnclaimedPrizeCountOk() (*int32, bool)`

GetUnclaimedPrizeCountOk returns a tuple with the UnclaimedPrizeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnclaimedPrizeCount

`func (o *GiveawayWinners) SetUnclaimedPrizeCount(v int32)`

SetUnclaimedPrizeCount sets UnclaimedPrizeCount field to given value.

### HasUnclaimedPrizeCount

`func (o *GiveawayWinners) HasUnclaimedPrizeCount() bool`

HasUnclaimedPrizeCount returns a boolean if a field has been set.

### GetOnlyNewMembers

`func (o *GiveawayWinners) GetOnlyNewMembers() bool`

GetOnlyNewMembers returns the OnlyNewMembers field if non-nil, zero value otherwise.

### GetOnlyNewMembersOk

`func (o *GiveawayWinners) GetOnlyNewMembersOk() (*bool, bool)`

GetOnlyNewMembersOk returns a tuple with the OnlyNewMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyNewMembers

`func (o *GiveawayWinners) SetOnlyNewMembers(v bool)`

SetOnlyNewMembers sets OnlyNewMembers field to given value.

### HasOnlyNewMembers

`func (o *GiveawayWinners) HasOnlyNewMembers() bool`

HasOnlyNewMembers returns a boolean if a field has been set.

### GetWasRefunded

`func (o *GiveawayWinners) GetWasRefunded() bool`

GetWasRefunded returns the WasRefunded field if non-nil, zero value otherwise.

### GetWasRefundedOk

`func (o *GiveawayWinners) GetWasRefundedOk() (*bool, bool)`

GetWasRefundedOk returns a tuple with the WasRefunded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasRefunded

`func (o *GiveawayWinners) SetWasRefunded(v bool)`

SetWasRefunded sets WasRefunded field to given value.

### HasWasRefunded

`func (o *GiveawayWinners) HasWasRefunded() bool`

HasWasRefunded returns a boolean if a field has been set.

### GetPrizeDescription

`func (o *GiveawayWinners) GetPrizeDescription() string`

GetPrizeDescription returns the PrizeDescription field if non-nil, zero value otherwise.

### GetPrizeDescriptionOk

`func (o *GiveawayWinners) GetPrizeDescriptionOk() (*string, bool)`

GetPrizeDescriptionOk returns a tuple with the PrizeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrizeDescription

`func (o *GiveawayWinners) SetPrizeDescription(v string)`

SetPrizeDescription sets PrizeDescription field to given value.

### HasPrizeDescription

`func (o *GiveawayWinners) HasPrizeDescription() bool`

HasPrizeDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


