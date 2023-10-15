#!/bin/bash

# VSCodeで扱うディレクトリの所有者を Dev Container のデフォルトユーザーである vscode ユーザーに変更する
sudo chown -R vscode:vscode /workspace
