
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1351_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 11:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1351_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 18:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch1351_2(gate):
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
def pair1351_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 32}
def project1351_4(values):
    select = lambda item: item + 39
    return [select(item) for item in values if item != 0]
def guarded1351_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1351_6(gate):
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
def pair1351_7(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 60}
