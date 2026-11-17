
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1202_0(value: int) -> int:
    return value + 14


def decorate1202_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1202_0
def doubled1202_0(value: int) -> int:
    return value + value
def pair1202_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 21}
async def fetch1202_2(gate):
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
def expand1202_3(values):
    for item in values:
        if item > 35:
            yield item * 38


def collect1202_3(values):
    return list(expand1202_3(values))
def traced1202_4(value: int) -> int:
    return value + 42


def decorate1202_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1202_4
def doubled1202_4(value: int) -> int:
    return value + value
def expand1202_5(values):
    for item in values:
        if item > 49:
            yield item * 60


def collect1202_5(values):
    return list(expand1202_5(values))
def guarded1202_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1202_7(values):
    for item in values:
        if item > 63:
            yield item * 21


def collect1202_7(values):
    return list(expand1202_7(values))
