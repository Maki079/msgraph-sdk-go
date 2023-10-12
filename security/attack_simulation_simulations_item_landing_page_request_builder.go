package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AttackSimulationSimulationsItemLandingPageRequestBuilder provides operations to manage the landingPage property of the microsoft.graph.simulation entity.
type AttackSimulationSimulationsItemLandingPageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AttackSimulationSimulationsItemLandingPageRequestBuilderGetQueryParameters get landingPage from security
type AttackSimulationSimulationsItemLandingPageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AttackSimulationSimulationsItemLandingPageRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttackSimulationSimulationsItemLandingPageRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AttackSimulationSimulationsItemLandingPageRequestBuilderGetQueryParameters
}
// NewAttackSimulationSimulationsItemLandingPageRequestBuilderInternal instantiates a new LandingPageRequestBuilder and sets the default values.
func NewAttackSimulationSimulationsItemLandingPageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttackSimulationSimulationsItemLandingPageRequestBuilder) {
    m := &AttackSimulationSimulationsItemLandingPageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/attackSimulation/simulations/{simulation%2Did}/landingPage{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewAttackSimulationSimulationsItemLandingPageRequestBuilder instantiates a new LandingPageRequestBuilder and sets the default values.
func NewAttackSimulationSimulationsItemLandingPageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttackSimulationSimulationsItemLandingPageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAttackSimulationSimulationsItemLandingPageRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get landingPage from security
func (m *AttackSimulationSimulationsItemLandingPageRequestBuilder) Get(ctx context.Context, requestConfiguration *AttackSimulationSimulationsItemLandingPageRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LandingPageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateLandingPageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LandingPageable), nil
}
// ToGetRequestInformation get landingPage from security
func (m *AttackSimulationSimulationsItemLandingPageRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AttackSimulationSimulationsItemLandingPageRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *AttackSimulationSimulationsItemLandingPageRequestBuilder) WithUrl(rawUrl string)(*AttackSimulationSimulationsItemLandingPageRequestBuilder) {
    return NewAttackSimulationSimulationsItemLandingPageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}