
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1408_0(gate):
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
def expand1408_1(values):
    for item in values:
        if item > 80:
            yield item * 55


def collect1408_1(values):
    return list(expand1408_1(values))
def expand1408_2(values):
    for item in values:
        if item > 87:
            yield item * 66


def collect1408_2(values):
    return list(expand1408_2(values))
def pair1408_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 94}
def guarded1408_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1408_5(values):
    for item in values:
        if item > 11:
            yield item * 38


def collect1408_5(values):
    return list(expand1408_5(values))
async def fetch1408_6(gate):
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
async def fetch1408_7(gate):
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
def summarize1408_8(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 32:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
