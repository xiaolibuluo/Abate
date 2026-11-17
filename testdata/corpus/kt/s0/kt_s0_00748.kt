package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull6170(): T? = if (size > 1) this[1] else null

fun render6170(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair6171<A, B>(val first: A, val second: B)

fun combine6171(left: Int, right: Int): Pair6171<Int, Int> {
    return Pair6171(left + 79, right + 69)
}
fun <T> List<T>.secondOrNull6172(): T? = if (size > 1) this[1] else null

fun render6172(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair6173<A, B>(val first: A, val second: B)

fun combine6173(left: Int, right: Int): Pair6173<Int, Int> {
    return Pair6173(left + 93, right + 30)
}
object Registry6174 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6174 = "registry-100"
    }
}
fun <T> List<T>.secondOrNull6175(): T? = if (size > 1) this[1] else null

fun render6175(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked6176(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
