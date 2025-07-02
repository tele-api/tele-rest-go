# PostCreateChatSubscriptionInviteLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**PostCreateChatSubscriptionInviteLinkRequestChatId**](PostCreateChatSubscriptionInviteLinkRequestChatId.md) |  | 
**Name** | Pointer to **string** | Invite link name; 0-32 characters | [optional] 
**SubscriptionPeriod** | **int32** | The number of seconds the subscription will be active for before the next payment. Currently, it must always be 2592000 (30 days). | 
**SubscriptionPrice** | **int32** | The amount of Telegram Stars a user must pay initially and after each subsequent subscription period to be a member of the chat; 1-10000 | 

## Methods

### NewPostCreateChatSubscriptionInviteLinkRequest

`func NewPostCreateChatSubscriptionInviteLinkRequest(chatId PostCreateChatSubscriptionInviteLinkRequestChatId, subscriptionPeriod int32, subscriptionPrice int32, ) *PostCreateChatSubscriptionInviteLinkRequest`

NewPostCreateChatSubscriptionInviteLinkRequest instantiates a new PostCreateChatSubscriptionInviteLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCreateChatSubscriptionInviteLinkRequestWithDefaults

`func NewPostCreateChatSubscriptionInviteLinkRequestWithDefaults() *PostCreateChatSubscriptionInviteLinkRequest`

NewPostCreateChatSubscriptionInviteLinkRequestWithDefaults instantiates a new PostCreateChatSubscriptionInviteLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *PostCreateChatSubscriptionInviteLinkRequest) GetChatId() PostCreateChatSubscriptionInviteLinkRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *PostCreateChatSubscriptionInviteLinkRequest) GetChatIdOk() (*PostCreateChatSubscriptionInviteLinkRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *PostCreateChatSubscriptionInviteLinkRequest) SetChatId(v PostCreateChatSubscriptionInviteLinkRequestChatId)`

SetChatId sets ChatId field to given value.


### GetName

`func (o *PostCreateChatSubscriptionInviteLinkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostCreateChatSubscriptionInviteLinkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostCreateChatSubscriptionInviteLinkRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostCreateChatSubscriptionInviteLinkRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubscriptionPeriod

`func (o *PostCreateChatSubscriptionInviteLinkRequest) GetSubscriptionPeriod() int32`

GetSubscriptionPeriod returns the SubscriptionPeriod field if non-nil, zero value otherwise.

### GetSubscriptionPeriodOk

`func (o *PostCreateChatSubscriptionInviteLinkRequest) GetSubscriptionPeriodOk() (*int32, bool)`

GetSubscriptionPeriodOk returns a tuple with the SubscriptionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriod

`func (o *PostCreateChatSubscriptionInviteLinkRequest) SetSubscriptionPeriod(v int32)`

SetSubscriptionPeriod sets SubscriptionPeriod field to given value.


### GetSubscriptionPrice

`func (o *PostCreateChatSubscriptionInviteLinkRequest) GetSubscriptionPrice() int32`

GetSubscriptionPrice returns the SubscriptionPrice field if non-nil, zero value otherwise.

### GetSubscriptionPriceOk

`func (o *PostCreateChatSubscriptionInviteLinkRequest) GetSubscriptionPriceOk() (*int32, bool)`

GetSubscriptionPriceOk returns a tuple with the SubscriptionPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPrice

`func (o *PostCreateChatSubscriptionInviteLinkRequest) SetSubscriptionPrice(v int32)`

SetSubscriptionPrice sets SubscriptionPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


