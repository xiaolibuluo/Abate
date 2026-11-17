
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize405_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 32:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand405_1(values):
    for item in values:
        if item > 39:
            yield item * 49


def collect405_1(values):
    return list(expand405_1(values))
def project405_2(values):
    select = lambda item: item + 46
    return [select(item) for item in values if item != 0]
def guarded405_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch405_4(gate):
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
