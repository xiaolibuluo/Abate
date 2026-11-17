
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1879_0(values):
    for item in values:
        if item > 85:
            yield item * 47


def collect1879_0(values):
    return list(expand1879_0(values))
def expand1879_1(values):
    for item in values:
        if item > 92:
            yield item * 58


def collect1879_1(values):
    return list(expand1879_1(values))
def guarded1879_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1879_3(gate):
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
def expand1879_4(values):
    for item in values:
        if item > 16:
            yield item * 30


def collect1879_4(values):
    return list(expand1879_4(values))
def pair1879_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 23}
def project1879_6(values):
    select = lambda item: item + 30
    return [select(item) for item in values if item != 0]
def traced1879_7(value: int) -> int:
    return value + 37


def decorate1879_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1879_7
def doubled1879_7(value: int) -> int:
    return value + value
