package Test

import ("testing"
		giDevice "github.com/electricbubble/gidevice"
		"log"
		"fmt"
		"os"
		"image"
		"image/jpeg"
		"image/png"
		"time"
)

/*
	设备信息
*/
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
		fmt.Println(dev.Properties().SerialNumber)
		fmt.Println(dev.Properties().ProductID)
		fmt.Println(dev.Properties().DeviceID)
	}
}

/*
	更改设备经纬度，用于测试定位功能
*/ 
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
	
	// 查询城市经纬度工具：https://api.map.baidu.com/lbsapi/getpoint/index.html
	if err = d.SimulateLocationUpdate(-118.272054,34.018309, giDevice.CoordinateSystemBD09); err != nil {
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
/*
	快速截图
*/
func TestScreen(t *testing.T){
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

    raw, err := d.Screenshot()
    if err != nil {
        log.Fatalln(err)
    }

    img, format, err := image.Decode(raw)
    if err != nil {
        log.Fatalln(err)
    }
    // userHomeDir, _ := os.UserHomeDir()
    file, err := os.Create("../../static/image/"+time.Now().Format("2006-01-02 03:04:05")+"." + format)
    if err != nil {
        log.Fatalln(err)
    }
    defer func() { _ = file.Close() }()
    switch format {
    case "png":
        err = png.Encode(file, img)
    case "jpeg":
        err = jpeg.Encode(file, img, nil)
    }
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println(file.Name())
}