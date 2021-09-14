package dfs

import (
	"reflect"
)

//func CalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
//	result := make([]float64,0)
//	for i:= range queries {
//		v := make(map[int]bool, len(equations))
//		result = append(result, calcDfs(equations, values, queries[i][0], queries[i][1], v))
//	}
//	for i:= range result {
//		if result[i] == 0 {
//			result[i] = -1.0
//		}
//	}
//	return result
//}

func calcDfs(equations [][]string, values []float64, start string, end string, v map[int]bool) float64 {
	e1 := make([]string, 0)
	e1 = append(e1, start, end)
	e2 := make([]string, 0)
	e2 = append(e2, end, start)
	for i := range equations {
		if v[i] {
			continue
		}
		if reflect.DeepEqual(e1, equations[i]) {
			return values[i]
		}
		if reflect.DeepEqual(e2, equations[i]) {
			return 1 / values[i]
		}
		if start == equations[i][0] {
			if end == equations[i][0] {
				return 1.0
			}
			v[i] = true
			return values[i] * calcDfs(equations, values, equations[i][1], end, v)
		}
		if start == equations[i][1] {
			if end == equations[i][1] {
				return 1.0
			}
			v[i] = true
			return 1 / values[i] * calcDfs(equations, values, equations[i][0], end, v)
		}
	}
	return 0
}

type key struct {
	a, b string
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	hash := map[key]float64{}
	for i, equation := range equations {
		a, b := equation[0], equation[1]
		hash[key{a, b}] = values[i]
		hash[key{b, a}] = 1 / values[i]
		hash[key{a, a}] = 1.0
		hash[key{b, b}] = 1.0
	}
	for k, v := range hash {
		graphDfs(k.a, k.b, v, hash)
	}
	var result []float64
	for _, query := range queries {
		v, ok := hash[key{query[0], query[1]}]
		if ok {
			result = append(result, v)
		} else {
			result = append(result, -1.0)
		}
	}
	return result
}

func graphDfs(a, b string, v float64, hash map[key]float64) {
	for k, vk := range hash {
		// a/b * b(k.a)/c(k.b)
		if b == k.a && a != k.b {
			newKey := key{a, k.b}
			_, ok := hash[newKey]
			if !ok {
				hash[newKey] = v * vk
				graphDfs(a, k.b, v*vk, hash)
			}
		}
	}
}
