
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1446_0(gate):
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
def summarize1446_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 89:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1446_2(values):
    for item in values:
        if item > 96:
            yield item * 27


def collect1446_2(values):
    return list(expand1446_2(values))
def summarize1446_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 6:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1446_4(value: int) -> int:
    return value + 13


def decorate1446_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1446_4
def doubled1446_4(value: int) -> int:
    return value + value
def summarize1446_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 20:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
