
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1503_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1503_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 54:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1503_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 61:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1503_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair1503_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 75}
def project1503_5(values):
    select = lambda item: item + 82
    return [select(item) for item in values if item != 0]
async def fetch1503_6(gate):
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
def traced1503_7(value: int) -> int:
    return value + 96


def decorate1503_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1503_7
def doubled1503_7(value: int) -> int:
    return value + value
def guarded1503_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
