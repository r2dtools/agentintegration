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
	Email, ServerName string
	Subjects          []string
}
