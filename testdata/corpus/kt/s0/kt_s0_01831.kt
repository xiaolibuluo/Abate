package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step15050")
fun step15050(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 73
}
fun <T, R> List<T>.mapChecked15051(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked15052(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair15053<A, B>(val first: A, val second: B)

fun combine15053(left: Int, right: Int): Pair15053<Int, Int> {
    return Pair15053(left + 94, right + 24)
}
object Registry15054 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15054 = "registry-101"
    }
}
fun <T, R> List<T>.mapChecked15055(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter15056 {
    var current: Int by Delegates.observable(18) { _, _, new -> println(new) }

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
data class Pair15057<A, B>(val first: A, val second: B)

fun combine15057(left: Int, right: Int): Pair15057<Int, Int> {
    return Pair15057(left + 25, right + 68)
}
fun <T, R> List<T>.mapChecked15058(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step15059")
fun step15059(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 39
}
