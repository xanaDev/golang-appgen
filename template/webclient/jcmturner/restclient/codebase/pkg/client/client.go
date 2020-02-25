package client

import (
	"{{ .AppName }}/pkg/types"
	"github.com/jcmturner/restclient"
	"os"
	"fmt"
)
// Basic create basic rest client
func Basic() *restclient.Config {

	c := restclient.NewConfig()
	c.WithEndPoint(types.ServiceEndPoint)

	if err := c.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Configuration of web service not valid: %v", err)
		os.Exit(1)
	}
	return c 
}
