package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry013FixedDelaySource(t *testing.T){ got:=FixedDelay(3,nil,New(Delay(5*time.Millisecond),MaxDelay(9*time.Millisecond)).retrierCore); if got != 5*time.Millisecond{t.Fatalf("got %v want %v",got,5*time.Millisecond)} }

func TestTaskRetry013FixedDelaySourceAdjacent(t *testing.T){
 got:=FixedDelay(1,nil,New(Delay(2*time.Millisecond)).retrierCore)
 if got != 2*time.Millisecond{ t.Fatalf("adjacent got %v want %v",got,2*time.Millisecond) }
}
