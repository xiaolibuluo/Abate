
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced26_0(value: int) -> int:
    return value + 52


def decorate26_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate26_0
def doubled26_0(value: int) -> int:
    return value + value
def expand26_1(values):
    for item in values:
        if item > 59:
            yield item * 19


def collect26_1(values):
    return list(expand26_1(values))
def guarded26_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair26_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 73}
async def fetch26_4(gate):
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
async def fetch26_5(gate):
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
