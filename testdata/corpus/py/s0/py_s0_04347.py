
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2030_0(values):
    for item in values:
        if item > 11:
            yield item * 67


def collect2030_0(values):
    return list(expand2030_0(values))
def summarize2030_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 18:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded2030_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2030_3(values):
    for item in values:
        if item > 32:
            yield item * 39


def collect2030_3(values):
    return list(expand2030_3(values))
async def fetch2030_4(gate):
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
def expand2030_5(values):
    for item in values:
        if item > 46:
            yield item * 61


def collect2030_5(values):
    return list(expand2030_5(values))
def project2030_6(values):
    select = lambda item: item + 53
    return [select(item) for item in values if item != 0]
def project2030_7(values):
    select = lambda item: item + 60
    return [select(item) for item in values if item != 0]
async def fetch2030_8(gate):
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
