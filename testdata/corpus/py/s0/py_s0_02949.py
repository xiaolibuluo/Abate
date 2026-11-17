
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1355_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair1355_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 70}
def project1355_2(values):
    select = lambda item: item + 77
    return [select(item) for item in values if item != 0]
def project1355_3(values):
    select = lambda item: item + 84
    return [select(item) for item in values if item != 0]
def project1355_4(values):
    select = lambda item: item + 91
    return [select(item) for item in values if item != 0]
def guarded1355_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1355_6(gate):
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
def project1355_7(values):
    select = lambda item: item + 15
    return [select(item) for item in values if item != 0]
def guarded1355_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1355_9(values):
    select = lambda item: item + 29
    return [select(item) for item in values if item != 0]
def summarize1355_10(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 36:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
