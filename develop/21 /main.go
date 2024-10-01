package main

import "fmt"

type JSON map[string]interface{}
type XML map[string]string

// Struct that keeps normalization parameters for every dimension
type NormalizationParameters struct {
	x_min, y_min, z_min float64
	x_max, y_max, z_max float64
}

// This function takes data in JSON format, normalizing it, and also returns it in JSON format
func (l *NormalizationParameters) NormalizeData(data JSON) JSON {
	fmt.Println("Normalizing data...")
	return data
}

// This struct represents a 3dmodel that can be drawn wireframe like
type Module struct {
	Vertices             []float64
	Polygons             []int
	CountVertices        int
	CountPolygons        int
	CountPolygonsIndices int
}

// This function getting data (vertices and polygons) from model.
func (m *Module) GetData() JSON {
	fmt.Println("Getting necessary data...")
	return JSON{"Vertices": m.Vertices, "Polygons": m.Polygons}
}

// This function loads data from specified path. path should have .obj format
func (m *Module) LoadData(path string) JSON {
	fmt.Println("Loading data...", path)
	return JSON{}
}

// Drawer representation. Every drawer should work with JSON-like data
type Drawer interface {
	Draw(data JSON)
}

// Concrete Drawer - QTDrawer.
type QTDrawer struct {
	color   [3]int
	opacity int8
}

// Some random functions
func (d *QTDrawer) SetColor() {
	fmt.Println("QT setting color")
}

// Some random functions
func (d *QTDrawer) SetOpacity() {
	fmt.Println("QT setting opacity")
}

// Draw function
func (d *QTDrawer) Draw(data JSON) {
	d.SetColor()
	d.SetOpacity()
	fmt.Println("drawing")
}

// Now we have a new module that works with XML data. That is bad, because every other block we created, works with JSON. We should do something about that. For example, adapter would the problem (it will convert XML to JSON for compatibility reasons)
type NewModule struct {
	Data XML
}

// Method for getting data
func (m *NewModule) GetData() XML {
	fmt.Println("Getting necessary data...")
	return m.Data
}

// Method for loading data
func (m *NewModule) LoadData(path string) XML {
	fmt.Println("Loading data...", path)
	return XML{}
}

// Adapter
type Adapter struct {
	Module NewModule
}

// Converting method
func (a *Adapter) GetData() JSON {
	a.Module.GetData()
	// Converting algorithm
	return JSON{}
}

func main() {

	// #1. OK

	// model := Module{}
	// data := model.GetData()
	// normalizer := NormalizationParameters{}
	// data = normalizer.NormalizeData(data)
	// var drawerModule Drawer = &QTDrawer{}
	// drawerModule.Draw(data)

	// #2. OK

	// modelAdap := Adapter{NewModule{}}
	// data := modelAdap.GetData()
	// normalizer := NormalizationParameters{}
	// data = normalizer.NormalizeData(data)
	// var drawerModule Drawer = &QTDrawer{}
	// drawerModule.Draw(data)

}
