package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry006UntilSucceeded(t *testing.T){ got:=New(UntilSucceeded()).attempts; if got != uint(0){t.Fatalf("got %v want %v",got,uint(0))} }
