
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch365_0(gate):
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
def expand365_1(values):
    for item in values:
        if item > 101:
            yield item * 13


def collect365_1(values):
    return list(expand365_1(values))
async def fetch365_2(gate):
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
def summarize365_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 18:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced365_4(value: int) -> int:
    return value + 25


def decorate365_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate365_4
def doubled365_4(value: int) -> int:
    return value + value
