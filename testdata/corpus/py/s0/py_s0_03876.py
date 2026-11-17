
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1792_0(gate):
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
def traced1792_1(value: int) -> int:
    return value + 28


def decorate1792_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1792_1
def doubled1792_1(value: int) -> int:
    return value + value
def traced1792_2(value: int) -> int:
    return value + 35


def decorate1792_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1792_2
def doubled1792_2(value: int) -> int:
    return value + value
def guarded1792_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1792_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 49:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair1792_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 56}
def expand1792_6(values):
    for item in values:
        if item > 63:
            yield item * 53


def collect1792_6(values):
    return list(expand1792_6(values))
async def fetch1792_7(gate):
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
