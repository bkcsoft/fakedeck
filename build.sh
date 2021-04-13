#!/bin/sh

go build -tags "gtk_3_22,pango_1_42" -o ./vlcdeck ./cmd/vlc_fakedeck/
