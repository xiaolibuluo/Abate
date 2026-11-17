
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2617_0(values):
    for item in values:
        if item > 76:
            yield item * 28


def collect2617_0(values):
    return list(expand2617_0(values))
def summarize2617_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 83:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize2617_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 90:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch2617_3(gate):
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
def expand2617_4(values):
    for item in values:
        if item > 7:
            yield item * 11


def collect2617_4(values):
    return list(expand2617_4(values))
def traced2617_5(value: int) -> int:
    return value + 14


def decorate2617_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2617_5
def doubled2617_5(value: int) -> int:
    return value + value
def guarded2617_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project2617_7(values):
    select = lambda item: item + 28
    return [select(item) for item in values if item != 0]
