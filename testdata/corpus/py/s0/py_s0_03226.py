
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1490_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1490_1(values):
    for item in values:
        if item > 79:
            yield item * 19


def collect1490_1(values):
    return list(expand1490_1(values))
def summarize1490_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 86:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch1490_3(gate):
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
def traced1490_4(value: int) -> int:
    return value + 100


def decorate1490_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1490_4
def doubled1490_4(value: int) -> int:
    return value + value
def guarded1490_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
