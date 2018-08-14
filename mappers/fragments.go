package mappers

import (
	"sync"

	"github.com/minhajuddinkhan/gatewaygo/models"
	"github.com/minhajuddinkhan/gatewaygo/queue"
	"github.com/minhajuddinkhan/gatewaygo/targets"
)

//GetMappedFragments GetMappedFragments
func GetMappedFragments(subscription models.Subscriptions, data []byte) ([]queue.Fragment, error) {

	finished := make(chan bool, 1)
	errChannel := make(chan error)
	nsqFragmentCh := make(chan queue.Fragment)

	var wg sync.WaitGroup
	wg.Add(len(subscription.Endpoint.DependentURLs))

	for _, dependentURL := range subscription.Endpoint.DependentURLs {

		go func(dependentURL models.Endpoints, b []byte, destinationCode string) {
			defer wg.Done()
			target, _ := targets.GetTarget(
				subscription.Source.Name,
				dependentURL.Event.DataModel.Name,
				dependentURL.Event.Name,
				subscription.Source.AuthParams)

			res, err := target.ToFHIR(b, destinationCode)
			if err != nil {
				errChannel <- err
				return
			}
			nsqFragmentCh <- queue.Fragment{
				DataModel:  dependentURL.Event.DataModel.Name,
				EndpointID: dependentURL.ID,
				Data:       res,
				Endpoint:   dependentURL,
			}

		}(dependentURL, data, subscription.Source.DestinationCode)
	}
	go func() { wg.Wait(); close(finished) }()

	fragments := []queue.Fragment{}
	wait := true
	for wait {
		select {
		case f := <-nsqFragmentCh:
			fragments = append(fragments, f)
		case <-finished:
			wait = false
			return fragments, nil
		case err := <-errChannel:
			wait = true
			return nil, err
		}

	}
	return fragments, nil

}
