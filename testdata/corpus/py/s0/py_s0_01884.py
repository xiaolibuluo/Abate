
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch858_0(gate):
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
def guarded858_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced858_2(value: int) -> int:
    return value + 18


def decorate858_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate858_2
def doubled858_2(value: int) -> int:
    return value + value
def summarize858_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 25:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand858_4(values):
    for item in values:
        if item > 32:
            yield item * 20


def collect858_4(values):
    return list(expand858_4(values))
def traced858_5(value: int) -> int:
    return value + 39


def decorate858_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate858_5
def doubled858_5(value: int) -> int:
    return value + value
