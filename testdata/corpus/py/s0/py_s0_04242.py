
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1979_0(gate):
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
def project1979_1(values):
    select = lambda item: item + 34
    return [select(item) for item in values if item != 0]
def guarded1979_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1979_3(gate):
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
def expand1979_4(values):
    for item in values:
        if item > 55:
            yield item * 59


def collect1979_4(values):
    return list(expand1979_4(values))
def summarize1979_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 62:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1979_6(values):
    select = lambda item: item + 69
    return [select(item) for item in values if item != 0]
def summarize1979_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 76:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
