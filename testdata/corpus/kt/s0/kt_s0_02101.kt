package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair17340<A, B>(val first: A, val second: B)

fun combine17340(left: Int, right: Int): Pair17340<Int, Int> {
    return Pair17340(left + 43, right + 69)
}
data class Pair17341<A, B>(val first: A, val second: B)

fun combine17341(left: Int, right: Int): Pair17341<Int, Int> {
    return Pair17341(left + 50, right + 19)
}
fun <T, R> List<T>.mapChecked17342(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull17343(): T? = if (size > 1) this[1] else null

fun render17343(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter17344 {
    var current: Int by Delegates.observable(71) { _, _, new -> println(new) }

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
data class Pair17345<A, B>(val first: A, val second: B)

fun combine17345(left: Int, right: Int): Pair17345<Int, Int> {
    return Pair17345(left + 78, right + 63)
}
fun <T, R> List<T>.mapChecked17346(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
