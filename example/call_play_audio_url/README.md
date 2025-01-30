# Simple example for call and play audio

1. Create .env
```sh
CALL_PLAY_URL_EXTENSION=
CALL_PLAY_URL_PASSWORD=
CALL_PLAY_URL_SIP_DOMAIN=
CALL_PLAY_URL_SIP_PORT=
CALL_PLAY_URL_CALLEE_NUMBER=
CALL_PLAY_URL_AUDIO_FILE_URL=
```

2. Export env by
```sh
export $(grep -v '^#' .env | xargs)
```

3. Run script
```sh
go run .
```

