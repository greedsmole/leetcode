
var pow []int64
var cache map[int64][]int
var mod int64 = (1 << 31) - 1

func hash(str string) (ret int64) {
	ret = 0
	newstr := []byte(str)
	sort.SliceStable(newstr, func(i, j int) bool {
		return newstr[i] < newstr[j]
	})
	for i := range newstr {
		ret = (ret + int64((int64(newstr[i]-'a'+1)*pow[i])%mod)) % mod
	}
	return
}

func groupAnagrams(strs []string) [][]string {
	pow = make([]int64, 100)
	k := int64(7)
	for i := 0; i < 100; i++ {
		k = k * k % mod
		pow[i] = k
	}
	cache = make(map[int64][]int)
	for i := range strs {
		cache[hash(strs[i])] = append(cache[hash(strs[i])], i)
	}

	ans := make([][]string, len(cache))
	j := 0
	for _, element := range cache {
		for i := range element {
			// append(ans[i], strs[element[i]])
			ans[j] = append(ans[j], strs[element[i]])
		}
		j++
	}
	return ans
}
