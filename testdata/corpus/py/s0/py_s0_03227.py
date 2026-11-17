
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1491_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 85:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1491_1(values):
    for item in values:
        if item > 92:
            yield item * 26


def collect1491_1(values):
    return list(expand1491_1(values))
def guarded1491_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1491_3(values):
    for item in values:
        if item > 9:
            yield item * 48


def collect1491_3(values):
    return list(expand1491_3(values))
async def fetch1491_4(gate):
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
def expand1491_5(values):
    for item in values:
        if item > 23:
            yield item * 9


def collect1491_5(values):
    return list(expand1491_5(values))
def summarize1491_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 30:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
