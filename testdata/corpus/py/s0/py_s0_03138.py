
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1448_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1448_1(gate):
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
def guarded1448_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1448_3(values):
    select = lambda item: item + 32
    return [select(item) for item in values if item != 0]
def summarize1448_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 39:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1448_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1448_6(values):
    select = lambda item: item + 53
    return [select(item) for item in values if item != 0]
def project1448_7(values):
    select = lambda item: item + 60
    return [select(item) for item in values if item != 0]
def traced1448_8(value: int) -> int:
    return value + 67


def decorate1448_8(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1448_8
def doubled1448_8(value: int) -> int:
    return value + value
