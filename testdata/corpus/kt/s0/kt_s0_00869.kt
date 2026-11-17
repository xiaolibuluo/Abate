package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair7120<A, B>(val first: A, val second: B)

fun combine7120(left: Int, right: Int): Pair7120<Int, Int> {
    return Pair7120(left + 46, right + 52)
}
class Counter7121 {
    var current: Int by Delegates.observable(53) { _, _, new -> println(new) }

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
data class Pair7122<A, B>(val first: A, val second: B)

fun combine7122(left: Int, right: Int): Pair7122<Int, Int> {
    return Pair7122(left + 60, right + 13)
}
fun <T> List<T>.secondOrNull7123(): T? = if (size > 1) this[1] else null

fun render7123(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair7124<A, B>(val first: A, val second: B)

fun combine7124(left: Int, right: Int): Pair7124<Int, Int> {
    return Pair7124(left + 74, right + 35)
}
data class Pair7125<A, B>(val first: A, val second: B)

fun combine7125(left: Int, right: Int): Pair7125<Int, Int> {
    return Pair7125(left + 81, right + 46)
}
class Counter7126 {
    var current: Int by Delegates.observable(88) { _, _, new -> println(new) }

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
