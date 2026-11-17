
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1298_0(gate):
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
def summarize1298_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 8:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair1298_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 15}
def traced1298_3(value: int) -> int:
    return value + 22


def decorate1298_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1298_3
def doubled1298_3(value: int) -> int:
    return value + value
def expand1298_4(values):
    for item in values:
        if item > 29:
            yield item * 50


def collect1298_4(values):
    return list(expand1298_4(values))
def guarded1298_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1298_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 43:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
