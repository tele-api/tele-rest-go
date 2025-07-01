# RefundStarPaymentPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | Identifier of the user whose payment will be refunded | 
**TelegramPaymentChargeId** | **string** | Telegram payment identifier | 

## Methods

### NewRefundStarPaymentPostRequest

`func NewRefundStarPaymentPostRequest(userId int32, telegramPaymentChargeId string, ) *RefundStarPaymentPostRequest`

NewRefundStarPaymentPostRequest instantiates a new RefundStarPaymentPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundStarPaymentPostRequestWithDefaults

`func NewRefundStarPaymentPostRequestWithDefaults() *RefundStarPaymentPostRequest`

NewRefundStarPaymentPostRequestWithDefaults instantiates a new RefundStarPaymentPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *RefundStarPaymentPostRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RefundStarPaymentPostRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RefundStarPaymentPostRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetTelegramPaymentChargeId

`func (o *RefundStarPaymentPostRequest) GetTelegramPaymentChargeId() string`

GetTelegramPaymentChargeId returns the TelegramPaymentChargeId field if non-nil, zero value otherwise.

### GetTelegramPaymentChargeIdOk

`func (o *RefundStarPaymentPostRequest) GetTelegramPaymentChargeIdOk() (*string, bool)`

GetTelegramPaymentChargeIdOk returns a tuple with the TelegramPaymentChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramPaymentChargeId

`func (o *RefundStarPaymentPostRequest) SetTelegramPaymentChargeId(v string)`

SetTelegramPaymentChargeId sets TelegramPaymentChargeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


