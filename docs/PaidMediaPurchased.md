# PaidMediaPurchased

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**User**](User.md) |  | 
**PaidMediaPayload** | **string** | Bot-specified paid media payload | 

## Methods

### NewPaidMediaPurchased

`func NewPaidMediaPurchased(from User, paidMediaPayload string, ) *PaidMediaPurchased`

NewPaidMediaPurchased instantiates a new PaidMediaPurchased object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaidMediaPurchasedWithDefaults

`func NewPaidMediaPurchasedWithDefaults() *PaidMediaPurchased`

NewPaidMediaPurchasedWithDefaults instantiates a new PaidMediaPurchased object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *PaidMediaPurchased) GetFrom() User`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PaidMediaPurchased) GetFromOk() (*User, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PaidMediaPurchased) SetFrom(v User)`

SetFrom sets From field to given value.


### GetPaidMediaPayload

`func (o *PaidMediaPurchased) GetPaidMediaPayload() string`

GetPaidMediaPayload returns the PaidMediaPayload field if non-nil, zero value otherwise.

### GetPaidMediaPayloadOk

`func (o *PaidMediaPurchased) GetPaidMediaPayloadOk() (*string, bool)`

GetPaidMediaPayloadOk returns a tuple with the PaidMediaPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidMediaPayload

`func (o *PaidMediaPurchased) SetPaidMediaPayload(v string)`

SetPaidMediaPayload sets PaidMediaPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


