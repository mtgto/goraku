# goraku

GoでSlackボットを書くためのフレームワーク

# コンセプト

- Goの勉強がてらなので作者 (mtgto) 以外が使うのはまだおすすめできません
- ログはglogつかうかも。ひとまずlog
- botの機能はプラグイン形式でBotKitぽく
- アプリ全体でひとつのconfigをもつ (環境変数)
- nlopes/slackのデータ構造を必要なところは出してAPI叩くとかプラグインからなんでもできるようにしておく
- BotKitのDB形式を用意してプラグインから永続化できるようにしたい
- テストを書きやすいようにinterface使って実装は外から隠す

# 使い方

examples みてね

# License

MIT