
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand308_0(values):
    for item in values:
        if item > 32:
            yield item * 30


def collect308_0(values):
    return list(expand308_0(values))
def expand308_1(values):
    for item in values:
        if item > 39:
            yield item * 41


def collect308_1(values):
    return list(expand308_1(values))
async def fetch308_2(gate):
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
def traced308_3(value: int) -> int:
    return value + 53


def decorate308_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate308_3
def doubled308_3(value: int) -> int:
    return value + value
