package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry016NilOnRetry(t *testing.T){ got:=New(OnRetry(nil)).onRetry != nil; if got != true{t.Fatalf("got %v want %v",got,true)} }
