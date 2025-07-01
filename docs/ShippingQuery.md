# ShippingQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique query identifier | 
**From** | [**User**](User.md) |  | 
**InvoicePayload** | **string** | Bot-specified invoice payload | 
**ShippingAddress** | [**ShippingAddress**](ShippingAddress.md) |  | 

## Methods

### NewShippingQuery

`func NewShippingQuery(id string, from User, invoicePayload string, shippingAddress ShippingAddress, ) *ShippingQuery`

NewShippingQuery instantiates a new ShippingQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingQueryWithDefaults

`func NewShippingQueryWithDefaults() *ShippingQuery`

NewShippingQueryWithDefaults instantiates a new ShippingQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShippingQuery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShippingQuery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShippingQuery) SetId(v string)`

SetId sets Id field to given value.


### GetFrom

`func (o *ShippingQuery) GetFrom() User`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ShippingQuery) GetFromOk() (*User, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ShippingQuery) SetFrom(v User)`

SetFrom sets From field to given value.


### GetInvoicePayload

`func (o *ShippingQuery) GetInvoicePayload() string`

GetInvoicePayload returns the InvoicePayload field if non-nil, zero value otherwise.

### GetInvoicePayloadOk

`func (o *ShippingQuery) GetInvoicePayloadOk() (*string, bool)`

GetInvoicePayloadOk returns a tuple with the InvoicePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePayload

`func (o *ShippingQuery) SetInvoicePayload(v string)`

SetInvoicePayload sets InvoicePayload field to given value.


### GetShippingAddress

`func (o *ShippingQuery) GetShippingAddress() ShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *ShippingQuery) GetShippingAddressOk() (*ShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *ShippingQuery) SetShippingAddress(v ShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


