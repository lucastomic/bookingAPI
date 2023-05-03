package viewport

type IViewManager[T any, I any] interface {
	ParseView(T) I
}
