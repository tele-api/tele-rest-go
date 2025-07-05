# AnswerShippingQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingQueryId** | **string** | Unique identifier for the query to be answered | 
**Ok** | **bool** | Pass *True* if delivery to the specified address is possible and *False* if there are any problems (for example, if delivery to the specified address is not possible) | 
**ShippingOptions** | Pointer to [**[]ShippingOption**](ShippingOption.md) | Required if *ok* is *True*. A JSON-serialized array of available shipping options. | [optional] 
**ErrorMessage** | Pointer to **string** | Required if *ok* is *False*. Error message in human readable form that explains why it is impossible to complete the order (e.g. “Sorry, delivery to your desired address is unavailable”). Telegram will display this message to the user. | [optional] 

## Methods

### NewAnswerShippingQueryRequest

`func NewAnswerShippingQueryRequest(shippingQueryId string, ok bool, ) *AnswerShippingQueryRequest`

NewAnswerShippingQueryRequest instantiates a new AnswerShippingQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerShippingQueryRequestWithDefaults

`func NewAnswerShippingQueryRequestWithDefaults() *AnswerShippingQueryRequest`

NewAnswerShippingQueryRequestWithDefaults instantiates a new AnswerShippingQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShippingQueryId

`func (o *AnswerShippingQueryRequest) GetShippingQueryId() string`

GetShippingQueryId returns the ShippingQueryId field if non-nil, zero value otherwise.

### GetShippingQueryIdOk

`func (o *AnswerShippingQueryRequest) GetShippingQueryIdOk() (*string, bool)`

GetShippingQueryIdOk returns a tuple with the ShippingQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingQueryId

`func (o *AnswerShippingQueryRequest) SetShippingQueryId(v string)`

SetShippingQueryId sets ShippingQueryId field to given value.


### GetOk

`func (o *AnswerShippingQueryRequest) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *AnswerShippingQueryRequest) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *AnswerShippingQueryRequest) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetShippingOptions

`func (o *AnswerShippingQueryRequest) GetShippingOptions() []ShippingOption`

GetShippingOptions returns the ShippingOptions field if non-nil, zero value otherwise.

### GetShippingOptionsOk

`func (o *AnswerShippingQueryRequest) GetShippingOptionsOk() (*[]ShippingOption, bool)`

GetShippingOptionsOk returns a tuple with the ShippingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingOptions

`func (o *AnswerShippingQueryRequest) SetShippingOptions(v []ShippingOption)`

SetShippingOptions sets ShippingOptions field to given value.

### HasShippingOptions

`func (o *AnswerShippingQueryRequest) HasShippingOptions() bool`

HasShippingOptions returns a boolean if a field has been set.

### GetErrorMessage

`func (o *AnswerShippingQueryRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AnswerShippingQueryRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AnswerShippingQueryRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AnswerShippingQueryRequest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


