package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// out := make(chan interface{})
	// go func() {
	// 	defer close(out)
	// 	for {
	// 		select {
	// 		default:
	// 			out <- runStages(done, in, stages...)
	// 			// fmt.Println(i)
	// 		case <-done:
	// 			return
	// 		}
	// 	}
	// }()
	// Place your code here.
	// return runStages2(done, in, stages...)
	var out In = in
	for _, stage := range stages {
		out = proccessStage(done, stage(out))
	}
	return out
}

func runStages(done In, in In, stages ...Stage) Out {
	out := make(chan interface{})
	s := stages[0]
	go func() {
		defer close(out)
		ino := proccessStage(done, s(in))
		for {
			select {
			case v, ok := <-ino:
				if ok {
					out <- v
				} else {
					return
				}
			case <-done:
				return
			}
		}
	}()
	if len(stages) > 1 {
		return runStages(done, out, stages[1:]...)
	}
	return out
}

func runStages2(done In, in In, stages ...Stage) Out {
	var out In = in
	for _, stage := range stages {
		out = proccessStage(done, stage(out))
	}
	return out
}

func proccessStage(done In, in In) Out {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for i := range in {
			select {
			case out <- i:
			case <-done:
				return
			}
		}
	}()
	return out
}
