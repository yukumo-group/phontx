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

:: Configure the third-party packages
echo [3/%TOTAL_STEPS%] Configure third-party packages
:: Set lib path according to windows 32-bit or 64 bit
set LIB_PATH_32=%CURRENT_PATH%third-party\win32
set LIB_PATH_64=%CURRENT_PATH%third-party\win64
if %PROCESSOR_ARCHITECTURE%==AMD64 (
    echo 64 bit system
    set LIB_PATH=%LIB_PATH_64%
) else if exist PROCESSOR_ARCHITEW6432 (
    echo 64 bit system
    set LIB_PATH=%LIB_PATH_64%
) else (
    echo 32 bit system
    set LIB_PATH=%LIB_PATH_32%
)
:: Set the data here
if not exist %LIB_PATH% (
    echo The libraries does not exists, set it according to README
) else (
    xcopy %LIB_PATH% %BUILD_DIR%
)
