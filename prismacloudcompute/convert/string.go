package convert

func SchemaToStringSlice(in []interface{}) []string {
	ans := make([]string, 0, len(in))
	for _, val := range in {
		ans = append(ans, val.(string))
	}
	return ans
}
