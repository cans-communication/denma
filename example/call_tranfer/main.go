package main

import (
	"context"
	"fmt"
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
	Ivr          string `envconfig:"IVR" required:"true"`
}

func main() {

	var envCfg EnvCfg
	err := envconfig.Process("CALL_TRANFER", &envCfg)
	if err != nil {
		panic(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	d, err := denma.NewDenma(
		envCfg.Extension,
		envCfg.Password,
		envCfg.Domain,
		envCfg.Port,
	)
	if err != nil {
		panic(err)
	}

	callResult, err := d.TranferCall(
		ctx,
		envCfg.CalleeNumber,
		envCfg.Ivr,
	)
	if err != nil {

		panic(err)

	}
	fmt.Println("callee number: ", callResult.CalleeNumber)
	fmt.Println("status: ", callResult.Status)
	fmt.Println("called duration: ", callResult.Duration)
	fmt.Println("talk time: ", callResult.TalkTime)

}
