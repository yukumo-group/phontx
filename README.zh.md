# phontx

[English](README.md) · **中文** · [日本語](README.ja.md)

[![Test Workflow](https://github.com/yukumo-group/phontx/actions/workflows/test.yaml/badge.svg)](https://github.com/yukumo-group/phontx/actions/workflows/test.yaml)
[![Go Version](https://img.shields.io/github/go-mod/go-version/yukumo-group/phontx)](go.mod)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

`phontx` 是一个 Go 库兼命令行工具，用于 **AquesTalk 系列日语语音合成引擎的统一音色文件**。单个 `.phontx` 文件即可同时描述 **AquesTalk1**、**AquesTalk2** 和 **AquesTalk10** 三种引擎下的音色，CLI 还能通过 [wazero](https://wazero.io/) 运行 [Yukw](https://github.com/yukumo-group/Yukw) 引擎模块（Rust → WebAssembly），将该音色渲染为 WAV 文件——运行时不依赖 CGO、C 工具链或原生 SDK。

> `.phontx` 不是 `.phont`。`.phont` 是 AquesTalk2 的原生音色数据文件；`.phontx` 是 phontx 自己的容器，可保存三种引擎的预设，并可选择内嵌一个 `.phont` 文件。

## 功能特性

- **一份音色，三种引擎** — AquesTalk1/2/10 的预设和自定义设置保存在同一个文件中。
- **预设** — 内置 9 个 AquesTalk1 预设、16 个 AquesTalk2 预设和 8 个 AquesTalk10 预设（含 `custom`）。
- **AquesTalk10 自定义音色** — 选择基础素片（F1E / F2E / M1E），调节音高、重音、LMD 和 FSC。
- **内嵌 `.phont`** — 导入 AquesTalk2 的 `.phont` 音色文件，并将其保存在 `.phontx` 中。
- **可查看** — `readPhontx` 输出音色内容，`showEngineInfo` 报告其支持哪些引擎。
- **语音合成** — `synthAudio` 使用 [Yukw](https://github.com/yukumo-group/Yukw) WebAssembly 模块写出 WAV 文件，同一二进制可运行于 Windows、Linux 和 macOS。
- **库优先** — CLI 只是 `pkg/tone`、`pkg/aqdata`、`pkg/custom` 和 `pkg/synth` 等导出包之上的薄封装。
- **可移植格式** — `.phontx` 是由 [`tinylib/msgp`](https://github.com/tinylib/msgp) 生成的 [MessagePack](https://msgpack.org/) 映射，便于其他语言读取。

## 目录

- [环境要求](#环境要求)
- [构建与安装](#构建与安装)
- [CLI 用法](#cli-用法)
- [库用法](#库用法)
- [`.phontx` 格式](#phontx-格式)
- [预设与 AquesTalk10 自定义音色](#预设与-aquestalk10-自定义音色)
- [合成原理](#合成原理)
- [项目结构](#项目结构)
- [开发](#开发)
- [已知注意事项](#已知注意事项)
- [相关项目](#相关项目)
- [许可证与 AquesTalk](#许可证与-aquestalk)

## 环境要求

| 项目 | 说明 |
| --- | --- |
| **Go** | 1.27 或更高版本（见 [`go.mod`](go.mod)）——仅从源码构建时需要 |
| **引擎模块** | 当前工作目录下的 `aq_wasm.wasm`，由 [Yukw](https://github.com/yukumo-group/Yukw) 构建（见下文） |
| **测试** | `PATH` 中可通过 `python` 调用 Python 3.13（部分测试会运行 Python 片段以生成期望的字节数据） |
| **CGO / 原生 SDK** | 不需要。phontx 是纯 Go 实现；引擎模块由 Yukw 提供，并由 [wazero](https://wazero.io/) 执行 |
| **重新构建引擎（可选）** | 需要 Rust 及 `wasm32-unknown-unknown` 目标，以及 AQUEST 可重定位归档——仅当你自行从 Yukw 构建 `aq_wasm.wasm` 时需要 |

### `aq_wasm.wasm` 引擎模块

`phontx` 不链接 AquesTalk 库，也不包含 C 代码。其唯一的外部运行时依赖是由 [`Yukw`](https://github.com/yukumo-group/Yukw) 构建的引擎模块——Yukw 是面向 AquesTalk 的 Rust/WebAssembly 运行时。Yukw 将 AQUEST 引擎归档链接为单个 wasm 模块（`aq_wasm.wasm`，0 个导入，一个线性内存），并暴露 `pkg/synth` 通过 [wazero](https://wazero.io/) 调用的 ABI。

`cmd/main.go` 会加载 `./aq_wasm.wasm`（相对于**当前工作目录**），并在任何命令执行前编译一次。`pkg/synth` 使用的导出为：

- `aq_synthesize` — 根据语音记号和 40 字节参数块合成 WAV 缓冲区。
- `aq_error_message` — 将引擎错误码转换为错误信息。

由于每个子命令启动时都会初始化运行时，即使只运行 `readPhontx` 也必须存在该文件；在不含该文件的目录中启动 `phontx` 会立即失败。

仓库中已提交的 `pkg/synth/testdata/aq_wasm.wasm` 是测试套件所使用的 Yukw 模块预构建副本。本地试验时可直接复制：

```bash
cp pkg/synth/testdata/aq_wasm.wasm .
```

如需重新构建该模块，请使用 Yukw。它需要 Rust 的 `wasm32-unknown-unknown` 目标，以及由 [`aquestalk-wasm`](https://github.com/yukumo-group/aquestalk-wasm) 生成的可重定位 AQUEST 引擎归档：

```bash
git clone https://github.com/yukumo-group/Yukw.git
cd Yukw
just wasm
# 等价于：
# cargo build -p aq-wasm --target wasm32-unknown-unknown --release --lib
# 输出：target/wasm32-unknown-unknown/release/aq_wasm.wasm
```

将生成的文件以 `aq_wasm.wasm` 为名放入运行 phontx 的工作目录。这些前置条件属于 Yukw，而非 phontx。Yukw 目前提供的是引擎评估版；生产环境请使用获得 AQUEST 授权的构建。

## 构建与安装

```bash
git clone https://github.com/yukumo-group/phontx.git
cd phontx
go mod download

# 构建 CLI。
go build -trimpath -o phontx ./cmd

# 引擎模块必须位于工作目录。
cp pkg/synth/testdata/aq_wasm.wasm .

./phontx --help
```

在 Windows 上，`build.bat` 会执行构建并写出 `build\phontx.exe`：

```powershell
.\build.bat
Copy-Item pkg\synth\testdata\aq_wasm.wasm build\
cd build
.\phontx.exe --help
```

`build.bat` 只构建可执行文件，不会复制 `aq_wasm.wasm`。

## CLI 用法

```text
phontx
├── createPhontx        交互式向导，写出 .phontx 文件
├── readPhontx          输出 .phontx 文件内容
│   └── showEngineInfo  报告文件支持的引擎
└── synthAudio          从 .phontx 文件渲染 WAV
```

### `phontx createPhontx`

交互式向导。对每个引擎回答 `y`/`n`；对于 AquesTalk2，可在预设和 `.phont` 文件之间选择；对于 AquesTalk10，可在预设和自定义音色之间选择。最后输入输出目录和音色名称，phontx 会写出 `<目录>/<名称>.phontx`。

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

输出目录必须已存在；`SaveFile` 不会创建目录。

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

转储中的字段顺序可能变化。内嵌的 `.phont` 字节不会被打印。

### `phontx readPhontx showEngineInfo`

```console
$ phontx readPhontx showEngineInfo -f out/demo-tone.phontx
Engines supported by out/demo-tone.phontx:
Aquestalk1: Supported
Aquestalk2: Supported
Aquestalk10: Supported
```

### `phontx synthAudio`

引擎版本在运行时交互选择，且必须是该音色启用的引擎之一。

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

| 参数 | 短参数 | 必填 | 默认值 | 说明 |
| --- | --- | --- | --- | --- |
| `--filePath` | `-f` | 是 | | `.phontx` 文件路径 |
| `--targetText` | `-t` | 是 | | 要朗读的文本，需使用 AquesTalk 语音记号（假名）；不会进行汉字转换 |
| `--targetAudioFileName` | `-a` | 是 | | 输出文件名，不带 `.wav` 后缀 |
| `--targetAudioDirectory` | `-d` | 否 | 当前目录 | 输出目录 |
| `--targetSpeed` | `-s` | 否 | `-1` | 语速，限制在 50–300（100 为正常值） |
| `--targetVolume` | `-v` | 否 | `-1` | 音量，限制在 0–300（AquesTalk10 中 100 为正常值） |

使用参数默认值时，`-1` 会在限制前转换为 `uint32`，因而选中最大值而非合理默认值。请像上面的示例一样显式传入值（通常为 `100`）。输出目录必须已存在。

## 库用法

下面的示例构建一个适用于全部三种引擎的音色，保存后重新加载，并渲染 WAV 文件。它是完整程序；请放在一个依赖 `github.com/yukumo-group/phontx` 的模块中。

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

当工作目录中存在 `aq_wasm.wasm` 时，输出为：

```text
saved: out/my-tone.phontx
Aq1=true Aq2=true Aq10=true
wrote out/hello.wav
```

### API 概览

| 包 | 提供的内容 |
| --- | --- |
| `pkg/tone` | `Data`（`.phontx` 模型）、`NewData`、`NewDataFromFile`、各设置方法、`SaveFile`、`SaveFileWithFullPath`、`Equals`、`GetAllowedEngines`、`IsVersionAllowed`、`GetPresetID` |
| `pkg/aqdata` | `AquesVersion`（`AqVer1`、`AqVer2`、`AqVer10`）、泛型 `PresetList`，以及所有预设常量（`Aq1F1`、`Aq2AqDefo1`、`Aq10Custom` 等） |
| `pkg/custom` | `Aq10CustomTone` 及其 `WithBase`、`WithPitch`、`WithAccent`、`WithLMD`、`WithFSC`，以及 `Aq10Base` 值 `F1E`、`F2E`、`M1E` |
| `pkg/synth` | `GeneratorRuntime`（只编译一次 wasm）、`Generator`（单个模块实例）、`Param`、`NewSynthParam`、`PackData` |

一些值得了解的行为：

- `SaveFile(dir, name)` 总会追加 `.phontx` 并返回最终路径；`SaveFileWithFullPath` 则严格按给定路径写入。
- `SetAq10CustomTone` 还会将 `Aq10Preset` 设为 `custom`，以便引擎使用自定义参数。
- `SetPhontDataFromFile` 只接受扩展名为 `.phont` 的文件；其字节会原样嵌入 `.phontx`。
- `NewSynthParam` 会将语速限制为 50–300，将音量限制为 0–300。
- 一个 `Generator` 拥有一个 wasm 模块实例及其线性内存。每次合成（或每个 goroutine）创建一个，并在完成后关闭；`GeneratorRuntime` 可以长期存活并创建多个 generator。

## `.phontx` 格式

`.phontx` 文件是 MessagePack 编码的 `tone.Data` 结构体。文件没有头部或魔数：顶层值是以字段名为键的映射。所有字段都是可选的，`nil` 字段表示“该引擎未配置”。

| 字段 | 类型 | 含义 |
| --- | --- | --- |
| `Aq1Preset` | 预设 | AquesTalk1 音色预设（`Index` + `Name`） |
| `Aq2Preset` | 预设 | AquesTalk2 音色预设 |
| `PhontFileData` | 字节 | AquesTalk2 `.phont` 文件的原始内容 |
| `Aq10Preset` | 预设 | AquesTalk10 预设；使用 `Aq10CustomTone` 时为 `custom` |
| `Aq10CustomTone` | 对象 | AquesTalk10 自定义参数（见下文） |

引擎支持由这些字段推导：

- 设置 `Aq1Preset` 时启用 AquesTalk1。
- 设置 `Aq2Preset` 或 `PhontFileData` 时启用 AquesTalk2。
- 设置 `Aq10Preset` 或 `Aq10CustomTone` 时启用 AquesTalk10。

预设同时保存列表索引和名称：

```go
type Preset struct {
	Index uint32
	Name  string
}
```

从概念上看，文件如下所示（以 JSON 展示；实际编码为 MessagePack，且 `PhontFileData` 是二进制值，而非 base64）：

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

编码器/解码器由 `tinylib/msgp` 生成到 `*_gen.go` 文件中。解码时会跳过未知键，缺失字段保持零值，因此格式具有一定容忍度，但**没有版本号**——结构体变更即格式变更。修改带 `//go:generate msgp` 的结构体后，请重新生成代码：

```bash
go generate ./...
```

## 预设与 AquesTalk10 自定义音色

预设索引按 `pkg/aqdata/presets.go` 中的声明顺序分配。Go 常量命名模式为 `Aq1<Name>`、`Aq2<Name>` 和 `Aq10<Name>`。

### AquesTalk1（9 个音色）

| 索引 | 名称 | 常量 |
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

### AquesTalk2（16 个音色）

| 索引 | 名称 | 常量 |
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

### AquesTalk10（7 个音色 + custom）

| 索引 | 名称 | 常量 |
| --- | --- | --- |
| 0 | `custom` | `aqdata.Aq10Custom` |
| 1 | `f1` | `aqdata.Aq10F1` |
| 2 | `f2` | `aqdata.Aq10F2` |
| 3 | `f3` | `aqdata.Aq10F3` |
| 4 | `m1` | `aqdata.Aq10M1` |
| 5 | `m2` | `aqdata.Aq10M2` |
| 6 | `r1` | `aqdata.Aq10R1` |
| 7 | `r2` | `aqdata.Aq10R2` |

### AquesTalk10 自定义音色

`custom` 预设会启用逐音色参数。超出文档范围的值会被 `With*` 方法限制到范围内。

| 字段 | 范围 | 默认值 | 说明 |
| --- | --- | --- | --- |
| `Base` | `0` F1E、`1` F2E、`2` M1E | `F1E` | 基础音色素片 |
| `LMD` | 0–200 | 100 | 音高参数 1 |
| `FSC` | 50–200 | 100 | 音高参数 2（采样频率） |
| `Pitch` | 20–200 | 取决于基础素片：F1E 100、F2E 77、M1E 30 | `nil` 表示“使用基础素片默认值” |
| `Accent` | 0–200 | 取决于基础素片：F1E 100、F2E 150、M1E 100 | `nil` 表示“使用基础素片默认值” |

```go
customTone := custom.NewAq10CustomTone().
	WithBase(custom.F2E).
	WithPitch(90).
	WithAccent(140).
	WithLMD(110).
	WithFSC(120)

toneData := tone.NewData().SetAq10CustomTone(customTone)
```

上述参数仅在 AquesTalk10 预设为 `custom` 时生效。选择 `f1`...`r2` 时，引擎会使用预设自身的调音。

## 合成原理

1. `tone.NewDataFromFile` 读取文件并将 MessagePack 映射反序列化为 `tone.Data`。
2. `synth.NewGenerator` 从编译后的 Yukw `aq_wasm.wasm` 实例化一个新的 WebAssembly 模块；每个 generator 拥有独立的线性内存。
3. `Generator.Generate` 执行以下步骤：
   - 验证音色启用了所请求的引擎版本；
   - 将内嵌的 `.phont` 字节（如有）写入模块内存偏移 0；
   - 打包引擎参数并写入 `.phont` 字节之后；
   - 写入 UTF-8 语音记号文本；
   - 调用 `aq_synthesize`；若返回的错误码非零，则调用 `aq_error_message` 转换为错误；
   - 从线性内存读回 WAV 字节并写入输出文件。
4. 各缓冲区保持 4 字节对齐（`internal/bts.NextOffset`），以便 wasm 侧访问自然对齐的字段。

`synth.PackData` 生成的参数块是一系列小端 `uint32`：

| 偏移 | 字段 |
| --- | --- |
| 0 | 版本（`1`、`2` 或 `10`） |
| 4 | 预设索引 |
| 8 | 语速 |
| 12 | 内嵌 phont 数据偏移（`0`） |
| 16 | AquesTalk10 基础素片 |
| 20 | 音量 |
| 24 | 音高 |
| 28 | 重音 |
| 32 | LMD |
| 36 | FSC |

总共 40 字节。内嵌 `.phont` 的长度会单独传给 `aq_synthesize`，文本则是独立的 UTF-8 缓冲区。

## 项目结构

```text
phontx/
├── cmd/                     CLI 入口和 Cobra 命令
│   ├── main.go              加载 aq_wasm.wasm，然后运行 CLI
│   ├── cmd.go               根命令和参数定义
│   ├── create_phontx.go     交互式 .phontx 构建器
│   ├── read_phontx.go       输出 .phontx 文件
│   ├── show_engine_info.go  showEngineInfo 子命令
│   ├── synth.go             synthAudio 命令
│   ├── prompt.go            提示辅助函数
│   └── welcome.go           欢迎横幅
├── internal/
│   ├── bts/                 字节/字符串辅助函数和偏移对齐
│   ├── python/              运行 Python 片段的测试辅助包
│   └── ui/                  交互提示、向导、结构体转储
├── pkg/
│   ├── aqdata/              引擎版本和预设列表
│   ├── custom/              AquesTalk10 自定义音色参数
│   ├── synth/               wazero 运行时和 WAV 合成
│   └── tone/                .phontx 数据模型和序列化/反序列化
├── build/                   构建输出和运行时资源（git 忽略）
├── build.bat                Windows 构建辅助脚本
└── go.mod
```

| 包 | 职责 |
| --- | --- |
| `pkg/tone` | `Data` 模型、加载/保存、引擎支持检查 |
| `pkg/aqdata` | `AquesVersion`、预设定义、泛型 `PresetList` |
| `pkg/custom` | `Aq10CustomTone` 与 `Aq10Base` |
| `pkg/synth` | `GeneratorRuntime`、`Generator`、`Param`、`PackData` |
| `internal/ui` | 交互提示、预设/自定义音色向导、`ToneToString`、`ShowEngineInfo` |
| `internal/bts` | 字节比较、空终止字符串读取、偏移对齐 |
| `internal/python` | `ExecutePython`，测试用它生成期望的字节数据 |

## 开发

```bash
go test ./...                     # 所有包
go test ./pkg/synth ./pkg/tone    # 部分包
go vet ./...
golint -set_exit_status ./...     # CI 执行的检查
```

- 部分测试会执行 Python（`internal/python` 以及 `pkg/synth` 中的参数打包测试），因此 `python` 必须能在 `PATH` 中解析到 Python 3.13。
- `pkg/synth` 的测试会使用 `pkg/synth/testdata/aq_wasm.wasm` 进行真实合成，并将输出写入 `t.TempDir()`。
- 代码生成：先执行 `go install github.com/tinylib/msgp@latest` 安装生成器，然后在仓库根目录运行 `go generate ./...`。不要手工编辑 `*_gen.go` 文件。
- CI（`.github/workflows/test.yaml`）在 Ubuntu 上运行，使用 `go.mod` 指定的 Go 版本和 Python 3.13，然后运行 `golint` 并将覆盖率上传到 Codecov。
- 发布流程（`.github/workflows/release.yaml`）在推送 `v*` 标签时运行 GoReleaser，并发布由 `./cmd` 构建的 `phontx` 二进制。

## 已知注意事项

- **每个命令都需要 `aq_wasm.wasm`**，不仅是 `synthAudio`，因为 `cmd/main.go` 会在 Cobra 运行前初始化 wazero 运行时。请从包含该模块的目录启动 `phontx`。
- **不要依赖 `-s`/`-v` 的默认值。** 它们默认为 `-1`，转换为 `uint32` 后变成 `4294967295`，随后被限制为 `300`。除非确实需要最高语速和音量，否则请传入 `-s 100 -v 100`。
- **交互输入以空白分隔。** 提示使用 `fmt.Scan`，因此无法输入包含空格的路径；是/否提示必须显式输入 `y` 或 `n`（按 Enter 不会接受显示的默认值）。
- **没有汉字转换。** `-t` 的文本会原样传给引擎；请提供 AquesTalk 语音记号（假名和语音符号）。本项目不包含 AqKanji2Koe。
- **`.phont` 导入要求严格。** `SetPhontDataFromFile` 会拒绝除 `.phont` 之外的扩展名（比较区分大小写）。
- **AquesTalk10 自定义参数需要 `custom` 预设。** 对于内置 AquesTalk10 预设，音高、重音、LMD 和 FSC 会被忽略。
- **`.phontx` 格式没有版本号。** msgp 会跳过未知键，但对 `tone.Data` 的结构性修改就是格式变更。

## 相关项目

- [Yukw](https://github.com/yukumo-group/Yukw) — 构建 `aq_wasm.wasm` 的 Rust/WebAssembly 运行时；phontx 使用其 `aq_synthesize` ABI。
- [aquestalk-wasm](https://github.com/yukumo-group/aquestalk-wasm) — 上游引擎项目，生成 Yukw 链接进模块的可重定位归档。

## 许可证与 AquesTalk

本仓库的源代码以 [MIT 许可证](LICENSE) 发布（Copyright © 2026 Yukumo Group）。

AquesTalk1、AquesTalk2 和 AquesTalk10 是 [AQUEST Corp.](https://www.a-quest.com/) 的商业产品。引擎二进制、音色数据和 WebAssembly 构建（包括 Yukw 构建的 `aq_wasm.wasm`）**不**受 MIT 许可证覆盖，也不由本仓库分发。Yukw 目前发布的是评估版；生产环境请获取 AQUEST 授权的引擎构建，并在合成或再分发音频时遵守 AQUEST 的条款。
