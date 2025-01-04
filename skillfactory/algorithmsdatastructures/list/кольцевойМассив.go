package main

import "fmt"

// IntRing - кольцевой массив
type IntRing struct {
	size  int
	data  []*int //указатель для удаления элемента
	Start int
	End   int
}

// NewIntRing - конструктор кольца
func NewIntRing(size int, start int) (*IntRing, error) {
	if start >= size {
		return nil, fmt.Errorf("Стартовая позиция не соответсвтвует размеру кольцевого массива")
	}
	return &IntRing{size, make([]*int, size, size), start, start}, nil
}

// Size - получение размера кольца
func (r IntRing) Size() int {
	return r.size
}

// пусто ли кольцо
func (r *IntRing) IsEmpty() bool {
	return r.Start == r.End
}

// достигнут ли конец
func (r IntRing) IsFull() bool {
	return (r.End < r.Start && r.Start-r.End == 1) || (r.Start == 0 && r.End == r.size-1)
}

// Read - чтение элемента из кольцевого массива
func (r *IntRing) Read() (int, error) {
	if !r.IsEmpty() {
		var el *int
		continuouslDataArray := r.getContinuousArray()
		// перебираем данные массива как непрерывную память, для согласования индексов
		// при переборе одновременно наращиваем поле индекса начального элемента кольцевого массива r.Start
		for i := 0; el == nil && i < len(continuouslDataArray); i, r.Start = i+1, r.Start+1 {
			el = r.data[r.Start]
		}
		// Если до этого массив был заполнен после успешного чтения сбрасываем этот флаг
		if r.IsFull() {
			r.isFull = false
		}
		return *el, nil
	}
	return 0, fmt.Errorf("Не удалось произвести чтение из кольцевого массива: массив пуст!")
}

// Write - добавление одного элемента в кольцевой массив
func (r *IntRing) Write(v int) error {
	if !r.IsFull() {
		r.data[r.End] = &v
		if r.End < r.Size()-1 {
			r.End++
		} else {
			// Если конец массива ушел за границу - наращиваем массив за счет начала
			r.End = 0
		}
		// Если конец массива стал указаывать на его начало
		// после операции записи - ставим флаг, что он полон
		if r.End == r.Start {
			r.isFull = true
		}
		return nil
	}
	return fmt.Errorf("Не удалось произвести запись в кольцевой массив: массив полон!")
}

// RemoveByIndex - удаление элемента кольцевого массива
func (r *IntRing) RemoveByIndex(index int) error {
	if index < 0 || index >= r.size {
		return fmt.Errorf("<%d> - Неверный индекс удаляемого элемента кольцевого массива", index)
	}
	if r.data[index] != nil {
		r.data[index] = nil
	}
	return nil
}

// Print - печать содержимого кольцевого массива
func (r IntRing) Print() {
	if r.IsEmpty() {
		fmt.Println("Кольцевой массив пуст!")
	} else {
		dataForPrint := r.getContinuousArray()
		for _, el := range dataForPrint {
			if el != nil {
				fmt.Printf("%d\t", *el)
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	ring, err := NewIntRing(5, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	ring.Print()
}
