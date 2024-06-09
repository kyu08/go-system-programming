# 第1章 Go言語で覗くシステムプログラミングの世界
- デバッガを使って`fmt.Println`の動作をシステムコールまで追った。[^1]
- 当然だけどデバッガを使うと変数の中身に何が入ってるかがわかって便利。


# 第2章 低レベルアクセスへの入口1：io.Writer
- OSでは、システムコールを ** ファイルディスクリプタ ** といわれるものに対して呼ぶ。
- ファイルディスクリプタ(file descriptor)は一種の識別子（数値）で、この数値を指定してシステムコールを呼び出すと、数値に対応するモノにアクセスできる。
- ファイルディスクリプタはOSがカーネルのレイヤーで用意している抽象化の仕組み。
- OSのカーネル内部のデータベースに、プロセスごとに実体が用意される。
- OSはプロセスが起動されるとまず3つの擬似ファイルを作成し、それぞれにファイルディスクリプタを割り当てる。
- 0が標準入力、1が標準出力、2が標準エラー出力。
- 以降はそのプロセスでファイルをオープンしたりソケットをオープンしたりするたびに1ずつ大きな数値が割り当てられていく。

`os.Create()`する処理をデバッガーで追ってみたところ`Sysfd int = 3`[^2]となっており上記の記述が確認できた。

<img width="1507" alt="image" src="https://github.com/kyu08/go-system-programming/assets/49891479/d35a0689-5188-4f28-ba7c-3a12009ed273">


# 第3章 低レベルアクセスへの入口1：io.Reader


[^1]: Neovimでのデバッガの環境構築は [nvim-dapでGolangのデバッグ環境構築](https://zenn.dev/saito9/articles/32c57f776dc369) を参考にした
[^2]: `Sysfd`の定義は golang/go/src/internal/poll/fd_unix.go#L23(https://github.com/golang/go/blob/c83b1a7013784098c2061ae7be832b2ab7241424/src/internal/poll/fd_unix.go#L23) にある。
