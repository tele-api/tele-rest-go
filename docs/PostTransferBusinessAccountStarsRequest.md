# PostTransferBusinessAccountStarsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection | 
**StarCount** | **int32** | Number of Telegram Stars to transfer; 1-10000 | 

## Methods

### NewPostTransferBusinessAccountStarsRequest

`func NewPostTransferBusinessAccountStarsRequest(businessConnectionId string, starCount int32, ) *PostTransferBusinessAccountStarsRequest`

NewPostTransferBusinessAccountStarsRequest instantiates a new PostTransferBusinessAccountStarsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTransferBusinessAccountStarsRequestWithDefaults

`func NewPostTransferBusinessAccountStarsRequestWithDefaults() *PostTransferBusinessAccountStarsRequest`

NewPostTransferBusinessAccountStarsRequestWithDefaults instantiates a new PostTransferBusinessAccountStarsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *PostTransferBusinessAccountStarsRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *PostTransferBusinessAccountStarsRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *PostTransferBusinessAccountStarsRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetStarCount

`func (o *PostTransferBusinessAccountStarsRequest) GetStarCount() int32`

GetStarCount returns the StarCount field if non-nil, zero value otherwise.

### GetStarCountOk

`func (o *PostTransferBusinessAccountStarsRequest) GetStarCountOk() (*int32, bool)`

GetStarCountOk returns a tuple with the StarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarCount

`func (o *PostTransferBusinessAccountStarsRequest) SetStarCount(v int32)`

SetStarCount sets StarCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


