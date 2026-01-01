# SuggestedPostRefunded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuggestedPostMessage** | Pointer to [**Message**](Message.md) |  | [optional] 
**Reason** | **string** | Reason for the refund. Currently, one of “post\\_deleted” if the post was deleted within 24 hours of being posted or removed from scheduled messages without being posted, or “payment\\_refunded” if the payer refunded their payment. | 

## Methods

### NewSuggestedPostRefunded

`func NewSuggestedPostRefunded(reason string, ) *SuggestedPostRefunded`

NewSuggestedPostRefunded instantiates a new SuggestedPostRefunded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestedPostRefundedWithDefaults

`func NewSuggestedPostRefundedWithDefaults() *SuggestedPostRefunded`

NewSuggestedPostRefundedWithDefaults instantiates a new SuggestedPostRefunded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggestedPostMessage

`func (o *SuggestedPostRefunded) GetSuggestedPostMessage() Message`

GetSuggestedPostMessage returns the SuggestedPostMessage field if non-nil, zero value otherwise.

### GetSuggestedPostMessageOk

`func (o *SuggestedPostRefunded) GetSuggestedPostMessageOk() (*Message, bool)`

GetSuggestedPostMessageOk returns a tuple with the SuggestedPostMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedPostMessage

`func (o *SuggestedPostRefunded) SetSuggestedPostMessage(v Message)`

SetSuggestedPostMessage sets SuggestedPostMessage field to given value.

### HasSuggestedPostMessage

`func (o *SuggestedPostRefunded) HasSuggestedPostMessage() bool`

HasSuggestedPostMessage returns a boolean if a field has been set.

### GetReason

`func (o *SuggestedPostRefunded) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SuggestedPostRefunded) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SuggestedPostRefunded) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


