package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry005AttemptsOption(t *testing.T){ got:=New(Attempts(3)).attempts; if got != uint(3){t.Fatalf("got %v want %v",got,uint(3))} }
