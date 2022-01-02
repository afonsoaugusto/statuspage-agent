# statuspage-agent

Agente para alimentar o status page da [statuspage.io](https://statuspage.io).
Iniciei este agente porém encontrei o projeto [gatus](https://github.com/TwiN/gatus) que fornece todas as funcionalidades necessárias para o agente (incluindo um sitema de regras).


## Estrutura do projeto

```txt
.
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE
├── local-api
│   ├── Dockerfile
│   ├── main.go
│   └── Makefile
├── main.go
├── Makefile
├── random-metric
│   ├── main.go
├── README.md
├── target.go
├── target_status.go
└── targets.yaml

2 directories, 18 files
```

## Curl example

```sh
export API_KEY=<>
export PAGE_CODE=rrb1yx04b76c

curl -i -X GET \
  -H "Authorization: OAuth ${API_KEY}" \
  https://api.statuspage.io/v1/pages/${PAGE_CODE}/incidents


curl https://api.statuspage.io/v1/pages/<>/components/<> \
  -H "Authorization: OAuth <>" \
  -X PATCH \
  -d "component[status]=degraded_performance" | jq .
```

## Confiuração do gatus para suportar a API do statuspage

```yaml
# Configuração do gatus
alerting:
  custom:
                              # Necessário inserir o page id da StatusPage
                              # Não consegui localizar uma variavel para passar o page id
                              # Diferente do componet id, que é o o Endpoint Name do gatus
    url: "https://api.statuspage.io/v1/pages/<PAGE_ID>/components/[ENDPOINT_NAME]/"
    headers:
      # Imporante passar o api token gerado no statuspage
      Authorization: "OAuth <API_TOKEN>"
      Content-Type: "application/x-www-form-urlencoded"
    method: "PATCH"
    body: component[status]=[ALERT_TRIGGERED_OR_RESOLVED]
    placeholders:
      ALERT_TRIGGERED_OR_RESOLVED:
        TRIGGERED: "degraded_performance"
        RESOLVED: "operational"    


# curl -v -X POST 'https://gatus.free.beeceptor.com/my/api/path' -H 'Content-Type: application/json' -d '{"data":"Hello Beeceptor"}'

endpoints:
  - name: xxxxxx (Componente ID)
    group: Local API
    url: "http://local-api:8081/health-check"
    client:
      timeout: 10s    
    interval: 10s
    conditions:
      - "[STATUS] == 200"
    # necessário a associação com o alerta customizado 
    alerts:
      - type: custom
        enabled: true
        failure-threshold: 2
        success-threshold: 2
        send-on-resolved: true
        description: "health check failed"
```
