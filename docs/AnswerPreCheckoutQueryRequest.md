# AnswerPreCheckoutQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreCheckoutQueryId** | **string** | Unique identifier for the query to be answered | 
**Ok** | **bool** | Specify *True* if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order. Use *False* if there are any problems. | 
**ErrorMessage** | Pointer to **string** | Required if *ok* is *False*. Error message in human readable form that explains the reason for failure to proceed with the checkout (e.g. \&quot;Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details. Please choose a different color or garment!\&quot;). Telegram will display this message to the user. | [optional] 

## Methods

### NewAnswerPreCheckoutQueryRequest

`func NewAnswerPreCheckoutQueryRequest(preCheckoutQueryId string, ok bool, ) *AnswerPreCheckoutQueryRequest`

NewAnswerPreCheckoutQueryRequest instantiates a new AnswerPreCheckoutQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerPreCheckoutQueryRequestWithDefaults

`func NewAnswerPreCheckoutQueryRequestWithDefaults() *AnswerPreCheckoutQueryRequest`

NewAnswerPreCheckoutQueryRequestWithDefaults instantiates a new AnswerPreCheckoutQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreCheckoutQueryId

`func (o *AnswerPreCheckoutQueryRequest) GetPreCheckoutQueryId() string`

GetPreCheckoutQueryId returns the PreCheckoutQueryId field if non-nil, zero value otherwise.

### GetPreCheckoutQueryIdOk

`func (o *AnswerPreCheckoutQueryRequest) GetPreCheckoutQueryIdOk() (*string, bool)`

GetPreCheckoutQueryIdOk returns a tuple with the PreCheckoutQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreCheckoutQueryId

`func (o *AnswerPreCheckoutQueryRequest) SetPreCheckoutQueryId(v string)`

SetPreCheckoutQueryId sets PreCheckoutQueryId field to given value.


### GetOk

`func (o *AnswerPreCheckoutQueryRequest) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *AnswerPreCheckoutQueryRequest) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *AnswerPreCheckoutQueryRequest) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetErrorMessage

`func (o *AnswerPreCheckoutQueryRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AnswerPreCheckoutQueryRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AnswerPreCheckoutQueryRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AnswerPreCheckoutQueryRequest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


