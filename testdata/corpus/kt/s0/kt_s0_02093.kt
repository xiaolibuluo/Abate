package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked17280(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair17281<A, B>(val first: A, val second: B)

fun combine17281(left: Int, right: Int): Pair17281<Int, Int> {
    return Pair17281(left + 69, right + 38)
}
fun <T> List<T>.secondOrNull17282(): T? = if (size > 1) this[1] else null

fun render17282(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter17283 {
    var current: Int by Delegates.observable(83) { _, _, new -> println(new) }

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
data class Pair17284<A, B>(val first: A, val second: B)

fun combine17284(left: Int, right: Int): Pair17284<Int, Int> {
    return Pair17284(left + 90, right + 10)
}
