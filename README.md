# AESO Go Client

## Summary
I found that I re-created this for a couple of random projects I was playing with more than once so I abstracted my logic to a library to share.

## Important Notes
I made this originally to help myself out but I intend on making it more solid over the coming weeks. Check back for a more stable release. I'm performing many changes including better error reporting and removing my old `log.Fatal` so stay tuned.

## Installation

`go get github.com/jimdhughes/go-aeso-client`

## MVP Usage
``` go
package main

import (
	"log"
	"time"

	"github.com/jimdhughes/go-aeso-client"
)

func main() {
	client, err := aeso.NewAesoApiService("Your-Private-API-Key-From-The-AESO")
	if err != nil {
		log.Fatal(err)
	}

	// get pool price for the last 24 hours
	poolPrice, err := client.GetPoolPrice(time.Now().Add(-24*time.Hour), time.Now())
	if err != nil {
		log.Printf("Error getting generation info: %s\n", err)
	}
	log.Println(poolPrice)
}
```

## API Coverage
### Swagger API
| Category | API | Completed? |
|:--|:--|:--|
| Pool Price Report | /v1.1/price/poolPrice | Yes |
| System Marginal Price Report | /v1.1/price/systemMarginalPrice | Yes |
| System Marginal Price Report | /v1.1/price/systemMarginalPrice/current | Yes |
| Pool Participant API | /v1/poolparticipantlist | Yes |
| Operating Reserve Offer Control Report | /v1/operatingReserveOfferControl | Yes |
| Metered Volume Report | /v1/meteredvolume/details | Yes |
|Energy Merit Order Report | /v1/meritOrder/energy | No |
| Actual Forecast Report | /v1/load/albertaInternalLoad | Yes |
| Current Supply Demand | /v1/csd/summary/current | Yes |
| Current Supply Demand | /v1/csd/generation/assets/current | Yes |
| Asset List API | /v1/assetlist | Yes |

### APIM APIs
*Not Started*