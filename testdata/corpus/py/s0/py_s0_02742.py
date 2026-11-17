
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1251_0(value: int) -> int:
    return value + 69


def decorate1251_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1251_0
def doubled1251_0(value: int) -> int:
    return value + value
def expand1251_1(values):
    for item in values:
        if item > 76:
            yield item * 54


def collect1251_1(values):
    return list(expand1251_1(values))
async def fetch1251_2(gate):
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
def summarize1251_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 90:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch1251_4(gate):
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
def pair1251_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 7}
def expand1251_6(values):
    for item in values:
        if item > 14:
            yield item * 48


def collect1251_6(values):
    return list(expand1251_6(values))
