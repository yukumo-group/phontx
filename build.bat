@echo off

set TOTAL_STEPS=2

:: Set environment such as paths
echo [1/%TOTAL_STEPS%] Set env
set CURRENT_PATH=%~dp0
echo - Current at: %CURRENT_PATH%

:: Build the program
echo [2/%TOTAL_STEPS%] Build
set BUILD_DIR=%CURRENT_PATH%\build
mkdir %BUILD_DIR%
cd cmd
go build -trimpath -o phontx.exe -tags=cmd .
echo - Saved to %CURRENT_PATH%cmd\phontx.exe
move %CURRENT_PATH%cmd\phontx.exe %BUILD_DIR%
cd ..
