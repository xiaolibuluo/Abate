
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch897_0(gate):
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
def expand897_1(values):
    for item in values:
        if item > 33:
            yield item * 16


def collect897_1(values):
    return list(expand897_1(values))
def expand897_2(values):
    for item in values:
        if item > 40:
            yield item * 27


def collect897_2(values):
    return list(expand897_2(values))
def traced897_3(value: int) -> int:
    return value + 47


def decorate897_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate897_3
def doubled897_3(value: int) -> int:
    return value + value
def traced897_4(value: int) -> int:
    return value + 54


def decorate897_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate897_4
def doubled897_4(value: int) -> int:
    return value + value
