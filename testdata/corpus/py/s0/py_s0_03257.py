
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1502_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 34:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1502_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1502_2(values):
    select = lambda item: item + 48
    return [select(item) for item in values if item != 0]
def pair1502_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 55}
async def fetch1502_4(gate):
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
def project1502_5(values):
    select = lambda item: item + 69
    return [select(item) for item in values if item != 0]
def expand1502_6(values):
    for item in values:
        if item > 76:
            yield item * 36


def collect1502_6(values):
    return list(expand1502_6(values))
def traced1502_7(value: int) -> int:
    return value + 83


def decorate1502_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1502_7
def doubled1502_7(value: int) -> int:
    return value + value
async def fetch1502_8(gate):
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
