package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry7450 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7450 = "registry-87"
    }
}
fun <T, R> List<T>.mapChecked7451(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair7452<A, B>(val first: A, val second: B)

fun combine7452(left: Int, right: Int): Pair7452<Int, Int> {
    return Pair7452(left + 101, right + 61)
}
object Registry7453 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7453 = "registry-11"
    }
}
