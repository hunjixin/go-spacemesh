package main

import (
	"bytes"
	"context"
	"log"

	"github.com/spacemeshos/go-spacemesh/signing"
	"github.com/spacemeshos/poet/service"
	"github.com/spacemeshos/poet/shared"
)

func main() {
	ctx := context.Background()

	signer, err := signing.NewEdSigner()
	if err != nil {
		log.Fatal(err)
		return
	}

	challenge := bytes.NewBufferString("xxxxxxxxxxxxxxxxxxxx").Bytes()
	challenge1 := bytes.NewBufferString("yyyyyyyyyyyyyyyyyy").Bytes()

	params := service.NewPowParams(challenge, 10)
	shared.FindSubmitPowNonce(ctx, params.Challenge, challenge1, signer.NodeID().Bytes(), params.Difficulty)
}
