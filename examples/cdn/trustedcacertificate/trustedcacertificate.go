package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cdn"
)

// testCACertificatePEM is a self-signed CA certificate generated solely for this
// example. It contains only the public certificate (no private key), so it is
// safe to commit. Trusted CA certificates are used to verify the origin when
// CDN connects to it over HTTPS.
const testCACertificatePEM = `-----BEGIN CERTIFICATE-----
MIIDYzCCAkugAwIBAgIUNt315IS68yUNhQvSV+qPme7huXgwDQYJKoZIhvcNAQEL
BQAwQTEiMCAGA1UEAwwZR2NvcmUgU0RLIEV4YW1wbGUgVGVzdCBDQTEbMBkGA1UE
CgwSR2NvcmUgU0RLIEV4YW1wbGVzMB4XDTI2MDUxOTE1NTMyOFoXDTM2MDUxNjE1
NTMyOFowQTEiMCAGA1UEAwwZR2NvcmUgU0RLIEV4YW1wbGUgVGVzdCBDQTEbMBkG
A1UECgwSR2NvcmUgU0RLIEV4YW1wbGVzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A
MIIBCgKCAQEAt/rXbgO4GtT+t3utz6YrcsF1rRdYYxlVjtSEsbsfw7sz1IElprT1
w9+Q3n4fmg0OaEF6yQWr9K0+SUiXWgFS9twaYFRUe6M8PANLAwyoWe0BT7gqkq3n
pF53uy0CuFE0HybC/3oiGpQeVvYDUtEr9I+c/rHwbSOPXSvJ8c86HqHCxNrxB21Q
8GDHBE47Er5LVbam6Cs8MG9nT5T10TkClDRckxoxCOcTX14Nf1Me+MTcDehcZ/mZ
vViVFNUaeTCTd9YXoZK1ok8YXDeD8hfGYsHsrHVhdLNYpXWo54DodnNvevw5ilXl
7bQ67WA8e2hpY01Odt/xUZQ9IRZs7lxsjwIDAQABo1MwUTAdBgNVHQ4EFgQU+exr
mFzjx+K9muACuEFypkxi3fkwHwYDVR0jBBgwFoAU+exrmFzjx+K9muACuEFypkxi
3fkwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEATTwh2j7DbF0v
4JOrSTaLUtWo7ugcy1H2nepA0aagZWNqoMDUmY3RLbgbz7HKBWS1lqjSJRQR8bRo
8qmSTbLYfX/NgERAcAgyv9UoGaTH3g5diCgr6jgy+t5qpRE4YNW0rWOO1RuE0/ly
U3tgSULD1n24IWLAQeLpz9ADZvXCZofBopuysro8yGvvz2BreqLAjFy8jTuPOoae
71/nXhQ+ycoxSZShJ5yIip4R8hTBPMBuIb3aLD/hNVO8qIl07DX7FrWQPiTaLZrF
KI0T0vo3PB1rEcmAQl/lNwbnIkyZMOYmGFi9FA01Uvf6QPi9u0GC+4NxWUHJjLju
mNIQqB91Uw==
-----END CERTIFICATE-----`

func main() {
	// No need to pass the API key explicitly — it will automatically be read
	// from the GCORE_API_KEY environment variable if omitted.
	client := gcore.NewClient()

	// Create a trusted CA certificate
	certificate := createTrustedCACertificate(&client)

	// List trusted CA certificates
	listTrustedCACertificates(&client)

	// Get the trusted CA certificate we just created
	getTrustedCACertificate(&client, certificate.ID)

	// Delete the trusted CA certificate
	deleteTrustedCACertificate(&client, certificate.ID)
}

func createTrustedCACertificate(client *gcore.Client) *cdn.CaCertificate {
	fmt.Println("\n=== CREATE TRUSTED CA CERTIFICATE ===")

	// The certificate name must be unique, so use a timestamp suffix.
	name := fmt.Sprintf("gcore-go-example-ca-%d", time.Now().Unix())

	certificate, err := client.CDN.TrustedCaCertificates.New(context.Background(), cdn.TrustedCaCertificateNewParams{
		Name:           name,
		SslCertificate: testCACertificatePEM,
	})
	if err != nil {
		log.Fatalf("Error creating trusted CA certificate: %v", err)
	}

	fmt.Printf("Created Trusted CA Certificate: ID=%d, Name=%s, Issuer=%s\n",
		certificate.ID, certificate.Name, certificate.CertIssuer)
	fmt.Println("=====================================")

	return certificate
}

func listTrustedCACertificates(client *gcore.Client) {
	fmt.Println("\n=== LIST TRUSTED CA CERTIFICATES ===")

	result, err := client.CDN.TrustedCaCertificates.List(context.Background(), cdn.TrustedCaCertificateListParams{})
	if err != nil {
		log.Fatalf("Error listing trusted CA certificates: %v", err)
	}

	certificates := result.AsPlainList()
	for i, certificate := range certificates {
		fmt.Printf("  %d. Trusted CA Certificate: ID=%d, Name=%s\n",
			i+1, certificate.ID, certificate.Name)
	}
	fmt.Println("====================================")
}

func getTrustedCACertificate(client *gcore.Client, certificateID int64) {
	fmt.Println("\n=== GET TRUSTED CA CERTIFICATE ===")

	certificate, err := client.CDN.TrustedCaCertificates.Get(context.Background(), certificateID)
	if err != nil {
		log.Fatalf("Error getting trusted CA certificate: %v", err)
	}

	fmt.Printf("Trusted CA Certificate: ID=%d, Name=%s, SubjectCN=%s, ValidUntil=%s\n",
		certificate.ID, certificate.Name, certificate.CertSubjectCn, certificate.ValidityNotAfter)
	fmt.Println("==================================")
}

func deleteTrustedCACertificate(client *gcore.Client, certificateID int64) {
	fmt.Println("\n=== DELETE TRUSTED CA CERTIFICATE ===")

	err := client.CDN.TrustedCaCertificates.Delete(context.Background(), certificateID)
	if err != nil {
		log.Fatalf("Error deleting trusted CA certificate: %v", err)
	}

	fmt.Printf("Trusted CA Certificate with ID %d successfully deleted\n", certificateID)
	fmt.Println("=====================================")
}
