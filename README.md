# Jaeger in action
## What is Jaeger?
https://github.com/ruthetum/study/tree/main/jaeger

## Installation
Jaeger를 운영환경에 적용할 때에는 all-in-one 보다는 Collector + Query + Agent 방식으로 적용

#### cert manager 설치
cert-manager 미설치 시 operator 작동 시 에러 발생
```shell
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.8.0/cert-manager.yaml
```

```shell
curl -LJ https://github.com/cert-manager/cert-manager/releases/download/v1.8.0/cert-manager.yaml >> cert-manager.yaml
```

#### namespace 생성
```shell
kubectl create namespace observability
```

#### operator 설치
```shell
kubectl create -f https://github.com/jaegertracing/jaeger-operator/releases/download/v1.43.0/jaeger-operator.yaml -n observability
```

```shell
curl -LJ https://github.com/jaegertracing/jaeger-operator/releases/download/v1.43.0/jaeger-operator.yaml >> operator.yaml
```

#### operator 적용
```shell
kubectl apply -f operator.yaml -n observability
```

#### jaeger 적용
```shell
kubectl apply -f jaeger.yaml -n observability
```

#### cluster role, cluster role binding 적용
sidecar 적용 시
```shell
kubectl create -f https://raw.githubusercontent.com/jaegertracing/jaeger-operator/master/deploy/cluster_role.yaml

kubectl create -f https://raw.githubusercontent.com/jaegertracing/jaeger-operator/master/deploy/cluster_role_binding.yaml
```