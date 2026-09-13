.PHONY: task_1 task_2 task_3

task_1:
	go test -v ./task_1 -count=1

task_2:
	go test -v ./task_2 -run '^TestFindMinMaxCompact$$' -count=1

task_3:
	go test -v ./task_3 -count=1
