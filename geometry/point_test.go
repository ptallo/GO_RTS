package geometry_test

import (
	"go_rts/geometry"
	"testing"
)

func Test_GivenZero_WhenConvertingCartoToIso_ThenReturnZero(t *testing.T) {
	point := geometry.NewPoint(0.0, 0.0)
	actual := geometry.CartoToIso(point)
	expected := geometry.NewPoint(0.0, 0.0)

	if !actual.Equals(expected) {
		t.Errorf("expected: %v doesn't equal the actual: %v", expected, actual)
	}
}

func Test_GivenPoint_WhenConvertingCartoToIso_ThenReturnsConvertedPoint(t *testing.T) {
	point := geometry.NewPoint(10.0, 12.0)
	actual := geometry.CartoToIso(point)
	expected := geometry.NewPoint(-2.0, 11.0)

	if !actual.Equals(expected) {
		t.Errorf("expected: %v doesn't equal the actual: %v", expected, actual)
	}
}

func Test_GivenZero_WhenConvertingIsoToCarto_ThenReturnsZero(t *testing.T) {
	point := geometry.NewPoint(0.0, 0.0)
	actual := geometry.IsoToCarto(point)
	expected := geometry.NewPoint(0.0, 0.0)

	if !actual.Equals(expected) {
		t.Errorf("expected: %v doesn't equal the actual: %v", expected, actual)
	}
}

func Test_GivenPoint_WhenConvertingIsoToCarto_ThenReturnsConvertedPoint(t *testing.T) {
	point := geometry.NewPoint(-2.0, 11.0)
	actual := geometry.IsoToCarto(point)
	expected := geometry.NewPoint(10.0, 12.0)

	if !actual.Equals(expected) {
		t.Errorf("expected: %v doesn't equal the actual: %v", expected, actual)
	}
}
