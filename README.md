# cpu_info for YaneuraOu

「やねうら王」（※）のどの実行ファイルを使うか、表示してくれるツールです。  
現在対応しているものは、以下です。  
SSE41, SSE42, AVX2, AVXVNNI, AVX512, AVX512VNNI, ZEN1, ZEN2, ZEN3  
  
（※） **やねうら王**  
GitHub - yaneurao/YaneuraOu:  
YaneuraOu is the World's Strongest Shogi engine (AI player),  
 WCSC29 1st winner, educational and USI compliant engine.  
[https://github.com/yaneurao/YaneuraOu](https://github.com/yaneurao/YaneuraOu)  

## cpu_info for YaneuraOuの開発者

末吉 竜介  
[https://github.com/sueyoshiryosuke](https://github.com/sueyoshiryosuke)  

## 動作確認、開発環境

Windows11  
go version go1.20.4 windows/amd64  

## 使い方

やねうら王の実行ファイル群があるところで、  
コンパイルされた実行ファイル（cpu_info-YO.exe）を  
コマンドプロンプト等で実行して、CPUの種類を確認する。  

## 入力と表示例

入力例  
```
> cpu_info-YO.exe
```
  
出力例  
```
Vendor ID: GenuineIntel
Model Name: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
Family Number: 6
Model Number: 165
Core Count: 6
Thread Count: 12

Recommended YaneuraOu exe on this PC:
AVX2
```

## コンパイルまでの手順

### プロジェクトの設定

現在のソースコードがあるディレクトリに`go.mod`ファイルがない場合、  
プロジェクトのフォルダに移動して、Goモジュールを初期化します。  

```
go mod init <プロジェクト名>
```

プロジェクト名はフォルダ名とは関係なく、任意の名前でよいです。  
例：  
```
go mod init cpu_info-YO
```

### 必要な外部パッケージ

`github.com/klauspost/cpuid/v2`をインストールします。  
```
go get github.com/klauspost/cpuid/v2
```

#### この外部パッケージの配布元

Releases · klauspost/cpuid · GitHub  
[https://github.com/klauspost/cpuid/releases](https://github.com/klauspost/cpuid/releases)

### コンパイル

Go言語のコンパイラはデフォルトでいくつかの最適化を行いますが、  
このコマンドはデバッグ情報を省略し、ファイルサイズを小さくする最適化を行います。  
  
```
go build -ldflags "-s -w" -o <exeファイル名> <Goのソースコードファイル名>
```
  
例：  
```
go build -ldflags "-s -w" -o cpu_info-YO.exe "cpu_info-YO.go"
```

## ライセンス

MITライセンス

