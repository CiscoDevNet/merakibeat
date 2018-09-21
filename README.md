# Merakibeat

This is elastic beats (https://www.elastic.co/products/beats) plugin for Meraki 
health API. 

Meraki exposes health API (https://dashboard.meraki.com/api_docs#wireless-health) to 
identify health of network, devices and client. Some of the key health parameters that 
Meraki health API exposes are 
- Connection Stats 
	- Success : Total number of successful connection
	- Assoc   : Number of connections in Association state
	- Auth	  : Number of connections in Authetication state
	- DHCP 	  : Number of connections in DHCP ip assignmnet stage
	- DNS 	  : Number of connections in DNS check stage   
- Latency Stats : Latency in miliseconds for packet transfer
    - latencyTime: Packet count
	- latencyTime is provided for 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024 & 2048 miliseconds.
	
This beats plugin polls the Meraki API for these health stats and allows sending Stats
to elasticsearch or any of ouput service supported by 
beats. (https://www.elastic.co/guide/en/logstash/current/output-plugins.html). 

This pipeline enables analysing mearki health data with other enterprise data like Point Of Sales, 
to identify relation between network status and revenue impact. 

![MerakiBeat pipeline](https://github.com/CiscoDevNet/merakibeat/blob/master/docs/media/merakibeat-pipeline.png)

 Note: POSBeat plugin mentioned in above diagram is to collect sales data, this is custom non-generic plugin can not be open sourced. POS simulator & posbeat docker images are listed in docker-compose for reference.  

Sample kibana dashboard included as part of docker compose is

![MerakiBeat sample dashboard](https://github.com/CiscoDevNet/merakibeat/blob/master/docs/media/merakibeat-dashboard.png)

# Configuring MerakiBeats plugin
Supports following plugin specific configs in merakibeat.yml file
-  period: Polling interval , recommended value 300s to 600s
-  merakihost: URL for meraki API endpoint in format, http://localhost:5050
-  merakikey: Meraki API key secret
-  merakinewtorkids: Netwroks IDs to be monitored by this plugin format, ["ABC", "XYZ"]
-  merakiorgid: ID of meraki oragnization

Following are fields to control sections that will be collected by data collector
'1' represent data will be collected and '0' means data will be skipped.
- nwconnstat: 1
- nwlatencystat: 1
- devconnstat: 0
- devlatencystat: 0
- clconnstat: 0
- cllatencystat: 0
	 
Following fields can be configured for enabling meraki API scan receiver
- scanEnable : 1    
- scanSecret : "simulator"
- scanValidator : "EnterYourValidator"

All these field are configured in merakibeat.yml config file

#### Meraki Scanning API configuration 
 - Meraki scanning receiver is exposed at http://\<baseurl\>:5001/api/v1/scanning/receiver/
	- \<baseurl\> could be http://publicip of host on which meraki beat is or user can configures dns name for assigned IP.
 - This endpoint should be registered in meraki dashboard details can found in this https://documentation.meraki.com/MR/Monitoring_and_Reporting/Scanning_API#Enable_Scanning_API document. 
 - Secret and validator from above registration process must be set in above merakibeat config file. 
	
## Running merakibeat
### As docker-composer (**Recommended**)
https://github.com/CiscoDevNet/merakibeat/blob/master/docker-compose/README.md

### As binary 
```
merakibeats -e -d *
```

## Building Source code
https://github.com/CiscoDevNet/merakibeat/blob/master/DEVELOPMENT.md

