
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project574_0(values):
    select = lambda item: item + 95
    return [select(item) for item in values if item != 0]
def guarded574_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced574_2(value: int) -> int:
    return value + 12


def decorate574_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate574_2
def doubled574_2(value: int) -> int:
    return value + value
def summarize574_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 19:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project574_4(values):
    select = lambda item: item + 26
    return [select(item) for item in values if item != 0]
def summarize574_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 33:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded574_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch574_7(gate):
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
