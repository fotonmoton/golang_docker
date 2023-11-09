package db

import "docker/warehouse"

type WarehouseInMemory struct {
	products []warehouse.Product
}

func NewWarehouseInMemoryState() *WarehouseInMemory {
	return &WarehouseInMemory{}
}

func (s *WarehouseInMemory) Save(p warehouse.Product) int {
	s.products = append(s.products, p)

	return 0
}

func (s *WarehouseInMemory) List() []warehouse.Product {
	return s.products
}
