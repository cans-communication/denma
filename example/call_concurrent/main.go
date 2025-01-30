package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/cans-communication/denma"
	"github.com/kelseyhightower/envconfig"
)

type EnvCfg struct {
	Extension0    string `envconfig:"EXTENSION_0" required:"true"`
	Extension1    string `envconfig:"EXTENSION_1" required:"true"`
	Password0     string `envconfig:"PASSWORD_0" required:"true"`
	Password1     string `envconfig:"PASSWORD_1" required:"true"`
	Domain        string `envconfig:"SIP_DOMAIN" required:"true"`
	Port          int64  `envconfig:"SIP_PORT" required:"true"`
	CalleeNumber0 string `envconfig:"CALLEE_NUMBER_0" required:"true"`
	CalleeNumber1 string `envconfig:"CALLEE_NUMBER_1" required:"true"`
	AudioFilePath string `envconfig:"AUDIO_FILE_PATH" required:"true"`
}

func main() {

	var envCfg EnvCfg
	err := envconfig.Process("CALL_CONCURRENT", &envCfg)
	if err != nil {
		panic(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	audioData, err := os.ReadFile(envCfg.AudioFilePath)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {

		defer wg.Done()

		d, err := denma.NewDenma(
			envCfg.Extension0,
			envCfg.Password0,
			envCfg.Domain,
			envCfg.Port,
		)
		if err != nil {
			panic(err)
		}

		callResult, err := d.CallAndPlayAudio(
			ctx,
			envCfg.CalleeNumber0,
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
	}()

	go func() {

		defer wg.Done()

		d, err := denma.NewDenma(
			envCfg.Extension1,
			envCfg.Password1,
			envCfg.Domain,
			envCfg.Port,
		)
		if err != nil {
			panic(err)
		}

		callResult, err := d.CallAndPlayAudio(
			ctx,
			envCfg.CalleeNumber1,
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
	}()

	wg.Wait()

}
