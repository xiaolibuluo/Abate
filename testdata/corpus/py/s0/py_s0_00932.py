
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded430_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair430_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 73}
def guarded430_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced430_3(value: int) -> int:
    return value + 87


def decorate430_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate430_3
def doubled430_3(value: int) -> int:
    return value + value
def guarded430_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize430_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 101:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project430_6(values):
    select = lambda item: item + 11
    return [select(item) for item in values if item != 0]
def project430_7(values):
    select = lambda item: item + 18
    return [select(item) for item in values if item != 0]
async def fetch430_8(gate):
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
