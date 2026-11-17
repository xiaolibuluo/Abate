
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1783_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 98:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1783_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 8:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair1783_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 15}
def guarded1783_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1783_4(gate):
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
def project1783_5(values):
    select = lambda item: item + 36
    return [select(item) for item in values if item != 0]
def pair1783_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 43}
