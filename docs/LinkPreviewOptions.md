# LinkPreviewOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDisabled** | Pointer to **bool** | *Optional*. *True*, if the link preview is disabled | [optional] 
**Url** | Pointer to **string** | *Optional*. URL to use for the link preview. If empty, then the first URL found in the message text will be used | [optional] 
**PreferSmallMedia** | Pointer to **bool** | *Optional*. *True*, if the media in the link preview is supposed to be shrunk; ignored if the URL isn&#39;t explicitly specified or media size change isn&#39;t supported for the preview | [optional] 
**PreferLargeMedia** | Pointer to **bool** | *Optional*. *True*, if the media in the link preview is supposed to be enlarged; ignored if the URL isn&#39;t explicitly specified or media size change isn&#39;t supported for the preview | [optional] 
**ShowAboveText** | Pointer to **bool** | *Optional*. *True*, if the link preview must be shown above the message text; otherwise, the link preview will be shown below the message text | [optional] 

## Methods

### NewLinkPreviewOptions

`func NewLinkPreviewOptions() *LinkPreviewOptions`

NewLinkPreviewOptions instantiates a new LinkPreviewOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkPreviewOptionsWithDefaults

`func NewLinkPreviewOptionsWithDefaults() *LinkPreviewOptions`

NewLinkPreviewOptionsWithDefaults instantiates a new LinkPreviewOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDisabled

`func (o *LinkPreviewOptions) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *LinkPreviewOptions) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *LinkPreviewOptions) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *LinkPreviewOptions) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetUrl

`func (o *LinkPreviewOptions) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LinkPreviewOptions) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LinkPreviewOptions) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LinkPreviewOptions) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetPreferSmallMedia

`func (o *LinkPreviewOptions) GetPreferSmallMedia() bool`

GetPreferSmallMedia returns the PreferSmallMedia field if non-nil, zero value otherwise.

### GetPreferSmallMediaOk

`func (o *LinkPreviewOptions) GetPreferSmallMediaOk() (*bool, bool)`

GetPreferSmallMediaOk returns a tuple with the PreferSmallMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferSmallMedia

`func (o *LinkPreviewOptions) SetPreferSmallMedia(v bool)`

SetPreferSmallMedia sets PreferSmallMedia field to given value.

### HasPreferSmallMedia

`func (o *LinkPreviewOptions) HasPreferSmallMedia() bool`

HasPreferSmallMedia returns a boolean if a field has been set.

### GetPreferLargeMedia

`func (o *LinkPreviewOptions) GetPreferLargeMedia() bool`

GetPreferLargeMedia returns the PreferLargeMedia field if non-nil, zero value otherwise.

### GetPreferLargeMediaOk

`func (o *LinkPreviewOptions) GetPreferLargeMediaOk() (*bool, bool)`

GetPreferLargeMediaOk returns a tuple with the PreferLargeMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferLargeMedia

`func (o *LinkPreviewOptions) SetPreferLargeMedia(v bool)`

SetPreferLargeMedia sets PreferLargeMedia field to given value.

### HasPreferLargeMedia

`func (o *LinkPreviewOptions) HasPreferLargeMedia() bool`

HasPreferLargeMedia returns a boolean if a field has been set.

### GetShowAboveText

`func (o *LinkPreviewOptions) GetShowAboveText() bool`

GetShowAboveText returns the ShowAboveText field if non-nil, zero value otherwise.

### GetShowAboveTextOk

`func (o *LinkPreviewOptions) GetShowAboveTextOk() (*bool, bool)`

GetShowAboveTextOk returns a tuple with the ShowAboveText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAboveText

`func (o *LinkPreviewOptions) SetShowAboveText(v bool)`

SetShowAboveText sets ShowAboveText field to given value.

### HasShowAboveText

`func (o *LinkPreviewOptions) HasShowAboveText() bool`

HasShowAboveText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


