# UniqueGiftInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gift** | [**UniqueGift**](UniqueGift.md) |  | 
**Origin** | **string** | Origin of the gift. Currently, either “upgrade” or “transfer” | 
**OwnedGiftId** | Pointer to **string** | *Optional*. Unique identifier of the received gift for the bot; only present for gifts received on behalf of business accounts | [optional] 
**TransferStarCount** | Pointer to **int32** | *Optional*. Number of Telegram Stars that must be paid to transfer the gift; omitted if the bot cannot transfer the gift | [optional] 

## Methods

### NewUniqueGiftInfo

`func NewUniqueGiftInfo(gift UniqueGift, origin string, ) *UniqueGiftInfo`

NewUniqueGiftInfo instantiates a new UniqueGiftInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniqueGiftInfoWithDefaults

`func NewUniqueGiftInfoWithDefaults() *UniqueGiftInfo`

NewUniqueGiftInfoWithDefaults instantiates a new UniqueGiftInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGift

`func (o *UniqueGiftInfo) GetGift() UniqueGift`

GetGift returns the Gift field if non-nil, zero value otherwise.

### GetGiftOk

`func (o *UniqueGiftInfo) GetGiftOk() (*UniqueGift, bool)`

GetGiftOk returns a tuple with the Gift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGift

`func (o *UniqueGiftInfo) SetGift(v UniqueGift)`

SetGift sets Gift field to given value.


### GetOrigin

`func (o *UniqueGiftInfo) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *UniqueGiftInfo) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *UniqueGiftInfo) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetOwnedGiftId

`func (o *UniqueGiftInfo) GetOwnedGiftId() string`

GetOwnedGiftId returns the OwnedGiftId field if non-nil, zero value otherwise.

### GetOwnedGiftIdOk

`func (o *UniqueGiftInfo) GetOwnedGiftIdOk() (*string, bool)`

GetOwnedGiftIdOk returns a tuple with the OwnedGiftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedGiftId

`func (o *UniqueGiftInfo) SetOwnedGiftId(v string)`

SetOwnedGiftId sets OwnedGiftId field to given value.

### HasOwnedGiftId

`func (o *UniqueGiftInfo) HasOwnedGiftId() bool`

HasOwnedGiftId returns a boolean if a field has been set.

### GetTransferStarCount

`func (o *UniqueGiftInfo) GetTransferStarCount() int32`

GetTransferStarCount returns the TransferStarCount field if non-nil, zero value otherwise.

### GetTransferStarCountOk

`func (o *UniqueGiftInfo) GetTransferStarCountOk() (*int32, bool)`

GetTransferStarCountOk returns a tuple with the TransferStarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferStarCount

`func (o *UniqueGiftInfo) SetTransferStarCount(v int32)`

SetTransferStarCount sets TransferStarCount field to given value.

### HasTransferStarCount

`func (o *UniqueGiftInfo) HasTransferStarCount() bool`

HasTransferStarCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


