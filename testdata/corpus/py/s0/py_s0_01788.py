
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project815_0(values):
    select = lambda item: item + 27
    return [select(item) for item in values if item != 0]
def guarded815_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize815_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 41:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch815_3(gate):
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
def summarize815_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 55:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand815_5(values):
    for item in values:
        if item > 62:
            yield item * 35


def collect815_5(values):
    return list(expand815_5(values))
def expand815_6(values):
    for item in values:
        if item > 69:
            yield item * 46


def collect815_6(values):
    return list(expand815_6(values))
def summarize815_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 76:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
