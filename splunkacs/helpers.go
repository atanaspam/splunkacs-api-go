package splunkacs

func (o HecTokenSpec) Equal(o2 HecTokenSpec) bool {
	a := o.DefaultHost == o2.DefaultHost &&
		o.DefaultIndex == o2.DefaultIndex &&
		o.DefaultSource == o2.DefaultSource &&
		o.DefaultSourcetype == o2.DefaultSourcetype &&
		o.Disabled == o2.Disabled &&
		o.Name == o2.Name &&
		o.UseACK == o2.UseACK

	return a && stringSlicesEqual(o.AllowedIndexes, o2.AllowedIndexes)
}

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
