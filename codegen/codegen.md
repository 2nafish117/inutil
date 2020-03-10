1. copy:
    a = b
    PushName "b"
    PopName "a"

2. add/sub/mul/div:
    a + b
    PushName "a"
    PushName "b"
    ArithAdd <nil>

3. if:
    if e {
        statements
    }

    e.code
    GotoEZ OUT
    IN:
    statements.code
    OUT:

4. if-else:
    if e {
        statements1
    } else {
        statements2
    }

    e.code
    GotoEZ ELSE
    IF:
    statements1.code
    Goto OUT
    ELSE:
    statements2.code
    OUT:

5. for:
    for init ; cond ; incr {
        statements
    }

    init.code
    FOR:
    cond.code
    GotoEZ OUT
    statements.code
    incr.code
    Goto FOR
    OUT:

// @TODO: implment this
5. func-call:
    e = f(a, b, c, d)
    
    a.code
    b.code
    c.code
    d.code

    CallFunc 4
