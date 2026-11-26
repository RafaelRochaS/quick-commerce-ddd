package order

import "log"

type PlaceOrderHandler struct {
	repository   Repository
	eventService Messenger
}

func NewPlaceOrderHandler(repository Repository, eventService Messenger) *PlaceOrderHandler {
	return &PlaceOrderHandler{repository: repository, eventService: eventService}
}

func (h *PlaceOrderHandler) Handle(cmd *PlaceOrderCmd) error {
	log.Println("Executing place order command...")
	err := h.repository.PlaceOrder(cmd.Order)

	if err != nil {
		return err
	}

	// if this has an error, needs a complimentary transaction to remove the placed order on the repository. Better to use Transactional Outbox
	err = h.eventService.SendMessage("order created", NewCreatedEvent(*cmd.Order).String())

	return err
}
