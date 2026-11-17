
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1118_0(value: int) -> int:
    return value + 86


def decorate1118_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1118_0
def doubled1118_0(value: int) -> int:
    return value + value
def pair1118_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 93}
def pair1118_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 100}
def summarize1118_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 10:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1118_4(values):
    for item in values:
        if item > 17:
            yield item * 10


def collect1118_4(values):
    return list(expand1118_4(values))
def pair1118_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 24}
def pair1118_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 31}
async def fetch1118_7(gate):
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
def summarize1118_8(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 45:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1118_9(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
