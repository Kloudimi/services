---
title: etas
servicename: etas
labels: 
- Readme
- Logistics
---
ETAs as a service. Provides single-pickup and multi-dropoff routes. Takes into account time and traffic.

# ETAs Service

Add accurate estimated time of arrival to apps for delivery tracking, routing directions and much more.

Current limitations:
• Only supports "Driving" (not walking, cycling)
• Does not optimize route

## Usage

The ETA service depends on the Google Maps API. Ensure you set the "google.maps.apikey" config value to your API key.

```shell
micro config set google.maps.apikey YOUR_API_KEY
```

Once set, run the service using `micro run github.com/micro/services/etas`.

```shell
$ micro call etas ETAs.Calculate $(cat example-req.json)
{
	"points": {
		"brentwood-station": {
			"estimated_arrival_time": "2020-12-15T11:01:29.429947Z",
			"estimated_departure_time": "2020-12-15T11:01:29.429947Z"
		},
		"nandos": {
			"estimated_arrival_time": "2020-12-15T10:54:38.429947Z",
			"estimated_departure_time": "2020-12-15T10:54:38.429947Z"
		},
		"shenfield-station": {
			"estimated_arrival_time": "2020-12-15T10:48:34.429947Z",
			"estimated_departure_time": "2020-12-15T10:48:34.429947Z"
		}
	}
}
```

## cURL


### ETAs Calculate
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Schema related to #/components/requestBodies/ETAsCalculateRequest not found
```shell
> curl 'https://api.m3o.com/etas/ETAs/Calculate' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d Schema related to #/components/requestBodies/ETAsCalculateRequest not found;
# Response
{
  "points": [
    {
      "key": "string",
      "value": {
        "estimated_arrival_time": "string",
        "estimated_departure_time": "string"
      }
    }
  ]
}
```


