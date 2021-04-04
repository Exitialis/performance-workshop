package options

type Option struct {
	Value *int64
	IntValues []int64
	RangeValue *RangeValue
}

type RangeValue struct {
	From int64
	To int64
}

func (o *OptionsService) GetOptionsByUserType(typ string) {

}
