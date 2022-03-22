package Test

import ("testing"
		giDevice "github.com/electricbubble/gidevice"
		"log"
		"fmt"
)

func Test(t *testing.T){
	usbmux, err := giDevice.NewUsbmux()
	if err != nil {
		log.Fatalln(err)
	}

	devices, err := usbmux.Devices()
	if err != nil {
		log.Fatal(err)
	}

	for _, dev := range devices {
		fmt.Println(dev.Properties().SerialNumber, dev.Properties().ProductID, dev.Properties().DeviceID)
	}
}


func TestSimulateLocation(t  *testing.T){
	usbmux, err := giDevice.NewUsbmux()
	if err != nil {
		log.Fatalln(err)
	}

	devices, err := usbmux.Devices()
	if err != nil {
		log.Fatalln(err)
	}

	if len(devices) == 0 {
		log.Fatalln("No Device")
	}

	d := devices[0]

	// https://api.map.baidu.com/lbsapi/getpoint/index.html
	if err = d.SimulateLocationUpdate(116.024067, 40.362639, giDevice.CoordinateSystemBD09); err != nil {
		log.Fatalln(err)
	}

	// https://developer.amap.com/tools/picker
	// https://lbs.qq.com/tool/getpoint/index.html
	// if err = d.SimulateLocationUpdate(120.116979, 30.252876, giDevice.CoordinateSystemGCJ02); err != nil {
	// 	log.Fatalln(err)
	// }

	// if err = d.SimulateLocationUpdate(121.499763, 31.239580,giDevice.CoordinateSystemWGS84); err != nil {
	// if err = d.SimulateLocationUpdate(121.499763, 31.239580); err != nil {
	// 	log.Fatalln(err)
	// }

	// err = d.SimulateLocationRecover()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}