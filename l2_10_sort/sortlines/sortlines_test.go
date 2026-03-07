package sortlines

import (
	"reflect"
	"testing"
)

func TestSortLines_String(t *testing.T) {

	lines := []string{
		"banana",
		"apple",
		"orange",
	}

	cfg := Config{}

	res := SortLines(lines, cfg)

	expected := []string{
		"apple",
		"banana",
		"orange",
	}

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected %v got %v", expected, res)
	}
}

func TestSortLines_Reverse(t *testing.T) {

	lines := []string{
		"a",
		"c",
		"b",
	}

	cfg := Config{
		Reverse: true,
	}

	res := SortLines(lines, cfg)

	expected := []string{
		"c",
		"b",
		"a",
	}

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected %v got %v", expected, res)
	}
}

func TestSortLines_Numeric(t *testing.T) {

	lines := []string{
		"10",
		"2",
		"1",
	}

	cfg := Config{
		Numeric: true,
	}

	res := SortLines(lines, cfg)

	expected := []string{
		"1",
		"2",
		"10",
	}

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected %v got %v", expected, res)
	}
}

func TestUnique(t *testing.T) {

	lines := []string{
		"a",
		"a",
		"b",
		"b",
		"c",
	}

	res := Unique(lines)

	expected := []string{
		"a",
		"b",
		"c",
	}

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected %v got %v", expected, res)
	}
}

func TestColumnSort(t *testing.T) {

	lines := []string{
		"a\t3",
		"b\t1",
		"c\t2",
	}

	cfg := Config{
		Column: 2,
		Numeric: true,
	}

	res := SortLines(lines, cfg)

	expected := []string{
		"b\t1",
		"c\t2",
		"a\t3",
	}

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected %v got %v", expected, res)
	}
}