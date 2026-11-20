# phontx

**English** · [中文](README.zh.md) · [日本語](README.ja.md)

[![Test Workflow](https://github.com/yukumo-group/phontx/actions/workflows/test.yaml/badge.svg)](https://github.com/yukumo-group/phontx/actions/workflows/test.yaml)
[![Go Version](https://img.shields.io/github/go-mod/go-version/yukumo-group/phontx)](go.mod)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

`phontx` is a Go library and CLI for **unified tone files for the AquesTalk family of Japanese text-to-speech engines**. A single `.phontx` file describes how a voice sounds on **AquesTalk1**, **AquesTalk2**, and **AquesTalk10** at once, and the CLI can render that tone to a WAV file by running the [Yukw](https://github.com/yukumo-group/Yukw) engine module (Rust → WebAssembly) through [wazero](https://wazero.io/) — no CGO, no C toolchain, and no native SDK at runtime.

> `.phontx` is not `.phont`. `.phont` is AquesTalk2's native voice data file. A `.phontx` is phontx's own container that can hold presets for all three engines and optionally embed a `.phont` file.

## Features

- **One tone, three engines** — presets and custom settings for AquesTalk1/2/10 live in the same file.
- **Presets** — 9 AquesTalk1 presets, 16 AquesTalk2 presets, and 8 AquesTalk10 presets (including `custom`) are defined.
- **Custom AquesTalk10 tone** — choose a base fragment (F1E / F2E / M1E) and tune pitch, accent, LMD, and FSC.
- **Embedded `.phont`** — import an AquesTalk2 `.phont` voice and carry it inside the `.phontx` file.
- **Inspectable** — `readPhontx` dumps a tone and `showEngineInfo` reports which engines it supports.
- **Synthesis** — `synthAudio` writes a WAV file using the [Yukw](https://github.com/yukumo-group/Yukw) WebAssembly module, so the same binary works on Windows, Linux, and macOS.
- **Library first** — the CLI is a thin layer over the exported `pkg/tone`, `pkg/aqdata`, `pkg/custom`, and `pkg/synth` packages.
- **Portable format** — `.phontx` files are [MessagePack](https://msgpack.org/) maps generated with [`tinylib/msgp`](https://github.com/tinylib/msgp), so they are easy to read from other languages.

## Contents

- [Requirements](#requirements)
- [Build and install](#build-and-install)
- [CLI usage](#cli-usage)
- [Library usage](#library-usage)
- [The `.phontx` format](#the-phontx-format)
- [Presets and AquesTalk10 custom tones](#presets-and-aquestalk10-custom-tones)
- [How synthesis works](#how-synthesis-works)
- [Project layout](#project-layout)
- [Development](#development)
- [Known caveats](#known-caveats)
- [Related projects](#related-projects)
- [License and AquesTalk](#license-and-aquestalk)

## Requirements

| Requirement | Notes |
| --- | --- |
| **Go** | 1.27 or newer (see [`go.mod`](go.mod)) — only needed to build from source |
| **Engine module** | `aq_wasm.wasm` in the current working directory, built from [Yukw](https://github.com/yukumo-group/Yukw) (see below) |
| **Tests** | Python 3.13 available as `python` on `PATH` (a few tests run Python snippets to build expected byte fixtures) |
| **CGO / native SDK** | Not required. phontx is pure Go; the engine module is provided by Yukw and executed by [wazero](https://wazero.io/) |
| **Rebuilding the engine (optional)** | Rust with the `wasm32-unknown-unknown` target and the relocatable AQUEST archives — only needed if you build `aq_wasm.wasm` from Yukw yourself |

### The `aq_wasm.wasm` engine module

`phontx` does not link against the AquesTalk libraries and contains no C code. Its only external runtime dependency is the engine module built by [`Yukw`](https://github.com/yukumo-group/Yukw), a Rust/WebAssembly runtime for AquesTalk. Yukw links the AQUEST engine archives into a single wasm module (`aq_wasm.wasm`, 0 imports, one linear memory) and exposes the ABI that `pkg/synth` calls through [wazero](https://wazero.io/).

`cmd/main.go` loads `./aq_wasm.wasm` — relative to the **current working directory** — and compiles it once before any command runs. The exports used by `pkg/synth` are:

- `aq_synthesize` — synthesize a WAV buffer from phonetic text and a 40-byte parameter block.
- `aq_error_message` — map an engine error code to an error token.

Because the runtime is initialized for every subcommand, the file must be present even for `readPhontx`; starting `phontx` from a directory without it fails immediately.

The checked-in `pkg/synth/testdata/aq_wasm.wasm` is a prebuilt copy of the Yukw module used by the test suite. For local experiments, copy it:

```bash
cp pkg/synth/testdata/aq_wasm.wasm .
```

To rebuild the module, use Yukw. It requires Rust with the `wasm32-unknown-unknown` target and the relocatable AQUEST engine archives produced by [`aquestalk-wasm`](https://github.com/yukumo-group/aquestalk-wasm):

```bash
git clone https://github.com/yukumo-group/Yukw.git
cd Yukw
just wasm
# equivalent:
# cargo build -p aq-wasm --target wasm32-unknown-unknown --release --lib
# output: target/wasm32-unknown-unknown/release/aq_wasm.wasm
```

Copy the resulting file into the working directory you will run phontx from, as `aq_wasm.wasm`. These prerequisites belong to Yukw, not to phontx. Yukw currently produces evaluation builds of the engines; use a licensed AQUEST build for production behavior.

## Build and install

```bash
git clone https://github.com/yukumo-group/phontx.git
cd phontx
go mod download

# Build the CLI.
go build -trimpath -o phontx ./cmd

# The engine module has to live in the working directory.
cp pkg/synth/testdata/aq_wasm.wasm .

./phontx --help
```

On Windows, `build.bat` performs the build step and writes `build\phontx.exe`:

```powershell
.\build.bat
Copy-Item pkg\synth\testdata\aq_wasm.wasm build\
cd build
.\phontx.exe --help
```

`build.bat` only builds the executable; it does not copy `aq_wasm.wasm`.

## CLI usage

```text
phontx
├── createPhontx        interactive wizard that writes a .phontx file
├── readPhontx          dump a .phontx file
│   └── showEngineInfo  report which engines the file supports
└── synthAudio          render a WAV file from a .phontx file
```

### `phontx createPhontx`

An interactive wizard. For each engine you answer `y`/`n`; for AquesTalk2 you choose between a preset and a `.phont` file; for AquesTalk10 you choose between a preset and a custom tone. Finally you enter the output directory and the tone name, and phontx writes `<directory>/<name>.phontx`.

```console
$ phontx createPhontx
Allows Aquestalk1 engine? (y/N):
y
Here are the available presets for engine Aquestalk1:
f1
f2
...
Input the preset you want to use:
f1
Allows Aquestalk2 engine? (y/N):
y
do you want to use preset as custom tone instead of phont file? (y/N):
y
...
Allows Aquestalk10 engine? (y/N):
y
do you want to use presets instead of custom tone (y/N):
n
index: 0
- base name: F1E

index: 1
- base name: F2E

index: 2
- base name: M1E

input the index for the base of the custome tone:
0
Input int value for LMD, pitch 1 range: (0-200)
100
Input int value for FSC, pitch 2 range: (50-200)
100
Do you want to adjust accent and pitch value which depends on the base setted? (y/N):
n
new tone configuration done, input the directory you want to save the phontx file in:
out
input the name of the tone:
demo-tone
file saved at out/demo-tone.phontx
```

The output directory must already exist; `SaveFile` does not create it.

### `phontx readPhontx`

```console
$ phontx readPhontx -f out/demo-tone.phontx
- ToneData:
    - Aq1Preset:
        - Preset:
            - Index: 0
            - Name: f1
    - Aq2Preset:
        - Preset:
            - Index: 1
            - Name: aq_defo1
    - Aq10Preset:
        - Preset:
            - Index: 0
            - Name: custom
    - Aq10CustomTone:
        - Base: 0
        - LMD: 100
        - FSC: 100
```

The order of fields in the dump may vary. Embedded `.phont` bytes are not printed.

### `phontx readPhontx showEngineInfo`

```console
$ phontx readPhontx showEngineInfo -f out/demo-tone.phontx
Engines supported by out/demo-tone.phontx:
Aquestalk1: Supported
Aquestalk2: Supported
Aquestalk10: Supported
```

### `phontx synthAudio`

The engine version is selected interactively, and must be one of the engines the tone enables.

```console
$ phontx synthAudio -f out/demo-tone.phontx -t "こんにちは" -a hello -d out -s 100 -v 100
select the version for generation:
Here are the available versions:
index of version 1: 1
index of version 2: 2
index of version 10: 10
Input the index of the version you want to use:
2
audio saved at out/hello.wav
```

| Flag | Short | Required | Default | Description |
| --- | --- | --- | --- | --- |
| `--filePath` | `-f` | yes | | Path to the `.phontx` file |
| `--targetText` | `-t` | yes | | Text to speak, in AquesTalk phonetic notation (kana); no kanji conversion is performed |
| `--targetAudioFileName` | `-a` | yes | | Output file name without the `.wav` suffix |
| `--targetAudioDirectory` | `-d` | no | current directory | Output directory |
| `--targetSpeed` | `-s` | no | `-1` | Speaking rate, clamped to 50–300 (100 is normal) |
| `--targetVolume` | `-v` | no | `-1` | Volume, clamped to 0–300 (100 is normal for AquesTalk10) |

With the flag defaults, `-1` is converted to `uint32` before clamping, which selects the maximum instead of a sensible default. Pass explicit values (usually `100`) as in the example above. The output directory must already exist.

## Library usage

The example below builds a tone for all three engines, saves it, loads it back, and renders a WAV file. It is a complete program; put it in a module that requires `github.com/yukumo-group/phontx`.

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/yukumo-group/phontx/pkg/aqdata"
	"github.com/yukumo-group/phontx/pkg/custom"
	"github.com/yukumo-group/phontx/pkg/synth"
	"github.com/yukumo-group/phontx/pkg/tone"
)

func main() {
	ctx := context.Background()

	// 1. Build a tone that works on all three engines.
	toneData := tone.NewData().
		SetAq1Preset(aqdata.Aq1F1).
		SetAq2Preset(aqdata.Aq2AqDefo1).
		SetAq10CustomTone(
			custom.NewAq10CustomTone().
				WithBase(custom.F1E).
				WithLMD(100).
				WithFSC(100),
		)

	// 2. Save it as out/my-tone.phontx (.phontx is appended for you).
	if err := os.MkdirAll("out", 0o755); err != nil {
		log.Fatal(err)
	}
	filePath, err := toneData.SaveFile("out", "my-tone")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("saved:", filePath)

	// 3. Load it back and inspect engine support.
	loaded, err := tone.NewDataFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	allowed := loaded.GetAllowedEngines()
	fmt.Printf("Aq1=%t Aq2=%t Aq10=%t\n", allowed.Aq1, allowed.Aq2, allowed.Aq10)

	// 4. Synthesize a WAV file with the WASM engine.
	generatorRuntime := synth.NewGeneratorRuntime("./aq_wasm.wasm")
	if err := generatorRuntime.InitializeRuntime(ctx); err != nil {
		log.Fatal(err)
	}
	defer generatorRuntime.CloseRuntime(ctx)

	generator, err := synth.NewGenerator(ctx, generatorRuntime)
	if err != nil {
		log.Fatal(err)
	}
	defer generator.CloseGenerator(ctx)

	param := synth.NewSynthParam(
		"こんにちは", // AquesTalk phonetic text
		100,         // speed: 50-300
		100,         // volume: 0-300
		"out/hello.wav",
	)
	if err := generator.Generate(ctx, aqdata.AqVer10, param, loaded); err != nil {
		log.Fatal(err)
	}
	fmt.Println("wrote out/hello.wav")
}
```

Output when `aq_wasm.wasm` is present in the working directory:

```text
saved: out/my-tone.phontx
Aq1=true Aq2=true Aq10=true
wrote out/hello.wav
```

### API overview

| Package | What it provides |
| --- | --- |
| `pkg/tone` | `Data` (the `.phontx` model), `NewData`, `NewDataFromFile`, setters, `SaveFile`, `SaveFileWithFullPath`, `Equals`, `GetAllowedEngines`, `IsVersionAllowed`, `GetPresetID` |
| `pkg/aqdata` | `AquesVersion` (`AqVer1`, `AqVer2`, `AqVer10`), the generic `PresetList`, and every preset constant (`Aq1F1`, `Aq2AqDefo1`, `Aq10Custom`, ...) |
| `pkg/custom` | `Aq10CustomTone` with `WithBase`, `WithPitch`, `WithAccent`, `WithLMD`, `WithFSC`, plus the `Aq10Base` values `F1E`, `F2E`, `M1E` |
| `pkg/synth` | `GeneratorRuntime` (compiles the wasm once), `Generator` (one module instance), `Param`, `NewSynthParam`, `PackData` |

A few behaviours worth knowing:

- `SaveFile(dir, name)` always appends `.phontx` and returns the final path. `SaveFileWithFullPath` writes exactly the path it is given.
- `SetAq10CustomTone` also sets `Aq10Preset` to `custom`, so the engine knows to use the custom parameters.
- `SetPhontDataFromFile` accepts only files whose extension is `.phont`; the bytes are embedded in the `.phontx` as-is.
- `NewSynthParam` clamps speed to 50–300 and volume to 0–300.
- A `Generator` owns a single wasm module instance and its linear memory. Create one per synthesis (or per goroutine) and close it when done; `GeneratorRuntime` can stay alive and hand out many generators.

## The `.phontx` format

A `.phontx` file is a MessagePack-encoded `tone.Data` struct. There is no header or magic number: the top-level value is a map keyed by field name. Every field is optional, and a `nil` field means "this engine is not configured".

| Field | Type | Meaning |
| --- | --- | --- |
| `Aq1Preset` | preset | AquesTalk1 voice preset (`Index` + `Name`) |
| `Aq2Preset` | preset | AquesTalk2 voice preset |
| `PhontFileData` | bytes | raw contents of an AquesTalk2 `.phont` file |
| `Aq10Preset` | preset | AquesTalk10 preset; `custom` when `Aq10CustomTone` is used |
| `Aq10CustomTone` | object | AquesTalk10 custom parameters (see below) |

Engine support is derived from those fields:

- AquesTalk1 is enabled when `Aq1Preset` is set.
- AquesTalk2 is enabled when `Aq2Preset` or `PhontFileData` is set.
- AquesTalk10 is enabled when `Aq10Preset` or `Aq10CustomTone` is set.

A preset stores both its list index and its name:

```go
type Preset struct {
	Index uint32
	Name  string
}
```

Conceptually, a file looks like this (shown as JSON; the actual encoding is MessagePack, and `PhontFileData` is a binary value rather than base64):

```json
{
  "Aq1Preset": { "Preset": { "Index": 0, "Name": "f1" } },
  "Aq2Preset": { "Preset": { "Index": 1, "Name": "aq_defo1" } },
  "PhontFileData": null,
  "Aq10Preset": { "Preset": { "Index": 0, "Name": "custom" } },
  "Aq10CustomTone": {
    "Base": 0,
    "LMD": 100,
    "FSC": 100,
    "Pitch": null,
    "Accent": null
  }
}
```

The encoder and decoder are generated by `tinylib/msgp` into `*_gen.go` files. Decoding skips keys it does not know and leaves missing fields at their zero values, so the format is tolerant but **not versioned** — changing a struct changes the format. Regenerate the code after editing a `//go:generate msgp` struct:

```bash
go generate ./...
```

## Presets and AquesTalk10 custom tones

Preset indices are assigned in declaration order in `pkg/aqdata/presets.go`. The Go constants follow the pattern `Aq1<Name>`, `Aq2<Name>`, and `Aq10<Name>`.

### AquesTalk1 (9 voices)

| Index | Name | Constant |
| --- | --- | --- |
| 0 | `f1` | `aqdata.Aq1F1` |
| 1 | `f2` | `aqdata.Aq1F2` |
| 2 | `f3` | `aqdata.Aq1F3` |
| 3 | `m1` | `aqdata.Aq1M1` |
| 4 | `m2` | `aqdata.Aq1M2` |
| 5 | `r1` | `aqdata.Aq1R1` |
| 6 | `dvd` | `aqdata.Aq1DVD` |
| 7 | `imd1` | `aqdata.Aq1IMD1` |
| 8 | `jgr` | `aqdata.Aq1JGR` |

### AquesTalk2 (16 voices)

| Index | Name | Constant |
| --- | --- | --- |
| 0 | `aq_yukkuri` | `aqdata.Aq2AqYukkuri` |
| 1 | `aq_defo1` | `aqdata.Aq2AqDefo1` |
| 2 | `aq_f1c` | `aqdata.Aq2AqF1c` |
| 3 | `aq_f3a` | `aqdata.Aq2AqF3a` |
| 4 | `aq_huskey` | `aqdata.Aq2AqHuskey` |
| 5 | `aq_m4b` | `aqdata.Aq2AqM4b` |
| 6 | `aq_mf1` | `aqdata.Aq2AqMf1` |
| 7 | `aq_rb2` | `aqdata.Aq2AqRb2` |
| 8 | `aq_rb3` | `aqdata.Aq2AqRb3` |
| 9 | `aq_rm` | `aqdata.Aq2AqRm` |
| 10 | `aq_robo` | `aqdata.Aq2AqRobo` |
| 11 | `aq_teto1` | `aqdata.Aq2AqTeto1` |
| 12 | `ar_f4` | `aqdata.Aq2ArF4` |
| 13 | `ar_m5` | `aqdata.Aq2ArM5` |
| 14 | `ar_mf2` | `aqdata.Aq2ArMF2` |
| 15 | `ar_rm3` | `aqdata.Aq2ArRM3` |

### AquesTalk10 (7 voices + custom)

| Index | Name | Constant |
| --- | --- | --- |
| 0 | `custom` | `aqdata.Aq10Custom` |
| 1 | `f1` | `aqdata.Aq10F1` |
| 2 | `f2` | `aqdata.Aq10F2` |
| 3 | `f3` | `aqdata.Aq10F3` |
| 4 | `m1` | `aqdata.Aq10M1` |
| 5 | `m2` | `aqdata.Aq10M2` |
| 6 | `r1` | `aqdata.Aq10R1` |
| 7 | `r2` | `aqdata.Aq10R2` |

### AquesTalk10 custom tone

The `custom` preset unlocks the per-tone parameters. Values outside the documented ranges are clamped by the `With*` methods.

| Field | Range | Default | Notes |
| --- | --- | --- | --- |
| `Base` | `0` F1E, `1` F2E, `2` M1E | `F1E` | base voice fragment |
| `LMD` | 0–200 | 100 | pitch parameter 1 |
| `FSC` | 50–200 | 100 | pitch parameter 2 (sampling frequency) |
| `Pitch` | 20–200 | depends on the base: F1E 100, F2E 77, M1E 30 | `nil` means "use the base default" |
| `Accent` | 0–200 | depends on the base: F1E 100, F2E 150, M1E 100 | `nil` means "use the base default" |

```go
customTone := custom.NewAq10CustomTone().
	WithBase(custom.F2E).
	WithPitch(90).
	WithAccent(140).
	WithLMD(110).
	WithFSC(120)

toneData := tone.NewData().SetAq10CustomTone(customTone)
```

The parameters above only apply while the AquesTalk10 preset is `custom`. Choosing `f1`...`r2` makes the engine use the preset's own tuning.

## How synthesis works

1. `tone.NewDataFromFile` reads the file and unmarshals the MessagePack map into `tone.Data`.
2. `synth.NewGenerator` instantiates a fresh WebAssembly module from the compiled Yukw `aq_wasm.wasm`; each generator gets its own linear memory.
3. `Generator.Generate` performs the following steps:
   - validates that the requested engine version is enabled by the tone,
   - writes the embedded `.phont` bytes (if any) at offset 0 of the module memory,
   - packs the engine parameters and writes them after the phont bytes,
   - writes the UTF-8 phonetic text,
   - calls `aq_synthesize`, and if the returned error code is non-zero, calls `aq_error_message` to turn it into an error,
   - reads the WAV bytes back out of linear memory and writes the output file.
4. Buffers are kept 4-byte aligned (`internal/bts.NextOffset`) so the wasm side sees naturally aligned fields.

The parameter block produced by `synth.PackData` is a sequence of little-endian `uint32` values:

| Offset | Field |
| --- | --- |
| 0 | version (`1`, `2`, or `10`) |
| 4 | preset index |
| 8 | speed |
| 12 | embedded phont data offset (`0`) |
| 16 | AquesTalk10 base |
| 20 | volume |
| 24 | pitch |
| 28 | accent |
| 32 | LMD |
| 36 | FSC |

That is 40 bytes in total. The embedded `.phont` length is passed separately to `aq_synthesize`, and the text is a separate UTF-8 buffer.

## Project layout

```text
phontx/
├── cmd/                     CLI entry point and Cobra commands
│   ├── main.go              loads aq_wasm.wasm, then runs the CLI
│   ├── cmd.go               root command and flag definitions
│   ├── create_phontx.go     interactive .phontx builder
│   ├── read_phontx.go       dump a .phontx file
│   ├── show_engine_info.go  showEngineInfo subcommand
│   ├── synth.go             synthAudio command
│   ├── prompt.go            prompt helpers
│   └── welcome.go           banner
├── internal/
│   ├── bts/                 byte/string helpers and offset alignment
│   ├── python/              test helper that runs Python snippets
│   └── ui/                  interactive prompts, wizards, struct dump
├── pkg/
│   ├── aqdata/              engine versions and preset lists
│   ├── custom/              AquesTalk10 custom tone parameters
│   ├── synth/               wazero runtime and WAV synthesis
│   └── tone/                .phontx data model and (de)serialization
├── build/                   build output and runtime assets (git-ignored)
├── build.bat                Windows build helper
└── go.mod
```

| Package | Responsibility |
| --- | --- |
| `pkg/tone` | `Data` model, load/save, engine support checks |
| `pkg/aqdata` | `AquesVersion`, preset definitions, generic `PresetList` |
| `pkg/custom` | `Aq10CustomTone` and `Aq10Base` |
| `pkg/synth` | `GeneratorRuntime`, `Generator`, `Param`, `PackData` |
| `internal/ui` | prompts, preset/custom-tone wizards, `ToneToString`, `ShowEngineInfo` |
| `internal/bts` | byte comparison, null-terminated string reading, offset alignment |
| `internal/python` | `ExecutePython`, used by tests to generate expected byte fixtures |

## Development

```bash
go test ./...                     # all packages
go test ./pkg/synth ./pkg/tone    # a subset
go vet ./...
golint -set_exit_status ./...     # what CI runs
```

- A few tests execute Python (`internal/python`, and the parameter-packing test in `pkg/synth`), so `python` must resolve to Python 3.13 on `PATH`.
- The `pkg/synth` tests perform real synthesis with `pkg/synth/testdata/aq_wasm.wasm` and write their output into `t.TempDir()`.
- Code generation: install the generator once with `go install github.com/tinylib/msgp@latest`, then run `go generate ./...` from the repository root. Do not hand-edit `*_gen.go` files.
- CI (`.github/workflows/test.yaml`) runs on Ubuntu with the Go version from `go.mod` and Python 3.13, then runs `golint` and uploads coverage to Codecov.
- Releases (`.github/workflows/release.yaml`) run GoReleaser when a `v*` tag is pushed and publish the `phontx` binary built from `./cmd`.

## Known caveats

- **`aq_wasm.wasm` is required for every command**, not only for `synthAudio`, because `cmd/main.go` initializes the wazero runtime before Cobra runs. Start `phontx` from the directory that contains the module.
- **Do not rely on the `-s`/`-v` defaults.** They default to `-1`, which becomes `4294967295` when converted to `uint32` and is then clamped to `300`. Pass `-s 100 -v 100` unless you actually want maximum speed and volume.
- **Interactive input is whitespace-delimited.** Prompts use `fmt.Scan`, so paths containing spaces cannot be entered, and yes/no prompts require an explicit `y` or `n` (pressing Enter does not accept the shown default).
- **No kanji conversion.** The text given to `-t` is passed to the engine as-is; supply AquesTalk phonetic text (kana and voice symbols). AqKanji2Koe is not bundled.
- **`.phont` import is strict.** `SetPhontDataFromFile` rejects any extension other than `.phont` (the comparison is case-sensitive).
- **AquesTalk10 custom parameters need the `custom` preset.** Pitch, accent, LMD, and FSC are ignored for the built-in AquesTalk10 presets.
- **The `.phontx` format is not versioned.** msgp skips unknown keys, but structural changes to `tone.Data` are format changes.

## Related projects

- [Yukw](https://github.com/yukumo-group/Yukw) — the Rust/WebAssembly runtime that builds `aq_wasm.wasm`; phontx uses its `aq_synthesize` ABI.
- [aquestalk-wasm](https://github.com/yukumo-group/aquestalk-wasm) — the upstream engine project that produces the relocatable archives Yukw links into the module.

## License and AquesTalk

The source code in this repository is released under the [MIT License](LICENSE) (Copyright © 2026 Yukumo Group).

AquesTalk1, AquesTalk2, and AquesTalk10 are commercial products of [AQUEST Corp.](https://www.a-quest.com/). The engine binaries, voice data, and WebAssembly builds — including the Yukw-built `aq_wasm.wasm` — are **not** covered by the MIT license and are not distributed by this repository. Yukw currently ships evaluation builds; obtain a licensed AQUEST engine build for production, and follow AQUEST's terms when synthesizing or redistributing audio.
