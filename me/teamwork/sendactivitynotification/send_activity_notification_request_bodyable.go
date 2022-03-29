package sendactivitynotification

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// SendActivityNotificationRequestBodyable 
type SendActivityNotificationRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActivityType()(*string)
    GetChainId()(*int64)
    GetPreviewText()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ItemBodyable)
    GetTemplateParameters()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.KeyValuePairable)
    GetTopic()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TeamworkActivityTopicable)
    SetActivityType(value *string)()
    SetChainId(value *int64)()
    SetPreviewText(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ItemBodyable)()
    SetTemplateParameters(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.KeyValuePairable)()
    SetTopic(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TeamworkActivityTopicable)()
}