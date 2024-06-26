// Code generated by dto builder generator; DO NOT EDIT.

package sdk

import ()

func NewCreateSaml2SecurityIntegrationRequest(
	name AccountObjectIdentifier,
	Enabled bool,
	Saml2Issuer string,
	Saml2SsoUrl string,
	Saml2Provider string,
	Saml2X509Cert string,
) *CreateSaml2SecurityIntegrationRequest {
	s := CreateSaml2SecurityIntegrationRequest{}
	s.name = name
	s.Enabled = Enabled
	s.Saml2Issuer = Saml2Issuer
	s.Saml2SsoUrl = Saml2SsoUrl
	s.Saml2Provider = Saml2Provider
	s.Saml2X509Cert = Saml2X509Cert
	return &s
}

func (s *CreateSaml2SecurityIntegrationRequest) WithOrReplace(OrReplace *bool) *CreateSaml2SecurityIntegrationRequest {
	s.OrReplace = OrReplace
	return s
}

func (s *CreateSaml2SecurityIntegrationRequest) WithIfNotExists(IfNotExists *bool) *CreateSaml2SecurityIntegrationRequest {
	s.IfNotExists = IfNotExists
	return s
}

func (s *CreateSaml2SecurityIntegrationRequest) WithAllowedUserDomains(AllowedUserDomains []UserDomain) *CreateSaml2SecurityIntegrationRequest {
	s.AllowedUserDomains = AllowedUserDomains
	return s
}

func (s *CreateSaml2SecurityIntegrationRequest) WithAllowedEmailPatterns(AllowedEmailPatterns []EmailPattern) *CreateSaml2SecurityIntegrationRequest {
	s.AllowedEmailPatterns = AllowedEmailPatterns
	return s
}

func (s *CreateSaml2SecurityIntegrationRequest) WithSaml2SpInitiatedLoginPageLabel(Saml2SpInitiatedLoginPageLabel *string) *CreateSaml2SecurityIntegrationRequest {
	s.Saml2SpInitiatedLoginPageLabel = Saml2SpInitiatedLoginPageLabel
	return s
}

func (s *CreateSaml2SecurityIntegrationRequest) WithSaml2EnableSpInitiated(Saml2EnableSpInitiated *bool) *CreateSaml2SecurityIntegrationRequest {
	s.Saml2EnableSpInitiated = Saml2EnableSpInitiated
	return s
}

func (s *CreateSaml2SecurityIntegrationRequest) WithSaml2SnowflakeX509Cert(Saml2SnowflakeX509Cert *string) *CreateSaml2SecurityIntegrationRequest {
	s.Saml2SnowflakeX509Cert = Saml2SnowflakeX509Cert
	return s
}

func (s *CreateSaml2SecurityIntegrationRequest) WithSaml2SignRequest(Saml2SignRequest *bool) *CreateSaml2SecurityIntegrationRequest {
	s.Saml2SignRequest = Saml2SignRequest
	return s
}

func (s *CreateSaml2SecurityIntegrationRequest) WithSaml2RequestedNameidFormat(Saml2RequestedNameidFormat *string) *CreateSaml2SecurityIntegrationRequest {
	s.Saml2RequestedNameidFormat = Saml2RequestedNameidFormat
	return s
}

func (s *CreateSaml2SecurityIntegrationRequest) WithSaml2PostLogoutRedirectUrl(Saml2PostLogoutRedirectUrl *string) *CreateSaml2SecurityIntegrationRequest {
	s.Saml2PostLogoutRedirectUrl = Saml2PostLogoutRedirectUrl
	return s
}

func (s *CreateSaml2SecurityIntegrationRequest) WithSaml2ForceAuthn(Saml2ForceAuthn *bool) *CreateSaml2SecurityIntegrationRequest {
	s.Saml2ForceAuthn = Saml2ForceAuthn
	return s
}

func (s *CreateSaml2SecurityIntegrationRequest) WithSaml2SnowflakeIssuerUrl(Saml2SnowflakeIssuerUrl *string) *CreateSaml2SecurityIntegrationRequest {
	s.Saml2SnowflakeIssuerUrl = Saml2SnowflakeIssuerUrl
	return s
}

func (s *CreateSaml2SecurityIntegrationRequest) WithSaml2SnowflakeAcsUrl(Saml2SnowflakeAcsUrl *string) *CreateSaml2SecurityIntegrationRequest {
	s.Saml2SnowflakeAcsUrl = Saml2SnowflakeAcsUrl
	return s
}

