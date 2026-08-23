package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry002DefaultDelay(t *testing.T){ got:=New().delay; if got != 100*time.Millisecond{t.Fatalf("got %v want %v",got,100*time.Millisecond)} }

func TestTaskRetry002DefaultDelayAdjacent(t *testing.T){
 got:=New(Delay(25*time.Millisecond)).delay
 if got != 25*time.Millisecond{ t.Fatalf("adjacent got %v want %v",got,25*time.Millisecond) }
}
