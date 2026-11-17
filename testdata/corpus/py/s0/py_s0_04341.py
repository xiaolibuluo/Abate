
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2027_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded2027_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2027_2(values):
    for item in values:
        if item > 83:
            yield item * 68


def collect2027_2(values):
    return list(expand2027_2(values))
def project2027_3(values):
    select = lambda item: item + 90
    return [select(item) for item in values if item != 0]
async def fetch2027_4(gate):
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
def expand2027_5(values):
    for item in values:
        if item > 7:
            yield item * 40


def collect2027_5(values):
    return list(expand2027_5(values))
def summarize2027_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 14:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair2027_7(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 21}
