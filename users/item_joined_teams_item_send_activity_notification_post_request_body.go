package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemJoinedTeamsItemSendActivityNotificationPostRequestBody 
type ItemJoinedTeamsItemSendActivityNotificationPostRequestBody struct {
    // The activityType property
    activityType *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The chainId property
    chainId *int64
    // The previewText property
    previewText iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemBodyable
    // The recipient property
    recipient iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkNotificationRecipientable
    // The templateParameters property
    templateParameters []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValuePairable
    // The topic property
    topic iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkActivityTopicable
}
// NewItemJoinedTeamsItemSendActivityNotificationPostRequestBody instantiates a new ItemJoinedTeamsItemSendActivityNotificationPostRequestBody and sets the default values.
func NewItemJoinedTeamsItemSendActivityNotificationPostRequestBody()(*ItemJoinedTeamsItemSendActivityNotificationPostRequestBody) {
    m := &ItemJoinedTeamsItemSendActivityNotificationPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateItemJoinedTeamsItemSendActivityNotificationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemJoinedTeamsItemSendActivityNotificationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedTeamsItemSendActivityNotificationPostRequestBody(), nil
}
// GetActivityType gets the activityType property value. The activityType property
func (m *ItemJoinedTeamsItemSendActivityNotificationPostRequestBody) GetActivityType()(*string) {
    return m.activityType
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemJoinedTeamsItemSendActivityNotificationPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChainId gets the chainId property value. The chainId property
func (m *ItemJoinedTeamsItemSendActivityNotificationPostRequestBody) GetChainId()(*int64) {
    return m.chainId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemJoinedTeamsItemSendActivityNotificationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["activityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityType(val)
        }
        return nil
    }
    res["chainId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChainId(val)
        }
        return nil
    }
    res["previewText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewText(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemBodyable))
        }
        return nil
    }
    res["recipient"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamworkNotificationRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipient(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkNotificationRecipientable))
        }
        return nil
    }
    res["templateParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValuePairable)
            }
            m.SetTemplateParameters(res)
        }
        return nil
    }
    res["topic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamworkActivityTopicFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopic(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkActivityTopicable))
        }
        return nil
    }
    return res
}
// GetPreviewText gets the previewText property value. The previewText property
func (m *ItemJoinedTeamsItemSendActivityNotificationPostRequestBody) GetPreviewText()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemBodyable) {
    return m.previewText
}
// GetRecipient gets the recipient property value. The recipient property
func (m *ItemJoinedTeamsItemSendActivityNotificationPostRequestBody) GetRecipient()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkNotificationRecipientable) {
    return m.recipient
}
// GetTemplateParameters gets the templateParameters property value. The templateParameters property
func (m *ItemJoinedTeamsItemSendActivityNotificationPostRequestBody) GetTemplateParameters()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValuePairable) {
    return m.templateParameters
}
// GetTopic gets the topic property value. The topic property
func (m *ItemJoinedTeamsItemSendActivityNotificationPostRequestBody) GetTopic()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkActivityTopicable) {
    return m.topic
}
// Serialize serializes information the current object
func (m *ItemJoinedTeamsItemSendActivityNotificationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("activityType", m.GetActivityType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("chainId", m.GetChainId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("previewText", m.GetPreviewText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("recipient", m.GetRecipient())
        if err != nil {
            return err
        }
    }
    if m.GetTemplateParameters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTemplateParameters()))
        for i, v := range m.GetTemplateParameters() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("templateParameters", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("topic", m.GetTopic())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivityType sets the activityType property value. The activityType property
func (m *ItemJoinedTeamsItemSendActivityNotificationPostRequestBody) SetActivityType(value *string)() {
    m.activityType = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemJoinedTeamsItemSendActivityNotificationPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChainId sets the chainId property value. The chainId property
func (m *ItemJoinedTeamsItemSendActivityNotificationPostRequestBody) SetChainId(value *int64)() {
    m.chainId = value
}
// SetPreviewText sets the previewText property value. The previewText property
func (m *ItemJoinedTeamsItemSendActivityNotificationPostRequestBody) SetPreviewText(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemBodyable)() {
    m.previewText = value
}
// SetRecipient sets the recipient property value. The recipient property
func (m *ItemJoinedTeamsItemSendActivityNotificationPostRequestBody) SetRecipient(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkNotificationRecipientable)() {
    m.recipient = value
}
// SetTemplateParameters sets the templateParameters property value. The templateParameters property
func (m *ItemJoinedTeamsItemSendActivityNotificationPostRequestBody) SetTemplateParameters(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValuePairable)() {
    m.templateParameters = value
}
// SetTopic sets the topic property value. The topic property
func (m *ItemJoinedTeamsItemSendActivityNotificationPostRequestBody) SetTopic(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkActivityTopicable)() {
    m.topic = value
}
