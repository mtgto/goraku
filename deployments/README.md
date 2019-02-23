このディレクトリについて
====

Docker, Kubernetesで動かすためのスクリプトのサンプルを置いています。

# Dockerfile

ビルド

```console
$ docker build --file deployments/Dockerfile --tag mtgto/goraku-example .
```

実行

```console
$ docker run -d --env SLACK_BOT_TOKEN=XXXXXXXX mtgto/goraku-example
```

# Kubernetes

Kustomizeをインストールしてください。Docker hubに `mtgto/goraku-example` というイメージを作ってあるとして、

```console
$ cp deployments/kustomize/secret.ini.txt deployments/kustomize/secret.ini
$ vi deployments/kustomize/secret.ini
$ kustomize build deployments/kustomize | kubectl apply -f -
```
