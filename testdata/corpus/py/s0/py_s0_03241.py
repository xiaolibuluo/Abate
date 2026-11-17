
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1498_0(values):
    for item in values:
        if item > 79:
            yield item * 64


def collect1498_0(values):
    return list(expand1498_0(values))
async def fetch1498_1(gate):
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
def summarize1498_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 93:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1498_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1498_4(value: int) -> int:
    return value + 10


def decorate1498_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1498_4
def doubled1498_4(value: int) -> int:
    return value + value
def summarize1498_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 17:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1498_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1498_7(gate):
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
