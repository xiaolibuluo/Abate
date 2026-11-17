
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2122_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2122_1(value: int) -> int:
    return value + 50


def decorate2122_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2122_1
def doubled2122_1(value: int) -> int:
    return value + value
async def fetch2122_2(gate):
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
def summarize2122_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 64:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2122_4(values):
    for item in values:
        if item > 71:
            yield item * 23


def collect2122_4(values):
    return list(expand2122_4(values))
def project2122_5(values):
    select = lambda item: item + 78
    return [select(item) for item in values if item != 0]
async def fetch2122_6(gate):
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
