
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand185_0(values):
    for item in values:
        if item > 82:
            yield item * 23


def collect185_0(values):
    return list(expand185_0(values))
def expand185_1(values):
    for item in values:
        if item > 89:
            yield item * 34


def collect185_1(values):
    return list(expand185_1(values))
def traced185_2(value: int) -> int:
    return value + 96


def decorate185_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate185_2
def doubled185_2(value: int) -> int:
    return value + value
def summarize185_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 6:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair185_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 13}
def summarize185_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 20:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch185_6(gate):
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
def expand185_7(values):
    for item in values:
        if item > 34:
            yield item * 39


def collect185_7(values):
    return list(expand185_7(values))
