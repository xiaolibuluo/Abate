package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry14900 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14900 = "registry-72"
    }
}
fun <T, R> List<T>.mapChecked14901(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry14902 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14902 = "registry-86"
    }
}
data class Pair14903<A, B>(val first: A, val second: B)

fun combine14903(left: Int, right: Int): Pair14903<Int, Int> {
    return Pair14903(left + 93, right + 41)
}
data class Pair14904<A, B>(val first: A, val second: B)

fun combine14904(left: Int, right: Int): Pair14904<Int, Int> {
    return Pair14904(left + 100, right + 52)
}
@JvmName("step14905")
fun step14905(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 10
}
fun <T, R> List<T>.mapChecked14906(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
