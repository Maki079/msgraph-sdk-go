package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilder provides operations to manage the memberOf property of the microsoft.graph.device entity.
type AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilderGetQueryParameters groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
type AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilderGetQueryParameters
}
// Application casts the previous resource to application.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilder) Application()(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfApplicationRequestBuilder) {
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilderInternal instantiates a new MemberOfRequestBuilder and sets the default values.
func NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilder) {
    m := &AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/memberOf{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilder instantiates a new MemberOfRequestBuilder and sets the default values.
func NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilder) Count()(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfCountRequestBuilder) {
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Device casts the previous resource to device.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilder) Device()(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDeviceRequestBuilder) {
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/device-list-memberof?view=graph-rest-1.0
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable), nil
}
// Group casts the previous resource to group.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilder) Group()(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfGroupRequestBuilder) {
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilder) OrgContact()(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfOrgContactRequestBuilder) {
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilder) ServicePrincipal()(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfServicePrincipalRequestBuilder) {
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ToGetRequestInformation groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// User casts the previous resource to user.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfRequestBuilder) User()(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfUserRequestBuilder) {
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}