# ConvertGiftToStarsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection | 
**OwnedGiftId** | **string** | Unique identifier of the regular gift that should be converted to Telegram Stars | 

## Methods

### NewConvertGiftToStarsRequest

`func NewConvertGiftToStarsRequest(businessConnectionId string, ownedGiftId string, ) *ConvertGiftToStarsRequest`

NewConvertGiftToStarsRequest instantiates a new ConvertGiftToStarsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertGiftToStarsRequestWithDefaults

`func NewConvertGiftToStarsRequestWithDefaults() *ConvertGiftToStarsRequest`

NewConvertGiftToStarsRequestWithDefaults instantiates a new ConvertGiftToStarsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *ConvertGiftToStarsRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *ConvertGiftToStarsRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *ConvertGiftToStarsRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetOwnedGiftId

`func (o *ConvertGiftToStarsRequest) GetOwnedGiftId() string`

GetOwnedGiftId returns the OwnedGiftId field if non-nil, zero value otherwise.

### GetOwnedGiftIdOk

`func (o *ConvertGiftToStarsRequest) GetOwnedGiftIdOk() (*string, bool)`

GetOwnedGiftIdOk returns a tuple with the OwnedGiftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedGiftId

`func (o *ConvertGiftToStarsRequest) SetOwnedGiftId(v string)`

SetOwnedGiftId sets OwnedGiftId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


