package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull24550(): T? = if (size > 1) this[1] else null

fun render24550(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair24551<A, B>(val first: A, val second: B)

fun combine24551(left: Int, right: Int): Pair24551<Int, Int> {
    return Pair24551(left + 14, right + 64)
}
class Counter24552 {
    var current: Int by Delegates.observable(21) { _, _, new -> println(new) }

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
class Counter24553 {
    var current: Int by Delegates.observable(28) { _, _, new -> println(new) }

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
