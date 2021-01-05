package core_test

import (
	"go_rts/core"
	"go_rts/core/geometry"
	"go_rts/core/objects"
	"testing"
)

func Test_GivenIdenticalGameObjects_WhenCheckingForEquality_ThenReturnsTrue(t *testing.T) {
	tile1 := objects.NewTile("water", false, geometry.NewPoint(10.0, 10.0))
	tile2 := objects.NewTile("grass", true, geometry.NewPoint(10.0, 10.0))
	unit1 := objects.NewUnit(geometry.NewPoint(100.0, 200.0))
	unit2 := objects.NewUnit(geometry.NewPoint(100.0, 300.0))

	gameObjects1 := &core.GameObjects{
		Tiles:         []*objects.Tile{tile1, tile2},
		Units:         []*objects.Unit{unit1, unit2},
		SelectedUnits: []*objects.Unit{unit1},
	}

	gameObjects2 := &core.GameObjects{
		Tiles:         []*objects.Tile{tile1, tile2},
		Units:         []*objects.Unit{unit1, unit2},
		SelectedUnits: []*objects.Unit{unit1},
	}

	if !gameObjects1.Equals(gameObjects2) {
		t.Error("identical gameObjects should return true for equality")
	}
}

func Test_GivenIdenticalGameObjects_WithNonIdenticalPointers_WhenCheckingForEquality_ThenReturnsTrue(t *testing.T) {

	gameObjects1 := &core.GameObjects{
		Tiles: []*objects.Tile{
			objects.NewTile("water", false, geometry.NewPoint(10.0, 10.0)),
			objects.NewTile("grass", true, geometry.NewPoint(10.0, 10.0)),
		},
		Units: []*objects.Unit{
			objects.NewUnit(geometry.NewPoint(100.0, 200.0)),
			objects.NewUnit(geometry.NewPoint(100.0, 300.0)),
		},
		SelectedUnits: []*objects.Unit{
			objects.NewUnit(geometry.NewPoint(100.0, 200.0)),
			objects.NewUnit(geometry.NewPoint(100.0, 300.0)),
		},
	}

	gameObjects2 := &core.GameObjects{
		Tiles: []*objects.Tile{
			objects.NewTile("water", false, geometry.NewPoint(10.0, 10.0)),
			objects.NewTile("grass", true, geometry.NewPoint(10.0, 10.0)),
		},
		Units: []*objects.Unit{
			objects.NewUnit(geometry.NewPoint(100.0, 200.0)),
			objects.NewUnit(geometry.NewPoint(100.0, 300.0)),
		},
		SelectedUnits: []*objects.Unit{
			objects.NewUnit(geometry.NewPoint(100.0, 200.0)),
			objects.NewUnit(geometry.NewPoint(100.0, 300.0)),
		},
	}

	if !gameObjects1.Equals(gameObjects2) {
		t.Error("identical gameObjects should return true for equality")
	}
}

func Test_GivenNonIdenticalGameobjects_WhenCheckingForEquality_ThenReturnsFalse(t *testing.T) {
	tile1 := objects.NewTile("water", false, geometry.NewPoint(10.0, 10.0))
	tile2 := objects.NewTile("grass", true, geometry.NewPoint(10.0, 10.0))
	unit1 := objects.NewUnit(geometry.NewPoint(100.0, 200.0))
	unit2 := objects.NewUnit(geometry.NewPoint(100.0, 300.0))

	gameObjects1 := &core.GameObjects{
		Tiles:         []*objects.Tile{tile1, tile2},
		Units:         []*objects.Unit{unit1, unit2},
		SelectedUnits: []*objects.Unit{},
	}

	gameObjects2 := &core.GameObjects{
		Tiles:         []*objects.Tile{tile1},
		Units:         []*objects.Unit{unit1},
		SelectedUnits: []*objects.Unit{unit1},
	}

	if gameObjects1.Equals(gameObjects2) {
		t.Error("non-identical gameObjects should return false for equality")
	}
}

func Test_GivenGameObjects_ThenCanSerialize_ThenDeserialize(t *testing.T) {
	tile1 := objects.NewTile("water", false, geometry.NewPoint(10.0, 10.0))
	tile2 := objects.NewTile("grass", true, geometry.NewPoint(20.0, 20.0))
	tile3 := objects.NewTile("grass", true, geometry.NewPoint(30.0, 30.0))
	unit1 := objects.NewUnit(geometry.NewPoint(100.0, 200.0))
	unit2 := objects.NewUnit(geometry.NewPoint(100.0, 300.0))
	unit3 := objects.NewUnit(geometry.NewPoint(100.0, 400.0))

	gameObjects := &core.GameObjects{
		Tiles:         []*objects.Tile{tile1, tile2, tile3},
		Units:         []*objects.Unit{unit1, unit2, unit3},
		SelectedUnits: []*objects.Unit{unit1, unit2, unit3},
	}

	if !gameObjects.Deserialize(gameObjects.Serialize()).Equals(gameObjects) {
		t.Error("deserialized gameObjects should equal pre-serialized game objects")
	}

	gameObjects2 := &core.GameObjects{
		Tiles:         []*objects.Tile{tile1},
		Units:         []*objects.Unit{unit1},
		SelectedUnits: []*objects.Unit{unit1},
	}

	if !gameObjects2.Deserialize(gameObjects2.Serialize()).Equals(gameObjects2) {
		t.Error("deserialized gameObjects should equal pre-serialized game objects")
	}

	gameObjects3 := &core.GameObjects{
		Tiles:         []*objects.Tile{tile1, tile2},
		Units:         []*objects.Unit{unit1, unit2},
		SelectedUnits: []*objects.Unit{unit1, unit2},
	}

	if !gameObjects3.Deserialize(gameObjects3.Serialize()).Equals(gameObjects3) {
		t.Error("deserialized gameObjects should equal pre-serialized game objects")
	}
}
