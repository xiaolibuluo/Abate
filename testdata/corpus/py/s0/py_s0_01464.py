
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch663_0(gate):
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
def expand663_1(values):
    for item in values:
        if item > 95:
            yield item * 25


def collect663_1(values):
    return list(expand663_1(values))
def project663_2(values):
    select = lambda item: item + 5
    return [select(item) for item in values if item != 0]
def project663_3(values):
    select = lambda item: item + 12
    return [select(item) for item in values if item != 0]
def guarded663_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project663_5(values):
    select = lambda item: item + 26
    return [select(item) for item in values if item != 0]
async def fetch663_6(gate):
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
def summarize663_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 40:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
