
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand794_0(values):
    for item in values:
        if item > 45:
            yield item * 16


def collect794_0(values):
    return list(expand794_0(values))
def traced794_1(value: int) -> int:
    return value + 52


def decorate794_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate794_1
def doubled794_1(value: int) -> int:
    return value + value
async def fetch794_2(gate):
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
async def fetch794_3(gate):
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
