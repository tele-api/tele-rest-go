# UniqueGiftInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gift** | [**UniqueGift**](UniqueGift.md) |  | 
**Origin** | **string** | Origin of the gift. Currently, either “upgrade” for gifts upgraded from regular gifts, “transfer” for gifts transferred from other users or channels, or “resale” for gifts bought from other users | 
**LastResaleStarCount** | Pointer to **int32** | *Optional*. For gifts bought from other users, the price paid for the gift | [optional] 
**OwnedGiftId** | Pointer to **string** | *Optional*. Unique identifier of the received gift for the bot; only present for gifts received on behalf of business accounts | [optional] 
**TransferStarCount** | Pointer to **int32** | *Optional*. Number of Telegram Stars that must be paid to transfer the gift; omitted if the bot cannot transfer the gift | [optional] 
**NextTransferDate** | Pointer to **int32** | *Optional*. Point in time (Unix timestamp) when the gift can be transferred. If it is in the past, then the gift can be transferred now | [optional] 

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


### GetLastResaleStarCount

`func (o *UniqueGiftInfo) GetLastResaleStarCount() int32`

GetLastResaleStarCount returns the LastResaleStarCount field if non-nil, zero value otherwise.

### GetLastResaleStarCountOk

`func (o *UniqueGiftInfo) GetLastResaleStarCountOk() (*int32, bool)`

GetLastResaleStarCountOk returns a tuple with the LastResaleStarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResaleStarCount

`func (o *UniqueGiftInfo) SetLastResaleStarCount(v int32)`

SetLastResaleStarCount sets LastResaleStarCount field to given value.

### HasLastResaleStarCount

`func (o *UniqueGiftInfo) HasLastResaleStarCount() bool`

HasLastResaleStarCount returns a boolean if a field has been set.

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

### GetNextTransferDate

`func (o *UniqueGiftInfo) GetNextTransferDate() int32`

GetNextTransferDate returns the NextTransferDate field if non-nil, zero value otherwise.

### GetNextTransferDateOk

`func (o *UniqueGiftInfo) GetNextTransferDateOk() (*int32, bool)`

GetNextTransferDateOk returns a tuple with the NextTransferDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTransferDate

`func (o *UniqueGiftInfo) SetNextTransferDate(v int32)`

SetNextTransferDate sets NextTransferDate field to given value.

### HasNextTransferDate

`func (o *UniqueGiftInfo) HasNextTransferDate() bool`

HasNextTransferDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


