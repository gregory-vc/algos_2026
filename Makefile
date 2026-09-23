.PHONY: test task_1 task_2 task_3 task_4 task_5 task_6 task_7 task_8 task_9 task_10

task_1:
	go test -v ./task_1 -count=1

task_2:
	go test -v ./task_2 -run '^TestFindMinMaxCompact$$' -count=1

task_3:
	go test -v ./task_3 -count=1

task_4:
	go test -v ./task_4 -count=1

task_5:
	go test -v ./task_5 -count=1

task_6:
	go test -v ./task_6 -count=1

task_7:
	go test -v ./task_7 -count=1

task_8:
	go test -v ./task_8 -count=1

task_9:
	go test -v ./task_9 -count=1

task_10:
	go test -v ./task_10/... -count=1

test:
	go test -v ./... -count=1
