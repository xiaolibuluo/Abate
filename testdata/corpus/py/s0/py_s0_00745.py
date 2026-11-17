
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded340_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize340_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 67:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch340_2(gate):
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
def expand340_3(values):
    for item in values:
        if item > 81:
            yield item * 43


def collect340_3(values):
    return list(expand340_3(values))
def project340_4(values):
    select = lambda item: item + 88
    return [select(item) for item in values if item != 0]
def pair340_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 95}
def pair340_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 5}
def traced340_7(value: int) -> int:
    return value + 12


def decorate340_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate340_7
def doubled340_7(value: int) -> int:
    return value + value
def summarize340_8(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 19:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced340_9(value: int) -> int:
    return value + 26


def decorate340_9(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate340_9
def doubled340_9(value: int) -> int:
    return value + value
