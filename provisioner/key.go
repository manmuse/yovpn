package provisioner

import (
	"log"

	"github.com/digitalocean/godo"
	"golang.org/x/crypto/ssh"
)

const keyBaseName = "yovpn-"

func uploadPublicKey(client *godo.Client, publicKey ssh.PublicKey, id string) (*godo.Key, error) {
	createRequest := &godo.KeyCreateRequest{
		Name:      keyBaseName + id,
		PublicKey: string(ssh.MarshalAuthorizedKey(publicKey)),
	}
	key, _, err := client.Keys.Create(createRequest)
	return key, err
}

func deletePublicKey(client *godo.Client, key *godo.Key) {
	log.Printf("Deleting key with fingerprint %s", key.Fingerprint)
	client.Keys.DeleteByFingerprint(key.Fingerprint)
}
