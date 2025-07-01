# AnswerPreCheckoutQueryPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreCheckoutQueryId** | **string** | Unique identifier for the query to be answered | 
**Ok** | **bool** | Specify *True* if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order. Use *False* if there are any problems. | 
**ErrorMessage** | Pointer to **string** | Required if *ok* is *False*. Error message in human readable form that explains the reason for failure to proceed with the checkout (e.g. \&quot;Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details. Please choose a different color or garment!\&quot;). Telegram will display this message to the user. | [optional] 

## Methods

### NewAnswerPreCheckoutQueryPostRequest

`func NewAnswerPreCheckoutQueryPostRequest(preCheckoutQueryId string, ok bool, ) *AnswerPreCheckoutQueryPostRequest`

NewAnswerPreCheckoutQueryPostRequest instantiates a new AnswerPreCheckoutQueryPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerPreCheckoutQueryPostRequestWithDefaults

`func NewAnswerPreCheckoutQueryPostRequestWithDefaults() *AnswerPreCheckoutQueryPostRequest`

NewAnswerPreCheckoutQueryPostRequestWithDefaults instantiates a new AnswerPreCheckoutQueryPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreCheckoutQueryId

`func (o *AnswerPreCheckoutQueryPostRequest) GetPreCheckoutQueryId() string`

GetPreCheckoutQueryId returns the PreCheckoutQueryId field if non-nil, zero value otherwise.

### GetPreCheckoutQueryIdOk

`func (o *AnswerPreCheckoutQueryPostRequest) GetPreCheckoutQueryIdOk() (*string, bool)`

GetPreCheckoutQueryIdOk returns a tuple with the PreCheckoutQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreCheckoutQueryId

`func (o *AnswerPreCheckoutQueryPostRequest) SetPreCheckoutQueryId(v string)`

SetPreCheckoutQueryId sets PreCheckoutQueryId field to given value.


### GetOk

`func (o *AnswerPreCheckoutQueryPostRequest) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *AnswerPreCheckoutQueryPostRequest) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *AnswerPreCheckoutQueryPostRequest) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetErrorMessage

`func (o *AnswerPreCheckoutQueryPostRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AnswerPreCheckoutQueryPostRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AnswerPreCheckoutQueryPostRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AnswerPreCheckoutQueryPostRequest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


