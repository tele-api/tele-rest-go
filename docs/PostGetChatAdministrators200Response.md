# PostGetChatAdministrators200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | [default to true]
**Result** | [**[]ChatMember**](ChatMember.md) |  | 

## Methods

### NewPostGetChatAdministrators200Response

`func NewPostGetChatAdministrators200Response(ok bool, result []ChatMember, ) *PostGetChatAdministrators200Response`

NewPostGetChatAdministrators200Response instantiates a new PostGetChatAdministrators200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostGetChatAdministrators200ResponseWithDefaults

`func NewPostGetChatAdministrators200ResponseWithDefaults() *PostGetChatAdministrators200Response`

NewPostGetChatAdministrators200ResponseWithDefaults instantiates a new PostGetChatAdministrators200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *PostGetChatAdministrators200Response) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *PostGetChatAdministrators200Response) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *PostGetChatAdministrators200Response) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetResult

`func (o *PostGetChatAdministrators200Response) GetResult() []ChatMember`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PostGetChatAdministrators200Response) GetResultOk() (*[]ChatMember, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PostGetChatAdministrators200Response) SetResult(v []ChatMember)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


