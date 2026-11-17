package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step19280")
fun step19280(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 43
}
object Registry19281 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19281 = "registry-50"
    }
}
fun <T, R> List<T>.mapChecked19282(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step19283")
fun step19283(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 64
}
fun <T> List<T>.secondOrNull19284(): T? = if (size > 1) this[1] else null

fun render19284(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair19285<A, B>(val first: A, val second: B)

fun combine19285(left: Int, right: Int): Pair19285<Int, Int> {
    return Pair19285(left + 78, right + 18)
}
data class Pair19286<A, B>(val first: A, val second: B)

fun combine19286(left: Int, right: Int): Pair19286<Int, Int> {
    return Pair19286(left + 85, right + 29)
}
