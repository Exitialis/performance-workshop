package lru

func (c *Cache) GetCalltracking(phones []RealPhone) (phonesInCache map[RealPhone]VirtualPhone, phonesNotFoundInCache []RealPhone) {
	phonesInCache = make(map[RealPhone]VirtualPhone, len(phones))
	phonesNotFoundInCache = make([]RealPhone, 0, len(phones))

	if !config.calltracking.enabled {
		phonesNotFoundInCache = phones
		return
	}

	for _, realPhone := range phones {
		value, err := c.calltracking.Get(realPhone)
		if err == nil {
			phonesInCache[realPhone] = value.(VirtualPhone)
			continue
		}

		phonesNotFoundInCache = append(phonesNotFoundInCache, realPhone)
	}

	return
}

func (c *Cache) SetCalltracking(realPhone RealPhone, virtualPhone VirtualPhone) error {
	if config.calltracking.enabled {
		return c.calltracking.SetWithExpire(realPhone, virtualPhone, config.calltracking.ttl)
	}

	return nil
}

func (c *Cache) Stat() float64 {
	return c.calltracking.HitRate()
}

func (c *Cache) Len() int {
	return c.calltracking.Len(false)
}
