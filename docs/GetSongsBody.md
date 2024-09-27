# GetSongsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Song** | Pointer to **string** |  | [optional] 
**ReleaseDate** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** | Limit the number of songs returned | [optional] [default to 10]
**Offset** | Pointer to **int32** | Offset for pagination | [optional] [default to 0]

## Methods

### NewGetSongsBody

`func NewGetSongsBody() *GetSongsBody`

NewGetSongsBody instantiates a new GetSongsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSongsBodyWithDefaults

`func NewGetSongsBodyWithDefaults() *GetSongsBody`

NewGetSongsBodyWithDefaults instantiates a new GetSongsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSongsBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSongsBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSongsBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetSongsBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroup

`func (o *GetSongsBody) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GetSongsBody) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GetSongsBody) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GetSongsBody) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetSong

`func (o *GetSongsBody) GetSong() string`

GetSong returns the Song field if non-nil, zero value otherwise.

### GetSongOk

`func (o *GetSongsBody) GetSongOk() (*string, bool)`

GetSongOk returns a tuple with the Song field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSong

`func (o *GetSongsBody) SetSong(v string)`

SetSong sets Song field to given value.

### HasSong

`func (o *GetSongsBody) HasSong() bool`

HasSong returns a boolean if a field has been set.

### GetReleaseDate

`func (o *GetSongsBody) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *GetSongsBody) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *GetSongsBody) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *GetSongsBody) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetText

`func (o *GetSongsBody) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *GetSongsBody) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *GetSongsBody) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *GetSongsBody) HasText() bool`

HasText returns a boolean if a field has been set.

### GetLink

`func (o *GetSongsBody) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *GetSongsBody) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *GetSongsBody) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *GetSongsBody) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetLimit

`func (o *GetSongsBody) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetSongsBody) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetSongsBody) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetSongsBody) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *GetSongsBody) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetSongsBody) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetSongsBody) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetSongsBody) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


