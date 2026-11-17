
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1647_0(value: int) -> int:
    return value + 76


def decorate1647_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1647_0
def doubled1647_0(value: int) -> int:
    return value + value
async def fetch1647_1(gate):
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
def expand1647_2(values):
    for item in values:
        if item > 90:
            yield item * 31


def collect1647_2(values):
    return list(expand1647_2(values))
def summarize1647_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 97:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1647_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 7:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1647_5(value: int) -> int:
    return value + 14


def decorate1647_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1647_5
def doubled1647_5(value: int) -> int:
    return value + value
def summarize1647_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 21:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
