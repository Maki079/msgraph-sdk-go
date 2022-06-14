package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows81GeneralConfigurationable 
type Windows81GeneralConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountsBlockAddingNonMicrosoftAccountEmail()(*bool)
    GetApplyOnlyToWindows81()(*bool)
    GetBrowserBlockAutofill()(*bool)
    GetBrowserBlockAutomaticDetectionOfIntranetSites()(*bool)
    GetBrowserBlockEnterpriseModeAccess()(*bool)
    GetBrowserBlockJavaScript()(*bool)
    GetBrowserBlockPlugins()(*bool)
    GetBrowserBlockPopups()(*bool)
    GetBrowserBlockSendingDoNotTrackHeader()(*bool)
    GetBrowserBlockSingleWordEntryOnIntranetSites()(*bool)
    GetBrowserEnterpriseModeSiteListLocation()(*string)
    GetBrowserInternetSecurityLevel()(*InternetSiteSecurityLevel)
    GetBrowserIntranetSecurityLevel()(*SiteSecurityLevel)
    GetBrowserLoggingReportLocation()(*string)
    GetBrowserRequireFirewall()(*bool)
    GetBrowserRequireFraudWarning()(*bool)
    GetBrowserRequireHighSecurityForRestrictedSites()(*bool)
    GetBrowserRequireSmartScreen()(*bool)
    GetBrowserTrustedSitesSecurityLevel()(*SiteSecurityLevel)
    GetCellularBlockDataRoaming()(*bool)
    GetDiagnosticsBlockDataSubmission()(*bool)
    GetPasswordBlockPicturePasswordAndPin()(*bool)
    GetPasswordExpirationDays()(*int32)
    GetPasswordMinimumCharacterSetCount()(*int32)
    GetPasswordMinimumLength()(*int32)
    GetPasswordMinutesOfInactivityBeforeScreenTimeout()(*int32)
    GetPasswordPreviousPasswordBlockCount()(*int32)
    GetPasswordRequiredType()(*RequiredPasswordType)
    GetPasswordSignInFailureCountBeforeFactoryReset()(*int32)
    GetStorageRequireDeviceEncryption()(*bool)
    GetUpdatesRequireAutomaticUpdates()(*bool)
    GetUserAccountControlSettings()(*WindowsUserAccountControlSettings)
    GetWorkFoldersUrl()(*string)
    SetAccountsBlockAddingNonMicrosoftAccountEmail(value *bool)()
    SetApplyOnlyToWindows81(value *bool)()
    SetBrowserBlockAutofill(value *bool)()
    SetBrowserBlockAutomaticDetectionOfIntranetSites(value *bool)()
    SetBrowserBlockEnterpriseModeAccess(value *bool)()
    SetBrowserBlockJavaScript(value *bool)()
    SetBrowserBlockPlugins(value *bool)()
    SetBrowserBlockPopups(value *bool)()
    SetBrowserBlockSendingDoNotTrackHeader(value *bool)()
    SetBrowserBlockSingleWordEntryOnIntranetSites(value *bool)()
    SetBrowserEnterpriseModeSiteListLocation(value *string)()
    SetBrowserInternetSecurityLevel(value *InternetSiteSecurityLevel)()
    SetBrowserIntranetSecurityLevel(value *SiteSecurityLevel)()
    SetBrowserLoggingReportLocation(value *string)()
    SetBrowserRequireFirewall(value *bool)()
    SetBrowserRequireFraudWarning(value *bool)()
    SetBrowserRequireHighSecurityForRestrictedSites(value *bool)()
    SetBrowserRequireSmartScreen(value *bool)()
    SetBrowserTrustedSitesSecurityLevel(value *SiteSecurityLevel)()
    SetCellularBlockDataRoaming(value *bool)()
    SetDiagnosticsBlockDataSubmission(value *bool)()
    SetPasswordBlockPicturePasswordAndPin(value *bool)()
    SetPasswordExpirationDays(value *int32)()
    SetPasswordMinimumCharacterSetCount(value *int32)()
    SetPasswordMinimumLength(value *int32)()
    SetPasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)()
    SetPasswordPreviousPasswordBlockCount(value *int32)()
    SetPasswordRequiredType(value *RequiredPasswordType)()
    SetPasswordSignInFailureCountBeforeFactoryReset(value *int32)()
    SetStorageRequireDeviceEncryption(value *bool)()
    SetUpdatesRequireAutomaticUpdates(value *bool)()
    SetUserAccountControlSettings(value *WindowsUserAccountControlSettings)()
    SetWorkFoldersUrl(value *string)()
}
