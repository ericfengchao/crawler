package dal

type Saver interface {
	ItemSaver() chan interface{}
}
