package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry001DefaultAttempts(t *testing.T){ got:=New().attempts; if got != uint(10){t.Fatalf("got %v want %v",got,uint(10))} }

func TestTaskRetry001DefaultAttemptsAdjacent(t *testing.T){
 got:=New(Attempts(4)).attempts
 if got != uint(4){ t.Fatalf("adjacent got %v want %v",got,uint(4)) }
}
