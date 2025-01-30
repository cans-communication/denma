package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"

	"github.com/cans-communication/denma"
	"github.com/kelseyhightower/envconfig"
)

type EnvCfg struct {
	Extension    string `envconfig:"EXTENSION" required:"true"`
	Password     string `envconfig:"PASSWORD" required:"true"`
	Domain       string `envconfig:"SIP_DOMAIN" required:"true"`
	Port         int64  `envconfig:"SIP_PORT" required:"true"`
	CalleeNumber string `envconfig:"CALLEE_NUMBER" required:"true"`
	AudioURL     string `envconfig:"AUDIO_FILE_URL" required:"true"`
}

func main() {

	var envCfg EnvCfg
	err := envconfig.Process("CALL_PLAY_URL", &envCfg)
	if err != nil {
		panic(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	resp, err := http.Get(envCfg.AudioURL)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic("status not OK")
	}

	audioData, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	d, err := denma.NewDenma(
		envCfg.Extension,
		envCfg.Password,
		envCfg.Domain,
		envCfg.Port,
	)
	if err != nil {
		panic(err)
	}

	callResult, err := d.CallAndPlayAudio(
		ctx,
		envCfg.CalleeNumber,
		audioData,
		"audio/wav",
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("callee number: ", callResult.CalleeNumber)
	fmt.Println("status: ", callResult.Status)
	fmt.Println("called duration: ", callResult.Duration)
	fmt.Println("talk time: ", callResult.TalkTime)

}
