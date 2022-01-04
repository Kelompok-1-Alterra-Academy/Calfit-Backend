package schedules_test

import (
	"CalFit/business/schedules"
	"CalFit/business/schedules/mocks"
	"time"
)

var repo mocks.Repository
var domain schedules.Domain
var usecase schedules.Usecase

func testSetup() {
	domain := schedules.Domain{
		Id:           1,
		TimeSchedule: "07.00-09.00",
		Duration:     120,
		SessionID:    1,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	usecase = schedules.NewSchedulesUsecase(&repo, time.Minute*1)
}
