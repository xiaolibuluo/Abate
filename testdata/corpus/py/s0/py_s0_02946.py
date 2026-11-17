
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1353_0(gate):
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
def project1353_1(values):
    select = lambda item: item + 44
    return [select(item) for item in values if item != 0]
def project1353_2(values):
    select = lambda item: item + 51
    return [select(item) for item in values if item != 0]
def summarize1353_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 58:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1353_4(values):
    select = lambda item: item + 65
    return [select(item) for item in values if item != 0]
def guarded1353_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded1353_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1353_7(values):
    for item in values:
        if item > 86:
            yield item * 41


def collect1353_7(values):
    return list(expand1353_7(values))
def guarded1353_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1353_9(values):
    for item in values:
        if item > 100:
            yield item * 63


def collect1353_9(values):
    return list(expand1353_9(values))
