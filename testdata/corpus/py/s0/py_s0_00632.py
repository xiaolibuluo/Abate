
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded289_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch289_1(gate):
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
def expand289_2(values):
    for item in values:
        if item > 90:
            yield item * 41


def collect289_2(values):
    return list(expand289_2(values))
def summarize289_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 97:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch289_4(gate):
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
def traced289_5(value: int) -> int:
    return value + 14


def decorate289_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate289_5
def doubled289_5(value: int) -> int:
    return value + value
async def fetch289_6(gate):
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
def pair289_7(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 28}
def expand289_8(values):
    for item in values:
        if item > 35:
            yield item * 46


def collect289_8(values):
    return list(expand289_8(values))
