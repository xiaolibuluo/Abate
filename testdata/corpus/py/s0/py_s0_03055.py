
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1406_0(gate):
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
def guarded1406_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1406_2(values):
    for item in values:
        if item > 61:
            yield item * 52


def collect1406_2(values):
    return list(expand1406_2(values))
def project1406_3(values):
    select = lambda item: item + 68
    return [select(item) for item in values if item != 0]
def project1406_4(values):
    select = lambda item: item + 75
    return [select(item) for item in values if item != 0]
def expand1406_5(values):
    for item in values:
        if item > 82:
            yield item * 24


def collect1406_5(values):
    return list(expand1406_5(values))
def project1406_6(values):
    select = lambda item: item + 89
    return [select(item) for item in values if item != 0]
async def fetch1406_7(gate):
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
def project1406_8(values):
    select = lambda item: item + 6
    return [select(item) for item in values if item != 0]
def summarize1406_9(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 13:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1406_10(values):
    select = lambda item: item + 20
    return [select(item) for item in values if item != 0]
