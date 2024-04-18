package collections

type List[T any] interface {
	Stack[T]
	// Delete entfernt das Element bei "index"
	Delete(index int)
	// Get gibt das Element bei "index" zurück
	Get(index int) T
	// Insert fügt ein Element bei "index" ein
	Insert(index int, value T)
	// Len gibt die Anzahl der Elemente zurück
	Len() int
}
