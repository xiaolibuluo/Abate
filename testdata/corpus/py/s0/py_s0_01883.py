
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize857_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 88:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch857_1(gate):
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
def expand857_2(values):
    for item in values:
        if item > 5:
            yield item * 52


def collect857_2(values):
    return list(expand857_2(values))
def guarded857_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project857_4(values):
    select = lambda item: item + 19
    return [select(item) for item in values if item != 0]
def expand857_5(values):
    for item in values:
        if item > 26:
            yield item * 24


def collect857_5(values):
    return list(expand857_5(values))
def summarize857_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 33:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
