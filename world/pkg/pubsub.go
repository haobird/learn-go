package pkg

type PubSub interface {
	Subscribe(event string, handler interface{})
}
