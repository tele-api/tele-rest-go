# ChatBoost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoostId** | **string** | Unique identifier of the boost | 
**AddDate** | **int32** | Point in time (Unix timestamp) when the chat was boosted | 
**ExpirationDate** | **int32** | Point in time (Unix timestamp) when the boost will automatically expire, unless the booster&#39;s Telegram Premium subscription is prolonged | 
**Source** | [**ChatBoostSource**](ChatBoostSource.md) |  | 

## Methods

### NewChatBoost

`func NewChatBoost(boostId string, addDate int32, expirationDate int32, source ChatBoostSource, ) *ChatBoost`

NewChatBoost instantiates a new ChatBoost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatBoostWithDefaults

`func NewChatBoostWithDefaults() *ChatBoost`

NewChatBoostWithDefaults instantiates a new ChatBoost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoostId

`func (o *ChatBoost) GetBoostId() string`

GetBoostId returns the BoostId field if non-nil, zero value otherwise.

### GetBoostIdOk

`func (o *ChatBoost) GetBoostIdOk() (*string, bool)`

GetBoostIdOk returns a tuple with the BoostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoostId

`func (o *ChatBoost) SetBoostId(v string)`

SetBoostId sets BoostId field to given value.


### GetAddDate

`func (o *ChatBoost) GetAddDate() int32`

GetAddDate returns the AddDate field if non-nil, zero value otherwise.

### GetAddDateOk

`func (o *ChatBoost) GetAddDateOk() (*int32, bool)`

GetAddDateOk returns a tuple with the AddDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddDate

`func (o *ChatBoost) SetAddDate(v int32)`

SetAddDate sets AddDate field to given value.


### GetExpirationDate

`func (o *ChatBoost) GetExpirationDate() int32`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ChatBoost) GetExpirationDateOk() (*int32, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ChatBoost) SetExpirationDate(v int32)`

SetExpirationDate sets ExpirationDate field to given value.


### GetSource

`func (o *ChatBoost) GetSource() ChatBoostSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ChatBoost) GetSourceOk() (*ChatBoostSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ChatBoost) SetSource(v ChatBoostSource)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


