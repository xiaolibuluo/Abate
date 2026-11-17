
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch872_0(gate):
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
def summarize872_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 96:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced872_2(value: int) -> int:
    return value + 6


def decorate872_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate872_2
def doubled872_2(value: int) -> int:
    return value + value
def pair872_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 13}
def expand872_4(values):
    for item in values:
        if item > 20:
            yield item * 57


def collect872_4(values):
    return list(expand872_4(values))
def pair872_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 27}
async def fetch872_6(gate):
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
async def fetch872_7(gate):
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
def expand872_8(values):
    for item in values:
        if item > 48:
            yield item * 40


def collect872_8(values):
    return list(expand872_8(values))
