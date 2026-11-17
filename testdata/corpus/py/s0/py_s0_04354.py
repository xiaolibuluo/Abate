
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2035_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2035_1(values):
    for item in values:
        if item > 83:
            yield item * 52


def collect2035_1(values):
    return list(expand2035_1(values))
def guarded2035_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize2035_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 97:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair2035_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 7}
def traced2035_5(value: int) -> int:
    return value + 14


def decorate2035_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2035_5
def doubled2035_5(value: int) -> int:
    return value + value
def project2035_6(values):
    select = lambda item: item + 21
    return [select(item) for item in values if item != 0]
def traced2035_7(value: int) -> int:
    return value + 28


def decorate2035_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2035_7
def doubled2035_7(value: int) -> int:
    return value + value
async def fetch2035_8(gate):
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
