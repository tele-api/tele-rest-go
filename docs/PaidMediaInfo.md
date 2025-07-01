# PaidMediaInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StarCount** | **int32** | The number of Telegram Stars that must be paid to buy access to the media | 
**PaidMedia** | [**[]PaidMedia**](PaidMedia.md) | Information about the paid media | 

## Methods

### NewPaidMediaInfo

`func NewPaidMediaInfo(starCount int32, paidMedia []PaidMedia, ) *PaidMediaInfo`

NewPaidMediaInfo instantiates a new PaidMediaInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaidMediaInfoWithDefaults

`func NewPaidMediaInfoWithDefaults() *PaidMediaInfo`

NewPaidMediaInfoWithDefaults instantiates a new PaidMediaInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStarCount

`func (o *PaidMediaInfo) GetStarCount() int32`

GetStarCount returns the StarCount field if non-nil, zero value otherwise.

### GetStarCountOk

`func (o *PaidMediaInfo) GetStarCountOk() (*int32, bool)`

GetStarCountOk returns a tuple with the StarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarCount

`func (o *PaidMediaInfo) SetStarCount(v int32)`

SetStarCount sets StarCount field to given value.


### GetPaidMedia

`func (o *PaidMediaInfo) GetPaidMedia() []PaidMedia`

GetPaidMedia returns the PaidMedia field if non-nil, zero value otherwise.

### GetPaidMediaOk

`func (o *PaidMediaInfo) GetPaidMediaOk() (*[]PaidMedia, bool)`

GetPaidMediaOk returns a tuple with the PaidMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidMedia

`func (o *PaidMediaInfo) SetPaidMedia(v []PaidMedia)`

SetPaidMedia sets PaidMedia field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


