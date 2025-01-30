# Simple example for call and play audio

1. Create .env
```sh
CALL_PLAY_FILE_EXTENSION=
CALL_PLAY_FILE_PASSWORD=
CALL_PLAY_FILE_SIP_DOMAIN=
CALL_PLAY_FILE_SIP_PORT=
CALL_PLAY_FILE_CALLEE_NUMBER=
CALL_PLAY_FILE_AUDIO_FILE_PATH=
```

2. Export env by
```sh
export $(grep -v '^#' .env | xargs)
```

3. Run script
```sh
go run .
```

