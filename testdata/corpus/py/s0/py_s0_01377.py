
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch626_0(gate):
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
async def fetch626_1(gate):
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
def guarded626_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair626_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 16}
def expand626_4(values):
    for item in values:
        if item > 23:
            yield item * 43


def collect626_4(values):
    return list(expand626_4(values))
def project626_5(values):
    select = lambda item: item + 30
    return [select(item) for item in values if item != 0]
def summarize626_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 37:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced626_7(value: int) -> int:
    return value + 44


def decorate626_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate626_7
def doubled626_7(value: int) -> int:
    return value + value
def summarize626_8(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 51:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
