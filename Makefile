.PHONY: task_1 task_2

task_1:
	go test -v ./task_1 -count=1

task_2:
	go test -v ./task_2 -run '^TestFindMinMaxCompact$$' -count=1
