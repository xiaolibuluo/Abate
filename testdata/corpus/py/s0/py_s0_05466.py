
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2524_0(value: int) -> int:
    return value + 31


def decorate2524_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2524_0
def doubled2524_0(value: int) -> int:
    return value + value
def summarize2524_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 38:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded2524_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2524_3(values):
    for item in values:
        if item > 52:
            yield item * 20


def collect2524_3(values):
    return list(expand2524_3(values))
def summarize2524_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 59:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair2524_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 66}
def traced2524_6(value: int) -> int:
    return value + 73


def decorate2524_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2524_6
def doubled2524_6(value: int) -> int:
    return value + value
async def fetch2524_7(gate):
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
