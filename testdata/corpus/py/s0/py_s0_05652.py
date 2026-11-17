
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2602_0(gate):
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
async def fetch2602_1(gate):
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
def expand2602_2(values):
    for item in values:
        if item > 89:
            yield item * 67


def collect2602_2(values):
    return list(expand2602_2(values))
async def fetch2602_3(gate):
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
def traced2602_4(value: int) -> int:
    return value + 6


def decorate2602_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2602_4
def doubled2602_4(value: int) -> int:
    return value + value
