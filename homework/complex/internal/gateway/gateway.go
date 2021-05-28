package phones_gateway

import (
	"context"
	"github.com/exitialis/workshop/homework/complex/internal/cache"
	"github.com/exitialis/workshop/homework/complex/internal/service/calltracking"
	"github.com/exitialis/workshop/homework/complex/internal/service/phones"
)

type PhoneService struct {
	phonesService       PhonesService
	calltrackingService CalltrackingService
	lru                 *lru.Cache
}

func NewPhonesGateway(
	phoneService PhonesService,
	calltrackingService CalltrackingService,
	lru *lru.Cache,
	ctx context.Context,
) *PhoneService {
	service := &PhoneService{
		phonesService:       phoneService,
		calltrackingService: calltrackingService,
		lru:                 lru,
	}
	go service.fillCache(ctx)

	return service
}

func (p *PhoneService) GetPhone(ctx context.Context, in GetPhoneIn) (*GetPhoneOut, error) {
	var realPhone string
	realPhonesInCache, _ := p.lru.GetPhone([]int64{in.PhoneID})
	if realPhoneInCache, ok := realPhonesInCache[in.PhoneID]; ok {
		realPhone = string(realPhoneInCache)
	} else {
		var err error
		realPhone, err = p.phonesService.GetPhone(phones.GetPhoneIn{PhoneID: in.PhoneID})
		if err != nil || realPhone == "" {
			return nil, err
		}
		err = p.lru.SetPhone(in.PhoneID, lru.RealPhone(realPhone))
	}

	if p.calltrackingService.Check(in.UserID, realPhone) {
		phonesInCache, _ := p.lru.GetCalltracking([]lru.RealPhone{lru.RealPhone(realPhone)})
		if phone, ok := phonesInCache[lru.RealPhone(realPhone)]; ok {
			return &GetPhoneOut{
				Phone: string(phone),
				Type:  TypeCalltracking,
			}, nil
		}

		phone, err := p.calltrackingService.GetPhone(
			ctx,
			calltracking.GetPhoneIn{
				RealPhone: realPhone,
				UserID:    in.UserID,
				ItemID:    in.ItemID,
			},
		)

		if phone != "" && err == nil {
			_ = p.lru.SetCalltracking(lru.RealPhone(realPhone), lru.VirtualPhone(phone))

			return &GetPhoneOut{
				Phone: phone,
				Type:  TypeCalltracking,
			}, nil
		}
	}

	return &GetPhoneOut{
		Phone: realPhone,
		Type:  TypeBase,
	}, nil
}


func (p *PhoneService) Validate(in GetPhoneIn) error {
	if in.PhoneID <= 0 {
		return NewValidationError("Передан некорректный phoneId")
	}

	if in.UserID <= 0 {
		return NewValidationError("Передан некорректный userId")
	}

	if in.CategoryID <= 0 {
		return NewValidationError("Передан некорректный categoryId")
	}

	if in.ItemID <= 0 {
		return NewValidationError("Передан некорректный itemId")
	}

	if len(in.PhoneDisplayLoc) <= 4 {
		return NewValidationError("Передан некорректный phoneDisplayLoc")
	}

	return nil
}
