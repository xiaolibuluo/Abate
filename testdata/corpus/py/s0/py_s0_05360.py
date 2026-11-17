
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2475_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded2475_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project2475_2(values):
    select = lambda item: item + 87
    return [select(item) for item in values if item != 0]
def project2475_3(values):
    select = lambda item: item + 94
    return [select(item) for item in values if item != 0]
async def fetch2475_4(gate):
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
def pair2475_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 11}
async def fetch2475_6(gate):
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
def expand2475_7(values):
    for item in values:
        if item > 25:
            yield item * 26


def collect2475_7(values):
    return list(expand2475_7(values))
def expand2475_8(values):
    for item in values:
        if item > 32:
            yield item * 37


def collect2475_8(values):
    return list(expand2475_8(values))
def summarize2475_9(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 39:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
