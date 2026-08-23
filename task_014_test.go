package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry014ZeroRandomDelay(t *testing.T){ got:=RandomDelay(1,nil,New(MaxJitter(0)).retrierCore); if got != time.Duration(0){t.Fatalf("got %v want %v",got,time.Duration(0))} }

func TestTaskRetry014ZeroRandomDelayAdjacent(t *testing.T){
 got:=RandomDelay(9,nil,New(MaxJitter(0)).retrierCore)
 if got != time.Duration(0){ t.Fatalf("adjacent got %v want %v",got,time.Duration(0)) }
}
