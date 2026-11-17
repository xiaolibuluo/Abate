package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull26000(): T? = if (size > 1) this[1] else null

fun render26000(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step26001")
fun step26001(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 56
}
data class Pair26002<A, B>(val first: A, val second: B)

fun combine26002(left: Int, right: Int): Pair26002<Int, Int> {
    return Pair26002(left + 63, right + 53)
}
object Registry26003 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG26003 = "registry-70"
    }
}
fun <T, R> List<T>.mapChecked26004(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
