# InlineQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for this query | 
**From** | [**User**](User.md) |  | 
**Query** | **string** | Text of the query (up to 256 characters) | 
**Offset** | **string** | Offset of the results to be returned, can be controlled by the bot | 
**ChatType** | Pointer to **string** | *Optional*. Type of the chat from which the inline query was sent. Can be either “sender” for a private chat with the inline query sender, “private”, “group”, “supergroup”, or “channel”. The chat type should be always known for requests sent from official clients and most third-party clients, unless the request was sent from a secret chat | [optional] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 

## Methods

### NewInlineQuery

`func NewInlineQuery(id string, from User, query string, offset string, ) *InlineQuery`

NewInlineQuery instantiates a new InlineQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineQueryWithDefaults

`func NewInlineQueryWithDefaults() *InlineQuery`

NewInlineQueryWithDefaults instantiates a new InlineQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineQuery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineQuery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineQuery) SetId(v string)`

SetId sets Id field to given value.


### GetFrom

`func (o *InlineQuery) GetFrom() User`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *InlineQuery) GetFromOk() (*User, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *InlineQuery) SetFrom(v User)`

SetFrom sets From field to given value.


### GetQuery

`func (o *InlineQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *InlineQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *InlineQuery) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetOffset

`func (o *InlineQuery) GetOffset() string`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *InlineQuery) GetOffsetOk() (*string, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *InlineQuery) SetOffset(v string)`

SetOffset sets Offset field to given value.


### GetChatType

`func (o *InlineQuery) GetChatType() string`

GetChatType returns the ChatType field if non-nil, zero value otherwise.

### GetChatTypeOk

`func (o *InlineQuery) GetChatTypeOk() (*string, bool)`

GetChatTypeOk returns a tuple with the ChatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatType

`func (o *InlineQuery) SetChatType(v string)`

SetChatType sets ChatType field to given value.

### HasChatType

`func (o *InlineQuery) HasChatType() bool`

HasChatType returns a boolean if a field has been set.

### GetLocation

`func (o *InlineQuery) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InlineQuery) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InlineQuery) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *InlineQuery) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


