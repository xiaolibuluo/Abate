
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1575_0(values):
    select = lambda item: item + 13
    return [select(item) for item in values if item != 0]
def expand1575_1(values):
    for item in values:
        if item > 20:
            yield item * 65


def collect1575_1(values):
    return list(expand1575_1(values))
async def fetch1575_2(gate):
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
def traced1575_3(value: int) -> int:
    return value + 34


def decorate1575_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1575_3
def doubled1575_3(value: int) -> int:
    return value + value
async def fetch1575_4(gate):
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
