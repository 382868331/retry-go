package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry001DefaultAttempts(t *testing.T){ got:=New().attempts; if got != uint(10){t.Fatalf("got %v want %v",got,uint(10))} }
