package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked10910(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked10911(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step10912")
fun step10912(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 40
}
object Registry10913 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10913 = "registry-47"
    }
}
data class Pair10914<A, B>(val first: A, val second: B)

fun combine10914(left: Int, right: Int): Pair10914<Int, Int> {
    return Pair10914(left + 54, right + 65)
}
