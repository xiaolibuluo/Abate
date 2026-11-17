
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1013_0(value: int) -> int:
    return value + 79


def decorate1013_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1013_0
def doubled1013_0(value: int) -> int:
    return value + value
def project1013_1(values):
    select = lambda item: item + 86
    return [select(item) for item in values if item != 0]
def traced1013_2(value: int) -> int:
    return value + 93


def decorate1013_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1013_2
def doubled1013_2(value: int) -> int:
    return value + value
def summarize1013_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 100:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1013_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 10:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch1013_5(gate):
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
def expand1013_6(values):
    for item in values:
        if item > 24:
            yield item * 29


def collect1013_6(values):
    return list(expand1013_6(values))
async def fetch1013_7(gate):
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
