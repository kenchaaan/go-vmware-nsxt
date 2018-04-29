/*
 * NSX API
 *
 * VMware NSX REST API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package loadbalancer

type ClientSslProfileBinding struct {

	// authentication depth is used to set the verification depth in the client certificates chain.
	CertificateChainDepth int64 `json:"certificate_chain_depth,omitempty"`

	// client authentication mode
	ClientAuth string `json:"client_auth,omitempty"`

	// If client auth type is REQUIRED, client certificate must be signed by one of the trusted Certificate Authorities (CAs), also referred to as root CAs, whose self signed certificates are specified.
	ClientAuthCaIds []string `json:"client_auth_ca_ids,omitempty"`

	// A Certificate Revocation List (CRL) can be specified in the client-side SSL profile binding to disallow compromised client certificates.
	ClientAuthCrlIds []string `json:"client_auth_crl_ids,omitempty"`

	// A default certificate should be specified which will be used if the server does not host multiple hostnames on the same IP address or if the client does not support SNI extension.
	DefaultCertificateId string `json:"default_certificate_id"`

	// Client-side SSL profile binding allows multiple certificates, for different hostnames, to be bound to the same virtual server.
	SniCertificateIds []string `json:"sni_certificate_ids,omitempty"`

	// Client SSL profile defines reusable, application-independent client side SSL properties.
	SslProfileId string `json:"ssl_profile_id,omitempty"`
}
