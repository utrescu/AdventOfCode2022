package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Data struct {
	number int
	list   []*Data

	isNumber bool
	parent   *Data
}

func (d Data) ToString() string {
	result := ""

	if d.isNumber {
		return fmt.Sprintf(" %d ", d.number)
	}
	result = "["
	for _, v := range d.list {
		result += v.ToString()
	}
	result = result + "]"

	return result
}

func (d Data) ToStringSortable() string {
	value := d.ToString()
	// El 10 no s'ordena bé com a string, el canvio perquè s'ordeni bé
	value = strings.Replace(value, "10", string('9'+1), -1)
	// Espais fora
	value = strings.Replace(value, " ", "", -1)
	// Cadena buida ha d'anar al davant. L'espai està davant
	value = strings.Replace(value, "[]", " ", -1)
	// Els símbols no serveixen per res
	value = strings.Replace(value, "[", "", -1)
	value = strings.Replace(value, "]", "", -1)
	return value
}

func (d *Data) AddNumber(value int) {
	d.list = append(d.list, &Data{number: value, isNumber: true})
}

func (d *Data) AddList() *Data {
	newData := Data{list: []*Data{}, isNumber: false, parent: d}
	d.list = append(d.list, &newData)
	return &newData
}

func Load(signal []string) Data {
	data := Data{parent: nil}
	current := &data
	for _, v := range signal[1:] {
		switch v {
		case "[":
			current = current.AddList()
		case "]":
			if current.parent != nil {
				current = current.parent
			}

		default:
			number, _ := strconv.Atoi(v)
			current.AddNumber(number)

		}
	}
	return data
}

const (
	INCORRECT      = 0
	CORRECT        = 1
	INDETERMINATED = 2
)

type Signal struct {
	datas [2]Data
}

func (s Signal) ToString() string {
	return s.datas[0].ToString() + "\n" + s.datas[1].ToString()
}

func (s Signal) IsValid() bool {
	return Compare(s.datas[0], s.datas[1]) == CORRECT
}

func Compare(left Data, right Data) int {
	if left.isNumber && right.isNumber {
		leftint := left.number
		rightint := right.number
		if leftint < rightint {
			return CORRECT
		} else if leftint > rightint {
			return INCORRECT
		}
	} else {

		newLeft := left.list
		newRight := right.list

		if left.isNumber {
			newLeft = []*Data{{number: left.number, isNumber: true}}
		}
		if right.isNumber {
			newRight = append(newRight, &Data{number: right.number, isNumber: true})
		}

		// si un no té elements i l'altre si, caput
		for i := 0; i < len(newLeft); i++ {

			if i > len(newRight)-1 {
				return INCORRECT
			}
			result := Compare(*newLeft[i], *newRight[i])
			if result != INDETERMINATED {
				return result
			}

		}
		if len(newLeft) < len(newRight) {
			return CORRECT
		}
	}
	return INDETERMINATED
}

func readLines(path string) ([]Signal, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	signals := []Signal{}
	scanner := bufio.NewScanner(file)
	current := 0
	signal := Signal{}
	for scanner.Scan() {

		line := scanner.Text()
		if len(line) == 0 {
			signals = append(signals, signal)
			signal = Signal{}
			current = 0
		} else {
			tmp := strings.Split(line, ",")
			values := []string{}
			for _, v := range tmp {
				part := v
				if strings.Contains(part, "[") || strings.Contains(v, "]") {
					toAdd := []string{}
					for strings.HasPrefix(part, "[") {
						toAdd = append(toAdd, string(part[0]))
						part = part[1:]
					}

					number := ""
					for i := range part {
						if part[i] != ']' {
							number = fmt.Sprintf("%s%s", number, string(part[i]))
						} else {
							if number != "" {
								toAdd = append(toAdd, number)
								number = ""
							}
							toAdd = append(toAdd, "]")
						}
					}
					if number != "" {
						toAdd = append(toAdd, number)
						number = ""
					}
					values = append(values, toAdd...)
				} else {
					values = append(values, v)
				}
			}

			data := Load(values)
			signal.datas[current] = data
			current = (current + 1) % 2
		}
	}
	signals = append(signals, signal)

	return signals, scanner.Err()
}

func Part1(signals []Signal) int {
	correct := 0
	for i, signal := range signals {

		if signal.IsValid() {
			correct += i + 1
		}
	}
	return correct
}

func Part2(signals []string, locate []string) int {
	sort.Strings(signals)
	result := 1
	numbersFound := 0
	for i, v := range signals {
		if v == locate[numbersFound] {
			result *= i + 1
			numbersFound++
			if numbersFound >= len(locate) {
				break
			}
		}
	}
	return result
}

func main() {
	signals, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	fmt.Println("Part 1:", Part1(signals))

	signalsToAdd := []string{
		"2", // "[[ 2 ]]",
		"6", // "[[ 6 ]]",
	}

	allSignals := []string{}

	for _, signal := range signals {
		allSignals = append(allSignals, signal.datas[0].ToStringSortable())
		allSignals = append(allSignals, signal.datas[1].ToStringSortable())
	}

	allSignals = append(allSignals, signalsToAdd...)

	decoderkey := Part2(allSignals, signalsToAdd)
	fmt.Println("Part 2:", decoderkey)
}
