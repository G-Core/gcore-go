// Custom code. This file is not generated and is preserved across codegen runs.
// It adds the GetAccountKey and GetToken getters on the Azure Blob auth config
// union of the update params. Generated code calls them, but the generator does
// not emit them.

package cdn

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
