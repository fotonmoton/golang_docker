package warehouse

type WarehouseState interface {
	Save(Product) int
	List() []Product
}

type Warehouse struct {
	state WarehouseState
}

func NewWarehouse(state WarehouseState) *Warehouse {
	return &Warehouse{state: state}
}

func (w *Warehouse) AddProduct(p Product) {
	w.state.Save(p)
}

func (w *Warehouse) ListProducts() []Product {
	return w.state.List()
}
