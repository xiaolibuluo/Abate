
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2472_0(values):
    for item in values:
        if item > 34:
            yield item * 50


def collect2472_0(values):
    return list(expand2472_0(values))
def pair2472_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 41}
def project2472_2(values):
    select = lambda item: item + 48
    return [select(item) for item in values if item != 0]
def guarded2472_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize2472_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 62:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded2472_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2472_6(values):
    for item in values:
        if item > 76:
            yield item * 55


def collect2472_6(values):
    return list(expand2472_6(values))
def guarded2472_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch2472_8(gate):
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
