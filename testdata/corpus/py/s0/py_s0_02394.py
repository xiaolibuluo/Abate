
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1104_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1104_1(value: int) -> int:
    return value + 8


def decorate1104_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1104_1
def doubled1104_1(value: int) -> int:
    return value + value
def pair1104_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 15}
def expand1104_3(values):
    for item in values:
        if item > 22:
            yield item * 23


def collect1104_3(values):
    return list(expand1104_3(values))
def expand1104_4(values):
    for item in values:
        if item > 29:
            yield item * 34


def collect1104_4(values):
    return list(expand1104_4(values))
def guarded1104_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1104_6(gate):
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
def summarize1104_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 50:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
