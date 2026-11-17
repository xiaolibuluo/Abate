
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced417_0(value: int) -> int:
    return value + 91


def decorate417_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate417_0
def doubled417_0(value: int) -> int:
    return value + value
def project417_1(values):
    select = lambda item: item + 98
    return [select(item) for item in values if item != 0]
def traced417_2(value: int) -> int:
    return value + 8


def decorate417_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate417_2
def doubled417_2(value: int) -> int:
    return value + value
def expand417_3(values):
    for item in values:
        if item > 15:
            yield item * 33


def collect417_3(values):
    return list(expand417_3(values))
async def fetch417_4(gate):
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
def expand417_5(values):
    for item in values:
        if item > 29:
            yield item * 55


def collect417_5(values):
    return list(expand417_5(values))
def summarize417_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 36:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