func (s *CreateSaml2SecurityIntegrationRequest) WithComment(Comment *string) *CreateSaml2SecurityIntegrationRequest {
	s.Comment = Comment
	return s
}

func NewCreateScimSecurityIntegrationRequest(
	name AccountObjectIdentifier,
	Enabled bool,
	ScimClient ScimSecurityIntegrationScimClientOption,
	RunAsRole ScimSecurityIntegrationRunAsRoleOption,
) *CreateScimSecurityIntegrationRequest {
	s := CreateScimSecurityIntegrationRequest{}
	s.name = name
	s.Enabled = Enabled
	s.ScimClient = ScimClient
	s.RunAsRole = RunAsRole
	return &s
}

func (s *CreateScimSecurityIntegrationRequest) WithOrReplace(OrReplace *bool) *CreateScimSecurityIntegrationRequest {
	s.OrReplace = OrReplace
	return s
}

func (s *CreateScimSecurityIntegrationRequest) WithIfNotExists(IfNotExists *bool) *CreateScimSecurityIntegrationRequest {
	s.IfNotExists = IfNotExists
	return s
}

func (s *CreateScimSecurityIntegrationRequest) WithNetworkPolicy(NetworkPolicy *AccountObjectIdentifier) *CreateScimSecurityIntegrationRequest {
	s.NetworkPolicy = NetworkPolicy
	return s
}

func (s *CreateScimSecurityIntegrationRequest) WithSyncPassword(SyncPassword *bool) *CreateScimSecurityIntegrationRequest {
	s.SyncPassword = SyncPassword
	return s
}

func (s *CreateScimSecurityIntegrationRequest) WithComment(Comment *string) *CreateScimSecurityIntegrationRequest {
	s.Comment = Comment
	return s
}

func NewAlterSaml2SecurityIntegrationRequest(
	name AccountObjectIdentifier,
) *AlterSaml2SecurityIntegrationRequest {
	s := AlterSaml2SecurityIntegrationRequest{}
	s.name = name
	return &s
}

func (s *AlterSaml2SecurityIntegrationRequest) WithIfExists(IfExists *bool) *AlterSaml2SecurityIntegrationRequest {
	s.IfExists = IfExists
	return s
}

func (s *AlterSaml2SecurityIntegrationRequest) WithSetTags(SetTags []TagAssociation) *AlterSaml2SecurityIntegrationRequest {
	s.SetTags = SetTags
	return s
}

func (s *AlterSaml2SecurityIntegrationRequest) WithUnsetTags(UnsetTags []ObjectIdentifier) *AlterSaml2SecurityIntegrationRequest {
	s.UnsetTags = UnsetTags
	return s
}

func (s *AlterSaml2SecurityIntegrationRequest) WithSet(Set *Saml2IntegrationSetRequest) *AlterSaml2SecurityIntegrationRequest {
	s.Set = Set
	return s
}

func (s *AlterSaml2SecurityIntegrationRequest) WithUnset(Unset *Saml2IntegrationUnsetRequest) *AlterSaml2SecurityIntegrationRequest {
	s.Unset = Unset
	return s
}

func (s *AlterSaml2SecurityIntegrationRequest) WithRefreshSaml2SnowflakePrivateKey(RefreshSaml2SnowflakePrivateKey *bool) *AlterSaml2SecurityIntegrationRequest {
	s.RefreshSaml2SnowflakePrivateKey = RefreshSaml2SnowflakePrivateKey
	return s
}

func NewSaml2IntegrationSetRequest() *Saml2IntegrationSetRequest {
	return &Saml2IntegrationSetRequest{}
}

func (s *Saml2IntegrationSetRequest) WithEnabled(Enabled *bool) *Saml2IntegrationSetRequest {
	s.Enabled = Enabled
	return s
}

func (s *Saml2IntegrationSetRequest) WithSaml2Issuer(Saml2Issuer *string) *Saml2IntegrationSetRequest {
	s.Saml2Issuer = Saml2Issuer
	return s
}

func (s *Saml2IntegrationSetRequest) WithSaml2SsoUrl(Saml2SsoUrl *string) *Saml2IntegrationSetRequest {
	s.Saml2SsoUrl = Saml2SsoUrl
	return s
}

