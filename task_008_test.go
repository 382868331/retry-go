package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry008DelayOption(t *testing.T){ got:=New(Delay(8*time.Millisecond)).delay; if got != 8*time.Millisecond{t.Fatalf("got %v want %v",got,8*time.Millisecond)} }

func TestTaskRetry008DelayOptionAdjacent(t *testing.T){
 got:=New(Delay(0)).delay
 if got != time.Duration(0){ t.Fatalf("adjacent got %v want %v",got,time.Duration(0)) }
}
