
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1652_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1652_1(gate):
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
def summarize1652_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 58:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch1652_3(gate):
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
def summarize1652_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 72:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair1652_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 79}
def expand1652_6(values):
    for item in values:
        if item > 86:
            yield item * 49


def collect1652_6(values):
    return list(expand1652_6(values))
def expand1652_7(values):
    for item in values:
        if item > 93:
            yield item * 60


def collect1652_7(values):
    return list(expand1652_7(values))
def summarize1652_8(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 100:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
