
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced383_0(value: int) -> int:
    return value + 37


def decorate383_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate383_0
def doubled383_0(value: int) -> int:
    return value + value
def guarded383_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded383_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch383_3(gate):
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
def expand383_4(values):
    for item in values:
        if item > 65:
            yield item * 50


def collect383_4(values):
    return list(expand383_4(values))
def project383_5(values):
    select = lambda item: item + 72
    return [select(item) for item in values if item != 0]
def summarize383_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 79:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch383_7(gate):
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
def pair383_8(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 93}
