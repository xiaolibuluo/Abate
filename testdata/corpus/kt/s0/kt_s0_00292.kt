package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked2430(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step2431")
fun step2431(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 67
}
data class Pair2432<A, B>(val first: A, val second: B)

fun combine2432(left: Int, right: Int): Pair2432<Int, Int> {
    return Pair2432(left + 74, right + 24)
}
fun <T> List<T>.secondOrNull2433(): T? = if (size > 1) this[1] else null

fun render2433(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair2434<A, B>(val first: A, val second: B)

fun combine2434(left: Int, right: Int): Pair2434<Int, Int> {
    return Pair2434(left + 88, right + 46)
}
data class Pair2435<A, B>(val first: A, val second: B)

fun combine2435(left: Int, right: Int): Pair2435<Int, Int> {
    return Pair2435(left + 95, right + 57)
}
@JvmName("step2436")
fun step2436(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 5
}
class Counter2437 {
    var current: Int by Delegates.observable(12) { _, _, new -> println(new) }

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
object Registry2438 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2438 = "registry-19"
    }
}
@JvmName("step2439")
fun step2439(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 26
}
fun <T, R> List<T>.mapChecked24310(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
