package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftAuthenticatorAuthenticationMethodConfiguration 
type MicrosoftAuthenticatorAuthenticationMethodConfiguration struct {
    AuthenticationMethodConfiguration
    // A collection of users or groups who are enabled to use the authentication method.
    includeTargets []MicrosoftAuthenticatorAuthenticationMethodTargetable
}
// NewMicrosoftAuthenticatorAuthenticationMethodConfiguration instantiates a new MicrosoftAuthenticatorAuthenticationMethodConfiguration and sets the default values.
func NewMicrosoftAuthenticatorAuthenticationMethodConfiguration()(*MicrosoftAuthenticatorAuthenticationMethodConfiguration) {
    m := &MicrosoftAuthenticatorAuthenticationMethodConfiguration{
        AuthenticationMethodConfiguration: *NewAuthenticationMethodConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.microsoftAuthenticatorAuthenticationMethodConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMicrosoftAuthenticatorAuthenticationMethodConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftAuthenticatorAuthenticationMethodConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftAuthenticatorAuthenticationMethodConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftAuthenticatorAuthenticationMethodConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethodConfiguration.GetFieldDeserializers()
    res["includeTargets"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMicrosoftAuthenticatorAuthenticationMethodTargetFromDiscriminatorValue , m.SetIncludeTargets)
    return res
}
// GetIncludeTargets gets the includeTargets property value. A collection of users or groups who are enabled to use the authentication method.
func (m *MicrosoftAuthenticatorAuthenticationMethodConfiguration) GetIncludeTargets()([]MicrosoftAuthenticatorAuthenticationMethodTargetable) {
    return m.includeTargets
}
// Serialize serializes information the current object
func (m *MicrosoftAuthenticatorAuthenticationMethodConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationMethodConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetIncludeTargets() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetIncludeTargets())
        err = writer.WriteCollectionOfObjectValues("includeTargets", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIncludeTargets sets the includeTargets property value. A collection of users or groups who are enabled to use the authentication method.
func (m *MicrosoftAuthenticatorAuthenticationMethodConfiguration) SetIncludeTargets(value []MicrosoftAuthenticatorAuthenticationMethodTargetable)() {
    m.includeTargets = value
}
