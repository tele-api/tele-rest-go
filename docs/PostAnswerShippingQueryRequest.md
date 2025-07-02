# PostAnswerShippingQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingQueryId** | **string** | Unique identifier for the query to be answered | 
**Ok** | **bool** | Pass *True* if delivery to the specified address is possible and *False* if there are any problems (for example, if delivery to the specified address is not possible) | 
**ShippingOptions** | Pointer to [**[]ShippingOption**](ShippingOption.md) | Required if *ok* is *True*. A JSON-serialized array of available shipping options. | [optional] 
**ErrorMessage** | Pointer to **string** | Required if *ok* is *False*. Error message in human readable form that explains why it is impossible to complete the order (e.g. “Sorry, delivery to your desired address is unavailable”). Telegram will display this message to the user. | [optional] 

## Methods

### NewPostAnswerShippingQueryRequest

`func NewPostAnswerShippingQueryRequest(shippingQueryId string, ok bool, ) *PostAnswerShippingQueryRequest`

NewPostAnswerShippingQueryRequest instantiates a new PostAnswerShippingQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAnswerShippingQueryRequestWithDefaults

`func NewPostAnswerShippingQueryRequestWithDefaults() *PostAnswerShippingQueryRequest`

NewPostAnswerShippingQueryRequestWithDefaults instantiates a new PostAnswerShippingQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShippingQueryId

`func (o *PostAnswerShippingQueryRequest) GetShippingQueryId() string`

GetShippingQueryId returns the ShippingQueryId field if non-nil, zero value otherwise.

### GetShippingQueryIdOk

`func (o *PostAnswerShippingQueryRequest) GetShippingQueryIdOk() (*string, bool)`

GetShippingQueryIdOk returns a tuple with the ShippingQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingQueryId

`func (o *PostAnswerShippingQueryRequest) SetShippingQueryId(v string)`

SetShippingQueryId sets ShippingQueryId field to given value.


### GetOk

`func (o *PostAnswerShippingQueryRequest) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *PostAnswerShippingQueryRequest) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *PostAnswerShippingQueryRequest) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetShippingOptions

`func (o *PostAnswerShippingQueryRequest) GetShippingOptions() []ShippingOption`

GetShippingOptions returns the ShippingOptions field if non-nil, zero value otherwise.

### GetShippingOptionsOk

`func (o *PostAnswerShippingQueryRequest) GetShippingOptionsOk() (*[]ShippingOption, bool)`

GetShippingOptionsOk returns a tuple with the ShippingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingOptions

`func (o *PostAnswerShippingQueryRequest) SetShippingOptions(v []ShippingOption)`

SetShippingOptions sets ShippingOptions field to given value.

### HasShippingOptions

`func (o *PostAnswerShippingQueryRequest) HasShippingOptions() bool`

HasShippingOptions returns a boolean if a field has been set.

### GetErrorMessage

`func (o *PostAnswerShippingQueryRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PostAnswerShippingQueryRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PostAnswerShippingQueryRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PostAnswerShippingQueryRequest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


