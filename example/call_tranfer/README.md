# Simple example for call and play audio

1. Create .env

```sh

CALL_TRANFER_IVR=
CALL_TRANFER_EXTENSION=
CALL_TRANFER_PASSWORD=
CALL_TRANFER_SIP_DOMAIN=
CALL_TRANFER_SIP_PORT=
CALL_TRANFER_CALLEE_NUMBER=
```

2. Export env by

```sh
export $(grep -v '^#' .env | xargs)
```

3. Run script

```sh
go run .
```
