# Hailo Server

Hailo server implements the Hailo [developer API](https://developer.hailoapp.com) for etas and neardrivers

## Getting started

1. Install Consul

	Consul is the default registry/discovery for go-micro apps. It's however pluggable.
	[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

2. Run Consul
	```
	$ consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul
	```

3. Get your an API key from Hailo developer site - [doc](https://developer.hailoapp.com/docs)

4. Download and start the service

	```shell
	go get github.com/micro/hailo-srv
	hailo-srv --api_token=YOUR_API_TOKEN
	```

	OR as a docker container

	```shell
	docker run microhq/hailo-srv --api_token=YOUR_API_TOKEN --registry_address=YOUR_REGISTRY_ADDRESS
	```

## The API
Hailo server implements the [Hailo Developer API](https://developer.hailoapp.com/docs) as RPC.

### API Status
```shell
$ micro query go.micro.srv.hailo Api.Status
{
	"status": "OK"
}
```

### Auth Test
```shell
$ micro query go.micro.srv.hailo Auth.Test
{
	"status": "OK"
}
```

### Eta
```shell
$ micro query go.micro.srv.hailo Drivers.Eta '{"latitude": 51.510761,  "longitude": -0.1174437 }'
{
	"etas": [
		{
			"count": 15,
			"eta": 3,
			"service_type": "regular"
		}
	]
}
```

### Near Drivers
```shell
$ micro query go.micro.srv.hailo Drivers.Near '{"latitude": 51.510761,  "longitude": -0.1174437 }'
{
	"drivers": [
		{
			"latitude": 51.5115966796875,
			"longitude": -0.1189914345741272,
			"service_type": "regular"
		},
		{
			"latitude": 51.511436462402344,
			"longitude": -0.11928340792655945,
			"service_type": "regular"
		},
		{
			"latitude": 51.51133728027344,
			"longitude": -0.11941815912723541,
			"service_type": "regular"
		},
		{
			"latitude": 51.512271881103516,
			"longitude": -0.11949911713600159,
			"service_type": "regular"
		},
		{
			"latitude": 51.5108528137207,
			"longitude": -0.1206439882516861,
			"service_type": "regular"
		},
		{
			"latitude": 51.510643005371094,
			"longitude": -0.12094280123710632,
			"service_type": "regular"
		},
		{
			"latitude": 51.512943267822266,
			"longitude": -0.1183713972568512,
			"service_type": "regular"
		},
		{
			"latitude": 51.51228332519531,
			"longitude": -0.12024059891700745,
			"service_type": "regular"
		},
		{
			"latitude": 51.513214111328125,
			"longitude": -0.11675159633159637,
			"service_type": "regular"
		},
		{
			"latitude": 51.51347351074219,
			"longitude": -0.1176547035574913,
			"service_type": "regular"
		},
		{
			"latitude": 51.512630462646484,
			"longitude": -0.12129580229520798,
			"service_type": "regular"
		},
		{
			"latitude": 51.51241683959961,
			"longitude": -0.12172360718250275,
			"service_type": "regular"
		},
		{
			"latitude": 51.51288986206055,
			"longitude": -0.12159369885921478,
			"service_type": "regular"
		},
		{
			"latitude": 51.507408142089844,
			"longitude": -0.11589550226926804,
			"service_type": "regular"
		},
		{
			"latitude": 51.51336669921875,
			"longitude": -0.11337850242853165,
			"service_type": "regular"
		}
	]
}
```

