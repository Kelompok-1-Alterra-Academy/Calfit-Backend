package schedules

type SchedulesUsecase struct {
	schedulesRepo Repository
}

func NewSchedulesUsecase(repo Repository) Usecase {
	return &SchedulesUsecase{
		schedulesRepo: repo,
	}
}

func (schedulesUseCase *SchedulesUsecase) Insert(schedules Domain) (Domain, error) {
	res, err := schedulesUseCase.schedulesRepo.Insert(schedules)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}

func (schedulesUseCase *SchedulesUsecase) Get(schedules Domain) ([]Domain, error) {
	res, err := schedulesUseCase.schedulesRepo.Get(schedules)
	if err != nil {
		return []Domain{}, err
	}
	return res, nil
}

func (schedulesUseCase *SchedulesUsecase) Update(schedules Domain) (Domain, error) {
	return Domain{}, nil
}

func (schedulesUseCase *SchedulesUsecase) Delete(schedules Domain) (Domain, error) {
	return Domain{}, nil
}
