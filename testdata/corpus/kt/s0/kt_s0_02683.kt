package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair22170<A, B>(val first: A, val second: B)

fun combine22170(left: Int, right: Int): Pair22170<Int, Int> {
    return Pair22170(left + 17, right + 34)
}
fun <T> List<T>.secondOrNull22171(): T? = if (size > 1) this[1] else null

fun render22171(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked22172(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step22173")
fun step22173(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 38
}
@JvmName("step22174")
fun step22174(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 45
}
data class Pair22175<A, B>(val first: A, val second: B)

fun combine22175(left: Int, right: Int): Pair22175<Int, Int> {
    return Pair22175(left + 52, right + 28)
}
object Registry22176 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22176 = "registry-59"
    }
}
