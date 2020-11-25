// Package builder example pattern.
package builder

// BuildProcess ...
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// VehicleProduct ...
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// ManufacturingDirector ...
type ManufacturingDirector struct {
	builder BuildProcess
}

// SetBuilder ...
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

// Construct builds to us in the desired sequence.
func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

// BikeBuilder example.
type BikeBuilder struct {
	v VehicleProduct
}

// SetWheels ...
func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2

	return b
}

// SetSeats ...
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2

	return b
}

// SetStructure ...
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"

	return b
}

// GetVehicle ...
func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

// CarBuilder example.
type CarBuilder struct {
	v VehicleProduct
}

// SetWheels ...
func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4

	return c
}

// SetSeats ...
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5

	return c
}

// SetStructure ...
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"

	return c
}

// GetVehicle ...
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}
