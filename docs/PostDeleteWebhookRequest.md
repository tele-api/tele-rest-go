# PostDeleteWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DropPendingUpdates** | Pointer to **bool** | Pass *True* to drop all pending updates | [optional] 

## Methods

### NewPostDeleteWebhookRequest

`func NewPostDeleteWebhookRequest() *PostDeleteWebhookRequest`

NewPostDeleteWebhookRequest instantiates a new PostDeleteWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostDeleteWebhookRequestWithDefaults

`func NewPostDeleteWebhookRequestWithDefaults() *PostDeleteWebhookRequest`

NewPostDeleteWebhookRequestWithDefaults instantiates a new PostDeleteWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDropPendingUpdates

`func (o *PostDeleteWebhookRequest) GetDropPendingUpdates() bool`

GetDropPendingUpdates returns the DropPendingUpdates field if non-nil, zero value otherwise.

### GetDropPendingUpdatesOk

`func (o *PostDeleteWebhookRequest) GetDropPendingUpdatesOk() (*bool, bool)`

GetDropPendingUpdatesOk returns a tuple with the DropPendingUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropPendingUpdates

`func (o *PostDeleteWebhookRequest) SetDropPendingUpdates(v bool)`

SetDropPendingUpdates sets DropPendingUpdates field to given value.

### HasDropPendingUpdates

`func (o *PostDeleteWebhookRequest) HasDropPendingUpdates() bool`

HasDropPendingUpdates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


