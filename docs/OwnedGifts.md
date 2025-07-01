# OwnedGifts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** | The total number of gifts owned by the user or the chat | 
**Gifts** | [**[]OwnedGift**](OwnedGift.md) | The list of gifts | 
**NextOffset** | Pointer to **string** | *Optional*. Offset for the next request. If empty, then there are no more results | [optional] 

## Methods

### NewOwnedGifts

`func NewOwnedGifts(totalCount int32, gifts []OwnedGift, ) *OwnedGifts`

NewOwnedGifts instantiates a new OwnedGifts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnedGiftsWithDefaults

`func NewOwnedGiftsWithDefaults() *OwnedGifts`

NewOwnedGiftsWithDefaults instantiates a new OwnedGifts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *OwnedGifts) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OwnedGifts) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OwnedGifts) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetGifts

`func (o *OwnedGifts) GetGifts() []OwnedGift`

GetGifts returns the Gifts field if non-nil, zero value otherwise.

### GetGiftsOk

`func (o *OwnedGifts) GetGiftsOk() (*[]OwnedGift, bool)`

GetGiftsOk returns a tuple with the Gifts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGifts

`func (o *OwnedGifts) SetGifts(v []OwnedGift)`

SetGifts sets Gifts field to given value.


### GetNextOffset

`func (o *OwnedGifts) GetNextOffset() string`

GetNextOffset returns the NextOffset field if non-nil, zero value otherwise.

### GetNextOffsetOk

`func (o *OwnedGifts) GetNextOffsetOk() (*string, bool)`

GetNextOffsetOk returns a tuple with the NextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOffset

`func (o *OwnedGifts) SetNextOffset(v string)`

SetNextOffset sets NextOffset field to given value.

### HasNextOffset

`func (o *OwnedGifts) HasNextOffset() bool`

HasNextOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


