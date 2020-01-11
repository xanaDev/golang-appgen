package os
import (
	"runtime"
	"fmt"
	
)



//GetOsDetails s
func GetOsDetails(cpu,name bool)(string){
	var result string
	if cpu{
	result +=	fmt.Sprintf("No of cpu cores : %d" ,runtime.NumCPU())
	}
	if name{
	result +=	fmt.Sprintf("Os Name : %s" ,runtime.GOOS)
	}
	if !cpu && !name{
	result+=	fmt.Sprintf("Os Name : %s \nNo of cpu cores : %d " ,runtime.GOOS,runtime.NumCPU())
	}
return result
}