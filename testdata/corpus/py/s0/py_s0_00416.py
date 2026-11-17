
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize188_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 24:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded188_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand188_2(values):
    for item in values:
        if item > 38:
            yield item * 66


def collect188_2(values):
    return list(expand188_2(values))
def summarize188_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 45:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project188_4(values):
    select = lambda item: item + 52
    return [select(item) for item in values if item != 0]
def summarize188_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 59:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch188_6(gate):
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
def project188_7(values):
    select = lambda item: item + 73
    return [select(item) for item in values if item != 0]
def project188_8(values):
    select = lambda item: item + 80
    return [select(item) for item in values if item != 0]
