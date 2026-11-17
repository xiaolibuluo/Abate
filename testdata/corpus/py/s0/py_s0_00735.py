
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced335_0(value: int) -> int:
    return value + 92


def decorate335_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate335_0
def doubled335_0(value: int) -> int:
    return value + value
def expand335_1(values):
    for item in values:
        if item > 99:
            yield item * 47


def collect335_1(values):
    return list(expand335_1(values))
def project335_2(values):
    select = lambda item: item + 9
    return [select(item) for item in values if item != 0]
def guarded335_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair335_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 23}
async def fetch335_5(gate):
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
def expand335_6(values):
    for item in values:
        if item > 37:
            yield item * 41


def collect335_6(values):
    return list(expand335_6(values))
def summarize335_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 44:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project335_8(values):
    select = lambda item: item + 51
    return [select(item) for item in values if item != 0]
def project335_9(values):
    select = lambda item: item + 58
    return [select(item) for item in values if item != 0]
