package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry010MaxJitterTarget(t *testing.T){ got:=New(MaxJitter(6*time.Millisecond)).maxJitter; if got != 6*time.Millisecond{t.Fatalf("got %v want %v",got,6*time.Millisecond)} }

func TestTaskRetry010MaxJitterTargetAdjacent(t *testing.T){
 got:=New(Delay(time.Millisecond),MaxJitter(4*time.Millisecond)).delay
 if got != time.Millisecond{ t.Fatalf("adjacent got %v want %v",got,time.Millisecond) }
}
