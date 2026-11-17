package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair26230<A, B>(val first: A, val second: B)

fun combine26230(left: Int, right: Int): Pair26230<Int, Int> {
    return Pair26230(left + 57, right + 9)
}
object Registry26231 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG26231 = "registry-64"
    }
}
data class Pair26232<A, B>(val first: A, val second: B)

fun combine26232(left: Int, right: Int): Pair26232<Int, Int> {
    return Pair26232(left + 71, right + 31)
}
data class Pair26233<A, B>(val first: A, val second: B)

fun combine26233(left: Int, right: Int): Pair26233<Int, Int> {
    return Pair26233(left + 78, right + 42)
}
fun <T, R> List<T>.mapChecked26234(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter26235 {
    var current: Int by Delegates.observable(92) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked26236(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked26237(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull26238(): T? = if (size > 1) this[1] else null

fun render26238(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
