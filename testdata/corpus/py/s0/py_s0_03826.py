
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1764_0(gate):
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
def expand1764_1(values):
    for item in values:
        if item > 52:
            yield item * 46


def collect1764_1(values):
    return list(expand1764_1(values))
def traced1764_2(value: int) -> int:
    return value + 59


def decorate1764_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1764_2
def doubled1764_2(value: int) -> int:
    return value + value
