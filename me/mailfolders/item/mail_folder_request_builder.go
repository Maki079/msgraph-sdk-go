package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i23a948d948931da4c2473f25bf2771887399012591e170bf2c38ffb2a3ba66c0 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages"
    i41238d4f3e944195a28128ba7c7e506c78ae7944f63ee5b99f6ccf43fc959a1f "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/childfolders"
    ia3a6564ed24f76e12dff691849f92f63255d225c1ca3736c45221460427057ed "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/move"
    iad30d3cef76cd9f7b9285b76f6f4477abd7036d9f2c35bc5ed1522f0af02c78a "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/singlevalueextendedproperties"
    id3d1b1d14374855b32699cf86c45aa2624debfc4187f8f34b6f03ceb6bcdc5ac "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messagerules"
    ief5060c34a3b1d4aad14d434f772d2eb07c6f99ebfee6ebac51aeec75a2ea7fb "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/multivalueextendedproperties"
    if16b5f206adf92e3678b287eee1edc09a825e313dddd8774b3a749501630578c "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/copy"
    i355444ee85a0200730d90e442a49ca5aee393c423237c41182037176c7a79d49 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messagerules/item"
    i90c07370d4369a4b6aee4e8f38d2071aa7e2def7f82f01fe6ce1274c5065e166 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/singlevalueextendedproperties/item"
    ia104b6ddf37a2b233498082fb22da19da091795fbd96aa1548ddd53087f11025 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/childfolders/item"
    ieb796547151eb3ed23dc629e3abe93c6b6c3f96fa2ab5ff99d5ec88c1e8ff28a "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/multivalueextendedproperties/item"
    ife199debe7c0cb4dd460b2e4c2e997cb1496b1b2299d0849f63f7a2d76121d8e "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages/item"
)

type MailFolderRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type MailFolderRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *MailFolderRequestBuilder) ChildFolders()(*i41238d4f3e944195a28128ba7c7e506c78ae7944f63ee5b99f6ccf43fc959a1f.ChildFoldersRequestBuilder) {
    return i41238d4f3e944195a28128ba7c7e506c78ae7944f63ee5b99f6ccf43fc959a1f.NewChildFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) ChildFoldersById(id string)(*ia104b6ddf37a2b233498082fb22da19da091795fbd96aa1548ddd53087f11025.MailFolderRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["mailFolder_id1"] = id
    }
    return ia104b6ddf37a2b233498082fb22da19da091795fbd96aa1548ddd53087f11025.NewMailFolderRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewMailFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MailFolderRequestBuilder) {
    m := &MailFolderRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/mailFolders/{mailFolder_id}{?select,expand}";
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
func NewMailFolderRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MailFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMailFolderRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MailFolderRequestBuilder) Copy()(*if16b5f206adf92e3678b287eee1edc09a825e313dddd8774b3a749501630578c.CopyRequestBuilder) {
    return if16b5f206adf92e3678b287eee1edc09a825e313dddd8774b3a749501630578c.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *MailFolderRequestBuilder) CreateGetRequestInformation(q func (value *MailFolderRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(MailFolderRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *MailFolderRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MailFolder, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *MailFolderRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *MailFolderRequestBuilder) Get(q func (value *MailFolderRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MailFolder, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewMailFolder() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MailFolder), nil
}
func (m *MailFolderRequestBuilder) MessageRules()(*id3d1b1d14374855b32699cf86c45aa2624debfc4187f8f34b6f03ceb6bcdc5ac.MessageRulesRequestBuilder) {
    return id3d1b1d14374855b32699cf86c45aa2624debfc4187f8f34b6f03ceb6bcdc5ac.NewMessageRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) MessageRulesById(id string)(*i355444ee85a0200730d90e442a49ca5aee393c423237c41182037176c7a79d49.MessageRuleRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["messageRule_id"] = id
    }
    return i355444ee85a0200730d90e442a49ca5aee393c423237c41182037176c7a79d49.NewMessageRuleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) Messages()(*i23a948d948931da4c2473f25bf2771887399012591e170bf2c38ffb2a3ba66c0.MessagesRequestBuilder) {
    return i23a948d948931da4c2473f25bf2771887399012591e170bf2c38ffb2a3ba66c0.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) MessagesById(id string)(*ife199debe7c0cb4dd460b2e4c2e997cb1496b1b2299d0849f63f7a2d76121d8e.MessageRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["message_id"] = id
    }
    return ife199debe7c0cb4dd460b2e4c2e997cb1496b1b2299d0849f63f7a2d76121d8e.NewMessageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) Move()(*ia3a6564ed24f76e12dff691849f92f63255d225c1ca3736c45221460427057ed.MoveRequestBuilder) {
    return ia3a6564ed24f76e12dff691849f92f63255d225c1ca3736c45221460427057ed.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) MultiValueExtendedProperties()(*ief5060c34a3b1d4aad14d434f772d2eb07c6f99ebfee6ebac51aeec75a2ea7fb.MultiValueExtendedPropertiesRequestBuilder) {
    return ief5060c34a3b1d4aad14d434f772d2eb07c6f99ebfee6ebac51aeec75a2ea7fb.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ieb796547151eb3ed23dc629e3abe93c6b6c3f96fa2ab5ff99d5ec88c1e8ff28a.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ieb796547151eb3ed23dc629e3abe93c6b6c3f96fa2ab5ff99d5ec88c1e8ff28a.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MailFolder, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *MailFolderRequestBuilder) SingleValueExtendedProperties()(*iad30d3cef76cd9f7b9285b76f6f4477abd7036d9f2c35bc5ed1522f0af02c78a.SingleValueExtendedPropertiesRequestBuilder) {
    return iad30d3cef76cd9f7b9285b76f6f4477abd7036d9f2c35bc5ed1522f0af02c78a.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i90c07370d4369a4b6aee4e8f38d2071aa7e2def7f82f01fe6ce1274c5065e166.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i90c07370d4369a4b6aee4e8f38d2071aa7e2def7f82f01fe6ce1274c5065e166.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}