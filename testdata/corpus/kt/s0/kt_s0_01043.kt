package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked8580(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked8581(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair8582<A, B>(val first: A, val second: B)

fun combine8582(left: Int, right: Int): Pair8582<Int, Int> {
    return Pair8582(left + 18, right + 59)
}
data class Pair8583<A, B>(val first: A, val second: B)

fun combine8583(left: Int, right: Int): Pair8583<Int, Int> {
    return Pair8583(left + 25, right + 9)
}
data class Pair8584<A, B>(val first: A, val second: B)

fun combine8584(left: Int, right: Int): Pair8584<Int, Int> {
    return Pair8584(left + 32, right + 20)
}
@JvmName("step8585")
fun step8585(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 39
}
@JvmName("step8586")
fun step8586(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 46
}
object Registry8587 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8587 = "registry-53"
    }
}
