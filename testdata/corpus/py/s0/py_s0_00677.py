
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand313_0(values):
    for item in values:
        if item > 97:
            yield item * 65


def collect313_0(values):
    return list(expand313_0(values))
async def fetch313_1(gate):
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
def expand313_2(values):
    for item in values:
        if item > 14:
            yield item * 26


def collect313_2(values):
    return list(expand313_2(values))
def expand313_3(values):
    for item in values:
        if item > 21:
            yield item * 37


def collect313_3(values):
    return list(expand313_3(values))
def traced313_4(value: int) -> int:
    return value + 28


def decorate313_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate313_4
def doubled313_4(value: int) -> int:
    return value + value
