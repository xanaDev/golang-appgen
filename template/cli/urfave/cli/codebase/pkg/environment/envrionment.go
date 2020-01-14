package environment
import(
	"os"
)
func Get()([]string){
	return os.Environ()
}
