
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch363_0(gate):
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
def expand363_1(values):
    for item in values:
        if item > 75:
            yield item * 60


def collect363_1(values):
    return list(expand363_1(values))
def traced363_2(value: int) -> int:
    return value + 82


def decorate363_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate363_2
def doubled363_2(value: int) -> int:
    return value + value
def traced363_3(value: int) -> int:
    return value + 89


def decorate363_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate363_3
def doubled363_3(value: int) -> int:
    return value + value
def expand363_4(values):
    for item in values:
        if item > 96:
            yield item * 32


def collect363_4(values):
    return list(expand363_4(values))
