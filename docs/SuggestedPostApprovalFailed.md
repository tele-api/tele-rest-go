# SuggestedPostApprovalFailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuggestedPostMessage** | Pointer to [**Message**](Message.md) |  | [optional] 
**Price** | [**SuggestedPostPrice**](SuggestedPostPrice.md) |  | 

## Methods

### NewSuggestedPostApprovalFailed

`func NewSuggestedPostApprovalFailed(price SuggestedPostPrice, ) *SuggestedPostApprovalFailed`

NewSuggestedPostApprovalFailed instantiates a new SuggestedPostApprovalFailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestedPostApprovalFailedWithDefaults

`func NewSuggestedPostApprovalFailedWithDefaults() *SuggestedPostApprovalFailed`

NewSuggestedPostApprovalFailedWithDefaults instantiates a new SuggestedPostApprovalFailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggestedPostMessage

`func (o *SuggestedPostApprovalFailed) GetSuggestedPostMessage() Message`

GetSuggestedPostMessage returns the SuggestedPostMessage field if non-nil, zero value otherwise.

### GetSuggestedPostMessageOk

`func (o *SuggestedPostApprovalFailed) GetSuggestedPostMessageOk() (*Message, bool)`

GetSuggestedPostMessageOk returns a tuple with the SuggestedPostMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedPostMessage

`func (o *SuggestedPostApprovalFailed) SetSuggestedPostMessage(v Message)`

SetSuggestedPostMessage sets SuggestedPostMessage field to given value.

### HasSuggestedPostMessage

`func (o *SuggestedPostApprovalFailed) HasSuggestedPostMessage() bool`

HasSuggestedPostMessage returns a boolean if a field has been set.

### GetPrice

`func (o *SuggestedPostApprovalFailed) GetPrice() SuggestedPostPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SuggestedPostApprovalFailed) GetPriceOk() (*SuggestedPostPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SuggestedPostApprovalFailed) SetPrice(v SuggestedPostPrice)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


