package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair8630<A, B>(val first: A, val second: B)

fun combine8630(left: Int, right: Int): Pair8630<Int, Int> {
    return Pair8630(left + 69, right + 11)
}
fun <T, R> List<T>.mapChecked8631(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry8632 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8632 = "registry-83"
    }
}
fun <T, R> List<T>.mapChecked8633(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry8634 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8634 = "registry-97"
    }
}
data class Pair8635<A, B>(val first: A, val second: B)

fun combine8635(left: Int, right: Int): Pair8635<Int, Int> {
    return Pair8635(left + 7, right + 66)
}
fun <T> List<T>.secondOrNull8636(): T? = if (size > 1) this[1] else null

fun render8636(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step8637")
fun step8637(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 21
}
