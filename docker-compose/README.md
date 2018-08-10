#### Running docker compose
Docker compose start and configure following service following
- merakimock : Mock service that mocks Meraki health API for simulating connectionStats and LatencyStats
- merakibeat : Beats plugin that polls meraki health api and pushes data to elastic search
- elasticsearch : elasticsearch service.
- Kibana : Kibana service connected to Elastic search

```
cd docker-compose
docker-compose up -d
```
other commands
```
docker-compose ps
docker-compose logs -f
```
