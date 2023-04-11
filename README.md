# Jaeger in action
## What is Jaeger?
https://github.com/ruthetum/study/tree/main/jaeger

## Installation
Jaeger를 운영환경에 적용할 때에는 all-in-one 보다는 Collector + Query + Agent 방식으로 적용

### [Prerequisite] cert manager 설치 및 적용
cert-manager 미설치 시 operator 작동 시 에러 발생 (cert-manager 1.6.1 이상 필요)

install guide: https://cert-manager.io/v1.8-docs/installation/#default-static-install

#### install
```shell
curl -LJ kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.8.2/cert-manager.yaml >> cert-manager.yaml
```

#### apply
```shell
kubectl apply -f cert-manager.yaml
```

### [Prerequisite] namespace 생성
```shell
kubectl create namespace observability
```

### operator 설치 및 적용
#### install
```shell
curl -LJ https://github.com/jaegertracing/jaeger-operator/releases/download/v1.43.0/jaeger-operator.yaml >> operator.yaml
```

#### apply
CRD(Custom Resource Definition) 설치 (`apiVersion: jaegertracing.io/v1`)

```shell
kubectl create -f operator.yaml -n observability
```

### storage 세팅
[ElasticSearch 활용](./storage/elasticsearch) (or Cassandra)

### jaeger 적용
production strategy: https://www.jaegertracing.io/docs/1.43/operator/#production-strategy

```shell
kubectl apply -f jaeger.yaml -n observability
```

#### Deployment Strategy
- AllInOne(default)
  > This strategy is intended for development, testing, and demo purposes. 
  >
  > The main backend components, agent, collector and query service, are all packaged into a single executable which is configured (by default) to use in-memory storage. This strategy cannot be scaled beyond one replica.

  - 메모리 내 저장소를 사용하여 올인원 이미지(jaeger-agent , jaeger-collector , jaeger-query 및 Jaeger-UI 결합)를 단일 포드에 배포
  - 운영용이 아닌 개발/테스트/데모용 

- Production
  > The production strategy is intended (as the name suggests) for production environments, where long term storage of trace data is important, as well as a more scalable and highly available architecture is required. Each of the backend components is therefore separately deployed.

  - 운영용 배포 전략

- Streaming
  > The streaming strategy is designed to augment the production strategy by providing a streaming capability that effectively sits between the collector and the backend storage (Cassandra or Elasticsearch). This provides the benefit of reducing the pressure on the backend storage, under high load situations, and enables other trace post-processing capabilities to tap into the real time span data directly from the streaming platform (Kafka).
  
  - 

#### cluster role, cluster role binding 적용
sidecar 적용 시
```shell
kubectl create -f https://raw.githubusercontent.com/jaegertracing/jaeger-operator/master/deploy/cluster_role.yaml

kubectl create -f https://raw.githubusercontent.com/jaegertracing/jaeger-operator/master/deploy/cluster_role_binding.yaml
```

## Reference
- jaeger-getting-started: https://www.jaegertracing.io/docs/1.43/getting-started/
- what is jaeger: https://blog.advenoh.pe.kr/cloud/Jaeger%EC%97%90-%EB%8C%80%ED%95%9C-%EC%86%8C%EA%B0%9C/
- jaeger installation: https://blog.naver.com/alice_k106/221832024817
  - jaeger with istio: https://m.blog.naver.com/freepsw/221945686208
    - https://istio.io/latest/docs/tasks/observability/distributed-tracing/overview/
  - https://twofootdog.tistory.com/74
    - https://twofootdog.tistory.com/67
  - jaeger with golang: **https://litaro.tistory.com/entry/Jaeger-with-Go**
  - jaeger operator: **https://www.jaegertracing.io/docs/1.43/operator/**
- jaeger-k8s: https://blog.karsei.pe.kr/m/50
- jaeger-operator: https://blog.karsei.pe.kr/52
  - https://www.oss.kr/storage/app/public/oss/9f/ca/[Jaeger]%20Solution%20Guide.pdf