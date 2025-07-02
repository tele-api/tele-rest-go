# RefundStarPaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | Identifier of the user whose payment will be refunded | 
**TelegramPaymentChargeId** | **string** | Telegram payment identifier | 

## Methods

### NewRefundStarPaymentRequest

`func NewRefundStarPaymentRequest(userId int32, telegramPaymentChargeId string, ) *RefundStarPaymentRequest`

NewRefundStarPaymentRequest instantiates a new RefundStarPaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundStarPaymentRequestWithDefaults

`func NewRefundStarPaymentRequestWithDefaults() *RefundStarPaymentRequest`

NewRefundStarPaymentRequestWithDefaults instantiates a new RefundStarPaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *RefundStarPaymentRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RefundStarPaymentRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RefundStarPaymentRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetTelegramPaymentChargeId

`func (o *RefundStarPaymentRequest) GetTelegramPaymentChargeId() string`

GetTelegramPaymentChargeId returns the TelegramPaymentChargeId field if non-nil, zero value otherwise.

### GetTelegramPaymentChargeIdOk

`func (o *RefundStarPaymentRequest) GetTelegramPaymentChargeIdOk() (*string, bool)`

GetTelegramPaymentChargeIdOk returns a tuple with the TelegramPaymentChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramPaymentChargeId

`func (o *RefundStarPaymentRequest) SetTelegramPaymentChargeId(v string)`

SetTelegramPaymentChargeId sets TelegramPaymentChargeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


