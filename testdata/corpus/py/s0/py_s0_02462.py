
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1138_0(value: int) -> int:
    return value + 55


def decorate1138_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1138_0
def doubled1138_0(value: int) -> int:
    return value + value
def summarize1138_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 62:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1138_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1138_3(values):
    for item in values:
        if item > 76:
            yield item * 17


def collect1138_3(values):
    return list(expand1138_3(values))
async def fetch1138_4(gate):
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
