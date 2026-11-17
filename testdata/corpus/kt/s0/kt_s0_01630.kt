package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter13470 {
    var current: Int by Delegates.observable(56) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
data class Pair13471<A, B>(val first: A, val second: B)

fun combine13471(left: Int, right: Int): Pair13471<Int, Int> {
    return Pair13471(left + 63, right + 55)
}
class Counter13472 {
    var current: Int by Delegates.observable(70) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
data class Pair13473<A, B>(val first: A, val second: B)

fun combine13473(left: Int, right: Int): Pair13473<Int, Int> {
    return Pair13473(left + 77, right + 16)
}
fun <T> List<T>.secondOrNull13474(): T? = if (size > 1) this[1] else null

fun render13474(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair13475<A, B>(val first: A, val second: B)

fun combine13475(left: Int, right: Int): Pair13475<Int, Int> {
    return Pair13475(left + 91, right + 38)
}
