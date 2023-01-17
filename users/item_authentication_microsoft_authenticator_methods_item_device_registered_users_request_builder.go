package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilder provides operations to manage the registeredUsers property of the microsoft.graph.device entity.
type ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilderGetQueryParameters collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
type ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilderGetQueryParameters struct {
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
// ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilderGetQueryParameters
}
// AppRoleAssignment casts the previous resource to appRoleAssignment.
func (m *ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilder) AppRoleAssignment()(*ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersAppRoleAssignmentRequestBuilder) {
    return NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersAppRoleAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilderInternal instantiates a new RegisteredUsersRequestBuilder and sets the default values.
func NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilder) {
    m := &ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/registeredUsers{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilder instantiates a new RegisteredUsersRequestBuilder and sets the default values.
func NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilder) Count()(*ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersCountRequestBuilder) {
    return NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Endpoint casts the previous resource to endpoint.
func (m *ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilder) Endpoint()(*ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersEndpointRequestBuilder) {
    return NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/device-list-registeredusers?view=graph-rest-1.0
func (m *ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
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
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilder) ServicePrincipal()(*ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersServicePrincipalRequestBuilder) {
    return NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ToGetRequestInformation collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
func (m *ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilder) User()(*ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersUserRequestBuilder) {
    return NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
