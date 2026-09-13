.PHONY: task_1 task_2 task_3 task_4 task_5 task_6 task_7

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
