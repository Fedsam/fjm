package cert

import (
	"crypto/tls"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
)

func downloadCertificate(jdkPath, certName, hostname string) (string, error) {
	con, err := tls.Dial("tcp", fmt.Sprintf("%s:443", hostname), &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return "", err
	}
	defer func(con *tls.Conn) {
		err := con.Close()
		if err != nil {
			return
		}
	}(con)

	cert := con.ConnectionState().PeerCertificates[0]

	// Encode certificate to PEM
	pemData := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	})

	// Write to file
	filename := certName + ".crt"
	if err := os.WriteFile(filepath.Join(jdkPath, filename), pemData, 0644); err != nil {
		panic(err)
	}
	fmt.Printf("certificate %s downloaded.\n", certName)
	return filename, nil
}
