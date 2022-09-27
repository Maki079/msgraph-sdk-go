package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintTaskTrigger provides operations to manage the collection of agreementAcceptance entities.
type PrintTaskTrigger struct {
    Entity
    // The definition property
    definition PrintTaskDefinitionable
    // The event property
    event *PrintEvent
}
// NewPrintTaskTrigger instantiates a new printTaskTrigger and sets the default values.
func NewPrintTaskTrigger()(*PrintTaskTrigger) {
    m := &PrintTaskTrigger{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.printTaskTrigger";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePrintTaskTriggerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintTaskTriggerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintTaskTrigger(), nil
}
// GetDefinition gets the definition property value. The definition property
func (m *PrintTaskTrigger) GetDefinition()(PrintTaskDefinitionable) {
    return m.definition
}
// GetEvent gets the event property value. The event property
func (m *PrintTaskTrigger) GetEvent()(*PrintEvent) {
    return m.event
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintTaskTrigger) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["definition"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePrintTaskDefinitionFromDiscriminatorValue , m.SetDefinition)
    res["event"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParsePrintEvent , m.SetEvent)
    return res
}
// Serialize serializes information the current object
func (m *PrintTaskTrigger) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("definition", m.GetDefinition())
        if err != nil {
            return err
        }
    }
    if m.GetEvent() != nil {
        cast := (*m.GetEvent()).String()
        err = writer.WriteStringValue("event", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefinition sets the definition property value. The definition property
func (m *PrintTaskTrigger) SetDefinition(value PrintTaskDefinitionable)() {
    m.definition = value
}
// SetEvent sets the event property value. The event property
func (m *PrintTaskTrigger) SetEvent(value *PrintEvent)() {
    m.event = value
}
