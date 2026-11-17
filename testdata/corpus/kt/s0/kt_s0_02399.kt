package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step19700")
fun step19700(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 7
}
@JvmName("step19701")
fun step19701(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 14
}
data class Pair19702<A, B>(val first: A, val second: B)

fun combine19702(left: Int, right: Int): Pair19702<Int, Int> {
    return Pair19702(left + 21, right + 35)
}
fun <T, R> List<T>.mapChecked19703(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair19704<A, B>(val first: A, val second: B)

fun combine19704(left: Int, right: Int): Pair19704<Int, Int> {
    return Pair19704(left + 35, right + 57)
}
object Registry19705 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19705 = "registry-42"
    }
}
