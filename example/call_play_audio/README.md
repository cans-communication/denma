# Simple example for call and play audio

1. Create .env
```sh
CALL_EXAMPLE_EXTENSION=
CALL_EXAMPLE_PASSWORD=
CALL_EXAMPLE_SIP_DOMAIN=
CALL_EXAMPLE_SIP_PORT=
CALL_EXAMPLE_CALLEE_NUMBER=
CALL_EXAMPLE_AUDIO_FILE_PATH=
```

2. Export env by
```sh
export $(grep -v '^#' .env | xargs)
```

3. Run script
```sh
go run .
```

