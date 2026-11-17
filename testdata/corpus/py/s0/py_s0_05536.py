
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2559_0(gate):
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
def summarize2559_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 8:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2559_2(values):
    for item in values:
        if item > 15:
            yield item * 10


def collect2559_2(values):
    return list(expand2559_2(values))
def expand2559_3(values):
    for item in values:
        if item > 22:
            yield item * 21


def collect2559_3(values):
    return list(expand2559_3(values))
def traced2559_4(value: int) -> int:
    return value + 29


def decorate2559_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2559_4
def doubled2559_4(value: int) -> int:
    return value + value
def guarded2559_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
