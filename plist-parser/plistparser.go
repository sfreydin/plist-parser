package plistparser

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	// "reflect"
)

type Plist struct {
    Dict Dict  `xml:"dict"`
}

type Dict struct {
    Key []string `xml:"key"`
    String []string `xml:"string"`
}


func GetSystemVersion(filename string) string {
    result := map[string]interface{}{}
    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    dec := xml.NewDecoder(f)
    dec.Strict = false

    for {
            token, _ := dec.Token()
            if token == nil {
                break
            }
            switch start := token.(type) {
            case xml.StartElement:
                switch start.Name.Local {
                case "plist":
                    var pl Plist
                    err := dec.DecodeElement(&pl, &start)
                    if err != nil {
                        fmt.Println(err.Error())
                    }
                    result["plist"] = pl

                default:
                    fmt.Errorf("Unrecognized token")
                }
            }
        }
jsonDataByte := result["plist"]
b, err := json.Marshal(&jsonDataByte)
if err != nil {
    fmt.Println(err)
    return "error"
}    

var jsonData *Plist
json.Unmarshal([]byte(b), &jsonData)
var ret string
for i := 1; i < len(jsonData.Dict.Key); i++ {
    if jsonData.Dict.Key[i] == "ProductVersion"{
        ret = jsonData.Dict.String[i]
        break
    }
}
return ret

}
