package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//функция находит угол между часовой и минутной стрелкой

// с примером входных данных 19:00, какой угол необходимо вычислить? от 00 минуты до 19 часов?
// или от 19 часов до 00 минуты?
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	s := scanner.Text()

	parts := strings.Split(s, ":")
	h, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	minAngle := 360 * m / 60
	hourAngle := float64(360*h/12) + 30.0/60.0*float64(m) //каждые 60 минут угол часов сдвигается на 30 градусов
	result := hourAngle - float64(minAngle)
	if result > 360 { //не может быть угла больше 360 градусов в часах
		result = result - 360
	}
	if result < 0 {
		result = math.Abs(result)
	}
	fmt.Println(result)

	fmt.Println(float64(30*h) - 5.5*float64(m)) // короткое упрощенное выражение
	//без учета -360
}
