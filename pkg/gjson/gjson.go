// Package gjson
// @author moqi
// On 2022/10/18 18:26:30
// see: https://github.com/tidwall/gjson
package gjson

import (
	"github.com/tidwall/gjson"
)

func JsonStringToMap(dataString string, keyPath string, valuePath string) map[string]string {
	dataMap := make(map[string]string)

	resultArray := gjson.GetMany(dataString, keyPath, valuePath)

	keyArray := resultArray[0].Array()
	valueArray := resultArray[1].Array()

	for i, key := range keyArray {
		dataMap[key.Raw] = valueArray[i].Raw
	}

	return dataMap
}
