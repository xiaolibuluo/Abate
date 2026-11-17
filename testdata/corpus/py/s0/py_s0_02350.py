
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1081_0(value: int) -> int:
    return value + 90


def decorate1081_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1081_0
def doubled1081_0(value: int) -> int:
    return value + value
def expand1081_1(values):
    for item in values:
        if item > 97:
            yield item * 23


def collect1081_1(values):
    return list(expand1081_1(values))
async def fetch1081_2(gate):
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
