# GiveawayCompleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WinnerCount** | **int32** | Number of winners in the giveaway | 
**UnclaimedPrizeCount** | Pointer to **int32** | *Optional*. Number of undistributed prizes | [optional] 
**GiveawayMessage** | Pointer to [**Message**](Message.md) |  | [optional] 
**IsStarGiveaway** | Pointer to **bool** | *Optional*. *True*, if the giveaway is a Telegram Star giveaway. Otherwise, currently, the giveaway is a Telegram Premium giveaway. | [optional] [default to true]

## Methods

### NewGiveawayCompleted

`func NewGiveawayCompleted(winnerCount int32, ) *GiveawayCompleted`

NewGiveawayCompleted instantiates a new GiveawayCompleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiveawayCompletedWithDefaults

`func NewGiveawayCompletedWithDefaults() *GiveawayCompleted`

NewGiveawayCompletedWithDefaults instantiates a new GiveawayCompleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWinnerCount

`func (o *GiveawayCompleted) GetWinnerCount() int32`

GetWinnerCount returns the WinnerCount field if non-nil, zero value otherwise.

### GetWinnerCountOk

`func (o *GiveawayCompleted) GetWinnerCountOk() (*int32, bool)`

GetWinnerCountOk returns a tuple with the WinnerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnerCount

`func (o *GiveawayCompleted) SetWinnerCount(v int32)`

SetWinnerCount sets WinnerCount field to given value.


### GetUnclaimedPrizeCount

`func (o *GiveawayCompleted) GetUnclaimedPrizeCount() int32`

GetUnclaimedPrizeCount returns the UnclaimedPrizeCount field if non-nil, zero value otherwise.

### GetUnclaimedPrizeCountOk

`func (o *GiveawayCompleted) GetUnclaimedPrizeCountOk() (*int32, bool)`

GetUnclaimedPrizeCountOk returns a tuple with the UnclaimedPrizeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnclaimedPrizeCount

`func (o *GiveawayCompleted) SetUnclaimedPrizeCount(v int32)`

SetUnclaimedPrizeCount sets UnclaimedPrizeCount field to given value.

### HasUnclaimedPrizeCount

`func (o *GiveawayCompleted) HasUnclaimedPrizeCount() bool`

HasUnclaimedPrizeCount returns a boolean if a field has been set.

### GetGiveawayMessage

`func (o *GiveawayCompleted) GetGiveawayMessage() Message`

GetGiveawayMessage returns the GiveawayMessage field if non-nil, zero value otherwise.

### GetGiveawayMessageOk

`func (o *GiveawayCompleted) GetGiveawayMessageOk() (*Message, bool)`

GetGiveawayMessageOk returns a tuple with the GiveawayMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiveawayMessage

`func (o *GiveawayCompleted) SetGiveawayMessage(v Message)`

SetGiveawayMessage sets GiveawayMessage field to given value.

### HasGiveawayMessage

`func (o *GiveawayCompleted) HasGiveawayMessage() bool`

HasGiveawayMessage returns a boolean if a field has been set.

### GetIsStarGiveaway

`func (o *GiveawayCompleted) GetIsStarGiveaway() bool`

GetIsStarGiveaway returns the IsStarGiveaway field if non-nil, zero value otherwise.

### GetIsStarGiveawayOk

`func (o *GiveawayCompleted) GetIsStarGiveawayOk() (*bool, bool)`

GetIsStarGiveawayOk returns a tuple with the IsStarGiveaway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStarGiveaway

`func (o *GiveawayCompleted) SetIsStarGiveaway(v bool)`

SetIsStarGiveaway sets IsStarGiveaway field to given value.

### HasIsStarGiveaway

`func (o *GiveawayCompleted) HasIsStarGiveaway() bool`

HasIsStarGiveaway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


