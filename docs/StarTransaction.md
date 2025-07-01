# StarTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the transaction. Coincides with the identifier of the original transaction for refund transactions. Coincides with *SuccessfulPayment.telegram\\_payment\\_charge\\_id* for successful incoming payments from users. | 
**Amount** | **int32** | Integer amount of Telegram Stars transferred by the transaction | 
**NanostarAmount** | Pointer to **int32** | *Optional*. The number of 1/1000000000 shares of Telegram Stars transferred by the transaction; from 0 to 999999999 | [optional] 
**Date** | **int32** | Date the transaction was created in Unix time | 
**Source** | Pointer to [**TransactionPartner**](TransactionPartner.md) |  | [optional] 
**Receiver** | Pointer to [**TransactionPartner**](TransactionPartner.md) |  | [optional] 

## Methods

### NewStarTransaction

`func NewStarTransaction(id string, amount int32, date int32, ) *StarTransaction`

NewStarTransaction instantiates a new StarTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStarTransactionWithDefaults

`func NewStarTransactionWithDefaults() *StarTransaction`

NewStarTransactionWithDefaults instantiates a new StarTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StarTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StarTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StarTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetAmount

`func (o *StarTransaction) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *StarTransaction) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *StarTransaction) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetNanostarAmount

`func (o *StarTransaction) GetNanostarAmount() int32`

GetNanostarAmount returns the NanostarAmount field if non-nil, zero value otherwise.

### GetNanostarAmountOk

`func (o *StarTransaction) GetNanostarAmountOk() (*int32, bool)`

GetNanostarAmountOk returns a tuple with the NanostarAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNanostarAmount

`func (o *StarTransaction) SetNanostarAmount(v int32)`

SetNanostarAmount sets NanostarAmount field to given value.

### HasNanostarAmount

`func (o *StarTransaction) HasNanostarAmount() bool`

HasNanostarAmount returns a boolean if a field has been set.

### GetDate

`func (o *StarTransaction) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *StarTransaction) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *StarTransaction) SetDate(v int32)`

SetDate sets Date field to given value.


### GetSource

`func (o *StarTransaction) GetSource() TransactionPartner`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StarTransaction) GetSourceOk() (*TransactionPartner, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StarTransaction) SetSource(v TransactionPartner)`

SetSource sets Source field to given value.

### HasSource

`func (o *StarTransaction) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetReceiver

`func (o *StarTransaction) GetReceiver() TransactionPartner`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *StarTransaction) GetReceiverOk() (*TransactionPartner, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *StarTransaction) SetReceiver(v TransactionPartner)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *StarTransaction) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


