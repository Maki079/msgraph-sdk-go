package schedulechangerequest

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ieee4e59c57c0e9adc1ea4345b4942d58f31e5728a35c2c221f53850f2f04c2ed "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/schedulechangerequest/approve"
    iffdc112027db305c5e8673cd0e0cc837fb3f9c085722ce9321bf70bad6906019 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/schedulechangerequest/decline"
)

type ScheduleChangeRequestRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *ScheduleChangeRequestRequestBuilder) Approve()(*ieee4e59c57c0e9adc1ea4345b4942d58f31e5728a35c2c221f53850f2f04c2ed.ApproveRequestBuilder) {
    return ieee4e59c57c0e9adc1ea4345b4942d58f31e5728a35c2c221f53850f2f04c2ed.NewApproveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewScheduleChangeRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduleChangeRequestRequestBuilder) {
    m := &ScheduleChangeRequestRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/insights/shared/{sharedInsight_id}/lastSharedMethod/microsoft.graph.scheduleChangeRequest";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewScheduleChangeRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduleChangeRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScheduleChangeRequestRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ScheduleChangeRequestRequestBuilder) Decline()(*iffdc112027db305c5e8673cd0e0cc837fb3f9c085722ce9321bf70bad6906019.DeclineRequestBuilder) {
    return iffdc112027db305c5e8673cd0e0cc837fb3f9c085722ce9321bf70bad6906019.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}