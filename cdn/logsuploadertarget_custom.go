// Custom code. This file is not generated and is preserved across codegen runs.
// It adds the GetAccountKey and GetToken getters on the Azure Blob auth config
// unions. Generated code calls them, but the generator does not emit them.

package cdn

// GetAccountKey returns the account key if the AccountKey variant is set
func (u *LogsUploaderTargetNewParamsConfigAzureBlobConfigAuthConfigUnion) GetAccountKey() *string {
	if u.OfLogsUploaderTargetNewsConfigAzureBlobConfigAuthConfigAccountKey != nil {
		return &u.OfLogsUploaderTargetNewsConfigAzureBlobConfigAuthConfigAccountKey.AccountKey
	}
	return nil
}

// GetToken returns the token if the Token variant is set
func (u *LogsUploaderTargetNewParamsConfigAzureBlobConfigAuthConfigUnion) GetToken() *string {
	if u.OfLogsUploaderTargetNewsConfigAzureBlobConfigAuthConfigToken != nil {
		return (*string)(&u.OfLogsUploaderTargetNewsConfigAzureBlobConfigAuthConfigToken.Token)
	}
	return nil
}

// GetAccountKey returns the account key if the AccountKey variant is set
func (u *LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigUnion) GetAccountKey() *string {
	if u.OfLogsUploaderTargetUpdatesConfigAzureBlobConfigAuthConfigAccountKey != nil {
		return &u.OfLogsUploaderTargetUpdatesConfigAzureBlobConfigAuthConfigAccountKey.AccountKey
	}
	return nil
}

// GetToken returns the token if the Token variant is set
func (u *LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigUnion) GetToken() *string {
	if u.OfLogsUploaderTargetUpdatesConfigAzureBlobConfigAuthConfigToken != nil {
		return (*string)(&u.OfLogsUploaderTargetUpdatesConfigAzureBlobConfigAuthConfigToken.Token)
	}
	return nil
}

// GetAccountKey returns the account key if the AccountKey variant is set
func (u *LogsUploaderTargetReplaceParamsConfigAzureBlobConfigAuthConfigUnion) GetAccountKey() *string {
	if u.OfLogsUploaderTargetReplacesConfigAzureBlobConfigAuthConfigAccountKey != nil {
		return &u.OfLogsUploaderTargetReplacesConfigAzureBlobConfigAuthConfigAccountKey.AccountKey
	}
	return nil
}

// GetToken returns the token if the Token variant is set
func (u *LogsUploaderTargetReplaceParamsConfigAzureBlobConfigAuthConfigUnion) GetToken() *string {
	if u.OfLogsUploaderTargetReplacesConfigAzureBlobConfigAuthConfigToken != nil {
		return (*string)(&u.OfLogsUploaderTargetReplacesConfigAzureBlobConfigAuthConfigToken.Token)
	}
	return nil
}
