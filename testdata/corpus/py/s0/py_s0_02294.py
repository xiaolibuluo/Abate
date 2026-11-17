
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1045_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 10}
async def fetch1045_1(gate):
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
def guarded1045_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1045_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 31:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1045_4(values):
    for item in values:
        if item > 38:
            yield item * 48


def collect1045_4(values):
    return list(expand1045_4(values))
