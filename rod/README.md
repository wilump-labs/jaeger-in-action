# Hot R.O.D. - Rides on Demand

> 운영 환경에서는 올인원 도커 이미지로 실행하고 만약 컨테이너가 죽게 되면 단일 장애 원천 (single source of failure)이 되고, 결국 운영 서비스에 큰 영향을 미침
>
> 운영 환경의 경우에 개별 컨포넌트로 배포하는 걸 권장

## Quick start
### Jaeger 실행
```shell
docker run \
  --rm \
  --name jaeger \
  -p6831:6831/udp \
  -p16686:16686 \
  jaegertracing/all-in-one:latest
```
- http://localhost:16686 접속

### ROD sample 실행
```shell
# git clone https://github.com/jaegertracing/jaeger
# cd jaeger/examples/hotrod
# go run ./main.go all
docker run \
  --rm \
  --link jaeger \
  --env OTEL_EXPORTER_JAEGER_ENDPOINT=http://jaeger:14268/api/traces \
  -p9090-9093:8080-8083 \
  jaegertracing/example-hotrod:latest \
  all
```

