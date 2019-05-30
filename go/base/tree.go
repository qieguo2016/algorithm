package base

func BuildDictTree(arr []string) map[string]interface{} {
	tree := map[string]interface{}{}
	for _, chars := range arr {
		curr := tree
		for _, c := range chars {
			cs := string(c)
			if v, exist := curr[cs]; exist {
				curr = v.(map[string]interface{})
			} else {
				curr[cs] = map[string]interface{}{}
				curr = curr[cs].(map[string]interface{})
			}
		}
		curr["value"] = 1
	}
	return tree
}
