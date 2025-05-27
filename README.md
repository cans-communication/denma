# denma

VoIP Phone SDK, implemented by Go and Diago.

Denma is inpired from Den Den Mushi, communication device in anime.

# Quick start

```go
func main() {
	d, err := denma.NewDenma(
		extension,
		password,
		domain,
		port,
	)

	if err != nil {
		panic(err)
	}

	callResult, err := d.CallAndPlayAudio(
		ctx,
		calleeNumber,
		audioFilePath,
	)
	if err != nil {
		panic(err)
	}
}
```

# Example

1. [Call And Play Audio](example/call_play_audio/README.md)

# Quick start

Call Flow Steps
Denma initiates a call to the phone number associated with user.
When the user answers the call, Denma will immediately connect the call to the IVR system as configured.

```go
func main() {
	d, err := denma.NewDenma(
		extension,
		password,
		domain,
		port,
	)

	if err != nil {
		panic(err)
	}

	callResult, err := d.TranferCall(
		ctx,
		calleeNumber,
		tranferIvrNumber,
	)
	if err != nil {
		panic(err)
	}
}
```
