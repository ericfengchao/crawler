package dal

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			<-out
			itemCount++
			//log.Printf("Got #%d item %+v\n", item)
		}
	}()
	return out
}
