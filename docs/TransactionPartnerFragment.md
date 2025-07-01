# TransactionPartnerFragment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the transaction partner, always “fragment” | [default to "fragment"]
**WithdrawalState** | Pointer to [**RevenueWithdrawalState**](RevenueWithdrawalState.md) |  | [optional] 

## Methods

### NewTransactionPartnerFragment

`func NewTransactionPartnerFragment(type_ string, ) *TransactionPartnerFragment`

NewTransactionPartnerFragment instantiates a new TransactionPartnerFragment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionPartnerFragmentWithDefaults

`func NewTransactionPartnerFragmentWithDefaults() *TransactionPartnerFragment`

NewTransactionPartnerFragmentWithDefaults instantiates a new TransactionPartnerFragment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransactionPartnerFragment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionPartnerFragment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionPartnerFragment) SetType(v string)`

SetType sets Type field to given value.


### GetWithdrawalState

`func (o *TransactionPartnerFragment) GetWithdrawalState() RevenueWithdrawalState`

GetWithdrawalState returns the WithdrawalState field if non-nil, zero value otherwise.

### GetWithdrawalStateOk

`func (o *TransactionPartnerFragment) GetWithdrawalStateOk() (*RevenueWithdrawalState, bool)`

GetWithdrawalStateOk returns a tuple with the WithdrawalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalState

`func (o *TransactionPartnerFragment) SetWithdrawalState(v RevenueWithdrawalState)`

SetWithdrawalState sets WithdrawalState field to given value.

### HasWithdrawalState

`func (o *TransactionPartnerFragment) HasWithdrawalState() bool`

HasWithdrawalState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


