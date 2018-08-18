# Meraki network health status & sales data demo

This demo will build a dashboard which will show real Meraki network health data and mock point-of-sale(POS) data.
  
[Docker compose](https://docs.docker.com/compose/) start and configure the following service following

- merakibeat : Beats plugin that polls Meraki health API and pushes data to elastic search
- elasticsearch: elasticsearch service. 
- kibana: Kibana service connected to elasticsearch
- posbeat : Mock service that mocks point-of-sale(POS) transaction API for simulating POS transaction data (Total amount and items) 


### Perquisites
Install [Docker](https://www.docker.com/) and [Docker compose](https://docs.docker.com/compose/).

## Instruction
Open a terminal window at the current folder.

## Config
#### 1. Update config in `config/merakibeat.yml`:
```
  period: 720s  
  merakihost: http://api.meraki.com
  merakikey: <MERAKI-API-KEY>
  merakinewtorkids: []
  merakiorgid: <OrganizationID>

```

You can specify the networks you want to monitor in the `merakinewtorkids `

#### 2. Update config file permission and ownership
```
sudo chmod go-w config/merakibeat.yml
sudo chmod go-w config/posbeat.yml

sudo chown root:root config/merakibeat.yml
sudo chown root:root config/config/posbeat.yml

```

### Start docker-compose
```
docker-compose up -d
```

You can use this command to see status or logs:
```
docker-compose ps
docker-compose logs -f
```

## Kibana settings
After around 1 min you should able to access Kibana dashboard at [http://localhost:5601/](http://localhost:5601/).
After around 15 mins elasticsearch should start getting data from Meraki Health API and Mock POS service.

#### 1. Create indexes
When you start getting data then add these two indexes to Kibana: `meraki*` and `posbeat*`.

#### 2. Add script field
In the `Management > Index Patterns > Meraki*` view, click `Scripted fields` and add a new Scripted field.
Name : `connectionSuccessPercentage`
Script:
```
if (doc['success'].value > 0) {
    (doc['success'].value * 100 / (doc['assoc'].value + doc['auth'].value + doc['dhcp'].value + doc['dns'].value + doc['success'].value))
}
```

#### 3. Import predefined visualizations(charts) and dashboard config

In the `Management > Saved Objects` view, click `Import` and import the `config/Kibana_meraki_dashboard.json`. And config fields to the right index.

#### 4. Dashboard
Click `Dashboard` at the navigation panel and select `Dashboard`. You can change the **Time Range** at the right top to show data from different time window.








 