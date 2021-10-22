package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserAttributeValuesItem struct {
    additionalData map[string]interface{};
    isDefault *bool;
    name *string;
    value *string;
}
func NewUserAttributeValuesItem()(*UserAttributeValuesItem) {
    m := &UserAttributeValuesItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UserAttributeValuesItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UserAttributeValuesItem) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
func (m *UserAttributeValuesItem) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *UserAttributeValuesItem) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *UserAttributeValuesItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDefault(val)
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
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetValue(val)
        return nil
    }
    return res
}
func (m *UserAttributeValuesItem) IsNil()(bool) {
    return m == nil
}
func (m *UserAttributeValuesItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isDefault", m.GetIsDefault())
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
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *UserAttributeValuesItem) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UserAttributeValuesItem) SetIsDefault(value *bool)() {
    m.isDefault = value
}
func (m *UserAttributeValuesItem) SetName(value *string)() {
    m.name = value
}
func (m *UserAttributeValuesItem) SetValue(value *string)() {
    m.value = value
}