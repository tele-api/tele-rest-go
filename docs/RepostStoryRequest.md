# RepostStoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection | 
**FromChatId** | **int32** | Unique identifier of the chat which posted the story that should be reposted | 
**FromStoryId** | **int32** | Unique identifier of the story that should be reposted | 
**ActivePeriod** | **int32** | Period after which the story is moved to the archive, in seconds; must be one of &#x60;6 * 3600&#x60;, &#x60;12 * 3600&#x60;, &#x60;86400&#x60;, or &#x60;2 * 86400&#x60; | 
**PostToChatPage** | Pointer to **bool** | Pass *True* to keep the story accessible after it expires | [optional] 
**ProtectContent** | Pointer to **bool** | Pass *True* if the content of the story must be protected from forwarding and screenshotting | [optional] 

## Methods

### NewRepostStoryRequest

`func NewRepostStoryRequest(businessConnectionId string, fromChatId int32, fromStoryId int32, activePeriod int32, ) *RepostStoryRequest`

NewRepostStoryRequest instantiates a new RepostStoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepostStoryRequestWithDefaults

`func NewRepostStoryRequestWithDefaults() *RepostStoryRequest`

NewRepostStoryRequestWithDefaults instantiates a new RepostStoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *RepostStoryRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *RepostStoryRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *RepostStoryRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetFromChatId

`func (o *RepostStoryRequest) GetFromChatId() int32`

GetFromChatId returns the FromChatId field if non-nil, zero value otherwise.

### GetFromChatIdOk

`func (o *RepostStoryRequest) GetFromChatIdOk() (*int32, bool)`

GetFromChatIdOk returns a tuple with the FromChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromChatId

`func (o *RepostStoryRequest) SetFromChatId(v int32)`

SetFromChatId sets FromChatId field to given value.


### GetFromStoryId

`func (o *RepostStoryRequest) GetFromStoryId() int32`

GetFromStoryId returns the FromStoryId field if non-nil, zero value otherwise.

### GetFromStoryIdOk

`func (o *RepostStoryRequest) GetFromStoryIdOk() (*int32, bool)`

GetFromStoryIdOk returns a tuple with the FromStoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromStoryId

`func (o *RepostStoryRequest) SetFromStoryId(v int32)`

SetFromStoryId sets FromStoryId field to given value.


### GetActivePeriod

`func (o *RepostStoryRequest) GetActivePeriod() int32`

GetActivePeriod returns the ActivePeriod field if non-nil, zero value otherwise.

### GetActivePeriodOk

`func (o *RepostStoryRequest) GetActivePeriodOk() (*int32, bool)`

GetActivePeriodOk returns a tuple with the ActivePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePeriod

`func (o *RepostStoryRequest) SetActivePeriod(v int32)`

SetActivePeriod sets ActivePeriod field to given value.


### GetPostToChatPage

`func (o *RepostStoryRequest) GetPostToChatPage() bool`

GetPostToChatPage returns the PostToChatPage field if non-nil, zero value otherwise.

### GetPostToChatPageOk

`func (o *RepostStoryRequest) GetPostToChatPageOk() (*bool, bool)`

GetPostToChatPageOk returns a tuple with the PostToChatPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostToChatPage

`func (o *RepostStoryRequest) SetPostToChatPage(v bool)`

SetPostToChatPage sets PostToChatPage field to given value.

### HasPostToChatPage

`func (o *RepostStoryRequest) HasPostToChatPage() bool`

HasPostToChatPage returns a boolean if a field has been set.

### GetProtectContent

`func (o *RepostStoryRequest) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *RepostStoryRequest) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *RepostStoryRequest) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *RepostStoryRequest) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


