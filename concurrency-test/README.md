


## Initialize project 
go mod init workerpool_demo

## Run main.go

go run main.go

## Test

go test .

## Change these constants in worker_test.go to make the test succeed or fail.

	maxTasks := 5
	numTasks := 8
	maxDuration := 250 * time.Millisecond
	
## Additionally, try to change the syntax of the code to see how test behaves.

Substitute

	fmt.Println("All work completed.")

For

	fmt.Println("\nAll work completed.")
