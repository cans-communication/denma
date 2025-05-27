package denma

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/emiago/diago"
	"github.com/emiago/sipgo"
	"github.com/emiago/sipgo/sip"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Denma struct {
	Extension string
	Password  string
	DomainSIP string
	PortSIP   int64

	ua *sipgo.UserAgent
	dg *diago.Diago
}

type TransportOptions struct {
	Transport string
	BindHost  string
	BindPort  int64
}

type CallStatus string

var (
	MissedCall CallStatus = "missedcall"
	Answered   CallStatus = "answered"
)

type CallResult struct {
	CalleeNumber string
	Status       CallStatus
	Duration     time.Duration
	TalkTime     time.Duration
}

func NewDenma(
	extension string,
	password string,
	domain string,
	port int64,
	options ...TransportOptions,
) (*Denma, error) {

	// Setup logger
	//	NOTE: for silent logging in diago
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	log.Logger = zerolog.New(
		zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.StampMicro,
		},
	).
		With().
		Timestamp().
		Logger().
		Level(zerolog.WarnLevel)

	ua, err := sipgo.NewUA(
		sipgo.WithUserAgent(
			fmt.Sprintf("denma_%s", extension),
		),
		sipgo.WithUserAgentHostname(domain),
	)
	if err != nil {
		return nil, err
	}

	var transportOpt diago.Transport
	// set default
	if len(options) == 0 {
		transportOpt.BindHost = "0.0.0.0"
		transportOpt.BindPort = 55556
		transportOpt.Transport = "udp"
	}

	for _, opt := range options {
		transportOpt = diago.Transport{
			Transport: opt.Transport,
			BindHost:  opt.BindHost,
			BindPort:  int(opt.BindPort),
		}
	}

	dg := diago.NewDiago(
		ua,
		diago.WithTransport(transportOpt),
	)

	return &Denma{
		Extension: extension,
		Password:  password,
		DomainSIP: domain,
		PortSIP:   port,
		ua:        ua,
		dg:        dg,
	}, nil
}

func (d *Denma) CallAndPlayAudio(ctx context.Context, calleeNumber string, audioData []byte, mimeType string) (*CallResult, error) {

	var uri sip.Uri
	err := sip.ParseUri(
		fmt.Sprintf(
			"sip:%s@%s:%d",
			calleeNumber,
			d.DomainSIP,
			d.PortSIP,
		),
		&uri,
	)
	if err != nil {
		return nil, err
	}

	startTime := time.Now()
	sess, err := d.dg.Invite(
		ctx,
		uri,
		diago.InviteOptions{
			Username: d.Extension,
			Password: d.Password,
		},
	)
	// error from INVITE might come from user-busy
	if err != nil {
		return &CallResult{
			CalleeNumber: calleeNumber,
			Status:       MissedCall,
			Duration:     time.Since(startTime),
		}, nil
	}

	err = sess.Ack(ctx)
	if err != nil {
		return nil, err
	}

	startTalkTime := time.Now()
	pb, err := sess.PlaybackCreate()
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(audioData)
	_, err = pb.Play(
		reader,
		mimeType,
	)

	// error when stream audio and it close unexpectedly
	if err != nil {
		return &CallResult{
			CalleeNumber: calleeNumber,
			Status:       Answered,
			Duration:     time.Since(startTime),
			TalkTime:     time.Since(startTalkTime),
		}, nil
	}

	err = sess.Hangup(ctx)
	if err != nil {
		return &CallResult{
			CalleeNumber: calleeNumber,
			Status:       Answered,
			Duration:     time.Since(startTime),
			TalkTime:     time.Since(startTalkTime),
		}, nil
	}

	return &CallResult{
		CalleeNumber: calleeNumber,
		Status:       Answered,
		Duration:     time.Since(startTime),
		TalkTime:     time.Since(startTalkTime),
	}, nil
}

func (d *Denma) Close() error {
	return d.ua.Close()
}

func (d *Denma) TranferCall(ctx context.Context, calleeNumber string, tranferIVR string) (*CallResult, error) {

	var uri sip.Uri
	err := sip.ParseUri(
		fmt.Sprintf(
			"sip:%s@%s:%d",
			calleeNumber,
			d.DomainSIP,
			d.PortSIP,
		),
		&uri,
	)
	if err != nil {

		return nil, err
	}

	startTime := time.Now()

	sess, err := d.dg.Invite(
		ctx,
		uri,
		diago.InviteOptions{
			Username: d.Extension,
			Password: d.Password,
		},
	)

	if err != nil {
		return &CallResult{
			CalleeNumber: calleeNumber,
			Status:       MissedCall,
			Duration:     time.Since(startTime),
		}, nil
	}

	var referTo sip.Uri
	err = sip.ParseUri(
		fmt.Sprintf(
			"sip:%s@%s:%d",
			tranferIVR,
			d.DomainSIP,
			d.PortSIP,
		),
		&referTo,
	)
	if err != nil {

		return nil, err
	}

	fmt.Println("Start Refer ...")
	fmt.Println("State : ", sess.DialogSIP().LoadState())

	err = sess.Refer(ctx, referTo)
	if err != nil {
		return nil, err
	}

	startTalkTime := time.Now()
	err = sess.Ack(ctx)
	if err != nil {
		return nil, err
	}

	err = sess.Hangup(ctx)
	if err != nil {
		return &CallResult{
			CalleeNumber: calleeNumber,
			Status:       Answered,
			Duration:     time.Since(startTime),
			TalkTime:     time.Since(startTalkTime),
		}, nil
	}

	return &CallResult{
		CalleeNumber: calleeNumber,
		Status:       Answered,
		Duration:     time.Since(startTime),
		TalkTime:     time.Since(startTalkTime),
	}, nil
}
