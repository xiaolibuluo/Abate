package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair6630<A, B>(val first: A, val second: B)

fun combine6630(left: Int, right: Int): Pair6630<Int, Int> {
    return Pair6630(left + 88, right + 14)
}
data class Pair6631<A, B>(val first: A, val second: B)

fun combine6631(left: Int, right: Int): Pair6631<Int, Int> {
    return Pair6631(left + 95, right + 25)
}
fun <T> List<T>.secondOrNull6632(): T? = if (size > 1) this[1] else null

fun render6632(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry6633 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6633 = "registry-12"
    }
}
fun <T, R> List<T>.mapChecked6634(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair6635<A, B>(val first: A, val second: B)

fun combine6635(left: Int, right: Int): Pair6635<Int, Int> {
    return Pair6635(left + 26, right + 69)
}
fun <T> List<T>.secondOrNull6636(): T? = if (size > 1) this[1] else null

fun render6636(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
