package noise

import (
	"bytes"
	"testing"

	"github.com/libp2p/go-libp2p-core/crypto"
)

func TestEncryptAndDecrypt_InitToResp(t *testing.T) {
	initTransport := newTestTransport(t, crypto.Ed25519, 2048)
	respTransport := newTestTransport(t, crypto.Ed25519, 2048)

	initConn, respConn := connect(t, initTransport, respTransport)
	defer initConn.Close()
	defer respConn.Close()

	plaintext := []byte("helloworld")
	ciphertext, err := initConn.Encrypt(plaintext)
	if err != nil {
		t.Fatal(err)
	}

	result, err := respConn.Decrypt(ciphertext)
	if !bytes.Equal(plaintext, result) {
		t.Fatalf("got %x expected %x", result, plaintext)
	} else if err != nil {
		t.Fatal(err)
	}

	plaintext = []byte("goodbye")
	ciphertext, err = initConn.Encrypt(plaintext)
	if err != nil {
		t.Fatal(err)
	}

	result, err = respConn.Decrypt(ciphertext)
	if !bytes.Equal(plaintext, result) {
		t.Fatalf("got %x expected %x", result, plaintext)
	} else if err != nil {
		t.Fatal(err)
	}
}

func TestEncryptAndDecrypt_RespToInit(t *testing.T) {
	initTransport := newTestTransport(t, crypto.Ed25519, 2048)
	respTransport := newTestTransport(t, crypto.Ed25519, 2048)

	initConn, respConn := connect(t, initTransport, respTransport)
	defer initConn.Close()
	defer respConn.Close()

	plaintext := []byte("helloworld")
	ciphertext, err := respConn.Encrypt(plaintext)
	if err != nil {
		t.Fatal(err)
	}

	result, err := initConn.Decrypt(ciphertext)
	if !bytes.Equal(plaintext, result) {
		t.Fatalf("got %x expected %x", result, plaintext)
	} else if err != nil {
		t.Fatal(err)
	}
}
