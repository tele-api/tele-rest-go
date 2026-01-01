# SuggestedPostParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to [**SuggestedPostPrice**](SuggestedPostPrice.md) |  | [optional] 
**SendDate** | Pointer to **int32** | *Optional*. Proposed send date of the post. If specified, then the date must be between 300 second and 2678400 seconds (30 days) in the future. If the field is omitted, then the post can be published at any time within 30 days at the sole discretion of the user who approves it. | [optional] 

## Methods

### NewSuggestedPostParameters

`func NewSuggestedPostParameters() *SuggestedPostParameters`

NewSuggestedPostParameters instantiates a new SuggestedPostParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestedPostParametersWithDefaults

`func NewSuggestedPostParametersWithDefaults() *SuggestedPostParameters`

NewSuggestedPostParametersWithDefaults instantiates a new SuggestedPostParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *SuggestedPostParameters) GetPrice() SuggestedPostPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SuggestedPostParameters) GetPriceOk() (*SuggestedPostPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SuggestedPostParameters) SetPrice(v SuggestedPostPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SuggestedPostParameters) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSendDate

`func (o *SuggestedPostParameters) GetSendDate() int32`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *SuggestedPostParameters) GetSendDateOk() (*int32, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *SuggestedPostParameters) SetSendDate(v int32)`

SetSendDate sets SendDate field to given value.

### HasSendDate

`func (o *SuggestedPostParameters) HasSendDate() bool`

HasSendDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


