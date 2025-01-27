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
