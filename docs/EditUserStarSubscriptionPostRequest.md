# EditUserStarSubscriptionPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | Identifier of the user whose subscription will be edited | 
**TelegramPaymentChargeId** | **string** | Telegram payment identifier for the subscription | 
**IsCanceled** | **bool** | Pass *True* to cancel extension of the user subscription; the subscription must be active up to the end of the current subscription period. Pass *False* to allow the user to re-enable a subscription that was previously canceled by the bot. | 

## Methods

### NewEditUserStarSubscriptionPostRequest

`func NewEditUserStarSubscriptionPostRequest(userId int32, telegramPaymentChargeId string, isCanceled bool, ) *EditUserStarSubscriptionPostRequest`

NewEditUserStarSubscriptionPostRequest instantiates a new EditUserStarSubscriptionPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditUserStarSubscriptionPostRequestWithDefaults

`func NewEditUserStarSubscriptionPostRequestWithDefaults() *EditUserStarSubscriptionPostRequest`

NewEditUserStarSubscriptionPostRequestWithDefaults instantiates a new EditUserStarSubscriptionPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *EditUserStarSubscriptionPostRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EditUserStarSubscriptionPostRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EditUserStarSubscriptionPostRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetTelegramPaymentChargeId

`func (o *EditUserStarSubscriptionPostRequest) GetTelegramPaymentChargeId() string`

GetTelegramPaymentChargeId returns the TelegramPaymentChargeId field if non-nil, zero value otherwise.

### GetTelegramPaymentChargeIdOk

`func (o *EditUserStarSubscriptionPostRequest) GetTelegramPaymentChargeIdOk() (*string, bool)`

GetTelegramPaymentChargeIdOk returns a tuple with the TelegramPaymentChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramPaymentChargeId

`func (o *EditUserStarSubscriptionPostRequest) SetTelegramPaymentChargeId(v string)`

SetTelegramPaymentChargeId sets TelegramPaymentChargeId field to given value.


### GetIsCanceled

`func (o *EditUserStarSubscriptionPostRequest) GetIsCanceled() bool`

GetIsCanceled returns the IsCanceled field if non-nil, zero value otherwise.

### GetIsCanceledOk

`func (o *EditUserStarSubscriptionPostRequest) GetIsCanceledOk() (*bool, bool)`

GetIsCanceledOk returns a tuple with the IsCanceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCanceled

`func (o *EditUserStarSubscriptionPostRequest) SetIsCanceled(v bool)`

SetIsCanceled sets IsCanceled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


