
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1798_0(values):
    for item in values:
        if item > 99:
            yield item * 29


def collect1798_0(values):
    return list(expand1798_0(values))
def pair1798_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 9}
def pair1798_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 16}
def traced1798_3(value: int) -> int:
    return value + 23


def decorate1798_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1798_3
def doubled1798_3(value: int) -> int:
    return value + value
def summarize1798_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 30:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1798_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1798_6(values):
    for item in values:
        if item > 44:
            yield item * 34


def collect1798_6(values):
    return list(expand1798_6(values))
async def fetch1798_7(gate):
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
async def fetch1798_8(gate):
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
def guarded1798_9(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
