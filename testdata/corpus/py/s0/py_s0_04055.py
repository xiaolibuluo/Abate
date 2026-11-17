
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1883_0(values):
    for item in values:
        if item > 40:
            yield item * 14


def collect1883_0(values):
    return list(expand1883_0(values))
def pair1883_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 47}
async def fetch1883_2(gate):
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
def project1883_3(values):
    select = lambda item: item + 61
    return [select(item) for item in values if item != 0]
async def fetch1883_4(gate):
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
def guarded1883_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair1883_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 82}
def traced1883_7(value: int) -> int:
    return value + 89


def decorate1883_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1883_7
def doubled1883_7(value: int) -> int:
    return value + value
