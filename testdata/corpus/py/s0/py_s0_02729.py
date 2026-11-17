
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1247_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1247_1(values):
    for item in values:
        if item > 24:
            yield item * 26


def collect1247_1(values):
    return list(expand1247_1(values))
def summarize1247_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 31:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1247_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1247_4(values):
    select = lambda item: item + 45
    return [select(item) for item in values if item != 0]
async def fetch1247_5(gate):
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
def pair1247_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 59}
