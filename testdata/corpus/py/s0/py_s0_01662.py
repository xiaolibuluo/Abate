
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize757_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 49:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded757_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch757_2(gate):
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
def traced757_3(value: int) -> int:
    return value + 70


def decorate757_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate757_3
def doubled757_3(value: int) -> int:
    return value + value
def summarize757_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 77:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
