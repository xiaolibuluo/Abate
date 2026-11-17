
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2129_0(value: int) -> int:
    return value + 37


def decorate2129_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2129_0
def doubled2129_0(value: int) -> int:
    return value + value
def traced2129_1(value: int) -> int:
    return value + 44


def decorate2129_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2129_1
def doubled2129_1(value: int) -> int:
    return value + value
def project2129_2(values):
    select = lambda item: item + 51
    return [select(item) for item in values if item != 0]
def guarded2129_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded2129_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize2129_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 72:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch2129_6(gate):
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
def guarded2129_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