func (s *Saml2IntegrationSetRequest) WithSaml2Provider(Saml2Provider *string) *Saml2IntegrationSetRequest {
	s.Saml2Provider = Saml2Provider
	return s
}

func (s *Saml2IntegrationSetRequest) WithSaml2X509Cert(Saml2X509Cert *string) *Saml2IntegrationSetRequest {
	s.Saml2X509Cert = Saml2X509Cert
	return s
}

func (s *Saml2IntegrationSetRequest) WithAllowedUserDomains(AllowedUserDomains []UserDomain) *Saml2IntegrationSetRequest {
	s.AllowedUserDomains = AllowedUserDomains
	return s
}

func (s *Saml2IntegrationSetRequest) WithAllowedEmailPatterns(AllowedEmailPatterns []EmailPattern) *Saml2IntegrationSetRequest {
	s.AllowedEmailPatterns = AllowedEmailPatterns
	return s
}

func (s *Saml2IntegrationSetRequest) WithSaml2SpInitiatedLoginPageLabel(Saml2SpInitiatedLoginPageLabel *string) *Saml2IntegrationSetRequest {
	s.Saml2SpInitiatedLoginPageLabel = Saml2SpInitiatedLoginPageLabel
	return s
}

func (s *Saml2IntegrationSetRequest) WithSaml2EnableSpInitiated(Saml2EnableSpInitiated *bool) *Saml2IntegrationSetRequest {
	s.Saml2EnableSpInitiated = Saml2EnableSpInitiated
	return s
}

func (s *Saml2IntegrationSetRequest) WithSaml2SnowflakeX509Cert(Saml2SnowflakeX509Cert *string) *Saml2IntegrationSetRequest {
	s.Saml2SnowflakeX509Cert = Saml2SnowflakeX509Cert
	return s
}

func (s *Saml2IntegrationSetRequest) WithSaml2SignRequest(Saml2SignRequest *bool) *Saml2IntegrationSetRequest {
	s.Saml2SignRequest = Saml2SignRequest
	return s
}

func (s *Saml2IntegrationSetRequest) WithSaml2RequestedNameidFormat(Saml2RequestedNameidFormat *string) *Saml2IntegrationSetRequest {
	s.Saml2RequestedNameidFormat = Saml2RequestedNameidFormat
	return s
}

func (s *Saml2IntegrationSetRequest) WithSaml2PostLogoutRedirectUrl(Saml2PostLogoutRedirectUrl *string) *Saml2IntegrationSetRequest {
	s.Saml2PostLogoutRedirectUrl = Saml2PostLogoutRedirectUrl
	return s
}

func (s *Saml2IntegrationSetRequest) WithSaml2ForceAuthn(Saml2ForceAuthn *bool) *Saml2IntegrationSetRequest {
	s.Saml2ForceAuthn = Saml2ForceAuthn
	return s
}

func (s *Saml2IntegrationSetRequest) WithSaml2SnowflakeIssuerUrl(Saml2SnowflakeIssuerUrl *string) *Saml2IntegrationSetRequest {
	s.Saml2SnowflakeIssuerUrl = Saml2SnowflakeIssuerUrl
	return s
}

func (s *Saml2IntegrationSetRequest) WithSaml2SnowflakeAcsUrl(Saml2SnowflakeAcsUrl *string) *Saml2IntegrationSetRequest {
	s.Saml2SnowflakeAcsUrl = Saml2SnowflakeAcsUrl
	return s
}

func (s *Saml2IntegrationSetRequest) WithComment(Comment *string) *Saml2IntegrationSetRequest {
	s.Comment = Comment
	return s
}

func NewSaml2IntegrationUnsetRequest() *Saml2IntegrationUnsetRequest {
	return &Saml2IntegrationUnsetRequest{}
}

func (s *Saml2IntegrationUnsetRequest) WithSaml2ForceAuthn(Saml2ForceAuthn *bool) *Saml2IntegrationUnsetRequest {
	s.Saml2ForceAuthn = Saml2ForceAuthn
	return s
}

func (s *Saml2IntegrationUnsetRequest) WithSaml2RequestedNameidFormat(Saml2RequestedNameidFormat *bool) *Saml2IntegrationUnsetRequest {
	s.Saml2RequestedNameidFormat = Saml2RequestedNameidFormat
	return s
}

