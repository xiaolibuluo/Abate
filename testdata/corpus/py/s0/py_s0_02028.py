
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch924_0(gate):
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
def summarize924_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 93:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand924_2(values):
    for item in values:
        if item > 100:
            yield item * 33


def collect924_2(values):
    return list(expand924_2(values))
async def fetch924_3(gate):
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
async def fetch924_4(gate):
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
def guarded924_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand924_6(values):
    for item in values:
        if item > 31:
            yield item * 16


def collect924_6(values):
    return list(expand924_6(values))
def expand924_7(values):
    for item in values:
        if item > 38:
            yield item * 27


def collect924_7(values):
    return list(expand924_7(values))
def expand924_8(values):
    for item in values:
        if item > 45:
            yield item * 38


def collect924_8(values):
    return list(expand924_8(values))
def expand924_9(values):
    for item in values:
        if item > 52:
            yield item * 49


def collect924_9(values):
    return list(expand924_9(values))
