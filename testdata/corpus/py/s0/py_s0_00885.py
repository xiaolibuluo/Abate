
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch404_0(gate):
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
def traced404_1(value: int) -> int:
    return value + 26


def decorate404_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate404_1
def doubled404_1(value: int) -> int:
    return value + value
def expand404_2(values):
    for item in values:
        if item > 33:
            yield item * 53


def collect404_2(values):
    return list(expand404_2(values))
def summarize404_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 40:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
