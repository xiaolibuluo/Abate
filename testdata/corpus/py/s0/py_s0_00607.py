
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand277_0(values):
    for item in values:
        if item > 17:
            yield item * 57


def collect277_0(values):
    return list(expand277_0(values))
async def fetch277_1(gate):
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
def expand277_2(values):
    for item in values:
        if item > 31:
            yield item * 18


def collect277_2(values):
    return list(expand277_2(values))
def traced277_3(value: int) -> int:
    return value + 38


def decorate277_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate277_3
def doubled277_3(value: int) -> int:
    return value + value
def pair277_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 45}
def expand277_5(values):
    for item in values:
        if item > 52:
            yield item * 51


def collect277_5(values):
    return list(expand277_5(values))
def traced277_6(value: int) -> int:
    return value + 59


def decorate277_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate277_6
def doubled277_6(value: int) -> int:
    return value + value
async def fetch277_7(gate):
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
