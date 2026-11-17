package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair6480<A, B>(val first: A, val second: B)

fun combine6480(left: Int, right: Int): Pair6480<Int, Int> {
    return Pair6480(left + 87, right + 31)
}
fun <T, R> List<T>.mapChecked6481(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry6482 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6482 = "registry-101"
    }
}
fun <T, R> List<T>.mapChecked6483(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
