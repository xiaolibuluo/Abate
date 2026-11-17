
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1291_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1291_1(gate):
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
def project1291_2(values):
    select = lambda item: item + 21
    return [select(item) for item in values if item != 0]
def expand1291_3(values):
    for item in values:
        if item > 28:
            yield item * 51


def collect1291_3(values):
    return list(expand1291_3(values))
def guarded1291_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1291_5(values):
    for item in values:
        if item > 42:
            yield item * 12


def collect1291_5(values):
    return list(expand1291_5(values))
def summarize1291_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 49:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
