package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilder casts the previous resource to endpoint.
type AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilderGetQueryParameters get the item of type microsoft.graph.directoryObject as microsoft.graph.endpoint
type AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilderGetQueryParameters
}
// NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilderInternal instantiates a new EndpointRequestBuilder and sets the default values.
func NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilder) {
    m := &AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}/microsoft.graph.endpoint{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilder instantiates a new EndpointRequestBuilder and sets the default values.
func NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.endpoint
func (m *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Endpointable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEndpointFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Endpointable), nil
}
// ToGetRequestInformation get the item of type microsoft.graph.directoryObject as microsoft.graph.endpoint
func (m *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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