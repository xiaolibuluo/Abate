
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1588_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1588_1(value: int) -> int:
    return value + 92


def decorate1588_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1588_1
def doubled1588_1(value: int) -> int:
    return value + value
def pair1588_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 99}
def summarize1588_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 9:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch1588_4(gate):
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
def guarded1588_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1588_6(values):
    select = lambda item: item + 30
    return [select(item) for item in values if item != 0]
