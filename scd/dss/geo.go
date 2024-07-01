package dss

import (
	"errors"
	"time"
)

type Geo interface {
	Intersects(a []Volume4D, b []Volume4D) error
}

type GeomHelper struct {
}

func (g GeomHelper) Intersects(a []Volume4D, b []Volume4D) error {
	for _, aVolume := range a {
		for _, bVolume := range b {
			if conflict := volumeIntersect(aVolume, bVolume); conflict != nil {
				return conflict
			}
		}
	}
	return nil
}

func volumeIntersect(a Volume4D, b Volume4D) error {
	timeConflict, err := timeIntersects(a, b)
	if err != nil {
		return err
	}
	if !timeConflict {
		return nil
	}

	if !heightIntersects(a, b) {
		return nil
	}

	//TODO Convert OutlineCircle into OutlinePolygon
	//TODO Check conflict between Polygons
	return errors.New("conflict between volumes")

}

func timeIntersects(a Volume4D, b Volume4D) (bool, error) {
	aStart, err := time.Parse(time.RFC3339, a.TimeStart.Value)
	if err != nil {
		return false, err
	}
	aEnd, err := time.Parse(time.RFC3339, a.TimeEnd.Value)
	if err != nil {
		return false, err
	}
	bStart, err := time.Parse(time.RFC3339, b.TimeStart.Value)
	if err != nil {
		return false, err
	}
	bEnd, err := time.Parse(time.RFC3339, b.TimeEnd.Value)
	if err != nil {
		return false, err
	}

	if bEnd.After(aStart) && aEnd.After(bStart) {
		return true, nil
	}
	return false, nil
}

func heightIntersects(a Volume4D, b Volume4D) bool {
	return a.Volume.AltitudeUpper.Value >= b.Volume.AltitudeLower.Value &&
		b.Volume.AltitudeUpper.Value >= a.Volume.AltitudeLower.Value
}
