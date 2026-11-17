
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1409_0(values):
    select = lambda item: item + 86
    return [select(item) for item in values if item != 0]
def project1409_1(values):
    select = lambda item: item + 93
    return [select(item) for item in values if item != 0]
async def fetch1409_2(gate):
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
def summarize1409_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 10:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch1409_4(gate):
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
def summarize1409_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 24:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1409_6(value: int) -> int:
    return value + 31


def decorate1409_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1409_6
def doubled1409_6(value: int) -> int:
    return value + value
def summarize1409_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 38:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1409_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
