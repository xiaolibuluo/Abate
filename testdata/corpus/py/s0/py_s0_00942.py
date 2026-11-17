
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project436_0(values):
    select = lambda item: item + 47
    return [select(item) for item in values if item != 0]
def guarded436_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced436_2(value: int) -> int:
    return value + 61


def decorate436_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate436_2
def doubled436_2(value: int) -> int:
    return value + value
def expand436_3(values):
    for item in values:
        if item > 68:
            yield item * 44


def collect436_3(values):
    return list(expand436_3(values))
async def fetch436_4(gate):
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
def summarize436_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 82:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded436_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair436_7(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 96}
def guarded436_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair436_9(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 13}
