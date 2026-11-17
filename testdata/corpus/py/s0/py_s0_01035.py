
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand480_0(values):
    for item in values:
        if item > 37:
            yield item * 14


def collect480_0(values):
    return list(expand480_0(values))
async def fetch480_1(gate):
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
def guarded480_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project480_3(values):
    select = lambda item: item + 58
    return [select(item) for item in values if item != 0]
def summarize480_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 65:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced480_5(value: int) -> int:
    return value + 72


def decorate480_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate480_5
def doubled480_5(value: int) -> int:
    return value + value
def traced480_6(value: int) -> int:
    return value + 79


def decorate480_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate480_6
def doubled480_6(value: int) -> int:
    return value + value
def expand480_7(values):
    for item in values:
        if item > 86:
            yield item * 30


def collect480_7(values):
    return list(expand480_7(values))
def pair480_8(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 93}
