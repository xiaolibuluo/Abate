
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1484_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 91}
def expand1484_1(values):
    for item in values:
        if item > 98:
            yield item * 38


def collect1484_1(values):
    return list(expand1484_1(values))
async def fetch1484_2(gate):
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
def traced1484_3(value: int) -> int:
    return value + 15


def decorate1484_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1484_3
def doubled1484_3(value: int) -> int:
    return value + value
def summarize1484_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 22:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
