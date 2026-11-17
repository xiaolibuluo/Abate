package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked10980(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked10981(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step10982")
fun step10982(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 34
}
data class Pair10983<A, B>(val first: A, val second: B)

fun combine10983(left: Int, right: Int): Pair10983<Int, Int> {
    return Pair10983(left + 41, right + 42)
}
data class Pair10984<A, B>(val first: A, val second: B)

fun combine10984(left: Int, right: Int): Pair10984<Int, Int> {
    return Pair10984(left + 48, right + 53)
}
@JvmName("step10985")
fun step10985(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 55
}
object Registry10986 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10986 = "registry-62"
    }
}
