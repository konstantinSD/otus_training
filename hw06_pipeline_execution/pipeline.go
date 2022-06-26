package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in
	for _, stage := range stages {
		if done != nil {
			out = doneCover(out, done)
		}
		out = stage(out)
	}
	return out
}

func doneCover(in In, done In) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for {
			select {
			case value, ok := <-in:
				if !ok {
					return
				}
				select {
				case out <- value:
				case <-done:
					return
				}
			case <-done:
				return
			}
		}
	}()
	return out
}
