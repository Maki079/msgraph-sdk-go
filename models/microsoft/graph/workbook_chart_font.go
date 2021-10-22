package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookChartFont struct {
    Entity
    bold *bool;
    color *string;
    italic *bool;
    name *string;
    size *float64;
    underline *string;
}
func NewWorkbookChartFont()(*WorkbookChartFont) {
    m := &WorkbookChartFont{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WorkbookChartFont) GetBold()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.bold
    }
}
func (m *WorkbookChartFont) GetColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
func (m *WorkbookChartFont) GetItalic()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.italic
    }
}
func (m *WorkbookChartFont) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *WorkbookChartFont) GetSize()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
func (m *WorkbookChartFont) GetUnderline()(*string) {
    if m == nil {
        return nil
    } else {
        return m.underline
    }
}
func (m *WorkbookChartFont) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["bold"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetBold(val)
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
    res["italic"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetItalic(val)
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
    res["size"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetSize(val)
        return nil
    }
    res["underline"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUnderline(val)
        return nil
    }
    return res
}
func (m *WorkbookChartFont) IsNil()(bool) {
    return m == nil
}
func (m *WorkbookChartFont) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("bold", m.GetBold())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("italic", m.GetItalic())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("underline", m.GetUnderline())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WorkbookChartFont) SetBold(value *bool)() {
    m.bold = value
}
func (m *WorkbookChartFont) SetColor(value *string)() {
    m.color = value
}
func (m *WorkbookChartFont) SetItalic(value *bool)() {
    m.italic = value
}
func (m *WorkbookChartFont) SetName(value *string)() {
    m.name = value
}
func (m *WorkbookChartFont) SetSize(value *float64)() {
    m.size = value
}
func (m *WorkbookChartFont) SetUnderline(value *string)() {
    m.underline = value
}