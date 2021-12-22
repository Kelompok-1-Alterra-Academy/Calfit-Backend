package schedules

type SchedulesUsecase struct {
	SchedulesRepo Repository
}

func NewSchedulesUsecase(repo Repository) Usecase {
	return &SchedulesUsecase{
		SchedulesRepo: repo,
	}
}

func (schedulesUseCase *SchedulesUsecase) Insert(schedules Domain) (Domain, error) {
	return Domain{}, nil
}

func (schedulesUseCase *SchedulesUsecase) Get(schedules Domain) ([]Domain, error) {
	return []Domain{}, nil
}

func (schedulesUseCase *SchedulesUsecase) Delete(schedules Domain) (Domain, error) {
	return Domain{}, nil
}
