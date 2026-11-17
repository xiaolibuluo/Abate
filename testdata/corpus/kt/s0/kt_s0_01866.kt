package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry15340 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15340 = "registry-62"
    }
}
object Registry15341 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15341 = "registry-69"
    }
}
fun <T, R> List<T>.mapChecked15342(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step15343")
fun step15343(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 83
}
data class Pair15344<A, B>(val first: A, val second: B)

fun combine15344(left: Int, right: Int): Pair15344<Int, Int> {
    return Pair15344(left + 90, right + 55)
}
@JvmName("step15345")
fun step15345(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 97
}
fun <T, R> List<T>.mapChecked15346(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
