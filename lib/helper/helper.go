package helper
import(
	"github.com/spf13/cast"
	"log"
)

func LogError(err error){
	if err != nil{
		log.Println(err)
	}
}

func ToString(value ...interface{}) string {
	return cast.ToString(value[0])
}

func ToStrings(value ...interface{}) interface{} {
	for i:=0;i<len(value);i++{
		value[i] = cast.ToString(value[i])
	}
	return value
}

func ToInt(value ...interface{}) int {
	return cast.ToInt(value[0])
}

func ToInts(value ...interface{}) interface{} {
	for i:=0;i<len(value);i++{
		value[i] = cast.ToInt(value[i])
	}
	return value
}

func ToInt64(value ...interface{}) int64 {
	return cast.ToInt64(value[0])
}

func ToInt64s(value ...interface{}) interface{} {
	for i:=0;i<len(value);i++{
		value[i] = cast.ToInt64(value[i])
	}
	return value
}

func ToUint(value ...interface{}) uint {
	return cast.ToUint(value[0])
}

func ToUints(value ...interface{}) interface{} {
	for i:=0;i<len(value);i++{
		value[i] = cast.ToUint(value[i])
	}
	return value
}

func ToBool(value ...interface{}) bool {
	return cast.ToBool(value[0])
}

func ToBools(value ...interface{}) interface{} {
	for i:=0;i<len(value);i++{
		value[i] = cast.ToBool(value[i])
	}
	return value
}