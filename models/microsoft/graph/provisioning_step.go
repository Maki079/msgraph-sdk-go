package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ProvisioningStep struct {
    additionalData map[string]interface{};
    description *string;
    details *DetailsInfo;
    name *string;
    provisioningStepType *ProvisioningStepType;
    status *ProvisioningResult;
}
func NewProvisioningStep()(*ProvisioningStep) {
    m := &ProvisioningStep{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ProvisioningStep) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ProvisioningStep) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *ProvisioningStep) GetDetails()(*DetailsInfo) {
    if m == nil {
        return nil
    } else {
        return m.details
    }
}
func (m *ProvisioningStep) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *ProvisioningStep) GetProvisioningStepType()(*ProvisioningStepType) {
    if m == nil {
        return nil
    } else {
        return m.provisioningStepType
    }
}
func (m *ProvisioningStep) GetStatus()(*ProvisioningResult) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *ProvisioningStep) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["details"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDetailsInfo() })
        if err != nil {
            return err
        }
        m.SetDetails(val.(*DetailsInfo))
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["provisioningStepType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningStepType)
        if err != nil {
            return err
        }
        cast := val.(ProvisioningStepType)
        m.SetProvisioningStepType(&cast)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningResult)
        if err != nil {
            return err
        }
        cast := val.(ProvisioningResult)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *ProvisioningStep) IsNil()(bool) {
    return m == nil
}
func (m *ProvisioningStep) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("details", m.GetDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetProvisioningStepType() != nil {
        cast := m.GetProvisioningStepType().String()
        err := writer.WriteStringValue("provisioningStepType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *ProvisioningStep) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ProvisioningStep) SetDescription(value *string)() {
    m.description = value
}
func (m *ProvisioningStep) SetDetails(value *DetailsInfo)() {
    m.details = value
}
func (m *ProvisioningStep) SetName(value *string)() {
    m.name = value
}
func (m *ProvisioningStep) SetProvisioningStepType(value *ProvisioningStepType)() {
    m.provisioningStepType = value
}
func (m *ProvisioningStep) SetStatus(value *ProvisioningResult)() {
    m.status = value
}