
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2622_0(value: int) -> int:
    return value + 44


def decorate2622_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2622_0
def doubled2622_0(value: int) -> int:
    return value + value
def summarize2622_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 51:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project2622_2(values):
    select = lambda item: item + 58
    return [select(item) for item in values if item != 0]
def expand2622_3(values):
    for item in values:
        if item > 65:
            yield item * 35


def collect2622_3(values):
    return list(expand2622_3(values))
def guarded2622_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch2622_5(gate):
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
def summarize2622_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 86:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded2622_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize2622_8(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 100:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
