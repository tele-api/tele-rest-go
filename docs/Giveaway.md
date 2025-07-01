# Giveaway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chats** | [**[]Chat**](Chat.md) | The list of chats which the user must join to participate in the giveaway | 
**WinnersSelectionDate** | **int32** | Point in time (Unix timestamp) when winners of the giveaway will be selected | 
**WinnerCount** | **int32** | The number of users which are supposed to be selected as winners of the giveaway | 
**OnlyNewMembers** | Pointer to **bool** | *Optional*. *True*, if only users who join the chats after the giveaway started should be eligible to win | [optional] [default to true]
**HasPublicWinners** | Pointer to **bool** | *Optional*. *True*, if the list of giveaway winners will be visible to everyone | [optional] [default to true]
**PrizeDescription** | Pointer to **string** | *Optional*. Description of additional giveaway prize | [optional] 
**CountryCodes** | Pointer to **[]string** | *Optional*. A list of two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes indicating the countries from which eligible users for the giveaway must come. If empty, then all users can participate in the giveaway. Users with a phone number that was bought on Fragment can always participate in giveaways. | [optional] 
**PrizeStarCount** | Pointer to **int32** | *Optional*. The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only | [optional] 
**PremiumSubscriptionMonthCount** | Pointer to **int32** | *Optional*. The number of months the Telegram Premium subscription won from the giveaway will be active for; for Telegram Premium giveaways only | [optional] 

## Methods

### NewGiveaway

`func NewGiveaway(chats []Chat, winnersSelectionDate int32, winnerCount int32, ) *Giveaway`

NewGiveaway instantiates a new Giveaway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiveawayWithDefaults

`func NewGiveawayWithDefaults() *Giveaway`

NewGiveawayWithDefaults instantiates a new Giveaway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChats

`func (o *Giveaway) GetChats() []Chat`

GetChats returns the Chats field if non-nil, zero value otherwise.

### GetChatsOk

`func (o *Giveaway) GetChatsOk() (*[]Chat, bool)`

GetChatsOk returns a tuple with the Chats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChats

`func (o *Giveaway) SetChats(v []Chat)`

SetChats sets Chats field to given value.


### GetWinnersSelectionDate

`func (o *Giveaway) GetWinnersSelectionDate() int32`

GetWinnersSelectionDate returns the WinnersSelectionDate field if non-nil, zero value otherwise.

### GetWinnersSelectionDateOk

`func (o *Giveaway) GetWinnersSelectionDateOk() (*int32, bool)`

GetWinnersSelectionDateOk returns a tuple with the WinnersSelectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnersSelectionDate

`func (o *Giveaway) SetWinnersSelectionDate(v int32)`

SetWinnersSelectionDate sets WinnersSelectionDate field to given value.


### GetWinnerCount

`func (o *Giveaway) GetWinnerCount() int32`

GetWinnerCount returns the WinnerCount field if non-nil, zero value otherwise.

### GetWinnerCountOk

`func (o *Giveaway) GetWinnerCountOk() (*int32, bool)`

GetWinnerCountOk returns a tuple with the WinnerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnerCount

`func (o *Giveaway) SetWinnerCount(v int32)`

SetWinnerCount sets WinnerCount field to given value.


### GetOnlyNewMembers

`func (o *Giveaway) GetOnlyNewMembers() bool`

GetOnlyNewMembers returns the OnlyNewMembers field if non-nil, zero value otherwise.

### GetOnlyNewMembersOk

`func (o *Giveaway) GetOnlyNewMembersOk() (*bool, bool)`

GetOnlyNewMembersOk returns a tuple with the OnlyNewMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyNewMembers

`func (o *Giveaway) SetOnlyNewMembers(v bool)`

SetOnlyNewMembers sets OnlyNewMembers field to given value.

### HasOnlyNewMembers

`func (o *Giveaway) HasOnlyNewMembers() bool`

HasOnlyNewMembers returns a boolean if a field has been set.

### GetHasPublicWinners

`func (o *Giveaway) GetHasPublicWinners() bool`

GetHasPublicWinners returns the HasPublicWinners field if non-nil, zero value otherwise.

### GetHasPublicWinnersOk

`func (o *Giveaway) GetHasPublicWinnersOk() (*bool, bool)`

GetHasPublicWinnersOk returns a tuple with the HasPublicWinners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPublicWinners

`func (o *Giveaway) SetHasPublicWinners(v bool)`

SetHasPublicWinners sets HasPublicWinners field to given value.

### HasHasPublicWinners

`func (o *Giveaway) HasHasPublicWinners() bool`

HasHasPublicWinners returns a boolean if a field has been set.

### GetPrizeDescription

`func (o *Giveaway) GetPrizeDescription() string`

GetPrizeDescription returns the PrizeDescription field if non-nil, zero value otherwise.

### GetPrizeDescriptionOk

`func (o *Giveaway) GetPrizeDescriptionOk() (*string, bool)`

GetPrizeDescriptionOk returns a tuple with the PrizeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrizeDescription

`func (o *Giveaway) SetPrizeDescription(v string)`

SetPrizeDescription sets PrizeDescription field to given value.

### HasPrizeDescription

`func (o *Giveaway) HasPrizeDescription() bool`

HasPrizeDescription returns a boolean if a field has been set.

### GetCountryCodes

`func (o *Giveaway) GetCountryCodes() []string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *Giveaway) GetCountryCodesOk() (*[]string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *Giveaway) SetCountryCodes(v []string)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *Giveaway) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### GetPrizeStarCount

`func (o *Giveaway) GetPrizeStarCount() int32`

GetPrizeStarCount returns the PrizeStarCount field if non-nil, zero value otherwise.

### GetPrizeStarCountOk

`func (o *Giveaway) GetPrizeStarCountOk() (*int32, bool)`

GetPrizeStarCountOk returns a tuple with the PrizeStarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrizeStarCount

`func (o *Giveaway) SetPrizeStarCount(v int32)`

SetPrizeStarCount sets PrizeStarCount field to given value.

### HasPrizeStarCount

`func (o *Giveaway) HasPrizeStarCount() bool`

HasPrizeStarCount returns a boolean if a field has been set.

### GetPremiumSubscriptionMonthCount

`func (o *Giveaway) GetPremiumSubscriptionMonthCount() int32`

GetPremiumSubscriptionMonthCount returns the PremiumSubscriptionMonthCount field if non-nil, zero value otherwise.

### GetPremiumSubscriptionMonthCountOk

`func (o *Giveaway) GetPremiumSubscriptionMonthCountOk() (*int32, bool)`

GetPremiumSubscriptionMonthCountOk returns a tuple with the PremiumSubscriptionMonthCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumSubscriptionMonthCount

`func (o *Giveaway) SetPremiumSubscriptionMonthCount(v int32)`

SetPremiumSubscriptionMonthCount sets PremiumSubscriptionMonthCount field to given value.

### HasPremiumSubscriptionMonthCount

`func (o *Giveaway) HasPremiumSubscriptionMonthCount() bool`

HasPremiumSubscriptionMonthCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


