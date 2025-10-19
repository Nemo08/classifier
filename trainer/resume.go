package trainer

import "github.com/Nemo08/classifier/net/feedforward"

func Resume(net feedforward.FeedforwardNetwork, resume *bool, dstmodel *string) {
	if resume != nil && *resume && dstmodel != nil {
		err := net.ReadZlibWeightsFromFile(*dstmodel)
		if err != nil {
			println(err.Error())
		}
	}
}
