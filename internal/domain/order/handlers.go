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

	err = h.eventService.SendMessage("order created", NewCreatedEvent(*cmd.Order).String())

	return err
}
