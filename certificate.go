package agentintegration

// Certificate represents certificate data
type Certificate struct {
	CN,
	ValidFrom,
	ValidTo string
	DNSNames,
	EmailAddresses, Organization []string
	IsCA, IsValid bool
	Issuer        Issuer
}

// Issuer represents base information about certificate issuer
type Issuer struct {
	CN           string
	Organization []string
}

// CertificateIssueRequestData contains data required to issue a certificate
type CertificateIssueRequestData struct {
	Email,
	ServerName,
	DocRoot,
	WebServer,
	ChallengeType string
	Subjects         []string
	AdditionalParams map[string]string
}

// GetAdditionalParam returns additional param
func (ctd *CertificateIssueRequestData) GetAdditionalParam(key string) string {
	if ctd.AdditionalParams == nil {
		ctd.AdditionalParams = make(map[string]string)
	}

	return ctd.AdditionalParams[key]
}

// CertificateUploadRequestData contains data required to upload a certificate
type CertificateUploadRequestData struct {
	ServerName,
	WebServer,
	PemCertificate string
}

type StorageCertificateNameList struct {
	CertNameList []string
}
