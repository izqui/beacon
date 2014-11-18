# Beacon

Advertise iBeacons using Go.

Only tested in Raspberry Pi and Banana Pi running Raspbian. 

### Usage
```go
	
	data, err := beacon.NewBeaconData("D9B9EC1F-3925-43D0-80A9-1E39D4CEA95C", 0, 0) // uuid, major, minor

	if err != nil {

		log.Fatal(err)
	}

	b := beacon.NewBeacon(data) 

	log.Fatal(b.StartAdvertising())
```

### Dependencies

Paypal's [Gatt](https://github.com/paypal/gatt)
