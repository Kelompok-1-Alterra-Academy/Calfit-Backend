package memberships

type MembershipsUsecase struct {
	membershipsRepo Repository
}

func NewMembershipsUsecase(repo Repository) Usecase {
	return &MembershipsUsecase{
		membershipsRepo: repo,
	}
}

func (membershipsUseCase *MembershipsUsecase) Insert(memberships Domain) (Domain, error) {
	res, err := membershipsUseCase.membershipsRepo.Insert(memberships)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}

func (membershipsUseCase *MembershipsUsecase) Get(memberships Domain) ([]Domain, error) {
	res, err := membershipsUseCase.membershipsRepo.Get(memberships)
	if err != nil {
		return []Domain{}, err
	}
	return res, nil
}

func (membershipsUseCase *MembershipsUsecase) Update(memberships Domain) (Domain, error) {
	res, err := membershipsUseCase.membershipsRepo.Update(memberships)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}

func (membershipsUseCase *MembershipsUsecase) Delete(memberships Domain) (Domain, error) {
	res, err := membershipsUseCase.membershipsRepo.Delete(memberships)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}
