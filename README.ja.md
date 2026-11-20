# phontx

[English](README.md) · [中文](README.zh.md) · **日本語**

[![Test Workflow](https://github.com/yukumo-group/phontx/actions/workflows/test.yaml/badge.svg)](https://github.com/yukumo-group/phontx/actions/workflows/test.yaml)
[![Go Version](https://img.shields.io/github/go-mod/go-version/yukumo-group/phontx)](go.mod)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

`phontx` は、**AquesTalk ファミリーの日本語音声合成エンジン向けの統一音色ファイル**を扱う Go ライブラリ兼 CLI です。1 つの `.phontx` ファイルで **AquesTalk1**、**AquesTalk2**、**AquesTalk10** の 3 エンジンにおける声の設定をまとめて記述でき、CLI は [wazero](https://wazero.io/) 経由で [Yukw](https://github.com/yukumo-group/Yukw) エンジンモジュール（Rust → WebAssembly）を実行し、その音色を WAV ファイルにレンダリングします。実行時に CGO、C ツールチェーン、ネイティブ SDK は不要です。

> `.phontx` は `.phont` ではありません。`.phont` は AquesTalk2 のネイティブ音色データファイルです。`.phontx` は phontx 独自のコンテナで、3 エンジン分のプリセットを保持し、必要に応じて `.phont` ファイルを埋め込むことができます。

## 機能

- **1 つの音色、3 つのエンジン** — AquesTalk1/2/10 のプリセットとカスタム設定を同じファイルに保存します。
- **プリセット** — AquesTalk1 が 9 個、AquesTalk2 が 16 個、AquesTalk10 が 8 個（`custom` を含む）定義されています。
- **AquesTalk10 カスタム音色** — 基本素片（F1E / F2E / M1E）を選び、ピッチ、アクセント、LMD、FSC を調整できます。
- **`.phont` の埋め込み** — AquesTalk2 の `.phont` 音色をインポートし、`.phontx` 内に保持します。
- **内容の確認** — `readPhontx` が音色をダンプし、`showEngineInfo` が対応エンジンを表示します。
- **音声合成** — `synthAudio` が [Yukw](https://github.com/yukumo-group/Yukw) WebAssembly モジュールで WAV ファイルを書き出すため、同じバイナリが Windows、Linux、macOS で動作します。
- **ライブラリ優先** — CLI はエクスポートされた `pkg/tone`、`pkg/aqdata`、`pkg/custom`、`pkg/synth` パッケージの薄いラッパーです。
- **可搬性の高い形式** — `.phontx` は [`tinylib/msgp`](https://github.com/tinylib/msgp) で生成された [MessagePack](https://msgpack.org/) マップなので、他の言語からも読みやすい形式です。

## 目次

- [必要要件](#必要要件)
- [ビルドとインストール](#ビルドとインストール)
- [CLI の使い方](#cli-の使い方)
- [ライブラリとしての使い方](#ライブラリとしての使い方)
- [`.phontx` フォーマット](#phontx-フォーマット)
- [プリセットと AquesTalk10 カスタム音色](#プリセットと-aquestalk10-カスタム音色)
- [音声合成の仕組み](#音声合成の仕組み)
- [プロジェクト構成](#プロジェクト構成)
- [開発](#開発)
- [既知の注意点](#既知の注意点)
- [関連プロジェクト](#関連プロジェクト)
- [ライセンスと AquesTalk](#ライセンスと-aquestalk)

## 必要要件

| 項目 | 内容 |
| --- | --- |
| **Go** | 1.27 以上（[`go.mod`](go.mod) を参照）— ソースからビルドする場合のみ必要 |
| **エンジンモジュール** | カレントディレクトリの `aq_wasm.wasm`（[Yukw](https://github.com/yukumo-group/Yukw) でビルド。後述） |
| **テスト** | `PATH` 上の `python` で Python 3.13 を実行できること（一部のテストが Python スニペットを実行して期待バイト列を生成します） |
| **CGO / ネイティブ SDK** | 不要。phontx は純粋な Go 実装で、エンジンモジュールは Yukw が提供し、[wazero](https://wazero.io/) が実行します |
| **エンジンの再ビルド（任意）** | Rust の `wasm32-unknown-unknown` ターゲットと AQUEST の再配置可能アーカイブ — Yukw から `aq_wasm.wasm` を自分でビルドする場合のみ必要 |

### `aq_wasm.wasm` エンジンモジュール

`phontx` は AquesTalk ライブラリにリンクせず、C コードも含みません。唯一の外部実行時依存は、[`Yukw`](https://github.com/yukumo-group/Yukw) がビルドするエンジンモジュールです。Yukw は AquesTalk 向けの Rust/WebAssembly ランタイムで、AQUEST のエンジンアーカイブを 1 つの wasm モジュール（`aq_wasm.wasm`、インポート 0、リニアメモリ 1 つ）にリンクし、`pkg/synth` が [wazero](https://wazero.io/) 経由で呼び出す ABI を公開します。

`cmd/main.go` は `./aq_wasm.wasm`（**カレントディレクトリ**基準）を読み込み、コマンド実行前に 1 回コンパイルします。`pkg/synth` が使用するエクスポートは次のとおりです。

- `aq_synthesize` — 音声記号列と 40 バイトのパラメータブロックから WAV バッファを合成します。
- `aq_error_message` — エンジンのエラーコードをエラーメッセージに変換します。

サブコマンドの実行前に必ずランタイムを初期化するため、`readPhontx` だけを実行する場合でもこのファイルが必要です。モジュールのないディレクトリから `phontx` を起動すると即座に失敗します。

リポジトリにコミットされている `pkg/synth/testdata/aq_wasm.wasm` は、テストスイートが使用する Yukw モジュールのビルド済みコピーです。ローカルで試す場合はこれをコピーしてください。

```bash
cp pkg/synth/testdata/aq_wasm.wasm .
```

モジュールを再ビルドするには Yukw を使用します。Rust の `wasm32-unknown-unknown` ターゲットと、[`aquestalk-wasm`](https://github.com/yukumo-group/aquestalk-wasm) が生成する再配置可能な AQUEST エンジンアーカイブが必要です。

```bash
git clone https://github.com/yukumo-group/Yukw.git
cd Yukw
just wasm
# 同等:
# cargo build -p aq-wasm --target wasm32-unknown-unknown --release --lib
# 出力: target/wasm32-unknown-unknown/release/aq_wasm.wasm
```

生成されたファイルを `aq_wasm.wasm` という名前で、phontx を実行するカレントディレクトリに配置してください。これらの前提条件は Yukw 側のもので、phontx には不要です。Yukw は現在エンジンの評価版を生成します。本番環境では AQUEST のライセンス済みビルドを使用してください。

## ビルドとインストール

```bash
git clone https://github.com/yukumo-group/phontx.git
cd phontx
go mod download

# CLI をビルド。
go build -trimpath -o phontx ./cmd

# エンジンモジュールはカレントディレクトリに置く必要があります。
cp pkg/synth/testdata/aq_wasm.wasm .

./phontx --help
```

Windows では `build.bat` がビルドを実行し、`build\phontx.exe` を出力します。

```powershell
.\build.bat
Copy-Item pkg\synth\testdata\aq_wasm.wasm build\
cd build
.\phontx.exe --help
```

`build.bat` は実行ファイルをビルドするだけで、`aq_wasm.wasm` はコピーしません。

## CLI の使い方

```text
phontx
├── createPhontx        .phontx ファイルを書き出す対話ウィザード
├── readPhontx          .phontx ファイルをダンプ
│   └── showEngineInfo  対応エンジンを表示
└── synthAudio          .phontx ファイルから WAV をレンダリング
```

### `phontx createPhontx`

対話式ウィザードです。各エンジンについて `y`/`n` で回答し、AquesTalk2 ではプリセットか `.phont` ファイルかを選択し、AquesTalk10 ではプリセットかカスタム音色かを選択します。最後に出力ディレクトリと音色名を入力すると、phontx が `<ディレクトリ>/<名前>.phontx` を書き出します。

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

出力ディレクトリは事前に存在している必要があります。`SaveFile` はディレクトリを作成しません。

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

ダンプ内のフィールド順は変わる場合があります。埋め込まれた `.phont` のバイト列は表示されません。

### `phontx readPhontx showEngineInfo`

```console
$ phontx readPhontx showEngineInfo -f out/demo-tone.phontx
Engines supported by out/demo-tone.phontx:
Aquestalk1: Supported
Aquestalk2: Supported
Aquestalk10: Supported
```

### `phontx synthAudio`

エンジンのバージョンは対話的に選択し、その音色で有効になっているエンジンのいずれかである必要があります。

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

| フラグ | 短縮 | 必須 | デフォルト | 説明 |
| --- | --- | --- | --- | --- |
| `--filePath` | `-f` | はい | | `.phontx` ファイルのパス |
| `--targetText` | `-t` | はい | | 読み上げるテキスト（AquesTalk の音声記号列・かな）。漢字変換は行いません |
| `--targetAudioFileName` | `-a` | はい | | 出力ファイル名（`.wav` 拡張子は不要） |
| `--targetAudioDirectory` | `-d` | いいえ | カレントディレクトリ | 出力ディレクトリ |
| `--targetSpeed` | `-s` | いいえ | `-1` | 話速。50–300 にクランプ（100 が標準） |
| `--targetVolume` | `-v` | いいえ | `-1` | 音量。0–300 にクランプ（AquesTalk10 では 100 が標準） |

フラグのデフォルトでは `-1` がクランプ前に `uint32` へ変換されるため、妥当なデフォルトではなく最大値が選ばれます。上の例のように明示的に値（通常は `100`）を指定してください。出力ディレクトリは事前に存在している必要があります。

## ライブラリとしての使い方

以下の例は 3 エンジンすべてで動作する音色を構築し、保存して読み戻し、WAV ファイルをレンダリングします。完全なプログラムです。`github.com/yukumo-group/phontx` を require するモジュールに配置してください。

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

カレントディレクトリに `aq_wasm.wasm` がある場合の出力:

```text
saved: out/my-tone.phontx
Aq1=true Aq2=true Aq10=true
wrote out/hello.wav
```

### API 概要

| パッケージ | 提供内容 |
| --- | --- |
| `pkg/tone` | `Data`（`.phontx` モデル）、`NewData`、`NewDataFromFile`、各種セッター、`SaveFile`、`SaveFileWithFullPath`、`Equals`、`GetAllowedEngines`、`IsVersionAllowed`、`GetPresetID` |
| `pkg/aqdata` | `AquesVersion`（`AqVer1`、`AqVer2`、`AqVer10`）、汎用 `PresetList`、すべてのプリセット定数（`Aq1F1`、`Aq2AqDefo1`、`Aq10Custom` など） |
| `pkg/custom` | `Aq10CustomTone` と `WithBase`、`WithPitch`、`WithAccent`、`WithLMD`、`WithFSC`、および `Aq10Base` の値 `F1E`、`F2E`、`M1E` |
| `pkg/synth` | `GeneratorRuntime`（wasm を 1 回コンパイル）、`Generator`（1 モジュールインスタンス）、`Param`、`NewSynthParam`、`PackData` |

知っておくと便利な動作:

- `SaveFile(dir, name)` は常に `.phontx` を付加して最終パスを返します。`SaveFileWithFullPath` は指定したパスにそのまま書き込みます。
- `SetAq10CustomTone` は `Aq10Preset` も `custom` に設定するため、エンジンはカスタムパラメータを使用します。
- `SetPhontDataFromFile` は拡張子が `.phont` のファイルのみ受け付けます。バイト列はそのまま `.phontx` に埋め込まれます。
- `NewSynthParam` は話速を 50–300、音量を 0–300 にクランプします。
- `Generator` は 1 つの wasm モジュールインスタンスとそのリニアメモリを所有します。合成ごと（または goroutine ごと）に作成し、完了後に閉じてください。`GeneratorRuntime` は存続させたまま、多数の generator を作成できます。

## `.phontx` フォーマット

`.phontx` ファイルは MessagePack でエンコードされた `tone.Data` 構造体です。ヘッダーやマジックナンバーはなく、トップレベルの値はフィールド名をキーとするマップです。すべてのフィールドは省略可能で、`nil` のフィールドは「そのエンジンは未設定」を意味します。

| フィールド | 型 | 意味 |
| --- | --- | --- |
| `Aq1Preset` | プリセット | AquesTalk1 の音声プリセット（`Index` + `Name`） |
| `Aq2Preset` | プリセット | AquesTalk2 の音声プリセット |
| `PhontFileData` | バイト列 | AquesTalk2 `.phont` ファイルの生データ |
| `Aq10Preset` | プリセット | AquesTalk10 プリセット。`Aq10CustomTone` 使用時は `custom` |
| `Aq10CustomTone` | オブジェクト | AquesTalk10 のカスタムパラメータ（後述） |

エンジン対応はこれらのフィールドから決まります。

- `Aq1Preset` が設定されていれば AquesTalk1 が有効。
- `Aq2Preset` または `PhontFileData` が設定されていれば AquesTalk2 が有効。
- `Aq10Preset` または `Aq10CustomTone` が設定されていれば AquesTalk10 が有効。

プリセットはリスト内のインデックスと名前の両方を保持します。

```go
type Preset struct {
	Index uint32
	Name  string
}
```

概念的には次のようなファイルです（JSON で表示。実際のエンコードは MessagePack で、`PhontFileData` は base64 ではなくバイナリ値です）。

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

エンコーダー/デコーダーは `tinylib/msgp` により `*_gen.go` に生成されます。デコード時は未知のキーをスキップし、欠落フィールドはゼロ値のままにするため、形式には一定の寛容さがありますが、**バージョン管理はされていません**。構造体の変更はフォーマットの変更になります。`//go:generate msgp` 構造体を編集したらコードを再生成してください。

```bash
go generate ./...
```

## プリセットと AquesTalk10 カスタム音色

プリセットのインデックスは `pkg/aqdata/presets.go` の宣言順に割り当てられます。Go 定数は `Aq1<Name>`、`Aq2<Name>`、`Aq10<Name>` のパターンに従います。

### AquesTalk1（9 音色）

| インデックス | 名前 | 定数 |
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

### AquesTalk2（16 音色）

| インデックス | 名前 | 定数 |
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

### AquesTalk10（7 音色 + custom）

| インデックス | 名前 | 定数 |
| --- | --- | --- |
| 0 | `custom` | `aqdata.Aq10Custom` |
| 1 | `f1` | `aqdata.Aq10F1` |
| 2 | `f2` | `aqdata.Aq10F2` |
| 3 | `f3` | `aqdata.Aq10F3` |
| 4 | `m1` | `aqdata.Aq10M1` |
| 5 | `m2` | `aqdata.Aq10M2` |
| 6 | `r1` | `aqdata.Aq10R1` |
| 7 | `r2` | `aqdata.Aq10R2` |

### AquesTalk10 カスタム音色

`custom` プリセットで音色ごとのパラメータが有効になります。範囲外の値は `With*` メソッドでクランプされます。

| フィールド | 範囲 | デフォルト | 備考 |
| --- | --- | --- | --- |
| `Base` | `0` F1E、`1` F2E、`2` M1E | `F1E` | 基本素片 |
| `LMD` | 0–200 | 100 | ピッチパラメータ 1 |
| `FSC` | 50–200 | 100 | ピッチパラメータ 2（サンプリング周波数） |
| `Pitch` | 20–200 | 基本素片に依存: F1E 100、F2E 77、M1E 30 | `nil` は「基本素片のデフォルトを使用」 |
| `Accent` | 0–200 | 基本素片に依存: F1E 100、F2E 150、M1E 100 | `nil` は「基本素片のデフォルトを使用」 |

```go
customTone := custom.NewAq10CustomTone().
	WithBase(custom.F2E).
	WithPitch(90).
	WithAccent(140).
	WithLMD(110).
	WithFSC(120)

toneData := tone.NewData().SetAq10CustomTone(customTone)
```

上記のパラメータは AquesTalk10 のプリセットが `custom` のときのみ有効です。`f1`...`r2` を選ぶと、エンジンはプリセット固有のチューニングを使用します。

## 音声合成の仕組み

1. `tone.NewDataFromFile` がファイルを読み込み、MessagePack マップを `tone.Data` にアンマーシャルします。
2. `synth.NewGenerator` がコンパイル済みの Yukw `aq_wasm.wasm` から新しい WebAssembly モジュールをインスタンス化します。各 generator は独自のリニアメモリを持ちます。
3. `Generator.Generate` は次の処理を行います。
   - 要求されたエンジンバージョンが音色で有効か検証する
   - 埋め込まれた `.phont` バイト列（あれば）をモジュールメモリのオフセット 0 に書き込む
   - エンジンパラメータをパックし、phont バイト列の後に書き込む
   - UTF-8 の音声記号列テキストを書き込む
   - `aq_synthesize` を呼び、返されたエラーコードが非ゼロなら `aq_error_message` でエラーに変換する
   - リニアメモリから WAV バイト列を読み戻し、出力ファイルに書き込む
4. バッファは 4 バイト境界に揃えられます（`internal/bts.NextOffset`）。wasm 側が自然に整列されたフィールドを参照できるようにするためです。

`synth.PackData` が生成するパラメータブロックは、リトルエンディアンの `uint32` の並びです。

| オフセット | フィールド |
| --- | --- |
| 0 | バージョン（`1`、`2`、`10`） |
| 4 | プリセットインデックス |
| 8 | 話速 |
| 12 | 埋め込み phont データのオフセット（`0`） |
| 16 | AquesTalk10 の基本素片 |
| 20 | 音量 |
| 24 | ピッチ |
| 28 | アクセント |
| 32 | LMD |
| 36 | FSC |

合計 40 バイトです。埋め込まれた `.phont` の長さは別途 `aq_synthesize` に渡され、テキストは独立した UTF-8 バッファになります。

## プロジェクト構成

```text
phontx/
├── cmd/                     CLI エントリポイントと Cobra コマンド
│   ├── main.go              aq_wasm.wasm を読み込んで CLI を実行
│   ├── cmd.go               ルートコマンドとフラグ定義
│   ├── create_phontx.go     対話式 .phontx ビルダー
│   ├── read_phontx.go       .phontx ファイルをダンプ
│   ├── show_engine_info.go  showEngineInfo サブコマンド
│   ├── synth.go             synthAudio コマンド
│   ├── prompt.go            プロンプト補助関数
│   └── welcome.go           ウェルカムバナー
├── internal/
│   ├── bts/                 バイト/文字列補助とオフセット整列
│   ├── python/              Python スニペットを実行するテスト補助
│   └── ui/                  対話プロンプト、ウィザード、構造体ダンプ
├── pkg/
│   ├── aqdata/              エンジンバージョンとプリセットリスト
│   ├── custom/              AquesTalk10 カスタム音色パラメータ
│   ├── synth/               wazero ランタイムと WAV 合成
│   └── tone/                .phontx データモデルとシリアライズ/デシリアライズ
├── build/                   ビルド出力と実行時アセット（git 管理外）
├── build.bat                Windows ビルド補助スクリプト
└── go.mod
```

| パッケージ | 役割 |
| --- | --- |
| `pkg/tone` | `Data` モデル、読み込み/保存、エンジン対応チェック |
| `pkg/aqdata` | `AquesVersion`、プリセット定義、汎用 `PresetList` |
| `pkg/custom` | `Aq10CustomTone` と `Aq10Base` |
| `pkg/synth` | `GeneratorRuntime`、`Generator`、`Param`、`PackData` |
| `internal/ui` | 対話プロンプト、プリセット/カスタム音色ウィザード、`ToneToString`、`ShowEngineInfo` |
| `internal/bts` | バイト比較、NUL 終端文字列の読み取り、オフセット整列 |
| `internal/python` | `ExecutePython`。テストで期待バイト列の生成に使用 |

## 開発

```bash
go test ./...                     # すべてのパッケージ
go test ./pkg/synth ./pkg/tone    # 一部のパッケージ
go vet ./...
golint -set_exit_status ./...     # CI が実行するチェック
```

- 一部のテストは Python を実行します（`internal/python` と `pkg/synth` のパラメータパックテスト）。そのため `python` が `PATH` 上で Python 3.13 に解決される必要があります。
- `pkg/synth` のテストは `pkg/synth/testdata/aq_wasm.wasm` を使って実際に合成し、出力を `t.TempDir()` に書き込みます。
- コード生成: まず `go install github.com/tinylib/msgp@latest` で生成器をインストールし、リポジトリルートで `go generate ./...` を実行します。`*_gen.go` は手動編集しないでください。
- CI（`.github/workflows/test.yaml`）は Ubuntu 上で `go.mod` の Go バージョンと Python 3.13 を使い、`golint` を実行してカバレッジを Codecov にアップロードします。
- リリース（`.github/workflows/release.yaml`）は `v*` タグの push 時に GoReleaser を実行し、`./cmd` からビルドした `phontx` バイナリを公開します。

## 既知の注意点

- **すべてのコマンドで `aq_wasm.wasm` が必要です**。`synthAudio` だけでなく、`cmd/main.go` が Cobra の実行前に wazero ランタイムを初期化するためです。モジュールのあるディレクトリから `phontx` を起動してください。
- **`-s`/`-v` のデフォルト値に依存しないでください。** デフォルトは `-1` で、`uint32` に変換すると `4294967295` になり、その後 `300` にクランプされます。最大の話速・音量を本当に必要としない限り、`-s 100 -v 100` を指定してください。
- **対話入力は空白区切りです。** プロンプトは `fmt.Scan` を使うため、スペースを含むパスは入力できません。yes/no プロンプトには明示的に `y` または `n` を入力する必要があります（Enter では表示されたデフォルトを受け付けません）。
- **漢字変換はありません。** `-t` のテキストはそのままエンジンに渡されます。AquesTalk の音声記号列（かなと音声記号）を指定してください。AqKanji2Koe は同梱されていません。
- **`.phont` のインポートは厳格です。** `SetPhontDataFromFile` は `.phont` 以外の拡張子を拒否します（大文字小文字を区別）。
- **AquesTalk10 のカスタムパラメータには `custom` プリセットが必要です。** 組み込みの AquesTalk10 プリセットでは、ピッチ、アクセント、LMD、FSC は無視されます。
- **`.phontx` フォーマットはバージョン管理されていません。** msgp は未知のキーをスキップしますが、`tone.Data` の構造変更はフォーマット変更になります。

## 関連プロジェクト

- [Yukw](https://github.com/yukumo-group/Yukw) — `aq_wasm.wasm` をビルドする Rust/WebAssembly ランタイム。phontx はその `aq_synthesize` ABI を使用します。
- [aquestalk-wasm](https://github.com/yukumo-group/aquestalk-wasm) — Yukw がモジュールにリンクする再配置可能アーカイブを生成する上流のエンジンプロジェクト。

## ライセンスと AquesTalk

本リポジトリのソースコードは [MIT ライセンス](LICENSE)（Copyright © 2026 Yukumo Group）で公開されています。

AquesTalk1、AquesTalk2、AquesTalk10 は [AQUEST Corp.](https://www.a-quest.com/) の商用製品です。エンジンバイナリ、音声データ、WebAssembly ビルド（Yukw がビルドした `aq_wasm.wasm` を含む）は MIT ライセンスの対象**外**であり、本リポジトリでは配布していません。Yukw は現在評価版を提供しています。本番環境では AQUEST のライセンス済みエンジンビルドを入手し、音声の合成や再配布時には AQUEST の条件に従ってください。
