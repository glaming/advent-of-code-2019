package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

type (
	xyz struct {
		x, y, z int
	}

	moon struct {
		position xyz
		velocity xyz
	}
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (x xyz) calcAbs() int {
	return abs(x.x) + abs(x.y) + abs(x.z)
}

func (m *moon) applyGravity(n moon) {
	xDiff := m.position.x - n.position.x
	if xDiff < 0 {
		m.velocity.x++
	}
	if xDiff > 0 {
		m.velocity.x--
	}

	yDiff := m.position.y - n.position.y
	if yDiff < 0 {
		m.velocity.y++
	}
	if yDiff > 0 {
		m.velocity.y--
	}

	zDiff := m.position.z - n.position.z
	if zDiff < 0 {
		m.velocity.z++
	}
	if zDiff > 0 {
		m.velocity.z--
	}
}

func (m moon) calculateEnergy() int {
	return m.velocity.calcAbs() * m.position.calcAbs()
}

func readMoons(filename string) ([]moon, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var ms []moon

	s := bufio.NewScanner(file)
	for s.Scan() {
		var m moon

		n, err := fmt.Sscanf(s.Text(), "<x=%d, y=%d, z=%d>", &m.position.x, &m.position.y, &m.position.z)
		if err != nil {
			return nil, err
		}
		if n != 3 {
			return nil, errors.New("parsed more than 3 values")
		}

		ms = append(ms, m)
	}

	return ms, s.Err()
}

func applyGravity(ms []moon) []moon {
	for i, m := range ms {
		for j, n := range ms {
			if j <= i {
				continue
			}
			ms[i].applyGravity(n)
			ms[j].applyGravity(m)
		}
	}
	return ms
}

func applyVelocity(ms []moon) []moon {
	for i, m := range ms {
		ms[i].position.x += m.velocity.x
		ms[i].position.y += m.velocity.y
		ms[i].position.z += m.velocity.z
	}
	return ms
}

func calculateEnergy(ms []moon) int {
	var energy int
	for _, m := range ms {
		energy += m.calculateEnergy()
	}
	return energy
}

func main() {
	ms, err := readMoons("day12/input.txt")
	if err != nil {
		log.Panic(err)
	}

	for i := 0; i < 1000; i++ {
		ms = applyGravity(ms)
		ms = applyVelocity(ms)
	}

	energy := calculateEnergy(ms)

	log.Printf("Total energy: %d", energy)
}