func (s *Saml2IntegrationUnsetRequest) WithSaml2PostLogoutRedirectUrl(Saml2PostLogoutRedirectUrl *bool) *Saml2IntegrationUnsetRequest {
	s.Saml2PostLogoutRedirectUrl = Saml2PostLogoutRedirectUrl
	return s
}

func (s *Saml2IntegrationUnsetRequest) WithComment(Comment *bool) *Saml2IntegrationUnsetRequest {
	s.Comment = Comment
	return s
}

func NewAlterScimSecurityIntegrationRequest(
	name AccountObjectIdentifier,
) *AlterScimSecurityIntegrationRequest {
	s := AlterScimSecurityIntegrationRequest{}
	s.name = name
	return &s
}

func (s *AlterScimSecurityIntegrationRequest) WithIfExists(IfExists *bool) *AlterScimSecurityIntegrationRequest {
	s.IfExists = IfExists
	return s
}

func (s *AlterScimSecurityIntegrationRequest) WithSetTags(SetTags []TagAssociation) *AlterScimSecurityIntegrationRequest {
	s.SetTags = SetTags
	return s
}

func (s *AlterScimSecurityIntegrationRequest) WithUnsetTags(UnsetTags []ObjectIdentifier) *AlterScimSecurityIntegrationRequest {
	s.UnsetTags = UnsetTags
	return s
}

func (s *AlterScimSecurityIntegrationRequest) WithSet(Set *ScimIntegrationSetRequest) *AlterScimSecurityIntegrationRequest {
	s.Set = Set
	return s
}

func (s *AlterScimSecurityIntegrationRequest) WithUnset(Unset *ScimIntegrationUnsetRequest) *AlterScimSecurityIntegrationRequest {
	s.Unset = Unset
	return s
}

func NewScimIntegrationSetRequest() *ScimIntegrationSetRequest {
	return &ScimIntegrationSetRequest{}
}

func (s *ScimIntegrationSetRequest) WithEnabled(Enabled *bool) *ScimIntegrationSetRequest {
	s.Enabled = Enabled
	return s
}

func (s *ScimIntegrationSetRequest) WithNetworkPolicy(NetworkPolicy *AccountObjectIdentifier) *ScimIntegrationSetRequest {
	s.NetworkPolicy = NetworkPolicy
	return s
}

func (s *ScimIntegrationSetRequest) WithSyncPassword(SyncPassword *bool) *ScimIntegrationSetRequest {
	s.SyncPassword = SyncPassword
	return s
}

func (s *ScimIntegrationSetRequest) WithComment(Comment *string) *ScimIntegrationSetRequest {
	s.Comment = Comment
	return s
}

func NewScimIntegrationUnsetRequest() *ScimIntegrationUnsetRequest {
	return &ScimIntegrationUnsetRequest{}
}

func (s *ScimIntegrationUnsetRequest) WithEnabled(Enabled *bool) *ScimIntegrationUnsetRequest {
	s.Enabled = Enabled
	return s
}

func (s *ScimIntegrationUnsetRequest) WithNetworkPolicy(NetworkPolicy *bool) *ScimIntegrationUnsetRequest {
	s.NetworkPolicy = NetworkPolicy
	return s
}

func (s *ScimIntegrationUnsetRequest) WithSyncPassword(SyncPassword *bool) *ScimIntegrationUnsetRequest {
	s.SyncPassword = SyncPassword
	return s
}

func (s *ScimIntegrationUnsetRequest) WithComment(Comment *bool) *ScimIntegrationUnsetRequest {
	s.Comment = Comment
	return s
}

func NewDropSecurityIntegrationRequest(
	name AccountObjectIdentifier,
) *DropSecurityIntegrationRequest {
	s := DropSecurityIntegrationRequest{}
	s.name = name
	return &s
}

func (s *DropSecurityIntegrationRequest) WithIfExists(IfExists *bool) *DropSecurityIntegrationRequest {
	s.IfExists = IfExists
	return s
}

func NewDescribeSecurityIntegrationRequest(
	name AccountObjectIdentifier,
) *DescribeSecurityIntegrationRequest {
	s := DescribeSecurityIntegrationRequest{}
	s.name = name
	return &s
}

func NewShowSecurityIntegrationRequest() *ShowSecurityIntegrationRequest {
	return &ShowSecurityIntegrationRequest{}
}

func (s *ShowSecurityIntegrationRequest) WithLike(Like *Like) *ShowSecurityIntegrationRequest {
	s.Like = Like
	return s
}
