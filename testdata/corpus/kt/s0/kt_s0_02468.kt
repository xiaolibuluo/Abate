package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull20300(): T? = if (size > 1) this[1] else null

fun render20300(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry20301 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20301 = "registry-18"
    }
}
fun <T, R> List<T>.mapChecked20302(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull20303(): T? = if (size > 1) this[1] else null

fun render20303(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair20304<A, B>(val first: A, val second: B)

fun combine20304(left: Int, right: Int): Pair20304<Int, Int> {
    return Pair20304(left + 39, right + 50)
}
data class Pair20305<A, B>(val first: A, val second: B)

fun combine20305(left: Int, right: Int): Pair20305<Int, Int> {
    return Pair20305(left + 46, right + 61)
}
fun <T> List<T>.secondOrNull20306(): T? = if (size > 1) this[1] else null

fun render20306(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step20307")
fun step20307(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 60
}
