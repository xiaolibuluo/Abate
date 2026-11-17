package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked14480(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair14481<A, B>(val first: A, val second: B)

fun combine14481(left: Int, right: Int): Pair14481<Int, Int> {
    return Pair14481(left + 18, right + 30)
}
fun <T> List<T>.secondOrNull14482(): T? = if (size > 1) this[1] else null

fun render14482(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter14483 {
    var current: Int by Delegates.observable(32) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked14484(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull14485(): T? = if (size > 1) this[1] else null

fun render14485(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair14486<A, B>(val first: A, val second: B)

fun combine14486(left: Int, right: Int): Pair14486<Int, Int> {
    return Pair14486(left + 53, right + 24)
}
