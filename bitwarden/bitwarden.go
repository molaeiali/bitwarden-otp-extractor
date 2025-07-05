package bitwarden

type BitwardenFolder struct {
	Id	string `json:"id"`
	Name	string `json:"name"`
}

type BitwardenItem struct {
	PasswordHistory	[]BitwardenItemPasswordHistory `json:"passwordHistory"`
	RevisionDate	string `json:"revisionDate"`
	CreationDate	string `json:"creationDate"`
	DeletedDate	string `json:"deletedDate"`
	Id	string `json:"id"`
	OrganizationId	string `json:"organizationId"`
	FolderId	string `json:"folderId"`
	Type	int64 `json:"type"`
	Reprompt	int64 `json:"reprompt"`
	Name	string `json:"name"`
	Notes	string `json:"notes"`
	Favorite	bool `json:"favorite"`
	Login	BitwardenItemLogin `json:"login"`
	CollectionIds	string `json:"collectionIds"`
}

type BitwardenItemPasswordHistory struct {
	LastUsedDate string `json:"lastUsedDate"`
	Password string `json:"password"`
}

type BitwardenItemLogin struct {
	Fido2Credentials	[]BitwardenItemLoginFido2Credentials `json:"fido2Credentials"`
	Uris	[]BitwardenItemLoginUri `json:"uris"`
	Username	string `json:"username"`
	Password	string `json:"password"`
	Totp	string `json:"totp"`
}

type BitwardenItemLoginUri struct {
	Match string `json:"match"`
	Uri string `json:"uri"`
}

type BitwardenItemLoginFido2Credentials struct {
	CredentialId string `json:"credentialId"`
	KeyType string `json:"keyType"`
	KeyAlgorithm string `json:"keyAlgorithm"`
	KeyCurve string `json:"keyCurve"`
	KeyValue string `json:"keyValue"`
	RpId string `json:"rpId"`
	UserHandle string `json:"userHandle"`
	UserName string `json:"userName"`
	Counter string `json:"counter"`
	RpName string `json:"rpName"`
	UserDisplayName string `json:"userDisplayName"`
	Discoverable string `json:"discoverable"`
	CreationDate string `json:"creationDate"`
}

type Bitwarden struct {
	Encrypted	bool `json:"encrypted"`
	Folders	[]BitwardenFolder	`json:"folders"`
	Items	[]BitwardenItem `json:"items"`
}
