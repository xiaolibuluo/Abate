
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize799_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 13:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch799_1(gate):
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
def pair799_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 27}
def expand799_3(values):
    for item in values:
        if item > 34:
            yield item * 23


def collect799_3(values):
    return list(expand799_3(values))
def traced799_4(value: int) -> int:
    return value + 41


def decorate799_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate799_4
def doubled799_4(value: int) -> int:
    return value + value
