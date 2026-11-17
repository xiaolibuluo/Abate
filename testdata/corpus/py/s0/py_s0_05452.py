
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2517_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project2517_1(values):
    select = lambda item: item + 44
    return [select(item) for item in values if item != 0]
def expand2517_2(values):
    for item in values:
        if item > 51:
            yield item * 21


def collect2517_2(values):
    return list(expand2517_2(values))
def expand2517_3(values):
    for item in values:
        if item > 58:
            yield item * 32


def collect2517_3(values):
    return list(expand2517_3(values))
def summarize2517_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 65:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize2517_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 72:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded2517_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch2517_7(gate):
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
