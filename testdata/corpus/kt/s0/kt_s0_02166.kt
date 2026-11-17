package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry17830 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17830 = "registry-98"
    }
}
fun <T, R> List<T>.mapChecked17831(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair17832<A, B>(val first: A, val second: B)

fun combine17832(left: Int, right: Int): Pair17832<Int, Int> {
    return Pair17832(left + 15, right + 68)
}
fun <T, R> List<T>.mapChecked17833(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry17834 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17834 = "registry-29"
    }
}
@JvmName("step17835")
fun step17835(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 36
}
fun <T, R> List<T>.mapChecked17836(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
