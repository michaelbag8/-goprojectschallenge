# Morsify

Morsify is a simple Morse code encoder and decoder written in Go.

## Features

- Encode text to Morse code
- Decode Morse code to text
- Case-insensitive input
- Replaces unsupported characters with ?
- Handles extra spaces during decoding

## Build

```bash
go build -o morsify
go run . encode SOS