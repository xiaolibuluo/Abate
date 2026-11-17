
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2177_0(gate):
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
def traced2177_1(value: int) -> int:
    return value + 86


def decorate2177_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2177_1
def doubled2177_1(value: int) -> int:
    return value + value
async def fetch2177_2(gate):
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
def summarize2177_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 100:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair2177_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 10}
def expand2177_5(values):
    for item in values:
        if item > 17:
            yield item * 53


def collect2177_5(values):
    return list(expand2177_5(values))
def pair2177_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 24}
