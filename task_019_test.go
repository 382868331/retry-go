package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry019TimerOption(t *testing.T){ got:=New(WithTimer(nil)).timer == nil; if got != true{t.Fatalf("got %v want %v",got,true)} }

func TestTaskRetry019TimerOptionAdjacent(t *testing.T){
 got:=New().timer != nil
 if got != true{ t.Fatalf("adjacent got %v want %v",got,true) }
}
