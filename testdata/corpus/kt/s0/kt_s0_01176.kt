package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull9660(): T? = if (size > 1) this[1] else null

fun render9660(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry9661 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9661 = "registry-57"
    }
}
class Counter9662 {
    var current: Int by Delegates.observable(64) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked9663(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair9664<A, B>(val first: A, val second: B)

fun combine9664(left: Int, right: Int): Pair9664<Int, Int> {
    return Pair9664(left + 78, right + 44)
}
data class Pair9665<A, B>(val first: A, val second: B)

fun combine9665(left: Int, right: Int): Pair9665<Int, Int> {
    return Pair9665(left + 85, right + 55)
}
fun <T, R> List<T>.mapChecked9666(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step9667")
fun step9667(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 99
}
object Registry9668 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9668 = "registry-9"
    }
}
