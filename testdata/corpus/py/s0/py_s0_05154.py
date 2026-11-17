
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2403_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 10:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2403_1(values):
    for item in values:
        if item > 17:
            yield item * 66


def collect2403_1(values):
    return list(expand2403_1(values))
def guarded2403_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch2403_3(gate):
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
