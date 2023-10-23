package jsons

import (
	"testing"
	"time"
)

type SmallJsonRaw struct {
	Method    string    `json:"method"`
	Timestamp int       `json:"timestamp"`
	Time      time.Time `json:"time"`
	Id        string    `json:"id"`
}

type SmallJson struct {
	//Method    string   `json:"method"`
	//Timestamp int      `json:"timestamp"`
	Time DateTime `json:"time"`
	//Id        string   `json:"id"`
}

var smallJson = `
{
    "time": "2023-08-07T00:22:33.000Z"
}
`

//var smallJson = `
//{
//    "method": "online",
//    "timestamp": 1604994124300,
//	"time": "2023-08-07T00:22:33.000Z",
//    "id": "0789daffded94a9daf8931496701e9e4",
//    "data": {
//        "deviceNo": "设备序列号",
//        "productVersion": "GD-2.2100",
//        "protocolVersion": "V1.00.000",
//        "deviceIP": "192.168.1.100",
//        "featVersion": {
//            "face": "xxxx",
//            "finger": "xxxx"
//        },
//        "algSDKVersion": {
//            "face": "xxxx",
//            "finger": "xxxx"
//        },
//        "deviceModel": "E86-1701-VF",
//        "deviceName": "E86",
//        "kernelVersion": "R110C-20230705",
//        "systemVersion": "107000203001"
//    }
//}
//`

func BenchmarkDateTime_ParseRaw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Parse[SmallJsonRaw](smallJson)
		if err != nil {
			panic(err)
		}

	}
}

func BenchmarkDateTime_Parse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Parse[SmallJson](smallJson)
		if err != nil {
			panic(err)
		}
	}
}
