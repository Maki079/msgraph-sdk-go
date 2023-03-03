package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// WorkbookFilterCriteria 
type WorkbookFilterCriteria struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWorkbookFilterCriteria instantiates a new workbookFilterCriteria and sets the default values.
func NewWorkbookFilterCriteria()(*WorkbookFilterCriteria) {
    m := &WorkbookFilterCriteria{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWorkbookFilterCriteriaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookFilterCriteriaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookFilterCriteria(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbookFilterCriteria) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *WorkbookFilterCriteria) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetColor gets the color property value. The color property
func (m *WorkbookFilterCriteria) GetColor()(*string) {
    val, err := m.GetBackingStore().Get("color")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCriterion1 gets the criterion1 property value. The criterion1 property
func (m *WorkbookFilterCriteria) GetCriterion1()(*string) {
    val, err := m.GetBackingStore().Get("criterion1")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCriterion2 gets the criterion2 property value. The criterion2 property
func (m *WorkbookFilterCriteria) GetCriterion2()(*string) {
    val, err := m.GetBackingStore().Get("criterion2")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDynamicCriteria gets the dynamicCriteria property value. The dynamicCriteria property
func (m *WorkbookFilterCriteria) GetDynamicCriteria()(*string) {
    val, err := m.GetBackingStore().Get("dynamicCriteria")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookFilterCriteria) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val)
        }
        return nil
    }
    res["criterion1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriterion1(val)
        }
        return nil
    }
    res["criterion2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriterion2(val)
        }
        return nil
    }
    res["dynamicCriteria"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDynamicCriteria(val)
        }
        return nil
    }
    res["filterOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilterOn(val)
        }
        return nil
    }
    res["icon"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookIconFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIcon(val.(WorkbookIconable))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["operator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperator(val)
        }
        return nil
    }
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValues(val.(Jsonable))
        }
        return nil
    }
    return res
}
// GetFilterOn gets the filterOn property value. The filterOn property
func (m *WorkbookFilterCriteria) GetFilterOn()(*string) {
    val, err := m.GetBackingStore().Get("filterOn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIcon gets the icon property value. The icon property
func (m *WorkbookFilterCriteria) GetIcon()(WorkbookIconable) {
    val, err := m.GetBackingStore().Get("icon")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookIconable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WorkbookFilterCriteria) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOperator gets the operator property value. The operator property
func (m *WorkbookFilterCriteria) GetOperator()(*string) {
    val, err := m.GetBackingStore().Get("operator")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetValues gets the values property value. The values property
func (m *WorkbookFilterCriteria) GetValues()(Jsonable) {
    val, err := m.GetBackingStore().Get("values")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Jsonable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WorkbookFilterCriteria) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("criterion1", m.GetCriterion1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("criterion2", m.GetCriterion2())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dynamicCriteria", m.GetDynamicCriteria())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("filterOn", m.GetFilterOn())
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("operator", m.GetOperator())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("values", m.GetValues())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbookFilterCriteria) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *WorkbookFilterCriteria) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetColor sets the color property value. The color property
func (m *WorkbookFilterCriteria) SetColor(value *string)() {
    err := m.GetBackingStore().Set("color", value)
    if err != nil {
        panic(err)
    }
}
// SetCriterion1 sets the criterion1 property value. The criterion1 property
func (m *WorkbookFilterCriteria) SetCriterion1(value *string)() {
    err := m.GetBackingStore().Set("criterion1", value)
    if err != nil {
        panic(err)
    }
}
// SetCriterion2 sets the criterion2 property value. The criterion2 property
func (m *WorkbookFilterCriteria) SetCriterion2(value *string)() {
    err := m.GetBackingStore().Set("criterion2", value)
    if err != nil {
        panic(err)
    }
}
// SetDynamicCriteria sets the dynamicCriteria property value. The dynamicCriteria property
func (m *WorkbookFilterCriteria) SetDynamicCriteria(value *string)() {
    err := m.GetBackingStore().Set("dynamicCriteria", value)
    if err != nil {
        panic(err)
    }
}
// SetFilterOn sets the filterOn property value. The filterOn property
func (m *WorkbookFilterCriteria) SetFilterOn(value *string)() {
    err := m.GetBackingStore().Set("filterOn", value)
    if err != nil {
        panic(err)
    }
}
// SetIcon sets the icon property value. The icon property
func (m *WorkbookFilterCriteria) SetIcon(value WorkbookIconable)() {
    err := m.GetBackingStore().Set("icon", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WorkbookFilterCriteria) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOperator sets the operator property value. The operator property
func (m *WorkbookFilterCriteria) SetOperator(value *string)() {
    err := m.GetBackingStore().Set("operator", value)
    if err != nil {
        panic(err)
    }
}
// SetValues sets the values property value. The values property
func (m *WorkbookFilterCriteria) SetValues(value Jsonable)() {
    err := m.GetBackingStore().Set("values", value)
    if err != nil {
        panic(err)
    }
}
// WorkbookFilterCriteriaable 
type WorkbookFilterCriteriaable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetColor()(*string)
    GetCriterion1()(*string)
    GetCriterion2()(*string)
    GetDynamicCriteria()(*string)
    GetFilterOn()(*string)
    GetIcon()(WorkbookIconable)
    GetOdataType()(*string)
    GetOperator()(*string)
    GetValues()(Jsonable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetColor(value *string)()
    SetCriterion1(value *string)()
    SetCriterion2(value *string)()
    SetDynamicCriteria(value *string)()
    SetFilterOn(value *string)()
    SetIcon(value WorkbookIconable)()
    SetOdataType(value *string)()
    SetOperator(value *string)()
    SetValues(value Jsonable)()
}
