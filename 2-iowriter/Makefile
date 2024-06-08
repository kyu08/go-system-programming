# `make new ARG=4-3`みたいな感じで呼ぶ
.PHONY: new
new:
	mkdir ${ARG}
	cd ${ARG} \
		&& go mod init github.com/kyu08/go-system-programming/${ARG} \
		&& touch main.go \
		&& echo "package main \n\nfunc main() {}" > main.go
