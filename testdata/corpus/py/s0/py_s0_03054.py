
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1405_0(gate):
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
def project1405_1(values):
    select = lambda item: item + 41
    return [select(item) for item in values if item != 0]
def pair1405_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 48}
def project1405_3(values):
    select = lambda item: item + 55
    return [select(item) for item in values if item != 0]
def guarded1405_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded1405_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1405_6(values):
    for item in values:
        if item > 76:
            yield item * 28


def collect1405_6(values):
    return list(expand1405_6(values))
def summarize1405_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 83:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1405_8(values):
    select = lambda item: item + 90
    return [select(item) for item in values if item != 0]
async def fetch1405_9(gate):
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
