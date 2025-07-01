# ForceReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceReply** | **bool** | Shows reply interface to the user, as if they manually selected the bot&#39;s message and tapped &#39;Reply&#39; | [default to true]
**InputFieldPlaceholder** | Pointer to **string** | *Optional*. The placeholder to be shown in the input field when the reply is active; 1-64 characters | [optional] 
**Selective** | Pointer to **bool** | *Optional*. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the *text* of the [Message](https://core.telegram.org/bots/api/#message) object; 2) if the bot&#39;s message is a reply to a message in the same chat and forum topic, sender of the original message. | [optional] 

## Methods

### NewForceReply

`func NewForceReply(forceReply bool, ) *ForceReply`

NewForceReply instantiates a new ForceReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForceReplyWithDefaults

`func NewForceReplyWithDefaults() *ForceReply`

NewForceReplyWithDefaults instantiates a new ForceReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceReply

`func (o *ForceReply) GetForceReply() bool`

GetForceReply returns the ForceReply field if non-nil, zero value otherwise.

### GetForceReplyOk

`func (o *ForceReply) GetForceReplyOk() (*bool, bool)`

GetForceReplyOk returns a tuple with the ForceReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceReply

`func (o *ForceReply) SetForceReply(v bool)`

SetForceReply sets ForceReply field to given value.


### GetInputFieldPlaceholder

`func (o *ForceReply) GetInputFieldPlaceholder() string`

GetInputFieldPlaceholder returns the InputFieldPlaceholder field if non-nil, zero value otherwise.

### GetInputFieldPlaceholderOk

`func (o *ForceReply) GetInputFieldPlaceholderOk() (*string, bool)`

GetInputFieldPlaceholderOk returns a tuple with the InputFieldPlaceholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFieldPlaceholder

`func (o *ForceReply) SetInputFieldPlaceholder(v string)`

SetInputFieldPlaceholder sets InputFieldPlaceholder field to given value.

### HasInputFieldPlaceholder

`func (o *ForceReply) HasInputFieldPlaceholder() bool`

HasInputFieldPlaceholder returns a boolean if a field has been set.

### GetSelective

`func (o *ForceReply) GetSelective() bool`

GetSelective returns the Selective field if non-nil, zero value otherwise.

### GetSelectiveOk

`func (o *ForceReply) GetSelectiveOk() (*bool, bool)`

GetSelectiveOk returns a tuple with the Selective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelective

`func (o *ForceReply) SetSelective(v bool)`

SetSelective sets Selective field to given value.

### HasSelective

`func (o *ForceReply) HasSelective() bool`

HasSelective returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


