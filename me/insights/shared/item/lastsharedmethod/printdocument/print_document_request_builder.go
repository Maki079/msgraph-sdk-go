package printdocument

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i6b060b538765d4070158e6569e9b2e5720bab93f21c266b99d3c5bb08eec9270 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/printdocument/createuploadsession"
)

type PrintDocumentRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func NewPrintDocumentRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintDocumentRequestBuilder) {
    m := &PrintDocumentRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/shared/{sharedInsight_id}/lastSharedMethod/microsoft.graph.printDocument";
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
func NewPrintDocumentRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintDocumentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintDocumentRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *PrintDocumentRequestBuilder) CreateUploadSession()(*i6b060b538765d4070158e6569e9b2e5720bab93f21c266b99d3c5bb08eec9270.CreateUploadSessionRequestBuilder) {
    return i6b060b538765d4070158e6569e9b2e5720bab93f21c266b99d3c5bb08eec9270.NewCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}