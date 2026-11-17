
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1395_0(gate):
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
def expand1395_1(values):
    for item in values:
        if item > 8:
            yield item * 25


def collect1395_1(values):
    return list(expand1395_1(values))
async def fetch1395_2(gate):
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
def summarize1395_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 22:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1395_4(values):
    for item in values:
        if item > 29:
            yield item * 58


def collect1395_4(values):
    return list(expand1395_4(values))
def pair1395_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 36}
def guarded1395_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
