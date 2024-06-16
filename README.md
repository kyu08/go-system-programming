# 第1章 Go言語で覗くシステムプログラミングの世界
- デバッガを使って`fmt.Println`の動作をシステムコールまで追った。[^1]
- 当然だがデバッガを使うと変数の中身に何が入ってるかがわかって便利。

# 第2章 低レベルアクセスへの入口1：io.Writer
- OSでは、システムコールを **ファイルディスクリプタ**と呼ばれるものに対して呼ぶ。
- ファイルディスクリプタ(file descriptor)は一種の識別子（数値）で、この数値を指定してシステムコールを呼び出すと、数値に対応するモノにアクセスできる。
- ファイルディスクリプタはOSがカーネルのレイヤーで用意している抽象化の仕組み。
- OSのカーネル内部のデータベースに、プロセスごとに実体が用意される。
- OSはプロセスが起動されるとまず3つの擬似ファイルを作成し、それぞれにファイルディスクリプタを割り当てる。
- 0が標準入力、1が標準出力、2が標準エラー出力。
- 以降はそのプロセスでファイルをオープンしたりソケットをオープンしたりするたびに1ずつ大きな数値が割り当てられていく。

`os.Create()`する処理をデバッガーで追ってみたところ`Sysfd int = 3`[^2]となっており上記の記述が確認できた。

<img width="1507" alt="image" src="https://github.com/kyu08/go-system-programming/assets/49891479/d35a0689-5188-4f28-ba7c-3a12009ed273">

# 第3章 低レベルアクセスへの入口1：io.Reader
- エンディアン変換
    - リトルエンディアンでは、10000という数値(`0x2710`)をメモリに格納するときに下位バイトから順に格納する。
    - ビッグエンディアンでは、上位バイトから順に格納する。
    - 現在主流のCPUではリトルエンディアンが採用されている。
    - ネットワーク上で転送されるデータの多くはビッグエンディアンが用いられている。
    - そのため多くの環境ではネットワークで受け取ったデータをリトルエンディアンに変換する必要がある。
- `io`パッケージのいくつかの関数 / 構造体 / インターフェースの使い方
    - `io.Pipe`[^3]
    - `io.LimitReader`: 先頭の`n`バイトだけ読み込む
    - `io.MultiReader`: 複数の`io.Reader`を1つの`io.Reader`にまとめる
    - `io.SectionReader`: `offset`と`n`を指定して一部のデータだけ読み込む

# 第4章 チャネル
普段goroutineを全然使わないので忘れていることが多かった。以下学んだことメモ。

- バッファなしチャネルでは、受け取り側が受信しないと、送信側もブロックされる。[^4]
- `for task := range tasks // tasksは任意のch`のように書くと、チャネルに値が入るたびにループが回り、チャネルがクローズされるまでループが回る。

# 第5章 システムコール


[^1]: Neovimでのデバッガの環境構築は [nvim-dapでGolangのデバッグ環境構築](https://zenn.dev/saito9/articles/32c57f776dc369) を参考にした
[^2]: `Sysfd`の定義は golang/go/src/internal/poll/fd_unix.go#L23(https://github.com/golang/go/blob/c83b1a7013784098c2061ae7be832b2ab7241424/src/internal/poll/fd_unix.go#L23) にある。
[^3]: `io.Pipe`の使いどころに関しては [Go言語のio.Pipeでファイルを効率よくアップロードする方法](https://medium.com/eureka-engineering/file-uploads-in-go-with-io-pipe-75519dfa647b) が大変参考になった。
[^4]: cf. [kyu08/go-system-programming/4-channel/unbufferedchannel/main.go#L8](https://github.com/kyu08/go-system-programming/blob/b9da4a0ce759b2df4ce884ab61248fb893b60bef/4-channel/unbufferedchannel/main.go#L8)
