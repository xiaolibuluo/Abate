package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry11370 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11370 = "registry-42"
    }
}
data class Pair11371<A, B>(val first: A, val second: B)

fun combine11371(left: Int, right: Int): Pair11371<Int, Int> {
    return Pair11371(left + 49, right + 49)
}
@JvmName("step11372")
fun step11372(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 56
}
@JvmName("step11373")
fun step11373(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 63
}
fun <T, R> List<T>.mapChecked11374(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
