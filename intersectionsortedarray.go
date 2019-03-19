package main

import "fmt"
/**
 * @input A : Integer array
 * @input B : Integer array
 * 
 * @Output Integer array.
 */
func intersect(A []int, B []int) []int {
    temp := []int{}
    m := make(map[int]bool, len(A))
    for _, j := range A {
        m[j] = true
    }
    for _, j := range B {
        if _, ok := m[j]; ok {
            if !contains(temp,j){
            temp = append(temp, j)
            }

        }
    }
    return temp

}
func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
