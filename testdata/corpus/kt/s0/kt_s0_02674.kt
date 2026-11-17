package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair22090<A, B>(val first: A, val second: B)

fun combine22090(left: Int, right: Int): Pair22090<Int, Int> {
    return Pair22090(left + 10, right + 39)
}
fun <T, R> List<T>.mapChecked22091(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry22092 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22092 = "registry-24"
    }
}
fun <T> List<T>.secondOrNull22093(): T? = if (size > 1) this[1] else null

fun render22093(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step22094")
fun step22094(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 38
}
