// Runtime : 46 ms
// Memory  : 9.88 MB
package mar25

type Ingredient struct {
	name      string
	available bool
	requires  []*Ingredient
}

func NewIngredient(name string) *Ingredient {
	return &Ingredient{
		name:      name,
		available: false,
		requires:  []*Ingredient{},
	}
}

func (this *Ingredient) IsAvailable(inStack map[*Ingredient]bool) bool {
	if inStack[this] {
		return false
	}

	if this.available {
		return true
	} else if len(this.requires) == 0 {
		return false
	}

	inStack[this] = true

	for _, item := range this.requires {
		if !item.IsAvailable(inStack) {
			return false
		}
	}
	this.available = true
	inStack[this] = false
	return true
}

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	ingredientMap := make(map[string]*Ingredient)
	answer := []string{}

	for i := 0; i < len(recipes); i++ {
		rname := recipes[i]
		ritem, ok := ingredientMap[rname]

		if !ok {
			ritem = NewIngredient(rname)
			ingredientMap[rname] = ritem
		}

		for _, iname := range ingredients[i] {
			iitem, ok := ingredientMap[iname]
			if !ok {
				iitem = NewIngredient(iname)
				ingredientMap[iname] = iitem
			}
			ritem.requires = append(ritem.requires, iitem)
		}
	}

	for _, name := range supplies {
		item, ok := ingredientMap[name]
		if ok {
			item.available = true
		}
	}

	for _, name := range recipes {
		item := ingredientMap[name]
		inStack := make(map[*Ingredient]bool)
		if item.IsAvailable(inStack) {
			answer = append(answer, item.name)
		}
	}

	return answer
}
