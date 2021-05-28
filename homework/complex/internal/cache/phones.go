package lru

func (c *Cache) GetPhone(phoneIDs []int64) (phonesInCache map[int64]RealPhone, phonesNotFoundInCache []int64) {
	phonesInCache = make(map[int64]RealPhone, len(phoneIDs))
	phonesNotFoundInCache = make([]int64, 0, len(phoneIDs))

	if !config.phones.enabled {
		phonesNotFoundInCache = phoneIDs
		return
	}

	for _, phoneID := range phoneIDs {
		value, err := c.phones.Get(phoneID)
		if err == nil {
			phonesInCache[phoneID] = value.(RealPhone)
			continue
		}

		phonesNotFoundInCache = append(phonesNotFoundInCache, phoneID)
	}

	return
}

func (c *Cache) SetPhone(phoneID int64, realPhone RealPhone) error {
	if config.phones.enabled {
		return c.phones.SetWithExpire(phoneID, realPhone, config.phones.ttl)
	}

	return nil
}

func (c *Cache) StatPhones() float64 {
	return c.phones.HitRate()
}

func (c *Cache) LenPhones() int {
	return c.phones.Len(false)
}
