package agentintegration

// Certificate represents certificate data
type Certificate struct {
	CN,
	ValidFrom,
	ValidTo string
	DNSNames,
	EmailAddresses,
	Organization,
	Province,
	Country,
	Locality []string
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
	WebServer,
	ChallengeType string
	Subjects         []string
	AdditionalParams map[string]string
	Assign           bool
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
	CertName,
	PemCertificate string
}

// CertificateAssignRequestData contains data required to assign a certificate to domain
type CertificateAssignRequestData struct {
	ServerName,
	WebServer,
	CertName string
}

type CertificatesResponseData struct {
	Certificates map[string]*Certificate
}

type CertificateDownloadResponseData struct {
	CertFileName,
	CertContent string
}

type CommonDirStatusRequestData struct {
	WebServer  string
	ServerName string
}

type CommonDirStatusResponseData struct {
	Status bool
}

type CommonDirChangeStatusRequestData struct {
	WebServer  string
	ServerName string
	Status     bool
}
