# Beacon

Advertise iBeacons using Go.

Only tested in Raspberry Pi and Banana Pi running Raspbian. 

### Usage
```go
	
	b, err := beacon.NewBeacon("D9B9EC1F-3925-43D0-80A9-1E39D4CEA95C", 0, 0) // uuid, major, minor
 
	if err != nil {

		log.Fatal(err)
	}

	log.Fatal(b.StartAdvertising())
```

### Dependencies

Paypal's [Gatt](https://github.com/paypal/gatt)
