# Merakibeat

This is elastic beats (https://www.elastic.co/products/beats) plugin for Meraki 
health API. 
Meraki exposes health API to identify health of network, devices and client. Some 
of the key health parameters that Meraki health API exposes are 
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


# Configuring MerakiBeats plugin
Supports following plugin specific configs
-  period: Polling interval , recommended value 300s to 600s
-  merakihost: URL for meraki API endpoint in format, http://localhost:5050
-  merakikey: Meraki API key secret
-  merakinewtorkids: Netwroks IDs to be monitored by this plugin format, ["ABC", "XYZ"]
-  merakiorgid: ID of meraki oragnization
	 
All these field are configured in merakibeat.yml config file

### Running merakibeats as binary 
```
merakibeats -e -d *
```

### Running merakibeats as docker image


