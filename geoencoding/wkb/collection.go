package wkb

import (
	"io"

	"github.com/spatial-go/geoos/space"
)

func readCollection(r io.Reader, order byteOrder, buf []byte) (space.Collection, error) {
	num, err := readUint32(r, order, buf[:4])
	if err != nil {
		return nil, err
	}

	alloc := num
	if alloc > maxMultiAlloc {
		// invalid data can come in here and allocate tons of memory.
		alloc = maxMultiAlloc
	}
	result := make(space.Collection, 0, alloc)

	d := NewDecoder(r)
	for i := 0; i < int(num); i++ {
		geom, err := d.Decode()
		if err != nil {
			return nil, err
		}
		result = append(result, geom)
	}
	return result, nil
}

func (e *Writer) writeCollection(c space.Collection) error {
	e.order.PutUint32(e.buf, geometryCollectionType)
	e.order.PutUint32(e.buf[4:], uint32(len(c)))
	_, err := e.w.Write(e.buf[:8])
	if err != nil {
		return err
	}

	for _, geom := range c {
		err := e.Encode(geom)
		if err != nil {
			return err
		}
	}
	return nil
}
