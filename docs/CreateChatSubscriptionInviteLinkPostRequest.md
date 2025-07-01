# CreateChatSubscriptionInviteLinkPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**CreateChatSubscriptionInviteLinkPostRequestChatId**](CreateChatSubscriptionInviteLinkPostRequestChatId.md) |  | 
**Name** | Pointer to **string** | Invite link name; 0-32 characters | [optional] 
**SubscriptionPeriod** | **int32** | The number of seconds the subscription will be active for before the next payment. Currently, it must always be 2592000 (30 days). | 
**SubscriptionPrice** | **int32** | The amount of Telegram Stars a user must pay initially and after each subsequent subscription period to be a member of the chat; 1-10000 | 

## Methods

### NewCreateChatSubscriptionInviteLinkPostRequest

`func NewCreateChatSubscriptionInviteLinkPostRequest(chatId CreateChatSubscriptionInviteLinkPostRequestChatId, subscriptionPeriod int32, subscriptionPrice int32, ) *CreateChatSubscriptionInviteLinkPostRequest`

NewCreateChatSubscriptionInviteLinkPostRequest instantiates a new CreateChatSubscriptionInviteLinkPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateChatSubscriptionInviteLinkPostRequestWithDefaults

`func NewCreateChatSubscriptionInviteLinkPostRequestWithDefaults() *CreateChatSubscriptionInviteLinkPostRequest`

NewCreateChatSubscriptionInviteLinkPostRequestWithDefaults instantiates a new CreateChatSubscriptionInviteLinkPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *CreateChatSubscriptionInviteLinkPostRequest) GetChatId() CreateChatSubscriptionInviteLinkPostRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *CreateChatSubscriptionInviteLinkPostRequest) GetChatIdOk() (*CreateChatSubscriptionInviteLinkPostRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *CreateChatSubscriptionInviteLinkPostRequest) SetChatId(v CreateChatSubscriptionInviteLinkPostRequestChatId)`

SetChatId sets ChatId field to given value.


### GetName

`func (o *CreateChatSubscriptionInviteLinkPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateChatSubscriptionInviteLinkPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateChatSubscriptionInviteLinkPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateChatSubscriptionInviteLinkPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubscriptionPeriod

`func (o *CreateChatSubscriptionInviteLinkPostRequest) GetSubscriptionPeriod() int32`

GetSubscriptionPeriod returns the SubscriptionPeriod field if non-nil, zero value otherwise.

### GetSubscriptionPeriodOk

`func (o *CreateChatSubscriptionInviteLinkPostRequest) GetSubscriptionPeriodOk() (*int32, bool)`

GetSubscriptionPeriodOk returns a tuple with the SubscriptionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriod

`func (o *CreateChatSubscriptionInviteLinkPostRequest) SetSubscriptionPeriod(v int32)`

SetSubscriptionPeriod sets SubscriptionPeriod field to given value.


### GetSubscriptionPrice

`func (o *CreateChatSubscriptionInviteLinkPostRequest) GetSubscriptionPrice() int32`

GetSubscriptionPrice returns the SubscriptionPrice field if non-nil, zero value otherwise.

### GetSubscriptionPriceOk

`func (o *CreateChatSubscriptionInviteLinkPostRequest) GetSubscriptionPriceOk() (*int32, bool)`

GetSubscriptionPriceOk returns a tuple with the SubscriptionPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPrice

`func (o *CreateChatSubscriptionInviteLinkPostRequest) SetSubscriptionPrice(v int32)`

SetSubscriptionPrice sets SubscriptionPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


