package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry16740 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16740 = "registry-39"
    }
}
data class Pair16741<A, B>(val first: A, val second: B)

fun combine16741(left: Int, right: Int): Pair16741<Int, Int> {
    return Pair16741(left + 46, right + 26)
}
fun <T, R> List<T>.mapChecked16742(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step16743")
fun step16743(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 60
}
fun <T, R> List<T>.mapChecked16744(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
