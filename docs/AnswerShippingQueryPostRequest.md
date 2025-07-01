# AnswerShippingQueryPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingQueryId** | **string** | Unique identifier for the query to be answered | 
**Ok** | **bool** | Pass *True* if delivery to the specified address is possible and *False* if there are any problems (for example, if delivery to the specified address is not possible) | 
**ShippingOptions** | Pointer to [**[]ShippingOption**](ShippingOption.md) | Required if *ok* is *True*. A JSON-serialized array of available shipping options. | [optional] 
**ErrorMessage** | Pointer to **string** | Required if *ok* is *False*. Error message in human readable form that explains why it is impossible to complete the order (e.g. “Sorry, delivery to your desired address is unavailable”). Telegram will display this message to the user. | [optional] 

## Methods

### NewAnswerShippingQueryPostRequest

`func NewAnswerShippingQueryPostRequest(shippingQueryId string, ok bool, ) *AnswerShippingQueryPostRequest`

NewAnswerShippingQueryPostRequest instantiates a new AnswerShippingQueryPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerShippingQueryPostRequestWithDefaults

`func NewAnswerShippingQueryPostRequestWithDefaults() *AnswerShippingQueryPostRequest`

NewAnswerShippingQueryPostRequestWithDefaults instantiates a new AnswerShippingQueryPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShippingQueryId

`func (o *AnswerShippingQueryPostRequest) GetShippingQueryId() string`

GetShippingQueryId returns the ShippingQueryId field if non-nil, zero value otherwise.

### GetShippingQueryIdOk

`func (o *AnswerShippingQueryPostRequest) GetShippingQueryIdOk() (*string, bool)`

GetShippingQueryIdOk returns a tuple with the ShippingQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingQueryId

`func (o *AnswerShippingQueryPostRequest) SetShippingQueryId(v string)`

SetShippingQueryId sets ShippingQueryId field to given value.


### GetOk

`func (o *AnswerShippingQueryPostRequest) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *AnswerShippingQueryPostRequest) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *AnswerShippingQueryPostRequest) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetShippingOptions

`func (o *AnswerShippingQueryPostRequest) GetShippingOptions() []ShippingOption`

GetShippingOptions returns the ShippingOptions field if non-nil, zero value otherwise.

### GetShippingOptionsOk

`func (o *AnswerShippingQueryPostRequest) GetShippingOptionsOk() (*[]ShippingOption, bool)`

GetShippingOptionsOk returns a tuple with the ShippingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingOptions

`func (o *AnswerShippingQueryPostRequest) SetShippingOptions(v []ShippingOption)`

SetShippingOptions sets ShippingOptions field to given value.

### HasShippingOptions

`func (o *AnswerShippingQueryPostRequest) HasShippingOptions() bool`

HasShippingOptions returns a boolean if a field has been set.

### GetErrorMessage

`func (o *AnswerShippingQueryPostRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AnswerShippingQueryPostRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AnswerShippingQueryPostRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AnswerShippingQueryPostRequest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


