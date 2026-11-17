package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair7610<A, B>(val first: A, val second: B)

fun combine7610(left: Int, right: Int): Pair7610<Int, Int> {
    return Pair7610(left + 101, right + 29)
}
fun <T> List<T>.secondOrNull7611(): T? = if (size > 1) this[1] else null

fun render7611(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry7612 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7612 = "registry-18"
    }
}
fun <T, R> List<T>.mapChecked7613(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull7614(): T? = if (size > 1) this[1] else null

fun render7614(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry7615 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7615 = "registry-39"
    }
}
data class Pair7616<A, B>(val first: A, val second: B)

fun combine7616(left: Int, right: Int): Pair7616<Int, Int> {
    return Pair7616(left + 46, right + 34)
}
