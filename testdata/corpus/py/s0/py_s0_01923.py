
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded875_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand875_1(values):
    for item in values:
        if item > 38:
            yield item * 45


def collect875_1(values):
    return list(expand875_1(values))
def summarize875_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 45:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch875_3(gate):
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
def project875_4(values):
    select = lambda item: item + 59
    return [select(item) for item in values if item != 0]
def traced875_5(value: int) -> int:
    return value + 66


def decorate875_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate875_5
def doubled875_5(value: int) -> int:
    return value + value
def guarded875_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair875_7(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 80}
def project875_8(values):
    select = lambda item: item + 87
    return [select(item) for item in values if item != 0]
async def fetch875_9(gate):
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
