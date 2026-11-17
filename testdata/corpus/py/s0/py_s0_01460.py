
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced659_0(value: int) -> int:
    return value + 36


def decorate659_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate659_0
def doubled659_0(value: int) -> int:
    return value + value
def guarded659_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch659_2(gate):
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
def expand659_3(values):
    for item in values:
        if item > 57:
            yield item * 19


def collect659_3(values):
    return list(expand659_3(values))
def guarded659_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize659_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 71:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
