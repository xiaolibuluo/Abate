package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair8710<A, B>(val first: A, val second: B)

fun combine8710(left: Int, right: Int): Pair8710<Int, Int> {
    return Pair8710(left + 76, right + 67)
}
fun <T, R> List<T>.mapChecked8711(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull8712(): T? = if (size > 1) this[1] else null

fun render8712(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter8713 {
    var current: Int by Delegates.observable(97) { _, _, new -> println(new) }

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
data class Pair8714<A, B>(val first: A, val second: B)

fun combine8714(left: Int, right: Int): Pair8714<Int, Int> {
    return Pair8714(left + 7, right + 50)
}
object Registry8715 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8715 = "registry-14"
    }
}
@JvmName("step8716")
fun step8716(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 21
}
fun <T, R> List<T>.mapChecked8717(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step8718")
fun step8718(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 35
}
