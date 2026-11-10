@echo off

set TOTAL_STEPS=3

:: Set environment such as pathes
echo [1/%TOTAL_STEPS%] Set env
set CURRENT_PATH=%~dp0
echo - Current at: %CURRENT_PATH%
set CGO_CFLAGS=-I%CURRENT_PATH%third-party
echo - CGO_CFLAGS: %CGO_CFLAGS%

:: Build the program
echo [2/%TOTAL_STEPS%] Build
set BUILD_DIR=%CURRENT_PATH%\build
mkdir %BUILD_DIR%
cd cmd
go build -trimpath -o phontx.exe -tags=cmd .
echo - Saved to %CURRENT_PATH%cmd\phontx.exe
move %CURRENT_PATH%cmd\phontx.exe %BUILD_DIR%
cd ..
