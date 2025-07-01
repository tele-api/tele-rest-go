# Dice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emoji** | **string** | Emoji on which the dice throw animation is based | 
**Value** | **int32** | Value of the dice, 1-6 for “🎲”, “🎯” and “🎳” base emoji, 1-5 for “🏀” and “⚽” base emoji, 1-64 for “🎰” base emoji | 

## Methods

### NewDice

`func NewDice(emoji string, value int32, ) *Dice`

NewDice instantiates a new Dice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiceWithDefaults

`func NewDiceWithDefaults() *Dice`

NewDiceWithDefaults instantiates a new Dice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmoji

`func (o *Dice) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *Dice) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *Dice) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.


### GetValue

`func (o *Dice) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Dice) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Dice) SetValue(v int32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


