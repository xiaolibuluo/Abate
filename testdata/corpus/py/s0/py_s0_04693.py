
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2173_0(values):
    select = lambda item: item + 27
    return [select(item) for item in values if item != 0]
async def fetch2173_1(gate):
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
def traced2173_2(value: int) -> int:
    return value + 41


def decorate2173_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2173_2
def doubled2173_2(value: int) -> int:
    return value + value
def project2173_3(values):
    select = lambda item: item + 48
    return [select(item) for item in values if item != 0]
def expand2173_4(values):
    for item in values:
        if item > 55:
            yield item * 14


def collect2173_4(values):
    return list(expand2173_4(values))
def pair2173_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 62}
def summarize2173_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 69:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded2173_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
