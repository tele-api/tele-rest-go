# DirectMessagePriceChanged

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreDirectMessagesEnabled** | **bool** | *True*, if direct messages are enabled for the channel chat; false otherwise | 
**DirectMessageStarCount** | Pointer to **int32** | *Optional*. The new number of Telegram Stars that must be paid by users for each direct message sent to the channel. Does not apply to users who have been exempted by administrators. Defaults to 0. | [optional] [default to 0]

## Methods

### NewDirectMessagePriceChanged

`func NewDirectMessagePriceChanged(areDirectMessagesEnabled bool, ) *DirectMessagePriceChanged`

NewDirectMessagePriceChanged instantiates a new DirectMessagePriceChanged object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectMessagePriceChangedWithDefaults

`func NewDirectMessagePriceChangedWithDefaults() *DirectMessagePriceChanged`

NewDirectMessagePriceChangedWithDefaults instantiates a new DirectMessagePriceChanged object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreDirectMessagesEnabled

`func (o *DirectMessagePriceChanged) GetAreDirectMessagesEnabled() bool`

GetAreDirectMessagesEnabled returns the AreDirectMessagesEnabled field if non-nil, zero value otherwise.

### GetAreDirectMessagesEnabledOk

`func (o *DirectMessagePriceChanged) GetAreDirectMessagesEnabledOk() (*bool, bool)`

GetAreDirectMessagesEnabledOk returns a tuple with the AreDirectMessagesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreDirectMessagesEnabled

`func (o *DirectMessagePriceChanged) SetAreDirectMessagesEnabled(v bool)`

SetAreDirectMessagesEnabled sets AreDirectMessagesEnabled field to given value.


### GetDirectMessageStarCount

`func (o *DirectMessagePriceChanged) GetDirectMessageStarCount() int32`

GetDirectMessageStarCount returns the DirectMessageStarCount field if non-nil, zero value otherwise.

### GetDirectMessageStarCountOk

`func (o *DirectMessagePriceChanged) GetDirectMessageStarCountOk() (*int32, bool)`

GetDirectMessageStarCountOk returns a tuple with the DirectMessageStarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectMessageStarCount

`func (o *DirectMessagePriceChanged) SetDirectMessageStarCount(v int32)`

SetDirectMessageStarCount sets DirectMessageStarCount field to given value.

### HasDirectMessageStarCount

`func (o *DirectMessagePriceChanged) HasDirectMessageStarCount() bool`

HasDirectMessageStarCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


