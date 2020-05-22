package main

// Expected result: the panic() will trigger all the defers of the functions from the stack in order until
// it's handled by the recover() in the jobExecutor defer.
func main() {
	result, err := jobExecutor()
	if err != nil {
		println("Job executor failed: ", err.Error())
	}

	println("Program finished. Result: ", result)
}

func jobExecutor() (int, error) {
	defer func() {
		if err := recover(); err != nil {
			println("Panic handled in jobExecutor defer. Cause: ", err.(string))
		} else {
			println("Executing job executor defer")
		}
	}()

	executeJob()
	println("Job executor finished")

	return 5, nil
}

func executeJob() {
	defer func() {
		println("Executing job defer")
	}()

	println("Executing job...")
	executeJobAction()

	println("Job finished")
}

func executeJobAction() {
	defer func() {
		println("Executing job action defer")
	}()

	println("Executing job action...")
	executeAnErrorProneOperation()

	println("Job action finished")
}

func executeAnErrorProneOperation() {
	defer func() {
		println("Executing error prone operation defer")
	}()

	println("Executing error prone operation...")
	println("PANIC!")
	panic("Unknown error happened when executing the error prone operation")
}
