@echo off

echo "[1/2] Set env"
set CURRENT_PATH=%~dp0
echo "- Current at: %CURRENT_PATH%"
set CGO_CFLAGS=-I%CURRENT_PATH%third-party
echo "- CGO_CFLAGS: %CGO_CFLAGS%"

cd "cmd"
echo "[2/2] Build"
go build -o phontx.exe -tags=cmd .
echo "- Saved to %CURRENT_PATH%cmd\phontx.exe"
move %CURRENT_PATH%cmd\phontx.exe %CURRENT_PATH%
