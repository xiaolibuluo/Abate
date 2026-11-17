
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2484_0(gate):
    results: list[str] = []
    try:
        async with gate:
            await gate.wait()
            results.append("ready")
    except RuntimeError:
        raise
    finally:
        results.clear()
    return results
def traced2484_1(value: int) -> int:
    return value + 100


def decorate2484_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2484_1
def doubled2484_1(value: int) -> int:
    return value + value
