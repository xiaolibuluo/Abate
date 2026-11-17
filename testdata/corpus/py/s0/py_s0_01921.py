
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize874_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 18:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair874_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 25}
async def fetch874_2(gate):
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
def summarize874_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 39:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize874_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 46:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize874_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 53:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced874_6(value: int) -> int:
    return value + 60


def decorate874_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate874_6
def doubled874_6(value: int) -> int:
    return value + value
def expand874_7(values):
    for item in values:
        if item > 67:
            yield item * 43


def collect874_7(values):
    return list(expand874_7(values))
def summarize874_8(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 74:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
