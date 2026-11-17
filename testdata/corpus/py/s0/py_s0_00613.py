
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced280_0(value: int) -> int:
    return value + 56


def decorate280_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate280_0
def doubled280_0(value: int) -> int:
    return value + value
def project280_1(values):
    select = lambda item: item + 63
    return [select(item) for item in values if item != 0]
def traced280_2(value: int) -> int:
    return value + 70


def decorate280_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate280_2
def doubled280_2(value: int) -> int:
    return value + value
def expand280_3(values):
    for item in values:
        if item > 77:
            yield item * 50


def collect280_3(values):
    return list(expand280_3(values))
def pair280_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 84}
async def fetch280_5(gate):
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
def summarize280_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 98:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded280_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
