package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in
	for i := 0; i < len(stages); i++ {
		select {
		case <-done:
			return out
		default:
			out = proccessStage(done, stages[i](out))
		}
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
