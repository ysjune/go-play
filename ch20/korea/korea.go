package korea

import "fmt"

type KoreaSender struct {
}

func (f *KoreaSender) Send(parcel string) {
	fmt.Printf("%v korea parcel\n", parcel)
}
