// Runtime : 31 ms
// Memory  : 17.66 MB
package feb25

type ProductOfNumbers struct {
	agg     []int
	nearest []int
}

func Constructor2() ProductOfNumbers {
	return ProductOfNumbers{
		agg:     []int{1},
		nearest: []int{-1},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.agg = []int{1}
		this.nearest = []int{-1}
	} else {
		n := len(this.agg)
		prod, near := this.agg[n-1]*num, this.nearest[n-1]
		this.agg = append(this.agg, prod)
		this.nearest = append(this.nearest, near)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.agg)
	if n-k <= 0 {
		return 0
	}
	return this.agg[n-1] / this.agg[n-1-k]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
