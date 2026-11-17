package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked2400(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry2401 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2401 = "registry-28"
    }
}
data class Pair2402<A, B>(val first: A, val second: B)

fun combine2402(left: Int, right: Int): Pair2402<Int, Int> {
    return Pair2402(left + 35, right + 64)
}
@JvmName("step2403")
fun step2403(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 42
}
data class Pair2404<A, B>(val first: A, val second: B)

fun combine2404(left: Int, right: Int): Pair2404<Int, Int> {
    return Pair2404(left + 49, right + 25)
}
data class Pair2405<A, B>(val first: A, val second: B)

fun combine2405(left: Int, right: Int): Pair2405<Int, Int> {
    return Pair2405(left + 56, right + 36)
}
class Counter2406 {
    var current: Int by Delegates.observable(63) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked2407(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair2408<A, B>(val first: A, val second: B)

fun combine2408(left: Int, right: Int): Pair2408<Int, Int> {
    return Pair2408(left + 77, right + 69)
}
class Counter2409 {
    var current: Int by Delegates.observable(84) { _, _, new -> println(new) }

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
