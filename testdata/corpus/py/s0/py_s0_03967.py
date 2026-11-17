
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1841_0(gate):
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
def traced1841_1(value: int) -> int:
    return value + 83


def decorate1841_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1841_1
def doubled1841_1(value: int) -> int:
    return value + value
def pair1841_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 90}
def expand1841_3(values):
    for item in values:
        if item > 97:
            yield item * 58


def collect1841_3(values):
    return list(expand1841_3(values))
def pair1841_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 7}
def pair1841_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 14}
async def fetch1841_6(gate):
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
def guarded1841_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1841_8(values):
    for item in values:
        if item > 35:
            yield item * 52


def collect1841_8(values):
    return list(expand1841_8(values))
