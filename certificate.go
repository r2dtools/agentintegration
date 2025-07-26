package agentintegration

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

type Issuer struct {
	CN           string
	Organization []string
}

type CertificateIssueRequestData struct {
	Email,
	ServerName,
	WebServer,
	ChallengeType string
	Subjects         []string
	AdditionalParams map[string]string
	Assign           bool
	PreventReload    bool
}

func (ctd *CertificateIssueRequestData) GetAdditionalParam(key string) string {
	if ctd.AdditionalParams == nil {
		ctd.AdditionalParams = make(map[string]string)
	}

	return ctd.AdditionalParams[key]
}

type CertificateUploadRequestData struct {
	ServerName     string
	WebServer      string
	CertName       string
	PemCertificate string
}

type CertificateAssignRequestData struct {
	ServerName  string
	WebServer   string
	CertName    string
	StorageType string
}

type CertificateDownloadRequestData struct {
	CertName    string
	StorageType string
}

type CertificateRemoveRequestData struct {
	CertName    string
	StorageType string
}

type CertificateInfoRequestData struct {
	CertName    string
	StorageType string
}

type CertificatesResponseData struct {
	Certificates map[string]*Certificate
}

type CertificateDownloadResponseData struct {
	CertFileName string
	CertContent  string
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
