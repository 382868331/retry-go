package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry011NilDelayType(t *testing.T){ got:=New(DelayType(nil)).delayType != nil; if got != true{t.Fatalf("got %v want %v",got,true)} }

func TestTaskRetry011NilDelayTypeAdjacent(t *testing.T){
 got:=New(DelayType(FixedDelay)).delayType != nil
 if got != true{ t.Fatalf("adjacent got %v want %v",got,true) }
}
