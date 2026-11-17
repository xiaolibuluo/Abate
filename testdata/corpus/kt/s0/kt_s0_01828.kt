package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull15030(): T? = if (size > 1) this[1] else null

fun render15030(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry15031 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15031 = "registry-54"
    }
}
class Counter15032 {
    var current: Int by Delegates.observable(61) { _, _, new -> println(new) }

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
class Counter15033 {
    var current: Int by Delegates.observable(68) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked15034(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step15035")
fun step15035(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 82
}
data class Pair15036<A, B>(val first: A, val second: B)

fun combine15036(left: Int, right: Int): Pair15036<Int, Int> {
    return Pair15036(left + 89, right + 43)
}
data class Pair15037<A, B>(val first: A, val second: B)

fun combine15037(left: Int, right: Int): Pair15037<Int, Int> {
    return Pair15037(left + 96, right + 54)
}
