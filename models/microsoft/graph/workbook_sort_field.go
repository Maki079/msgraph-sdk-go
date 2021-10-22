package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookSortField struct {
    additionalData map[string]interface{};
    ascending *bool;
    color *string;
    dataOption *string;
    icon *WorkbookIcon;
    key *int32;
    sortOn *string;
}
func NewWorkbookSortField()(*WorkbookSortField) {
    m := &WorkbookSortField{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *WorkbookSortField) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *WorkbookSortField) GetAscending()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.ascending
    }
}
func (m *WorkbookSortField) GetColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
func (m *WorkbookSortField) GetDataOption()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dataOption
    }
}
func (m *WorkbookSortField) GetIcon()(*WorkbookIcon) {
    if m == nil {
        return nil
    } else {
        return m.icon
    }
}
func (m *WorkbookSortField) GetKey()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
func (m *WorkbookSortField) GetSortOn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sortOn
    }
}
func (m *WorkbookSortField) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["ascending"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAscending(val)
        return nil
    }
    res["color"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetColor(val)
        return nil
    }
    res["dataOption"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDataOption(val)
        return nil
    }
    res["icon"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookIcon() })
        if err != nil {
            return err
        }
        m.SetIcon(val.(*WorkbookIcon))
        return nil
    }
    res["key"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetKey(val)
        return nil
    }
    res["sortOn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSortOn(val)
        return nil
    }
    return res
}
func (m *WorkbookSortField) IsNil()(bool) {
    return m == nil
}
func (m *WorkbookSortField) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("ascending", m.GetAscending())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dataOption", m.GetDataOption())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("icon", m.GetIcon())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sortOn", m.GetSortOn())
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
func (m *WorkbookSortField) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *WorkbookSortField) SetAscending(value *bool)() {
    m.ascending = value
}
func (m *WorkbookSortField) SetColor(value *string)() {
    m.color = value
}
func (m *WorkbookSortField) SetDataOption(value *string)() {
    m.dataOption = value
}
func (m *WorkbookSortField) SetIcon(value *WorkbookIcon)() {
    m.icon = value
}
func (m *WorkbookSortField) SetKey(value *int32)() {
    m.key = value
}
func (m *WorkbookSortField) SetSortOn(value *string)() {
    m.sortOn = value
}