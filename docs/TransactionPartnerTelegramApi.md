# TransactionPartnerTelegramApi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the transaction partner, always “telegram\\_api” | [default to "telegram_api"]
**RequestCount** | **int32** | The number of successful requests that exceeded regular limits and were therefore billed | 

## Methods

### NewTransactionPartnerTelegramApi

`func NewTransactionPartnerTelegramApi(type_ string, requestCount int32, ) *TransactionPartnerTelegramApi`

NewTransactionPartnerTelegramApi instantiates a new TransactionPartnerTelegramApi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionPartnerTelegramApiWithDefaults

`func NewTransactionPartnerTelegramApiWithDefaults() *TransactionPartnerTelegramApi`

NewTransactionPartnerTelegramApiWithDefaults instantiates a new TransactionPartnerTelegramApi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransactionPartnerTelegramApi) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionPartnerTelegramApi) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionPartnerTelegramApi) SetType(v string)`

SetType sets Type field to given value.


### GetRequestCount

`func (o *TransactionPartnerTelegramApi) GetRequestCount() int32`

GetRequestCount returns the RequestCount field if non-nil, zero value otherwise.

### GetRequestCountOk

`func (o *TransactionPartnerTelegramApi) GetRequestCountOk() (*int32, bool)`

GetRequestCountOk returns a tuple with the RequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCount

`func (o *TransactionPartnerTelegramApi) SetRequestCount(v int32)`

SetRequestCount sets RequestCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


