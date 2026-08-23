package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry016NilOnRetry(t *testing.T){ got:=New(OnRetry(nil)).onRetry != nil; if got != true{t.Fatalf("got %v want %v",got,true)} }

func TestTaskRetry016NilOnRetryAdjacent(t *testing.T){
 got:=func() bool { r:=New(); r.onRetry(0,nil); return true }()
 if got != true{ t.Fatalf("adjacent got %v want %v",got,true) }
}
