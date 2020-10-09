package agentintegration

// Certificate represents certificate data
type Certificate struct {
	CN,
	ValidFrom,
	ValidTo string
	DNSNames,
	EmailAddresses, Organization []string
	IsCA   bool
	Issuer Issuer
}

// Issuer represents base information about certificate issuer
type Issuer struct {
	CN           string
	Organization []string
}
